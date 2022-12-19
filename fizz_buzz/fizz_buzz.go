package main

import (
	"fmt"
	"strconv"
	"strings"
)

func fizz_buzz(fizz int, buzz int, max int) string{
	var output string
	for i := 1; i <= max; i++ {
		if i % fizz == 0 {
			output += "fizz"
		}
		if i % buzz == 0 {
			output += "buzz"
		}
		res := strings.HasSuffix(output, "z")
		if !res {
			output += strconv.Itoa(i)
		}
		output += "\n"
	}
	return output
}

func main(){
	fmt.Println(fizz_buzz(3, 5, 20))
}
