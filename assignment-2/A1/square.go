package assignment

type Square struct {
	side float64
}

func (sq *Square) AreaSq() float64 {
	return sq.side * sq.side
}

func (sq *Square) PerimeterSq() float64 {
	return 4 * sq.side
}

func (sq *Square) SetSide(s float64)  {
	sq.side = s
}