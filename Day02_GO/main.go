package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// var filename = "test_input"
var filename = "input"

var mapDraw = map[string]int{
	"A": 0, "B": 1, "C": 2,
	"X": 0, "Y": 1, "Z": 2,
}

func shapeSelectionScore(s string) (int, error) {
	val, exist := mapDraw[s]
	if exist {
		return val + 1, nil
	} else {
		return 0, errors.New("Unknown Input " + s)
	}
}

func outcomeScore(eStr, pStr string) (int, error) {
	e, existE := mapDraw[eStr]
	p, existP := mapDraw[pStr]
	if !existE {
		return 0, errors.New("Failed to map " + eStr)
	} else if !existP {
		return 0, errors.New("Failed to map " + pStr)
	}
	if e == p {
		// Tie
		return 3, nil
	} else if (e+1)%3 == p {
		// We win
		return 6, nil
	} else if e == (p+1)%3 {
		// Enemy win
		return 0, nil
	} else {
		// We "should not end up here"
		return 0, errors.New("Internal error")
	}
}

func valueOfLine(line string) (int, error) {
	s := strings.Split(line, " ")
	if len(s) < 2 {
		return 0, errors.New("Invalid line")
	}

	score1, err := outcomeScore(s[0], s[1])
	if err != nil {
		return 0, err
	}
	score2, err := shapeSelectionScore(s[1])
	if err != nil {
		return 0, err
	}
	return score1 + score2, nil
}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open input: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score := 0
	for scanner.Scan() {
		scoreInner, err := valueOfLine(scanner.Text())
		if err != nil {
			fmt.Println("Failed to parse line ", scanner.Text(), " because: ", err)
			os.Exit(1)
		}
		// fmt.Println("+ ", scoreInner, " for ", scanner.Text())
		score += scoreInner
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		os.Exit(1)
	}
	fmt.Println("Score: ", score)
}
