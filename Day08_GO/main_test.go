package main

import "testing"

func Test_Part1(t *testing.T) {
	f := makeForest("test_input")
	i := countVisibleTrees(f)
	if i != 21 {
		t.Error("Counted wrong. Expected 21 but got", i)
	}
}
