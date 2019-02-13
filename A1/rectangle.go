package assignment

type Rectangle struct {
	length float64
	width  float64
}

func (rect *Rectangle) AreaRect() float64 {
	return rect.length * rect.width
}

func (rect *Rectangle) Perimeter() float64 {
	return 2 * (rect.length + rect.width)
}

func (rect *Rectangle) SetLength(l float64)  {
	rect.length = l
}

func (rect *Rectangle) SetWidth(w float64)  {
	rect.width = w

}