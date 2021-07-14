package main

import (
	"context"
	"net/http"
)
import log "github.com/sirupsen/logrus"

func main()  {

	if err := app(); err != nil {
		log.WithError(err).Fatal("application failed with error")
	}
}

func app() error {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)

	go func() {
		errCh <- startServer(cancel)
	}()

	select {
	case <-ctx.Done():
		log.Info("Service shutdown by ctx.Done")
		return nil
	case err := <-errCh:
		return err
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request)  {
	v, err := w.Write([]byte("Just index"))
	log.Info("Index handler response: ", v, err)
}

func nfHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}


func startServer(cancel context.CancelFunc) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/not-found", nfHandler)
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		cancel()
	})

	if err := http.ListenAndServe("localhost:9876", mux); err != nil {
		return err
	}

	return nil
}