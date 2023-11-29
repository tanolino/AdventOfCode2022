package main

import (
	"fmt"
	"math"
)

type graph struct {
	valves []valveDef
	reach  [][]int
}

func contains(list []string, elem string) bool {
	for _, v := range list {
		if elem == v {
			return true
		}
	}
	return false
}

func intMin(i1, i2 int) int {
	if i1 < i2 {
		return i1
	} else {
		return i2
	}
}

func marchReachable(r1 []int, time int, r2 []int) {
	for i, v := range r2 {
		if v < math.MaxInt {
			r1[i] = intMin(r1[i], v+time)
		}
	}
}

func march(reach [][]int) {
	for i := range reach {
		for j := range reach {
			if i == j {
				continue
			} else if time := reach[i][j]; time < math.MaxInt {
				marchReachable(
					reach[i],
					time,
					reach[j],
				)
			}
		}
	}
}

func makeGraph(v []valveDef) graph {
	valveCount := len(v)
	reach := make([][]int, valveCount)
	for i := range reach {
		reach[i] = make([]int, valveCount)
		for j := range reach[i] {
			i_to_j := math.MaxInt
			if i == j {
				i_to_j = 0
			} else if contains(v[i].tunnels, v[j].name) {
				i_to_j = 1
			}
			reach[i][j] = i_to_j
		}
	}

	for range reach {
		march(reach)
	}

	return graph{
		valves: v,
		reach:  reach,
	}
}

func (g graph) String() string {
	res := ""
	for i, r := range g.reach {
		res += fmt.Sprintln(i, ":", r)
	}
	return res
}
