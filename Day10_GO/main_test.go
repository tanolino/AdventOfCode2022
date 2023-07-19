package main

import (
	"testing"
)

func Test_1(t *testing.T) {
	cmds := readCommands("test1")
	// Simple simulation on our own behalf
	cycleCount := 0
	x := 1
	for _, cmd := range cmds {
		cycleCount += cmd.cycles
		x += cmd.deltaX
	}
	if cycleCount != 5 {
		t.Error("Counted", cycleCount, " instead of 5")
	}
	if x != -1 {
		t.Error("X is", x, " instead of -1")
	}

	p := makeProcessor(cmds)
	p.cycles(cycleCount + 1)
	if p.regX != x {
		t.Error("Processor is off.", p.regX, "instead of", x)
	}
}

func Test_2(t *testing.T) {
	p := makeProcessor(readCommands("test2"))
	res := []int{420, 1140, 1800, 2940, 2880, 3960}
	p.cycles(20)
	sum := 0
	for _, v := range res {
		sig := p.signal()
		if sig != v {
			t.Error("Expected X to be", v,
				"but we have", sig, "\n", p,
			)
		}
		sum += sig
		p.cycles(40)
	}
	correctSum := 13140
	if sum != correctSum {
		t.Error("Sum missmatch:", sum, "instead of", correctSum)
	}
}

func Test_Part1(t *testing.T) {
	p := makeProcessor(readCommands("part1"))
	sum := 0
	p.cycles(20)
	for i := 0; i < 6; i++ {
		sig := p.signal()
		sum += sig
		p.cycles(40)
	}
	correctSum := 17940
	if sum != correctSum {
		t.Error("Sum missmatch:", sum, "instead of", correctSum)
	}
}
