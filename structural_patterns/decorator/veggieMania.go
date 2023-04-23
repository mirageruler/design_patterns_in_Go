package main

// Concrete component
type VeggeMania struct {
	pizza IPizza
}

func (p *VeggeMania) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 5
}
