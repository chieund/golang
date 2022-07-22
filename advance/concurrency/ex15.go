package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64 = 0

func addOne(w *sync.WaitGroup) {
	x = x + 1
	w.Done()
}

func addOne2(w *sync.WaitGroup) {
	atomic.AddInt64(&x, 1)
	w.Done()
}

func main() {
	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		go addOne2(&w)
	}

	w.Wait()
	fmt.Println("Value x:", x)
}
