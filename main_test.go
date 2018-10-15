package main

import "testing"

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Errorf("expecting 3, got %d", result)
	}
}

func TestSwap(t *testing.T) {
	result1, result2 := swap("hello", "world")
	if result1 != "world" && result2 != "hello" {
		t.Errorf("expecting world hello, got %[2]s %[1]s", result1, result2)
	}
}
