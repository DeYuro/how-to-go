package main

import (
	"errors"
	"fmt"
	"log"
)

type FighterType string

const (
	human FighterType = "human"
	ai    FighterType = "ai"
)

type Fighter interface {
	getHp() int
	restoreHp()
	takeDamage(damage int)
	void()
}

//Human value type receiver methods
type Human struct {
	Hp          int
	MaxHp       int
	RestoreStep int
}

func (h Human) getHp() int {
	return h.Hp
}

func (h Human) restoreHp() {
	if h.MaxHp < (h.Hp + h.RestoreStep) {
		h.Hp += h.RestoreStep
	}

	h.Hp = h.MaxHp
}

func (h Human) takeDamage(damage int) {
	if (h.Hp - damage) < 0 {
		h.Hp = 0
	}

	h.Hp -= damage
}

// void pointer receiver invalid for human{....}  but valid for &human{...} other methods valid also
func (h *Human) void() {

}

// AI pointer type receiver methods
type AI struct {
	Hp    int
	MaxHp int
}

func (a *AI) getHp() int {
	return a.Hp
}

func (a *AI) restoreHp() {
	a.Hp = a.MaxHp - 1
}

func (a *AI) takeDamage(damage int) {
	if (a.Hp - damage) < 0 {
		a.Hp = 0
	}

	a.Hp -= damage
}

// void value receiver valid for &AI{...}
func (a AI) void() {

}
func createFighter(fighterType FighterType, maxHp, restoreStep int) (Fighter, error) {
	switch fighterType {
	case human:
		return &Human{
			Hp:          maxHp,
			MaxHp:       maxHp,
			RestoreStep: restoreStep,
		}, nil
	case ai:
		return &AI{
			Hp:    maxHp,
			MaxHp: maxHp,
		}, nil
	}

	return nil, errors.New("wrong type")
}

func main() {
	player, err := createFighter(human, 100, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Human(value receiver) created and got have: %d\n", player.getHp())       // 100
	player.takeDamage(42)                                                                // pointless coz value receiver
	fmt.Printf("Human(value receiver) take damage and have HP: %d\n", player.getHp())    // 100
	player.restoreHp()                                                                   // pointless coz value receiver
	fmt.Printf("Human(value receiver) restore health and have HP: %d\n", player.getHp()) // 100

	bot, err := createFighter(ai, 1000, 10)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("AI(pointer receiver) created and got have: %d\n", bot.getHp()) //1000
	bot.takeDamage(444)
	fmt.Printf("AI(pointer receiver) take damage and have HP: %d\n", bot.getHp()) //556
	bot.restoreHp()
	fmt.Printf("AI(pointer receiver) restore health and have HP: %d\n", bot.getHp()) //999
}
