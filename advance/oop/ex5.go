package main

import "fmt"

type iAnimal interface {
	breadthe()
}

type animal struct {
}

func (a *animal) breadthe() {
	fmt.Println("Animal breate")
}

type iAquatic interface {
	iAnimal
	swim()
}

type aquatic struct {
	animal
}

func (a *aquatic) swim() {
	fmt.Println("Aquatic swim")
}

type iNonAquatic interface {
	iAnimal
	walk()
}

type nonAquatic struct {
	animal
}

func (a *nonAquatic) walk() {
	fmt.Println("Non-Aquatic walk")
}

type shark struct {
	aquatic
}

type lion struct {
	nonAquatic
}

func main() {
	share := &shark{}
	checkAquatic(share)
}

func checkAquatic(a iAquatic) {
	a.swim()
}
