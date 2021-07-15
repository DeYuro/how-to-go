package main

import (
	"context"
	"fmt"
	"time"
)

func main()  {
	fmt.Println("func end by context timeout early than by condition")
	withTimeout(2 * time.Second, 7 * time.Second)
	fmt.Println("func end by condition early than by context timeout")
	withTimeout(7 * time.Second, 2 * time.Second)
}

func withTimeout(ctxTimeout, condTimeout time.Duration)  {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)

	defer cancel()

	end := make(chan interface{})

	go func() {
		time.Sleep(condTimeout)
		end <- struct {}{}
	}()

	select {
	case <- ctx.Done():
		fmt.Println("Context timeout")
	case <- end:
		fmt.Println("End without context timeout")
	}
}
