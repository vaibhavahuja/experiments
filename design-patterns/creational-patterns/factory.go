package main

import (
	"errors"
	"fmt"
)

//Factory method
// factory which produces a struct and all the structs implement
// the common interface.

type IGun interface {
	getName() string
}

type Gun struct {
	name string
}

func (g *Gun) getName() string {
	return g.name
}

type Ak47 struct {
	Gun
}

func newAk47() IGun {
	return &Ak47{Gun{name: "AK47 Gun"}}
}

type Musket struct {
	Gun
}

func newMusket() IGun {
	return &Musket{Gun{name: "Musket Gun"}}
}

func GunFactory(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	} else if gunType == "musket" {
		return newMusket(), nil
	}
	return nil, errors.New("unable to find gun")

}

func main() {
	ak47, _ := GunFactory("ak47")
	musket, _ := GunFactory("musket")

	fmt.Println(ak47.getName())
	fmt.Println(musket.getName())
}
