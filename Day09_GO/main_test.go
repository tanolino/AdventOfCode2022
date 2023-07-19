package main

import "testing"

func Test_P1(t *testing.T) {
	cmds := readCmds("test_input")
	field := makeField()
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	poss := len(field.tailResult)
	correctPoss := 13
	if poss != correctPoss {
		t.Error("Wrong amount of positions. Expected",
			correctPoss, "but received", poss)
	}
}
