package main

import "fmt"

type MyInt int // khai bao mot kieu moi thuc chat van la kieu int

func (p *MyInt) Double() {
	*p = *p * 2
}

func main() {
	muoi := new(MyInt)
	*muoi = 10
	muoi.Double()
	fmt.Println(*muoi)
}
