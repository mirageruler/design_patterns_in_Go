package main

type IPizza interface {
	getPrice() int
}

type Pizza struct {
	price int
}

func (p Pizza) getPrice() int {
	return p.price
}

func newPizza() IPizza {
	return Pizza{
		price: 10,
	}
}
