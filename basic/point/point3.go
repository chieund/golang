package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

func (p *Person) IncreaseAge() {
	p.Age += 1
}

func (p Person) NotReallyIncreaseAge() {
	p.Age += 1
}

func main() {
	bod := &Person{"Bob", 46}
	bod.IncreaseAge()
	bod.NotReallyIncreaseAge()
	fmt.Println(bod)

	tom := new(Person)
	tom.Name = "Tom"
	tom.Age = 27
	tom.IncreaseAge()
	tom.NotReallyIncreaseAge()
	fmt.Println(tom)

	alice := Person{Name: "Alice", Age: 18}
	alice.IncreaseAge()
	fmt.Println(&alice)

	tuoi := 29
	age := &tuoi
	println(*age)
	println(age)
	*age += 1
	println(*age)
	println(tuoi)
}
