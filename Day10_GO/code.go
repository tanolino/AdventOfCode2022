package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type command struct {
	cycles int
	deltaX int
}

func makeCommand(s string) command {
	sp := strings.Split(s, " ")
	if sp[0] == "noop" && len(sp) == 1 {
		return command{
			cycles: 1,
			deltaX: 0,
		}
	} else if sp[0] == "addx" && len(sp) == 2 {
		dx, err := strconv.Atoi(sp[1])
		if err != nil {
			panic(err)
		}
		return command{
			cycles: 2,
			deltaX: dx,
		}
	} else {
		panic("Unknown command: " + s)
	}
}

func readCommands(progName string) []command {
	file, err := os.Open("programms/" + progName)
	if err != nil {
		panic(err)
	}

	res := make([]command, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		res = append(
			res,
			makeCommand(line),
		)
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
	return res
}
