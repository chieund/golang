package main

import (
	"fmt"
	"log"
)

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		log.Fatalln("runtime error: first name cannot be nil")
	}

	if lastName == nil {
		log.Fatalln("runtime error: last name cannnot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
}
