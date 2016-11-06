package calc

import "fmt"

func init() {
	fmt.Println("calc.init executed")
}

func Add(x, y int) int  {
	return x+y
}

func Substract(x, y int) int {
	return x-y
}