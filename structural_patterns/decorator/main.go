package main

import "fmt"

func main() {

	pizza := newPizza()
	// Add vegga
	pizzaWithVegga := &VeggeMania{
		pizza: pizza,
	}

	//Add cheese topping
	pizzaWithCheese := &CheeseTopping{
		pizza: pizzaWithVegga,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %d\n", pizzaWithCheeseAndTomato.getPrice())
}
