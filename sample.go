package main

import (
	"fmt"
	"math"
)

var i int = 10

func main() {
	getTypeInterface()
	//fmt.Print(getTypeConvert(3, 4))
	//getZeroValue()
	//getI()
	//fmt.Println(i)
	//fmt.Println(add(1, 3))
	//a, b := swap("hello", "world")
	//fmt.Println(a, b)
}

func add(x, y int) (result int) {
	result = x + y
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func getI() {
	i := 10 // can use := only inside function
	fmt.Println(i)
}

func getZeroValue() {
	var (
		i int
		j float64
		k bool
		l string
	)

	fmt.Printf("%v, %v, %v, %q\n", i, j, k, l)
}

func getTypeConvert(x, y int) uint {
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	return z
}

func getTypeInterface() {
	v := 12.00
	fmt.Printf("v is of type %T\n", v)
}
