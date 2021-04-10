package main

import "fmt"

func main() {

	var bigInt, smallInt = 1, 1

	fmt.Printf("Enter big integer : ")
	fmt.Scanf("%d\n", &bigInt)

	fmt.Printf("Enter small integer : ")
	fmt.Scanf("%d\n", &smallInt)

	div := bigInt % smallInt
	fmt.Println(div)

}
