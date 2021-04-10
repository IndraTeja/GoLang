package main

import "fmt"

func main() {
		x := "Hello world"
		fmt.Println(x)
		fmt.Println(&x)
		new(*x)
		fmt.Println(x)
}
