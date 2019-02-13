package main

import "fmt"

//func zero(x int) int {
//	return x = 0
//}

//func main() {
//	x := 5
//	//zero(x)
//	fmt.Println(&x) // x is still 5
//}

// pass by reference



//func zero(xPtr *int) {
//	*xPtr = 0
//}
//func main() {
//	x := 5
//	zero(&x)
//	fmt.Println(x) // x is 0
//}


func main() {
		xPtr := new(int)
		fmt.Println(xPtr)
		fmt.Println(*xPtr)
}