package main

import "fmt"

func convert_distance(x float64, y float64) float64{
	return x * y
}

func main(){
	km := 10.0
	m := 10.0
	ktm := 0.6213712
	mtk := 1.609344

	fmt.Printf("%gkm is %g miles\n", km, convert_distance(km, ktm))
	fmt.Printf("%g miles is %gkm\n", m, convert_distance(m, mtk))
}