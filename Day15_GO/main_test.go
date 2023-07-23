package main

import (
	"testing"
)

func Test1(t *testing.T) {
	s := readScan("test")
	res := checkScanline2(10, s)
	if res != 26 {
		t.Error("Instead of 26 blocked positions we got", res, "(2)")
	}
}

func Test2(t *testing.T) {
	s := readScan("test")
	res := findFreeSpace(s, 20)
	if res != 56000011 {
		t.Error("Wrong tuning feq", res, "instead of", 56000011)
	}
}
