package main

import "fmt"

func even(a int) (int, bool) {

	if a%2 == 0 {
		return a / 2, true
	} else {
		return a / 2, false
	}

}

func main() {

	var a int

	fmt.Scan("%d", &a)

	fmt.Println(even(a))

}
