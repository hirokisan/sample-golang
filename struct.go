package main

import (
	"fmt"
)

type Sample struct {
	X, Y int
}

//func (s *Sample) modifyValue() {
func (s Sample) modifyValue() {
	s.X = 10
	s.Y = 20
}

func (s *Sample) renderValue() {
	fmt.Println(s.X, s.Y)
}

func main() {
	s := Sample{X: 1, Y: 2}
	s.modifyValue()
	s.renderValue()
}
