package http

import (
	"net/http"
	"os"
	"path"
)

type wrapper struct {
	assets http.FileSystem
}

// Open returns file from http FileSystem by path,
// if file is not found fallbacks to /index.html.
func (w *wrapper) Open(name string) (http.File, error) {
	ret, err := w.assets.Open(name)
	if !os.IsNotExist(err) || path.Ext(name) != "" {
		return ret, err
	}

	return w.assets.Open("/index.html")
}
