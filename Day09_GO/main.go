package main

import "fmt"

func part1() {
	cmds := readCmds("input")
	field := makeField()
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	fmt.Println("Unique tail positions (Part 1): ", len(field.tailResult))
}

func part2() {
	cmds := readCmds("input")
	field := makeField()
	field.addKnots(8)
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	fmt.Println("Unique tail positions (Part 2): ", len(field.tailResult))
}

func main() {
	part1()
	part2()
}
