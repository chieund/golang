package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	values := map[string]string{}
	fmt.Println(values)
	values["foo"] = "baz"
	fmt.Println(values)

	jsonData, _ := json.Marshal(values)
	fmt.Println(jsonData)
	fmt.Println(bytes.NewBuffer(jsonData))
	req, err := http.NewRequest(http.MethodPost, "https://httpbin.org/post", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	println(req)
}
