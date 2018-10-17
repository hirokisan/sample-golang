package main

import "fmt"

var i int = 10

func main() {
	getZeroValue()
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
