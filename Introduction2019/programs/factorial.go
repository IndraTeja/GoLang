package main

import "fmt"

func factorial(n int) int {
	var fact = 1
	for i:=1; i<=n; i++ {
		fact = fact * i
	}
	return fact
}

func main()  {
	var num int
	fmt.Println("Enter a number : ")
	fmt.Scanf("%d",&num)
	fmt.Println(factorial(num))
}


