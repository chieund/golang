package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

func main() {
	var b = [...]int{5: 3, 1: 2}

	for index, value := range b {
		fmt.Println(index, value)
	}

	fmt.Println("============", cap(b), "========")

	fmt.Println("\n")
	var c = [...]int{1, 2, 4: 5, 6}
	for index, value := range c {
		fmt.Println(index, value)
	}

	// array string
	var s1 = [2]string{"hello", "world"}
	var s2 = [...]string{"Hello!", "world"}
	var s3 = [...]string{1: "Hello", 0: "world"}

	for index, value := range s1 {
		fmt.Println(index, value)
	}

	for index, value := range s2 {
		fmt.Println(index, value)
	}

	for index, value := range s3 {
		fmt.Println(index, value)
	}

	// array struct
	var line1 [2]image.Point
	var line2 = [...]image.Point{image.Point{X: 0, Y: 0}, image.Point{X: 1, Y: 1}}
	var line3 = [...]image.Point{{0, 0}, {1, 1}}

	// array func
	var decoder1 [2]func(io.Reader) (image.Image, error)
	var decoder2 = [...]func(io.Reader) (image.Image, error){
		png.Decode,
		jpeg.Decode,
	}

	// array interface
	var unknown1 [2]interface{}
	var unknonw2 = [...]interface{}{123, "Hello"}

	// array channel
	var chanList = [2]chan int{}
}
