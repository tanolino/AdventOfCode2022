package main

import "fmt"

func part1() {
	g := makeGrid("input")
	steps := g.marchToEnd()
	fmt.Println("Steps (1): ", steps)
}

func part2() {
	g := makeGrid("input")
	g.resetStartAtAllLowPoints()
	steps := g.marchToEnd()
	fmt.Println("Steps (2): ", steps)
}

func main() {
	part1()
	part2()
}
