package main

import "fmt"

func isBinary(arr []int, n int, find int) int {
	//mid := n/2
	var position int
	//first, last := 0,n
	for i:=0; i<n; i++ {
		if arr[i] == find {
			position = i + 1
			break
		}
	}
	return position
}


func main()  {
	// n reads and stores the length of array
	var n int
	fmt.Println("Enter length of array : ")
	fmt.Scanf("%d",&n)
	// arr of length n is created
	arr := make([]int, n)
	// ask user to enter a number to search in array
	for i:=0; i<n; i++{
		fmt.Println("Enter array numbers : ")
		fmt.Scanf("%d",&arr[i])
	}

	var find int
	fmt.Println("Enter the number you want to find : ")
	fmt.Scanf("%d",&find)
	fmt.Print("Found at position : ")
	fmt.Println(isBinary(arr, n, find))
}