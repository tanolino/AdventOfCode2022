package main

type simulate struct {
	graph    graph
	position int // I am here
	released int // steam released
	passed   int // minutes passed
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
		position: pos,
		released: 0,
	}
}

func (s simulate) checkReleaseOn(v int) (steam int, timePassed int) {
	timeToActivate := s.graph.reach[s.position][v] + 1
	runTime := (30 - timeToActivate) - s.passed
	return runTime * s.graph.valves[v].flow, timeToActivate
}

func (s *simulate) runScenario(val []int) int {
	for _, v := range val {
		rel, t := s.checkReleaseOn(v)
		// fmt.Println("Debug: ", rel, t, s.passed)
		if rel >= 0 {
			// fmt.Println("Open ", s.graph.valves[v].name, " for ", rel)
			s.released += rel
			s.passed += t
			s.position = v
		} else {
			break
		}
	}
	return s.released
}
