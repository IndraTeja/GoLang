package main

import (
	"log"
	"net/http"
)


func main()  {

	handlerAdd := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Incoming Request:", req.Method)

	}

	http.HandleFunc("/add", handlerAdd)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
