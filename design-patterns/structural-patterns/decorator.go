package main

import "fmt"

//this one is pretty straightforward I guess
// say i have a class A which does something
// I want to extend the behaviour of that class (decorate) it
// but don't want to modify the actual class, I use this pattern
// I don't want to break the code which actually uses these classes

type IPizza interface {
	getPrice() int
}

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
	return 15
}

//concrete decorator

type TomatoTopping struct {
	pizza IPizza
}

func (t *TomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 7
}

//concrete decorator

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {
	pizza := &VeggieMania{}
	pizzaWithCheese := &CheeseTopping{pizza}

	pizzaWithCheeseAndTomato := &TomatoTopping{pizzaWithCheese}

	fmt.Println("final price ", pizzaWithCheeseAndTomato.getPrice())
}
