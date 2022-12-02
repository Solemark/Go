package main

import "fmt"

func check_palindrome(str string) bool {
	c := len(str) - 1
	for i := 0; i < c; i++ {
		if(str[i] != str[c]){
			return false
		}
		c--
	}
	return true
}

func main(){
	fmt.Println("is DAD a palindrome?", check_palindrome("DAD"))
	fmt.Println("is Dad a palindrome?", check_palindrome("Dad"))
	fmt.Println("is 123456789987654321 a palindrome?", check_palindrome("123456789987654321"))
}