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
	if len(res) != 1 {
		panic("Rucksack has too many doubles: " + res)
	}
	return res
}

func (r rucksack) all() string {
	return r.left + r.right
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

func commonElements(r1, r2, r3 rucksack) string {
	res := ""
	r2all := r2.all()
	r3all := r3.all()
	for _, ru := range r1.all() {
		if strings.ContainsRune(r2all, ru) &&
			strings.ContainsRune(r3all, ru) &&
			!strings.ContainsRune(res, ru) {
			res += string(ru)
		}
	}
	if len(res) != 1 {
		panic("common Elements is len() != 1  " + res)
	}
	return res
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

func solvePart1(rucksacks []rucksack) {
	sumOfAllPrios := 0
	for _, r := range rucksacks {
		d := r.findDoubles()

		for _, ru := range d {
			sumOfAllPrios += runeValue(ru)
		}
	}
	fmt.Println("(Part 1) Prio sum: ", sumOfAllPrios)
}

func solvePart2(rucksacks []rucksack) {
	sumOfAllPrios := 0
	for i := 0; i+2 < len(rucksacks); i += 3 {
		common := commonElements(
			rucksacks[i],
			rucksacks[i+1],
			rucksacks[i+2],
		)

		for _, ru := range common {
			sumOfAllPrios += runeValue(ru)
		}
	}
	fmt.Println("(Part 2) Prio sum: ", sumOfAllPrios)
}

func main() {
	rucksacks := readInput()
	solvePart1(rucksacks)
	solvePart2(rucksacks)
}
