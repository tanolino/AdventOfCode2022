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

func Test_P1_Solution(t *testing.T) {
	cmds := readCmds("input")
	field := makeField()
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	poss := len(field.tailResult)
	correct := 6357
	if poss != correct {
		t.Error("Solution 1 failed with", poss, "instead of", correct)
	}
}

func Test_P2(t *testing.T) {
	cmds := readCmds("test_input2")
	field := makeField()
	field.addKnots(8)
	for _, cmd := range cmds {
		field.moveHeadByCmd(cmd)
	}
	poss := len(field.tailResult)
	correct := 36
	if poss != correct {
		t.Error("Solution 1 failed with", poss, "instead of", correct)
	}
}
