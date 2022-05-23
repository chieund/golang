package main

import "fmt"

type SomeStruct struct {
	Field string
}

func (s *SomeStruct) foo(field string) {
	s.Field = field
}

func main() {
	someStruct := new(SomeStruct)
	someStruct.foo("a")

	fmt.Println(someStruct.Field)
	someStruct.foo("b")
	fmt.Println(someStruct.Field)

	b := SomeStruct{}
	b.foo("a")
	fmt.Println(b.Field)
}
