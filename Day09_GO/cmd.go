package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type direction int

const (
	right direction = iota
	left
	up
	down
)

var mapToDir = map[string]direction{
	"R": right,
	"L": left,
	"U": up,
	"D": down,
}

type cmd struct {
	dir   direction
	steps int
}

func makeCmd(s string) cmd {
	sp := strings.Split(s, " ")
	dir, avail := mapToDir[sp[0]]
	if len(sp) != 2 {
		panic("Illegal format: " + s)
	}
	if !avail {
		panic("Failed to map " + sp[0])
	}
	num, err := strconv.Atoi(sp[1])
	if err != nil {
		panic(err)
	}
	return cmd{
		dir:   dir,
		steps: num,
	}
}

func readCmds(filename string) []cmd {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	res := make([]cmd, 0)
	for scan.Scan() {
		line := scan.Text()
		c := makeCmd(line)
		res = append(res, c)
	}
	return res
}
