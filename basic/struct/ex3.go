package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Book1 Books

	Book1.title = "Java Core"
	Book1.author = "Ha"
	Book1.subject = "Java Core Tutorial"
	Book1.book_id = 1212

	printInfo(&Book1)

	println(Book1)
}

func printInfo(book *Books) {
	fmt.Printf("title: %s", book.title)
}
