package main

import "fmt"

type printer interface {
	print()
}

type students struct {
	name, gender string
}

func (s students) print() {
	fmt.Println("User Name:", s.name, ", Gender:", s.gender)
}

func main() {
	s := students{"Indra","Male"}

	entities := []printer{
		s,
		&s,
	}

	s.name = "Jenifer"
	s.gender = "female"

	for _, e := range entities {
		e.print()
	}
}
