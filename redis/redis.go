package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type svc interface {
	checkFinish(ctx context.Context, errs chan<- error, cancelFunc context.CancelFunc)
	getClient() *redis.Client // just for manual, may use repo pattern instead
	demo(ctx context.Context, errs chan<- error)
}

type service struct {
	someDeps interface{}
	redis 	 *redis.Client
}

func newService() svc {
	return service{
		someDeps: nil,
		redis:    redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		}),
	}
}

func (s service) checkFinish(ctx context.Context, errs chan<- error, cancelFunc context.CancelFunc) {

	ticker := time.NewTicker(time.Second)

	cli := s.getClient()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			val, err := cli.Get(ctx, "finish").Result()
			fmt.Printf("Got <%v> as finish on tick\n", val)

			if err != nil && err.Error() != "redis: nil" {
				errs <- err
			}

			if val == "end" {
				cancelFunc()
			}
		}
	}
}

func (s service) getClient() *redis.Client {
	return s.redis
}

func main() {
	if err := app(); err != nil {
		log.Fatal("some err in app:", err)
	}
}

func app() error  {
	ctx, cancel := context.WithCancel(context.Background())
	errs := make(chan error)

	svc := newService()

	go svc.demo(ctx, errs)

	go svc.checkFinish(ctx, errs, cancel)
	select {
	case <-ctx.Done():
		fmt.Println("Finish by context")
		return nil
	case err := <-errs:
		return err
	}
}


func (s service) demo(ctx context.Context, errs chan<- error) {
	cli := s.getClient()

	err := cli.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		errs <- err
	}

	val, err := cli.Get(ctx, "foo").Result()
	if err != nil {
		errs <- err
	}

	val, err = cli.Get(ctx, "not exist").Result()
	if err != nil && err.Error() != "redis: nil" {
		errs <- err
	}

	fmt.Printf("key:'not exist', value:%s, error:%s\n", val, err.Error()) //error != nil

	time.Sleep(5 * time.Second)

	err = cli.Set(ctx, "finish", "end", time.Second * 5).Err()
	if err != nil {
		errs <- err
	}
}