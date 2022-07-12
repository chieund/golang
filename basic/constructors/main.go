package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Age  int
}

func NewEmployee(ID int, name string, age int) *Employee {
	return &Employee{
		ID:   ID,
		Name: name,
		Age:  age,
	}
}

func (e *Employee) GetId() int {
	return e.ID
}

func (e *Employee) GetName() string {
	return e.Name
}

func (e *Employee) GetAge() int {
	return e.Age
}

func (e *Employee) SetId(id int) {
	e.ID = id
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e *Employee) SetAge(age int) {
	e.Age = age
}

func main() {
	employee1 := Employee{}
	fmt.Println(employee1)

	employee2 := Employee{ID: 1, Name: "chieund", Age: 20}
	fmt.Println(employee2)

	employee3 := new(Employee)
	fmt.Println(*employee3)

	employee4 := NewEmployee(1, "Chieu2", 30)
	fmt.Println(employee4.GetId())
}
