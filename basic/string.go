package main
import (
	"fmt"
	"strings"
)

func main() {
	var greeting = "Hello world!"

	fmt.Printf("normal string:")
	fmt.Printf("%s", greeting)
	fmt.Println("\n")
	fmt.Println("hex bytes: ")
	for i:=0; i < len(greeting); i++ {
		fmt.Printf("%x ", greeting[i])
	}

	fmt.Printf("\n")
	const sampleText = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Printf("quoted string: ")
	fmt.Printf("%+q", sampleText)
	fmt.Printf("\n")

	greeting1 := []string{"Hello", "world!"}

	fmt.Println(strings.Join(greeting1, " "))
}