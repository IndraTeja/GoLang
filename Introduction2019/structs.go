package main

import "fmt"

type Circle struct {
		radius float64
		perimeter float64
}

var c Circle
// c := new(Circle)
// c := Circle{radius = 5.0, perimeter = 0.0}
// c := Circle{5.0, 0.0}

func main(){
		c.radius = 10.0
		c.perimeter = 2 * 3.14 * c.radius
		fmt.Println(c.perimeter)
}
