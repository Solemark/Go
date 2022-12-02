package main

import "fmt"

func reverse_array(arr []int) []int{
	x := 0
	y := 0
	c := len(arr)-1
	for i := 0; i < c; i++ {
		x = arr[i]
		y = arr[c]
		arr[i] = y
		arr[c] = x
		c--
	}
	return arr
}

func main(){
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(reverse_array(arr))
}