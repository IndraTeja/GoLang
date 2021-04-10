package main

import (
	"database/sql"
	"net/http"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "index.html")
}

func main(){

	db, err = sql.Open("mysql", "root:password@/chat")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)

}
