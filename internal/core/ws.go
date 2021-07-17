package core

import "net/http"

// WebSocket defines operations on working with websocket
// server. It also includes upstream handler which is used
// for proxying websocket connection with zero-copy upgrades.
type WebSocket interface {
	// Run starts the websocker server.
	Run() error

	// Stop stops the websocket server.
	Stop() error

	// UpstreamHandler returns an http.HandlerFunc to use in
	// custom router and it is responsible for proxying websocket
	// connections with zero-copy upgrades.
	UpstreamHandler() http.HandlerFunc

	// Broadcast emits message to all subscribers.
	Broadcast(sub string, data map[string]interface{})
}
