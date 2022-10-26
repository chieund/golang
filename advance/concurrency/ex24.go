package main

import (
	"fmt"
	"sync"
)

const (
	numberOfURLs    = 10000
	numberOfWorkers = 5
)

func crawlURL(queue <-chan int, wg *sync.WaitGroup, name string) {
	for v := range queue {
		fmt.Printf("Worker %s is crawling URL %d\n", name, v)
		//time.Sleep(time.Second)
	}

	fmt.Printf("Worker %s done and exit\n", name)
	wg.Done()
}

func startQueue() <-chan int {
	queue := make(chan int, 100)

	go func() {
		for i := 1; i <= numberOfURLs; i++ {
			queue <- i
			fmt.Printf("URL %d has been enqueued\n", i)
		}

		close(queue)
	}()

	return queue
}

func main() {
	var wg sync.WaitGroup
	queue := startQueue()
	wg.Add(numberOfWorkers)

	for i := 1; i <= numberOfWorkers; i++ {
		go crawlURL(queue, &wg, fmt.Sprintf("%d", i))
	}
	wg.Wait()
}
