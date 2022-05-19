package main

import "fmt"

func main() {
	var countryCaptitalMap map[string]string

	countryCaptitalMap = make(map[string]string)

	countryCaptitalMap["France"] = "Paris"
	countryCaptitalMap["Italy"] = "Rome"
	countryCaptitalMap["Japan"] = "Tokyo"

	for country := range countryCaptitalMap {
		fmt.Println("Capital of", country, "is", countryCaptitalMap[country])
	}

	fmt.Println("=============================")

	delete(countryCaptitalMap, "Franc11e")
	for country := range countryCaptitalMap {
		fmt.Println("Capital of", country, "is", countryCaptitalMap[country])
	}

	capital, ok := countryCaptitalMap["United States"]
	if ok {
		fmt.Println("Capital of United States is", capital)
	} else {
		fmt.Println("Captial of United States is not present")
	}
}
