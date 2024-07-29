package main

import "fmt"

type pizzaType string

const (
	cheesePizzaType pizzaType = "cheese"
	clamPizzaType   pizzaType = "clam"
)

type pizza interface {
	prepare()
}

type Dough interface {
}

// ThinCrustDough 北京披萨的薄面团
type ThinCrustDough struct {
}

type basePizza struct {
	dough Dough
}

func (b *basePizza) prepare() {
	fmt.Println("base prepare")
}

type cheesePizza struct {
	*basePizza
	IngredientFactory pizzaIngredientFactory
}

func (c *cheesePizza) prepare() {
	fmt.Println("preparing cheesePizza")
	c.dough = c.IngredientFactory.CreateDough()
}

type clamPizza struct {
	*basePizza
}

//
//func (c *clamPizza) prepare() {
//	fmt.Println("clam prepare")
//}
