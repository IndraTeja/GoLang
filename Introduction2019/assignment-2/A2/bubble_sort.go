package main

import "fmt"

func Swap() {
	for i := 1; i < n; i++ {
		for j:=1; j < n; j++ {
		if arr[j-1] > arr[j] {
			arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}

func BinarySort(arr []int, n int)  []int {


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
	fmt.Println(BinarySort(arr, n))
	}