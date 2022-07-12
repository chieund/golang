package main

import "fmt"

func main() {
	var a int = 20
	var ip *int

	ip = &a /* store address of a in pointer variable */

	fmt.Printf("Address of a variable: %x\n", &a)

	/* address stored in pointer variable */
	fmt.Printf("Address stored in ip variable: %x\n", ip)

	fmt.Printf("Address stored in ip variable: %x\n", &ip)

	/* access the value using the point */
	fmt.Printf("Value of *ip variable: %d", *ip)
}