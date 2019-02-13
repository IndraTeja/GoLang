package main

import "fmt"

func main() {

		var x bool
		x = (true && false) || (false && true) || !(false && false)
		fmt.Println(x)

		//a := true
		//b := false

		fmt.Println("true && false :", true && false)
		fmt.Println("false && true :", false && true)
		fmt.Println("true || false :", true || false)
}
