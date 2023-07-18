package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// var inputFile = "test_input"

var inputFile = "input"

type section struct {
	start int
	end   int
}

func makeSection(s string) section {
	v := strings.Split(s, "-")
	if len(v) != 2 {
		panic("Illegal range definition: " + s)
	}
	start, err := strconv.Atoi(v[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(v[1])
	if err != nil {
		panic(err)
	}
	return section{
		start: start,
		end:   end,
	}
}

func (s1 section) contains(s2 section) bool {
	return s1.start <= s2.start && s1.end >= s2.end
}

type pair struct {
	sec1 section
	sec2 section
}

func makePair(s string) pair {
	sp := strings.Split(s, ",")
	if len(sp) != 2 {
		panic("Illegal pair definition: " + s)
	}
	return pair{
		sec1: makeSection(sp[0]),
		sec2: makeSection(sp[1]),
	}
}

func (p pair) oneFullyContainsTheOther() bool {
	return p.sec1.contains(p.sec2) || p.sec2.contains(p.sec1)
}

func readInput() []pair {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	res := make([]pair, 0)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		res = append(res, makePair(line))
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}
	return res
}

func main() {
	input := readInput()

	counter := 0
	for _, p := range input {
		if p.oneFullyContainsTheOther() {
			counter++
		}
	}
	fmt.Println("Counted: ", counter)
}
