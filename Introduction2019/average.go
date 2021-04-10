package main

import "fmt"

func main() {
	arr := []float64{65, 89, 54, 81, 93}
	avg := average(arr)
	fmt.Println("Average : ", avg)
}

func average(x []float64) float64{
	total := 0.0
	for _,v := range x {
		total += v
	}
	return total / float64(len(x))
}