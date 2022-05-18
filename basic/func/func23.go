package main

import (
	"fmt"
)


func getSequence() func() int {
   i:= 0


   a:= func() int {
      i+=1
      return i;
   }

   return a
}

func main() {
   nextNumber := getSequence()
   fmt.Println(nextNumber)

   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
   fmt.Println(nextNumber())
}