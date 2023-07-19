package main

import "fmt"

func main() {
	cmds := readCmds("input")
	field := makeField()
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	fmt.Println("Unique tail positions: ", len(field.tailResult))
}
