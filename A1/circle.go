package assignment

type Circle struct {
	radius float64
}

func (c *Circle) CircleArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) CirclePerimeter() float64 {
	return 2 * 3.14 * c.radius
}

func (c *Circle) SetRadius(r float64)  {
		c.radius = r
}
