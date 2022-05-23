package main

import "fmt"

type base struct {
	color string
}

func (b *base) say() {
	fmt.Println("say function")
}

type child struct {
	base
	style string
}

func check(b base) {
	b.say()
}

func main() {
	base := base{color: "red"}
	child := &child{
		base:  base,
		style: "somestyle",
	}

	child.say()
	fmt.Println("Child color is " + child.color)
	check(base)
}
