package main

import "fmt"

var (
	CompanyName = "test"
	CompanyLocaltion = "somecity"
)

type Person struct {
	Name string
	Age int
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) getName() string {
	return p.Name
}

type Company struct {}

func GetPerson() *Person {
	p := &Person {
		Name: "Test",
		Age: 21,
	}

	fmt.Println("Model Package")
	fmt.Println(p.Name)
	fmt.Println(p.Age)
	return p
}

func getCompanyName() string {
	return CompanyName
}

func main() {
	p := &Person{
		Name: "test",
		Age: 21,
	}

	fmt.Println(p)
	c := &Company{}
	fmt.Println(c)

	fmt.Println(p.GetAge())
	fmt.Println(p.getName())

	fmt.Println(p.Name)
	fmt.Println(p.Age)

	person2 := GetPerson()
	fmt.Println(person2)

	companyName := getCompanyName()
	fmt.Println(companyName)

	fmt.Println(CompanyLocaltion)
	fmt.Println(companyName)
}