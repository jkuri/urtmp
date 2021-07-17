package http

import (
	"log"
	"net"
	"net/http"

	"github.com/jkuri/statik/fs"
	"github.com/jkuri/urtmp/internal/core"
	_ "github.com/jkuri/urtmp/internal/ui"
)

// server extends net/http Server with graceful shutdowns.
type server struct {
	*http.Server
	logger   *log.Logger
	api      *http.ServeMux
	listener net.Listener
	ws       core.WebSocket
}

// New creates a new HTTP server instance.
func New(api *http.ServeMux, ws core.WebSocket) core.HTTPServer {
	return &server{
		Server: &http.Server{},
		api:    api,
		logger: log.Default(),
		ws:     ws,
	}
}

// Run starts HTTP server instance and listens of specified port.
func (s *server) Run(addr string) error {
	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s.Handler = s.handler()

	s.logger.Printf("HTTP server listening on http://%s", addr)
	return s.Serve(s.listener)
}

// Stop terminates HTTP server and closes listener.
func (s *server) Stop() error {
	s.logger.Printf("stopping down HTTP server")
	return s.Close()
}

func (s *server) handler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", s.api)
	mux.Handle("/ws", s.ws.UpstreamHandler())
	mux.HandleFunc("/", s.ui())
	return mux
}

func (s *server) ui() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		root, _ := fs.New()
		fs := http.FileServer(&wrapper{root})
		fs.ServeHTTP(w, r)
	}
}
