package main

import "fmt"

func main() {
	scan := readScan(scan_test)
	g := makeGraph(scan)
	sim := makeSimulate(g)
	fmt.Println(sim)
}
