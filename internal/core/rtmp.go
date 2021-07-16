package core

import "net/http"

// Server defines operations for working with
// RTMP server.
type RTMPServer interface {
	// Run starts the RTMP server.
	Run(addr string) error

	// Stop stops the RTMP server.
	Stop() error

	// Handler returns an *http.ServeMux, the API router.
	Handler() *http.ServeMux
}
