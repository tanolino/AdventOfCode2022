package main

import "fmt"

func main() {
	//scan := readScan(scan_test)
	scan := readScan(scan_input)
	g := makeGraph(scan)
	vNum := make([]int, len(g.valves))
	for i := range g.valves {
		vNum[i] = i
	}
	fmt.Println(g)

	sim2 := makeSimulate(g)
	fmt.Println("DD BB JJ HH EE CC ",
		sim2.runScenario([]int{3, 1, 9, 7, 4, 2}),
	)
	// return
	// */
	maxVal := 0
	maxSet := make([]int, len(g.valves))
	perm := countPermut(len(g.valves))
	for i := 0; i < perm; i++ {
		sim := makeSimulate(g)
		set := permutate(vNum, i)
		released := sim.runScenario(set)
		if maxVal < released {
			maxVal = released
			fmt.Println("New Max", released, set)
			maxSet = set
		}
	}
	fmt.Print("Max steam ", maxVal, " with: ")
	for _, i := range maxSet {
		fmt.Print(g.valves[i].name, "(", i, ") ")
	}
	fmt.Println()
}
