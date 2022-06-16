package main

import "fmt"

type Employee struct {
	id   int
	name string
	age  int
}

// get ID
func (e *Employee) GetId() int {
	return e.id
}

// Get Name
func (e *Employee) GetName() string {
	return e.name
}

func (e *Employee) GetAge() int {
	return e.age
}

// Set ID
func (e *Employee) SetId(id int) {
	e.id = id
}

// Set Name
func (e *Employee) SetName(name string) {
	e.name = name
}

// set Age
func (e *Employee) SetAge(age int) {
	e.age = age
}

func main() {
	employee := Employee{}
	fmt.Println(employee)

	employee.SetId(1)
	employee.SetName("Test")
	employee.SetAge(10)

	fmt.Println(employee)

	fmt.Println(employee.GetId(), employee.GetName(), employee.GetAge())
}
