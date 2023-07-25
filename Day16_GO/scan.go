package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type inputFile string

const (
	scan_test  inputFile = "test"
	scan_input inputFile = "input"
)

type valveDef struct {
	name    string
	flow    int
	tunnels []string
}

func makeValveDef(s string) valveDef {
	sp := strings.Split(s, ";")
	s = strings.Replace(sp[0], "Valve ", "", 1)
	s = strings.Replace(s, " has flow rate=", " ", 1)
	fields := strings.Fields(s)
	name := fields[0]
	flow, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}

	s = strings.ReplaceAll(sp[1], ",", "")
	fields = strings.Fields(s)
	return valveDef{
		name:    name,
		flow:    flow,
		tunnels: fields[4:],
	}
}

func readScan(in inputFile) []valveDef {
	file, err := os.Open("scans/" + string(in))
	if err != nil {
		panic(err)
	}
	res := make([]valveDef, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		res = append(res, makeValveDef(line))
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
	return res
}
