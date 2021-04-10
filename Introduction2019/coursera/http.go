package main

import (
	"fmt"
	"net/http"
)

func main() {
	var resp example
	resp, err := http.Get("http://example.com/")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

}
