package main

import "fmt"

func main() {
	grid := generateGrid(readScan("input"))
	spawns := grid.spawnAll()
	grid.printGrid()
	// 645 too low
	fmt.Println("Spawned:", spawns)
}
