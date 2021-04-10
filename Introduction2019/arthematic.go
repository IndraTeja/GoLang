package main

import (
		"fmt"
		"reflect"
)

func main() {
		var a int32 = 32132 * 42452
		b := "Hello world"
		fmt.Println("32132 × 42452 = ", 32132 * 42452)
		fmt.Println(reflect.TypeOf(a))
		fmt.Println(len(b))
}
