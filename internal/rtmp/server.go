package rtmp

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/jkuri/rtmp-server/internal/core"
	"github.com/jkuri/rtmp-server/pkg/render"
	"github.com/nareix/joy5/format/flv"
	"github.com/nareix/joy5/format/rtmp"
)

type server struct {
	logger   *log.Logger
	listener net.Listener
	server   *rtmp.Server
	streams  *streams
	quitch   chan error
}

func New() core.RTMPServer {
	return &server{
		logger: log.Default(),
		quitch: make(chan error, 1),
	}
}

func (s *server) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.listener = listener

	s.logger.Printf("RTMP server listening on rtmp://%s", addr)

	s.server = rtmp.NewServer()
	s.server.LogEvent = func(c *rtmp.Conn, nc net.Conn, e int) {
		s.logger.Printf("rtmp event %s <-> %s event: %s", nc.LocalAddr(), nc.RemoteAddr(), rtmp.EventString[e])
	}

	s.streams = newStreams()

	s.server.HandleConn = func(c *rtmp.Conn, nc net.Conn) {
		stream, remove := s.streams.add(c.URL.Path, c.Publishing)
		defer remove()

		if c.Publishing {
			stream.setPub(c)
		} else {
			stream.addSub(c.CloseNotify(), c)
		}
	}

	go func() {
		for {
			nc, err := s.listener.Accept()
			if err != nil {
				s.logger.Printf("error accepting incoming rtmp connetion: %s", err.Error())
				break
			}
			go s.server.HandleNetConn(nc)
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

func (s *server) Handler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/live/", s.liveHandler())
	mux.HandleFunc("/api/v1/streams", s.streamsHandler())
	return mux
}

func (s *server) liveHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		splitted := strings.Split(r.URL.Path, "/")
		relPath := append(splitted[:1], splitted[3:]...)
		id := strings.Join(relPath, "/")

		if !s.streams.exists(id) {
			render.NotFoundError(w, "stream does not exists")
			return
		}

		stream, remove := s.streams.add(id, false)

		w.Header().Set("Content-Type", "video/x-flv")
		w.Header().Set("Transfer-Encoding", "chunked")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		flusher := w.(http.Flusher)
		flusher.Flush()
		muxer := flv.NewMuxer(writeFlusher{httpflusher: flusher, Writer: w})
		closech := make(chan bool)

		go func() {
			<-r.Context().Done()
			remove()
			closech <- true
		}()

		stream.addSub(closech, muxer)
	}
}

func (s *server) streamsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, http.StatusOK, s.streams.list())
	}
}
