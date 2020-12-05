package main

import "fmt"

func div(x, y float64) float64 {
	return x / y
}
//func main is mandatory duh else code won't run 
func main() {
	fmt.Println(div(42, 13))
}
