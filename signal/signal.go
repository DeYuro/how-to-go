package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	ctx, cancel := context.WithCancel(context.Background())

	go handleSignal(cancel)
	for {
		select {
		case <- ctx.Done():
			return
		}
	}
}

// kill USR2 [PID] while run
func handleSignal(cancel context.CancelFunc) {
	count := 3
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGUSR2)
	for range ch {
		fmt.Println("GOT SIGUSR2")
		if count > 0 {
			count--
			fmt.Println(count, " attempt left")
		} else {
			cancel()
		}
	}
}