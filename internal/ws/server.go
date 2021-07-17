package ws

import (
	"io"
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gobwas/ws"
	"github.com/jkuri/urtmp/internal/core"
)

const addr = "127.0.0.1:8081"

// server contains options and methods for running zero-copy
// websocket server on straight TCP connection, use in a
// combination with UpstreamHandler.
type server struct {
	mu        sync.RWMutex
	logger    *log.Logger
	ioTimeout time.Duration
	listener  net.Listener
	clients   []*client
	quitch    chan error
}

// New initializes and returns a new websocket server instance.
func New() core.WebSocket {
	return &server{
		logger:    log.Default(),
		ioTimeout: 100 * time.Millisecond,
		quitch:    make(chan error, 1),
	}
}

func (s *server) Run() error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener

	s.logger.Printf("websocket server listening on ws://%s", addr)

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				s.logger.Printf("error accepting incoming websocket connection: %s", err.Error())
				break
			}
			go s.handleConnection(conn)
		}
	}()

	return <-s.quitch
}

func (s *server) Stop() error {
	if err := s.listener.Close(); err != nil {
		s.quitch <- err
		return err
	}

	s.quitch <- nil
	return nil
}

func (s *server) UpstreamHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		peer, err := net.Dial("tcp", addr)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		if err := r.Write(peer); err != nil {
			w.WriteHeader(http.StatusBadGateway)
			return
		}
		hj, ok := w.(http.Hijacker)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		conn, _, err := hj.Hijack()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		go pipe(peer, conn)
		go pipe(conn, peer)
	})
}

func (s *server) Broadcast(sub string, data map[string]interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, c := range s.clients {
		for _, subscription := range c.subs {
			if subscription == sub {
				if err := c.send(sub, data); err != nil {
					s.logger.Printf("error sending message to client due: %s", err.Error())
				}
				break
			}
		}
	}
}

func (s *server) handleConnection(conn net.Conn) {
	var err error

	header := ws.HandshakeHeaderHTTP(http.Header{
		"X-uRTMP-Version": []string{"μRTMP"},
	})

	upgrader := ws.Upgrader{
		OnHost: func(host []byte) error {
			return nil
		},
		OnHeader: func(key, value []byte) error {
			return nil
		},
		OnBeforeUpgrade: func() (ws.HandshakeHeader, error) {
			return header, nil
		},
	}

	if err != nil {
		s.logger.Printf("websocket connection not upgraded: %s", err.Error())
		return
	}

	if _, err := upgrader.Upgrade(conn); err != nil {
		s.logger.Printf("error upgrading websocket connection %s: %s", nameConn(conn), err.Error())
		return
	}

	client := s.register(conn)
	if err := s.initClient(client); err != nil {
		s.unregister(client)
		s.logger.Printf("websocket user unregistered")
	}
}

func (s *server) register(conn net.Conn) *client {
	client := &client{
		conn: conn,
		c:    conn,
	}

	s.mu.Lock()
	s.clients = append(s.clients, client)
	s.mu.Unlock()

	s.logger.Printf("websocket user registered")

	return client
}

func (s *server) unregister(client *client) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, c := range s.clients {
		if c == client {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			break
		}
	}
}

func (s *server) initClient(client *client) error {
	for {
		msg, err := client.receive()
		if err != nil {
			return err
		}

		if msg == nil {
			continue
		}

		if msg.Type == "subscribe" {
			if data, ok := msg.Data["sub"].(string); ok {
				client.subscribe(data)
			}
		}

		if msg.Type == "unsubscribe" {
			if data, ok := msg.Data["sub"].(string); ok {
				client.unsubscribe(data)
			}
		}
	}
}

func pipe(c1 net.Conn, c2 net.Conn) {
	defer c1.Close()
	defer c2.Close()
	io.Copy(c1, c2)
}

func nameConn(conn net.Conn) string {
	return conn.LocalAddr().String() + " <> " + conn.RemoteAddr().String()
}
