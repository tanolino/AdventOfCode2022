package main

import (
	"testing"
)

func Test1(t *testing.T) {
	grid := generateGrid(readScan("test"))
	spawns := grid.spawnAll()
	if spawns != 24 {
		t.Error("Spawn mismatch", spawns, " != 24")
	}
}
