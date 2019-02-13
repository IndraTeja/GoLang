package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	custid int      `json:"custid"`
	username string `json:"username`
}

func main() {
	fmt.Println("GoLang MySQL Database")

	// Open up our database connection, on my localhost:3306
	// The database is called new_schema and table is mytable
	db, err := sql.Open("mysql", "roo:password@tcp(localhost:3306)/new_schema")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert into mytable
	insert, err := db.Query("INSERT INTO mytable VALUES ( 106, 'saran' )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	// Execute the query
	results, err := db.Query("SELECT * FROM mytable")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.custid, &tag.username)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		//fmt.Printf(tag.custid)
		log.Printf(tag.username)
	}

}
