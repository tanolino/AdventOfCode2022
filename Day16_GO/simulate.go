package main

import "fmt"

type simulate struct {
	graph    graph
	open     []bool // already opened
	position int    // I am here
	released int    // steam released
	passed   int    // minutes passed
}

func makeSimulate(g graph) simulate {
	pos := 0
	for i, v := range g.valves {
		if v.name == "AA" {
			pos = i
			break
		}
	}
	return simulate{
		graph:    g,
		open:     make([]bool, len(g.valves)),
		position: pos,
		released: 0,
	}
}

func (s simulate) checkReleaseOn(v int) int {
	timeToActivate := s.graph.reach[s.position][v] + 1
	runTime := 30 - timeToActivate - s.passed
	return runTime * s.graph.valves[v].flow
}

func (s simulate) String() string {
	res := ""
	for i, o := range s.open {
		if o {
			res += fmt.Sprintln(i, "open")
		} else {
			res += fmt.Sprintln(i, "get", s.checkReleaseOn(i))
		}
	}
	return res
}
