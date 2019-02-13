package main

import (
	"GoLang/assignment"
	"fmt"
)

func main() {

	circle := new(assignment.Circle)
	circle.SetRadius(5.0)
	rectangle := new(assignment.Rectangle)
	rectangle.SetLength(10.0)
	rectangle.SetWidth(7.0)
	square := new(assignment.Square)
	square.SetSide(6.0)

	fmt.Println("Circle area = ", circle.CircleArea())
	fmt.Println("Circle perimeter = ", circle.CirclePerimeter())
	fmt.Println("Square area = ", square.AreaSq())
	fmt.Println("Square perimeter = ", square.PerimeterSq())
	fmt.Println("Rectangle area = ", rectangle.AreaRect())
	fmt.Println("Rectangle perimeter = ", rectangle.Perimeter())
}
