package main

import "fmt"

//- GO Program to Check Whether a Number is Palindrome or Not

func isPalindrome(num string) bool {
	for i := 0; i <= len(num)/2 ; i++ {
		if num[i] != num[len(num) - i - 1] {
			return false
		}
	}
	return true
}

func main()  {
	var num string
	fmt.Printf("Enter a number : ")
	fmt.Scanf("%s",&num)
	fmt.Println(isPalindrome(num))
}
