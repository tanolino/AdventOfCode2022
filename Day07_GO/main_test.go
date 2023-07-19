package main

import (
	"testing"
)

func Test_Part1(t *testing.T) {
	rootDir := readInput("test_input")
	// rootDir.printOut("")
	sum := 0
	sumRight := 95437
	for _, d := range findDirsRecursive(
		rootDir, filterBelow100000,
	) {
		sum += d.getSize()
	}
	if sum != sumRight {
		t.Error("Sum missmatch ", sum, " != ", sumRight)
	}
}

func Test_Part2(t *testing.T) {
	rootDir := readInput("test_input")
	freeSpace := deviceSpace - rootDir.getSize()
	spaceNeeded := deviceSpaceForUpdate - freeSpace

	filter := func(d *myDir) bool {
		return d.getSize() >= spaceNeeded
	}

	var delDir *myDir
	var delDirSize int
	for _, d := range findDirsRecursive(
		rootDir, filter,
	) {
		dSize := d.getSize()
		if delDir == nil || delDirSize > dSize {
			delDir = d
			delDirSize = dSize
		}
	}
	correctDelDirSize := 24933642
	if delDirSize != correctDelDirSize {
		t.Error("Del dir missmatch ", delDirSize,
			" != ", correctDelDirSize)
	}
}
