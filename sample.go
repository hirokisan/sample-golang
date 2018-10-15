package main

import "fmt"

func main() {
	//fmt.Println(add(1, 3))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}
