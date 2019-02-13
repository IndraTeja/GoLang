package main

import "fmt"

func binarySort(arr []int, n int)  []int {
	for {
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
		break
	}
	return arr
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
	fmt.Print("sorted array : ")
	fmt.Println(binarySort(arr, n))
	}