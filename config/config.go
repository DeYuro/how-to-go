package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

type Config struct {
	Foo Foo `yaml:"foo"`
	Bar Bar `yaml:"bar"`
}

type Foo struct {
	Text     string        `yaml:"text"`
	Numbers  int           `yaml:"numbers"`
	Interval time.Duration `yaml:"interval"`
}

type Bar struct {
	Foo *Foo `yaml:"foo"`
}

func main() {
	if err := app(); err != nil {
		log.WithError(err).Fatal("application failed with error")
	}
}

func app() error {
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error)

	go func() {
		errCh <- start(cancel)
	}()

	select {
	case <-ctx.Done():
		log.Info("Service shutdown by ctx.Done")
		return nil
	case err := <-errCh:
		return err
	}
}

func start(cancel context.CancelFunc) error {
	pp.Println(loadConfig("./config.yml"))
	return nil
}

func loadConfig(pathToConfig string) *Config {
	config := new(Config)

	err := configor.Load(config, pathToConfig)
	if err != nil {
		fmt.Println("Failed to load configuration file " + err.Error())
		os.Exit(1)
	}

	return config
}
