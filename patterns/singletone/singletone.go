package main

import (
	"fmt"
	"github.com/how-to-go/patterns/singletone/native"
)

func main() {
	for i := 0; i < 10; i++ {
		go native.GetInstance()
	}
	fmt.Scanln()
}
