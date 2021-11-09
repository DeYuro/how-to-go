package main

import (
	"context"
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/k0kubun/pp"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
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
	fmt.Println("Use gopkg.in/yaml.v2")
	config, err := loadConfigYml("config/config.yml")
	if err != nil {
		return err
	}

	pp.Println(config)

	fmt.Printf("Now sleep %f seconds\n", config.Foo.Interval.Seconds())
	time.Sleep(config.Foo.Interval)
	fmt.Println("sleep over")


	fmt.Println("Use github.com/jinzhu/configor")
	config, err = loadConfigConfigor("config/config.yml")
	if err != nil {
		return err
	}

	pp.Println(config)

	fmt.Printf("Now sleep %f seconds\n", config.Foo.Interval.Seconds())
	time.Sleep(config.Foo.Interval)
	fmt.Println("sleep over")

	return nil
}

func loadConfigConfigor(pathToConfig string) (*Config, error) {
	config := new(Config)

	err := configor.Load(config, pathToConfig)
	if err != nil {
		return nil,err
	}

	return config, nil
}

func loadConfigYml(pathToConfig string) (*Config, error) {
	config := new(Config)
	file, err := os.Open(pathToConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}