package main

import "testing"

func Test_Part1(t *testing.T) {
	f := makeForest("test_input")
	i := f.countVisibleTrees()
	if i != 21 {
		t.Error("Counted wrong. Expected 21 but got", i)
	}
}

func Test_Part2(t *testing.T) {
	f := makeForest("test_input")
	i := f.maxScenicScore()
	if i != 8 {
		t.Error("Calculated wrong score. Expected 8 but got", i)
	}
}
