package main

import "errors"

type FighterType string

const (
	human FighterType = "human"
	ai    FighterType = "ai"
)

type Fighter interface {
	getHp() int
	RestoreHp()
}

type Human struct {
	Hp          int
	MaxHp       int
	RestoreStep int
}

func (h Human) getHp() int {
	return h.Hp
}

func (h Human) RestoreHp() {
	if h.MaxHp < (h.Hp + h.RestoreStep) {
		h.Hp += h.RestoreStep
	}

	h.Hp = h.MaxHp
}

type AI struct {
	Hp    int
	MaxHp int
}

func (a AI) getHp() int {
	return a.Hp
}

func (a AI) RestoreHp() {
	a.Hp = a.MaxHp
}

func createFighter(fighterType FighterType, maxHp, restoreStep int) (Fighter, error) {
	switch fighterType {
	case human:
		return Human{
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
