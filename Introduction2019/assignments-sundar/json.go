package main

import (
	"fmt"
	"gopkg.in/gin-gonic/gin.v1/json"
	"log"
)

type person struct {
	Name string `json:"name"`
	age string `json:"age"`
}

func main()  {

	person1 := person{
			Name:"Indra",
			age:"27",
	}

	fmt.Println(person1)

	var jsonData []byte
	jsonData, err := json.Marshal(person1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))
}
