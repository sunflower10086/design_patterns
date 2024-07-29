package main

type pizzaIngredientFactory interface {
	CreateDough() Dough
}

type beijingPizzaIngredientFactory struct {
}

func (n *beijingPizzaIngredientFactory) CreateDough() Dough {
	return ThinCrustDough{}
}
