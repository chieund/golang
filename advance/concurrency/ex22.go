package main

import (
	"fmt"
	"sync"
)

var (
	//mutex1   sync.Mutex
	balance int
)

func init() {
	balance = 100
}

func deposit(val int, wg *sync.WaitGroup) {
	//mutex1.Lock()             // lock
	balance += val
	//mutex1.Unlock()           // unlock
	wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup) {
	//mutex1.Lock()             // lock
	balance -= val
	//mutex1.Unlock()           // unlock
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go deposit(50, &wg)
	go withdraw(80, &wg)
	//go deposit(40, &wg)
	wg.Wait()
	fmt.Printf("Balance is: %d\n", balance)
}
