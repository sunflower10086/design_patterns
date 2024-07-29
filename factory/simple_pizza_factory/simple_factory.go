package main

type PizzaStory struct {
	PizzaCreator pizzaCreator
}

func (o *PizzaStory) OrderPizza(pizzaType pizzaType) {
	pizzaEnt := o.PizzaCreator.createPizza(pizzaType)

	pizzaEnt.prepare()
}

type pizzaCreator interface {
	createPizza(pizzaType pizzaType) pizza
}

type beijingPizzaStory struct {
}

func (b *beijingPizzaStory) createPizza(pizzaType pizzaType) pizza {
	var pizzaEnt pizza
	if pizzaType == "cheese" {
		pizzaEnt = &cheesePizza{
			IngredientFactory: &beijingPizzaIngredientFactory{},
		}
	} else if pizzaType == "clam" {
		pizzaEnt = &clamPizza{}
	}
	return pizzaEnt
}
