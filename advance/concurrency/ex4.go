package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

func Consumer(in <-chan int) {
	for v := range in {
		fmt.Println("lay ra", v)
	}
}

func main() {
	// hang doi
	ch := make(chan int, 64)

	// tao mot chuoi voi so voi boi so 3
	go Producer(3, ch)

	// tao mot chuo si voi boi so 5
	go Producer(5, ch)

	// tao consumer
	go Consumer(ch)

	// thoat ra sau khi chay trong mot khoang thoi gian nhat dinh
	//time.Sleep(5*time.Second)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}
