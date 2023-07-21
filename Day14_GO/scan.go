package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func makeCoord(s string) coord {
	sp := strings.Split(s, ",")
	if len(sp) != 2 {
		panic("Malformed coord " + s)
	}
	x, err := strconv.Atoi(sp[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(sp[1])
	if err != nil {
		panic(err)
	}
	return coord{
		x: x,
		y: y,
	}
}

func clampTo1n1(i int) int {
	if i <= -1 {
		return -1
	} else if i >= 1 {
		return 1
	} else {
		return 0
	}
}

// return all coordinates between points
func (c1 coord) interpolate(c2 coord) []coord {
	dX := clampTo1n1(c2.x - c1.x)
	dY := clampTo1n1(c2.y - c1.y)
	res := make([]coord, 0)
	if dX != 0 {
		for x := c1.x + dX; x != c2.x; x += dX {
			res = append(res,
				coord{
					x: x,
					y: c1.y,
				},
			)
		}
	} else if dY != 0 {
		for y := c1.y + dY; y != c2.y; y += dY {
			res = append(res,
				coord{
					x: c1.x,
					y: y,
				},
			)
		}
	}
	return res
}

func (left coord) sub(right coord) coord {
	return coord{
		x: left.x - right.x,
		y: left.y - right.y,
	}
}

type scanline []coord

func makeScanline(line string) scanline {
	line = strings.ReplaceAll(line, "->", "  ")
	fields := strings.Fields(line)
	if len(fields) == 0 {
		panic("Empty scanline: " + line)
	}
	res := make(scanline, 0)
	for _, f := range fields {
		res = append(res, makeCoord(f))
	}
	return res
}

// return a list for every place beeing blocked with this scanline
func (sl scanline) interpolate() []coord {
	res := make([]coord, 0)
	res = append(res, sl[0])
	for i := 1; i < len(sl); i++ {
		res = append(
			res,
			sl[i-1].interpolate(sl[i])...,
		)
		res = append(
			res,
			sl[i],
		)
	}
	return res
}

type scan []scanline

func makeScan(s *bufio.Scanner) scan {
	res := make(scan, 0)
	for s.Scan() {
		line := s.Text()
		res = append(res, makeScanline(line))
	}
	return res
}

func readScan(filename string) scan {
	file, err := os.Open("scans/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	res := makeScan(scan)
	if err := scan.Err(); err != nil {
		panic(err)
	}

	return res
}

// returns minimum x, y / maximum x, y
func (s scan) minmax() (coord, coord) {
	if len(s) == 0 {
		panic("no data")
	}
	min := coord{
		x: s[0][0].x,
		y: s[0][0].y,
	}
	max := min
	for sl := 0; sl < len(s); sl++ {
		for co := 0; co < len(s[sl]); co++ {
			tempX := s[sl][co].x
			if tempX < min.x {
				min.x = tempX
			} else if tempX > max.x {
				max.x = tempX
			}

			tempY := s[sl][co].y
			if tempY < min.y {
				min.y = tempY
			} else if tempY > max.y {
				max.y = tempY
			}
		}
	}
	return min, max
}
