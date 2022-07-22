package main

import (
	"fmt"
	"log"
)

func defFoo() {
	fmt.Println("defFoo() started")

	if r := recover(); r != nil {
		fmt.Println("WOHA! Program is panicking with value", r)
	}

	fmt.Println("defFoo() done")
}

func normMain() {
	fmt.Println("normMain() started")

	defer defFoo()
	log.Panicln("HELP")

	fmt.Println("normMain() done")
}

func main() {
	fmt.Println("main() started")
	normMain()

	fmt.Println("main() done")
}
