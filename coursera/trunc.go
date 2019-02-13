package main

import "fmt"

func main()  {
	var F float64

	fmt.Println("Enter a float value : ")
	_, err := fmt.Scanf("%f", &F)

	if err != nil {
		fmt.Println(err)
	}
	var n int
	n = int(F)
	fmt.Println("After truncation : ", n)
}
