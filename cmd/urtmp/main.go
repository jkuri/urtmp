package main

import (
	"log"
	"os"

	"github.com/jkuri/urtmp/internal/http"
	"github.com/jkuri/urtmp/internal/rtmp"
	"github.com/jkuri/urtmp/internal/ws"
)

func main() {
	errch := make(chan error, 1)
	ws := ws.New()
	rtmp := rtmp.New(ws)
	http := http.New(rtmp.Handler(), ws)

	go func() {
		if err := rtmp.Run("0.0.0.0:1935"); err != nil {
			errch <- err
		}
	}()

	go func() {
		if err := ws.Run(); err != nil {
			errch <- err
		}
	}()

	go func() {
		if err := http.Run("0.0.0.0:8080"); err != nil {
			errch <- err
		}
	}()

	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
