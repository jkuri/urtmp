package rtmp

import (
	"io"
	"net/http"
)

type writeFlusher struct {
	httpflusher http.Flusher
	io.Writer
}

func (w writeFlusher) Flush() error {
	w.httpflusher.Flush()
	return nil
}
