package main

import (
	"fmt"

	_ "github.com/chieund/go_module/v2"
)

func main() {
	msg := HelloTest("hello")
	fmt.Println(msg)
}
