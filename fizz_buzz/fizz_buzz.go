package main

import (
	"fmt"
	"strconv"
)

func fizz_buzz(fizz int, buzz int, max int) string{
	output := ""
	str := ""
	for i := 1; i <= max; i++ {
		str = ""
		if i % fizz == 0 {
			str += "fizz"
		}
		if i % buzz == 0 {
			str += "buzz"
		}
		if str == "" {
			str += strconv.Itoa(i)
		}
		str += "\n"
		output += str 
	}
	return output
}

func main(){
	fmt.Println(fizz_buzz(3, 5, 20))
}