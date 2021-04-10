package main

import "fmt"

func main()  {

	//declared slice with length 5
	slice := make([]int, 0, 5)
	fmt.Println(slice)
	//printing capacity and length
	fmt.Println(cap(slice))
	fmt.Println(len(slice))

	for i:=0; i<10; i++ {
		slice = append(slice, i)
		fmt.Printf("capacity %d, length %d, slice %v, address %p \n",cap(slice),len(slice), slice, &slice)
	}

}
