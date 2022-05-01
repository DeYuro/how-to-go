package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"golang.org/x/exp/constraints"
)

func main() {
	s, ss := "a", []string{"a", "b"}
	i, is := 1, []int{1, 2}

	println(contains(s, ss))
	println(contains(i, is))

	a, b := 1, 20
	println(max(a, b))
	as, bs := "lorem", "dorem"
	println(max(as, bs))

	printErr(InvalidNameError)
	printErr(ResourceNotFoundError)
	printErr(EmptyValueError)
	printErr(ServiceError)

	player, err := createFighter[uint](human, 100, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Human(value receiver) created and got have: %d\n", player.getHp())       // 100
	player.takeDamage(42)                                                                // pointless coz value receiver
	fmt.Printf("Human(value receiver) take damage and have HP: %d\n", player.getHp())    // 100
	player.restoreHp()                                                                   // pointless coz value receiver
	fmt.Printf("Human(value receiver) restore health and have HP: %d\n", player.getHp()) // 100

	bot, err := createFighter[uint](ai, 1000, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("AI(pointer receiver) created and got have: %d\n", bot.getHp()) //1000
	bot.takeDamage(444)
	fmt.Printf("AI(pointer receiver) take damage and have HP: %d\n", bot.getHp()) //556
	bot.restoreHp()
	fmt.Printf("AI(pointer receiver) restore health and have HP: %d\n", bot.getHp()) //999
}

func contains[T comparable](value T, sliceOfValue []T) bool {
	for _, v := range sliceOfValue {
		if v == value {
			return true
		}
	}

	return false
}

type CanCompare interface {
	constraints.Ordered
}

func max[T CanCompare](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type Err string
type InternalErr string

const (
	InvalidNameError Err = `InvalidName`
	EmptyValueError  Err = `EmptyValue`
)

const (
	ResourceNotFoundError InternalErr = `ResourceNotFound`
	ServiceError          InternalErr = `ServiceError`
)

type AllErr interface {
	Err | InternalErr
}

func printErr[T AllErr](error T) {
	println(error)
}
