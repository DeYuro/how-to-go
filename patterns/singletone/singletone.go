package main

import (
	"fmt"
	"github.com/how-to-go/patterns/singletone/initInstance"
)

func main() {
	for i := 0; i < 10; i++ {
		go initInstance.GetInstance()
	}
	fmt.Scanln()
}
