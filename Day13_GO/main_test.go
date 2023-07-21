package main

import (
	"testing"
)

func Test1(t *testing.T) {
	d := readData("test")
	sum := 0
	for i, v := range d {
		if v.compare() <= 0 {
			sum += i + 1
		}
	}
	corrSum := 13
	if sum != corrSum {
		t.Error("Wrong Sum:", sum, "instead of", corrSum)
	}
}
