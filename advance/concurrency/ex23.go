package main

import (
	"fmt"
	"sync"
)

var (
	balance1 int
)

func init() {
	balance1 = 100
}

func deposit1(val int, wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	//mutex1.Lock()             // lock
	balance1 += val
	//mutex1.Unlock()           // unlock
	<-ch
	wg.Done()
}

func withdraw1(val int, wg *sync.WaitGroup, ch chan bool) {
	//mutex1.Lock()             // lock
	ch <- true
	balance1 -= val
	//mutex1.Unlock()           // unlock
	ch <- true
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool, 1)
	wg.Add(2)
	go deposit1(50, &wg, ch)
	go withdraw1(80, &wg, ch)
	//go deposit(40, &wg)
	wg.Wait()
	fmt.Printf("Balance is: %d\n", balance1)
}
