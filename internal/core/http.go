package core

// HTTPServer defines operations for working with
// HTTP server.
type HTTPServer interface {
	// Run starts the HTTP server.
	Run(addr string) error

	// Stop stops the HTTP server.
	Stop() error
}
