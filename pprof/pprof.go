package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"sync"
)

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


func startServer(cancel context.CancelFunc) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/long", longHandler)
	mux.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		cancel()
	})
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	if err := http.ListenAndServe("localhost:9876", mux); err != nil {
		return err
	}

	return nil
}
func indexHandler(w http.ResponseWriter, r *http.Request)  {
	v, err := w.Write([]byte("Just index"))
	log.Info("Index handler response: ", v, err)
}

func longHandler(w http.ResponseWriter, r *http.Request) {

	var wg sync.WaitGroup
	rng := rand.Intn(1000000 - 500) +500

	for i := 0; i < rng; i++ {
		wg.Add(1)

		go func() {
			counter := 0
			for counter < math.MaxInt32 {
				counter += rand.Intn(10 - 1) + 1
			}
			wg.Done()
		}()
	}

	wg.Wait()

	w.WriteHeader(http.StatusOK)
}