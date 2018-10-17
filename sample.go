package main

import (
	"fmt"
	"math"
)

const Pi = 3.14

var i int = 10

func main() {
	fmt.Println(isFive(5))
	//getInfiniteLoop()
	//getLoop()
	//getConst()
	//getTypeInterface()
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

func getConst() {
	fmt.Println(Pi)
}

func getLoop() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
}

func getInfiniteLoop() {
	for {
	}
}

func isFive(x int) (result bool) {
	if x == 5 {
		result = true
	}
	return result
}
