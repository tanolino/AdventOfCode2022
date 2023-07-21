package main

import (
	"testing"
)

func Test1(t *testing.T) {
	d := readDataAsPairs("test")
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

func Test2(t *testing.T) {
	key := readDecoderKey("test")
	if key != 140 {
		t.Error("Key should be 140 but we got", key)
	}
}
