package main

import (
	"fmt"
	str "strings"
	"test_v2/numbers"
	"test_v2/strings"
	"test_v2/strings/greeting"
)

func main() {
	fmt.Println(numbers.IsPrime(19))

	fmt.Println(greeting.WelcomeText)

	fmt.Println(strings.Reverse("test"))

	fmt.Println(str.Count("te", "t"))
}
