package main

import (
	"fmt"
	"sync"
)

var x int64 = 0
var mutex = &sync.Mutex{}

func addOne2(w *sync.WaitGroup) {
	// lock
	mutex.Lock()
	x = x + 1

	// Unloc
	mutex.Unlock()
	w.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go addOne2(&w)
	}

	w.Wait()
	fmt.Println("Value x:", x)
}
