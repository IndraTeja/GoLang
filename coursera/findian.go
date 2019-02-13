package main

import (
	"fmt"
	"strings"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter a string: ")
	//str, err := reader.ReadString('\n')

	fmt.Println("Enter text: ")
	str := ""
	fmt.Scanln(&str)

	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(reflect.TypeOf(str))

	str = strings.ToLower(str)
	strPrefix := strings.HasPrefix(str, "i")
	strContains := strings.ContainsAny(str,"a")
	strSuffix := strings.HasSuffix(str, "n")

	if strPrefix && strContains && strSuffix {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

}
