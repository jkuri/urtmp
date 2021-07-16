package main

import (
	"log"
	"os"

	"github.com/jkuri/rtmp-server/internal/http"
	"github.com/jkuri/rtmp-server/internal/rtmp"
)

func main() {
	errch := make(chan error, 1)
	rtmp := rtmp.New()
	http := http.New(rtmp.Handler())

	go func() {
		if err := rtmp.Run("0.0.0.0:1935"); err != nil {
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
