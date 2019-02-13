package main

import "fmt"

var x = "Hello"

const y = "World"

func main() {

	f()
	fmt.Println(y)
	x = "Changed the value!"
	fmt.Println(x)

}


func f() {

	fmt.Println(x)

}