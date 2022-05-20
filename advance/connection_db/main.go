package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Test struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	fmt.Println("Go Mysql Tutorial")

	db, err := sql.Open("mysql", "laravel:laravel@tcp(127.0.0.1:33085)/laravel")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO test value('title')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	results, err := db.Query("SELECT id, title FROM test")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var test Test
		err = results.Scan(&test.ID, &test.Title)
		if err != nil {
			panic(err.Error())
		}

		log.Printf(test.Title)
	}

	var test1 Test
	err = db.QueryRow("select id, title from test where id = ?", 1).Scan(&test1.ID, &test1.Title)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(test1.ID)
	fmt.Println(test1.Title)

}
