package main

import (
	"errors"
	"fmt"
)

func test2(name string) (string, error) {
	if name == "" {
		return name, errors.New("name empty 2222")
	}

	return name, nil
}

func test(name string) (string, error) {
	b, err := test2(name)
	if err != nil {
		return name, log.Errorf("not connection db: %v", err)
	}

	return b, nil
}

func main() {
	a, err := test("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
}
