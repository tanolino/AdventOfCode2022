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

type Draw int

const (
	Rock = iota
	Paper
	Scissors
)

var mapDraw = map[string]Draw{
	"A": Rock, "B": Paper, "C": Scissors,
	"X": Rock, "Y": Paper, "Z": Scissors,
}

var reverseLookupDraw = map[Draw]string{
	Rock: "A", Paper: "B", Scissors: "C",
}

func getWin(draw Draw) Draw {
	return (draw + 1) % 3
}

func getLoose(draw Draw) Draw {
	return (draw + 2) % 3
}

func getTie(draw Draw) Draw {
	return draw
}

func shapeSelectionScore(s string) (int, error) {
	val, exist := mapDraw[s]
	if exist {
		return (int)(val) + 1, nil
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
	if getTie(e) == p {
		// Tie
		return 3, nil
	} else if getWin(e) == p {
		// We win
		return 6, nil
	} else if getLoose(e) == p {
		// Enemy win
		return 0, nil
	} else {
		// We "should not end up here"
		return 0, errors.New("Internal error")
	}
}

func sumUpScores(enemy, ours string) (int, error) {
	score1, err := outcomeScore(enemy, ours)
	if err != nil {
		return 0, err
	}
	score2, err := shapeSelectionScore(ours)
	if err != nil {
		return 0, err
	}
	return score1 + score2, nil
}

func valueOfLinePart1(line string) (int, error) {
	s := strings.Split(line, " ")
	if len(s) < 2 {
		return 0, errors.New("Invalid line")
	}

	return sumUpScores(s[0], s[1])
}

func valueOfLinePart2(line string) (int, error) {
	s := strings.Split(line, " ")
	if len(s) < 2 {
		return 0, errors.New("Invalid line")
	}
	enemyDraw, exist := mapDraw[s[0]]
	if !exist {
		return 0, errors.New("Unkown enemy pick: " + s[0])
	}

	var ourPick Draw
	if s[1] == "X" {
		// we need to loose
		ourPick = getLoose(enemyDraw)
	} else if s[1] == "Y" {
		// we need a draw
		ourPick = getTie(enemyDraw)
	} else if s[1] == "Z" {
		// we need to win
		ourPick = getWin(enemyDraw)
	} else {
		return 0, errors.New("Unknown input " + s[1])
	}

	ours := reverseLookupDraw[ourPick]

	return sumUpScores(s[0], ours)
}

func main() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open input: ", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	score1 := 0
	score2 := 0
	for scanner.Scan() {
		scoreInner1, err := valueOfLinePart1(scanner.Text())
		if err != nil {
			fmt.Println("Failed to parse line ", scanner.Text(), " because: ", err)
			os.Exit(1)
		}
		// fmt.Println("+ ", scoreInner1, " for ", scanner.Text())
		score1 += scoreInner1

		scoreInner2, err := valueOfLinePart2(scanner.Text())
		if err != nil {
			fmt.Println("Failed to parse line ", scanner.Text(), " because: ", err)
			os.Exit(1)
		}
		// fmt.Println("+ ", scoreInner2, " for ", scanner.Text())
		score2 += scoreInner2
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to scan: ", err)
		os.Exit(1)
	}
	fmt.Println("Score first: ", score1)
	fmt.Println("Score second: ", score2)
}
