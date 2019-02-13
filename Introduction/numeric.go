package main

import (
	"fmt"
	"unsafe"
	"reflect"
)

func main()  {

	a := 108888888.5
	b := 109 + 8i
	fmt.Println("size of a is ", unsafe.Sizeof(a))
	fmt.Println("Type of a is ", reflect.TypeOf(a))
	fmt.Println("size of  is ", unsafe.Sizeof(b))
	fmt.Println("Type of b is ", reflect.TypeOf(b))

}
