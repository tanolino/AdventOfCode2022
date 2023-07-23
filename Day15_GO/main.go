package main

import (
	"fmt"
	"math"
)

func checkScanlinePosWithBeacon(line, pos int, sc scan) bool {
	for _, sl := range sc {
		steppy := sl.manhattan()
		steppy -= intAbs(sl.pos.y - line)
		steppy -= intAbs(sl.pos.x - pos)
		if steppy < 0 {
			continue
		} else if sl.beacon.y == line && sl.beacon.x == pos {
			return false
		} else {
			return true
		}
	}
	return false
}

func checkScanlinePos(line, pos int, sc scan) (bool, int) {
	for _, sl := range sc {
		steppy := sl.manhattan()
		steppy -= intAbs(sl.pos.y - line)
		if steppy < intAbs(sl.pos.x-pos) {
			continue
		} else {
			return true, sl.pos.x + steppy
		}
	}
	return false, 0
}

func checkScanline2(line int, sc scan) int {
	// first we find the start end end point to scan
	start := math.MaxInt
	end := math.MinInt
	for _, sl := range sc {
		mh := sl.manhattan() + 1
		start = intMin(start, sl.pos.x-mh)
		end = intMax(end, sl.pos.x+mh)
	}

	res := 0
	for i := start; i < end; i++ {
		// fmt.Println("Progress:", start, i, end, res)
		if checkScanlinePosWithBeacon(line, i, sc) {
			res++
		}
	}
	return res
}

func findFreeSpace(sc scan, limit int) int {
	// Target  0 <= x, y <= 4000000
	for y := 0; y <= limit; y++ {
		for x := 0; x <= limit; x++ {
			if cond, off := checkScanlinePos(y, x, sc); cond {
				if off < x {
					panic("We mixed something up")
				}
				x = off
				continue
			}
			// Tuning Frequency is:
			return x*4000000 + y
		}
	}
	return -1
}

func main() {
	s := readScan("input")
	res := checkScanline2(2000000, s)
	fmt.Println("Blocked: ", res)

	res = findFreeSpace(s, 4000000)
	fmt.Println("Tuning freq: ", res)
}

func intMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func intMax(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
