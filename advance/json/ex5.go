package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("foo", "bar")
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error when sending request to the server")
		return
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Status)
	fmt.Println(string(responseBody))
}
