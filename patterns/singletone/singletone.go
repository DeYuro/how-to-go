package main

import (
	"fmt"
	"github.com/how-to-go/patterns/singletone/withSync"
)

func main() {
	for i := 0; i < 10; i++ {
		//go native.GetInstance()
		go withSync.GetInstance()
	}
	fmt.Scanln()
}
