package main

import "fmt"

type Class interface {
	findGender() string
}

type student struct {
	name, gender string
	age int
}

func (s student) findGender()  string{
	return s.gender
}

func printGender(c Class){
	fmt.Println(c.findGender())
}

func main(){

	s1 := student{name:"indra", gender:"male",age:27}
	s2 := student{name:"jenifer", gender:"female", age:30}

	printGender(s1)
	printGender(&s2)

}
