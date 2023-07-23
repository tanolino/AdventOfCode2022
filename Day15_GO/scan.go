package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func intAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

type coord struct {
	x, y int
}

func makeCoord(xStr, yStr string) coord {
	x, err := strconv.Atoi(xStr)
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(yStr)
	if err != nil {
		panic(err)
	}
	return coord{x, y}
}

type sensor struct {
	pos, beacon coord
}

func (s sensor) manhattan() int {
	return intAbs(s.pos.x-s.beacon.x) +
		intAbs(s.pos.y-s.beacon.y)
}

func makeSensor(s string) sensor {
	s = strings.Replace(s, "Sensor at x=", "", 1)
	s = strings.ReplaceAll(s, ", y=", " ")
	s = strings.Replace(s, ": closest beacon is at x=", " ", 1)
	fields := strings.Fields(s)

	return sensor{
		pos:    makeCoord(fields[0], fields[1]),
		beacon: makeCoord(fields[2], fields[3]),
	}
}

type scan []sensor

func readScan(filename string) scan {
	file, err := os.Open("scans/" + filename)
	if err != nil {
		panic(err)
	}
	res := make(scan, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		txt := scanner.Text()
		sens := makeSensor(txt)
		res = append(res, sens)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return res
}
