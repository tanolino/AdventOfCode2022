package main

import "testing"

func Test_Part1(t *testing.T) {
	rootDir := readInput("test_input")
	rootDir.printOut("")
	sum := 0
	sumRight := 95437
	for _, d := range findDirsRecursive(rootDir) {
		sum += d.getSize()
	}
	if sum != sumRight {
		t.Error("Sum missmatch ", sum, " != ", sumRight)
	}
}
