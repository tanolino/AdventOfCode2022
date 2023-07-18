package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// var inputFile string = "test_input"

var inputFile string = "input"

type rucksack struct {
	left  string
	right string
}

func makeRucksack(line string) rucksack {
	if len(line)%2 == 1 {
		panic("Failed to process line " + line)
	}
	half := len(line) / 2

	return rucksack{
		left:  line[:half],
		right: line[half:],
	}
}

func (r rucksack) findDoubles() string {
	var res string = ""
	for _, cL := range r.left {
		if strings.ContainsRune(r.right, cL) &&
			!strings.ContainsRune(res, cL) {
			res += string(cL)
		}
	}
	return res
}

func runeValue(c rune) int {
	if c >= 'a' && c <= 'z' {
		return ((int)(c) - (int)('a')) + 1
	} else if c >= 'A' && c <= 'Z' {
		return ((int)(c) - (int)('A')) + 27
	} else {
		panic("Can not determine value of " + string(c))
	}
}

func readInput() []rucksack {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := make([]rucksack, 0)
	for scanner.Scan() {
		res = append(res, makeRucksack(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return res
}

func main() {
	rucksacks := readInput()
	sumOfAllPrios := 0
	for _, r := range rucksacks {
		d := r.findDoubles()

		prio := 0
		for _, ru := range d {
			prio += runeValue(ru)
		}

		// fmt.Println("Rucksack ", r, " doubles: ", d, " prio: ", prio)
		sumOfAllPrios += prio
	}
	fmt.Println("Prio sum: ", sumOfAllPrios)
}
