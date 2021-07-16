package http

import (
	"log"
	"net"
	"net/http"

	"github.com/jkuri/rtmp-server/internal/core"
)

// server extends net/http Server with graceful shutdowns.
type server struct {
	*http.Server
	logger   *log.Logger
	listener net.Listener
}

// New creates a new HTTP server instance.
func New(handler http.HandlerFunc) core.HTTPServer {
	return &server{
		Server: &http.Server{Handler: handler},
		logger: log.Default(),
	}
}

// Run starts HTTP server instance and listens of specified port.
func (s server) Run(addr string) error {
	var err error
	s.listener, err = net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	s.logger.Printf("HTTP server listening on http://%s", addr)
	return s.Serve(s.listener)
}

// Stop terminates HTTP server and closes listener.
func (s server) Stop() error {
	s.logger.Printf("stopping down HTTP server")
	return s.Close()
}
