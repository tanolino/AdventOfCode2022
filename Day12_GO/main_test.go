package main

import "testing"

func Test_1(t *testing.T) {
	g := makeGrid("test1")
	steps := g.marchToEnd()
	if steps != 31 {
		t.Error("Wrong steps:", steps, "instead of 31")
	}
}

func Test_Part1(t *testing.T) {
	g := makeGrid("input")
	steps := g.marchToEnd()
	if steps != 490 {
		t.Error("Wrong steps:", steps, "instead of 490")
	}
}

func Test_2(t *testing.T) {
	g := makeGrid("test1")
	g.resetStartAtAllLowPoints()
	steps := g.marchToEnd()
	if steps != 29 {
		t.Error("Wrong steps:", steps, "instead of 29")
	}
}
