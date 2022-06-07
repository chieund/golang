package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee1 struct {
	Name string `json:"n"`
	Age int `json:"a"`
	salary int `json:"s"`
}

type employee2 struct {
	Name string
	Age int
	salary int
}

type PubSubMessage struct {
	Data []byte `json:"data"`
}

func main()  {
	data := []byte(string("hello world"))
	message := PubSubMessage{data}
	fmt.Printf("employee1 Struc: %#v\n", string(message.Data))

	e1Json := `{"n" : "john", "a": 21}`

	var e1Converted employee1

	fmt.Printf("employee1 Struc: %#v\n", []byte(e1Json))


	err := json.Unmarshal([]byte(e1Json), &e1Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling, Error: %s", err.Error())
	}

	fmt.Printf("employee1 Struc: %#v\n", e1Converted)

	e2Json := `{"Name": "John", "Age": 21}`
	var e2Converted employee2
	err = json.Unmarshal([]byte(e2Json), &e2Converted)
	if err != nil {
		log.Fatalf("Error occured during unmarshaling, Error: %s", err.Error())
	}

	fmt.Printf("employee2 struc: %#v\n", e2Converted)
}