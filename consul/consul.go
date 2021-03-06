package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const IP string = "192.168.22.154" // Local host IP

func main()  {
	if err := app(); err != nil {
		log.WithError(err).Fatal("application failed with error")
	}
}

func app() error {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 2)

	go func() {
		errCh <- startServer()
	}()
	go func() {
		consul()
		time.Sleep(5 * time.Minute)
		cancel()
	}()

	select {
	case <-ctx.Done():
		log.Info("Service shutdown by ctx.Done")
		return nil
	case err := <-errCh:
		return err
	}
}

func consulRegisterService(client *api.Client) error {
	agent := client.Agent()
	service := &api.AgentServiceRegistration{
		Kind:              "",
		ID:                "MyAwesomeService1",
		Name:              "MyAwesomeService",
		Tags:              []string{"my", "awesome", "service"},
		Address:           "192.168.1.44",
		Port:              8080,
		Check:             &api.AgentServiceCheck{
			Name:                           "HEALTH CHECK",
			Interval:                       "2s",
			Timeout:                        "1s",
			HTTP:                           fmt.Sprintf("http://%s:8080/health", IP),
		},
	}
	return agent.ServiceRegister(service)
}
func consulDeregisterService(client *api.Client) error {
	agent := client.Agent()
	return agent.ServiceDeregister("MyAwesomeService1")
}
func consul()  {
	client, err := api.NewClient(api.DefaultConfig())

	if err != nil {
		panic(err)
	}

	err = consulDeregisterService(client)
	if err != nil {
		panic(err)
	}
	err = consulRegisterService(client)
	if err != nil {
		panic(err)
	}

	kv := client.KV()

	p := &api.KVPair{Key: "FOO", Value: []byte("BAR")}
	_, err = kv.Put(p, nil)

	if err != nil {
		panic(err)
	}
}

func startServer() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", hcHandler)


	if err := http.ListenAndServe(fmt.Sprintf("%s:8080", IP), mux); err != nil {
		return err
	}

	return nil
}

func hcHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	w.WriteHeader(http.StatusOK)
}
