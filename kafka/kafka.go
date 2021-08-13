package main

import (
	"github.com/how-to-go/kafka/consumer"
	"github.com/how-to-go/kafka/producer"
)

func main() {
	go producer.Produce()
	go consumer.Consume()

	println("foo")
	for {

	}
}