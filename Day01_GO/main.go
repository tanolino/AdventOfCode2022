package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calories int
}

func readElfes() ([]elf, error) {
	file, err := os.Open("input")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	current := 0
	res := make([]elf, 0)
	for scanner.Scan() {
		txt := scanner.Text()
		if len(txt) > 0 {
			// we have a number
			num, err := strconv.Atoi(txt)
			if err != nil {
				return nil, err
			} else {
				current += num
			}
		} else {
			// empty line means pepare for next set
			res = append(res, elf{calories: current})
			current = 0
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if current > 0 {
		res = append(res, elf{calories: current})
	}
	return res, nil
}

// To sort the list of elfes
type ByCalories []elf

func (t ByCalories) Len() int           { return len(t) }
func (t ByCalories) Less(i, j int) bool { return t[i].calories < t[j].calories }
func (t ByCalories) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func main() {
	elfes, err := readElfes()
	if err != nil {
		fmt.Println("Failed to read input file: ", err)
		os.Exit(1)
	} else if len(elfes) < 3 {
		fmt.Println("Not enough elves")
		os.Exit(1)
	}

	sort.Stable(ByCalories(elfes))
	fmt.Println("Maximum kalories are: ", elfes[len(elfes)-1])

	highest3 := 0
	for i := 0; i < 3; i++ {
		highest3 += elfes[len(elfes)-i-1].calories
	}
	fmt.Println("Highest 3 summed up: ", highest3)
}
