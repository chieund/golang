package main

import "fmt"

type Duck interface {
	Quacks()
}

type Animal struct {
}

func (a Animal) Quacks() {
	fmt.Println("Animal Quacks")
}

func SomeDuck(duck Duck) {
	duck.Quacks()
}

func test() {
	return test()
}

func main() {
	animal := Animal{}
	SomeDuck(animal)
}
