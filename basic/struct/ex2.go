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
	var Book2 Books

	Book1.title = "Go Programming"
	Book1.author = "Mahesh Kumar"
	Book1.subject = "Go Progrmaming Tutorial"
	Book1.book_id = 121212

	Book2.title = "Go Programming"
	Book2.author = "Mahesh Kumar"
	Book2.subject = "Go Progrmaming Tutorial"
	Book2.book_id = 121212

	printInfo(Book1)
	printInfo(Book2)
}

func printInfo(book Books) {
	fmt.Printf("title: %s\n", book.title)
	fmt.Printf("author %s\n", book.author)
	fmt.Printf("subject: %s\n", book.subject)
	fmt.Printf("book id: %d\n", book.book_id)
}
