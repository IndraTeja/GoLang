package main

import "fmt"

//- Write a Simple program using Print, Printf and Println statements to print Floyd's Triangle up to N
// number of rows entered by user.
func main(){

	fmt.Printf("Enter a number: ")
	var n int
	var x int
	x = 1
	fmt.Scanf("%d",&n)

	for i := 1; i <= n; i++ {
		for j :=1; j <= i; j++ {
			fmt.Printf(" %d", x)
			x++
		}
		fmt.Println("")
	}
}



