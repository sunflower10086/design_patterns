package main

func main() {
	pizzaStoryEnt := PizzaStory{
		PizzaCreator: &beijingPizzaStory{},
	}

	pizzaStoryEnt.OrderPizza(cheesePizzaType)
}
