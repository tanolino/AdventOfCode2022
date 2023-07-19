package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// [row][col]
type forest [][]int

func makeForest(filename string) forest {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	res := make(forest, 0)
	for scan.Scan() {
		txt := scan.Text()
		row := make([]int, len(txt))
		for i := 0; i < len(txt); i++ {
			num, err := strconv.Atoi(txt[i : i+1])
			if err != nil {
				panic(err)
			}
			row[i] = num
		}

		res = append(res, row)
	}
	return res
}

func (f forest) isVisible(row, col int) bool {
	height := f[row][col]

	// How many sides are hiding it
	hidden := 0
	// Check row visibility
	for r := 0; r < row; r++ {
		if f[r][col] >= height {
			hidden++
			break
		}
	}
	for r := row + 1; r < len(f); r++ {
		if f[r][col] >= height {
			hidden++
			break
		}
	}

	// Check col visibility
	for c := 0; c < col; c++ {
		if f[row][c] >= height {
			hidden++
			break
		}
	}
	for c := col + 1; c < len(f[row]); c++ {
		if f[row][c] >= height {
			hidden++
			break
		}
	}

	return hidden < 4
}

var debugOut bool = false

func (f forest) countVisibleTrees() int {
	res := 0
	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			if f.isVisible(r, c) {
				res++
				if debugOut {
					fmt.Print("X")
				}
			} else if debugOut {
				fmt.Print("_")
			}
		}
		if debugOut {
			fmt.Println()
		}
	}
	return res
}

func (f forest) scenicScoreDir(r, c, dR, dC int) int {
	height := f[r][c]
	score := 0

	if dR == 0 && dC == 0 {
		panic("Delta 0 not allowed")
	}

	r += dR
	c += dC
	for r >= 0 && r < len(f) && c >= 0 && c < len(f[r]) {
		score++
		if f[r][c] >= height {
			break
		} else {
			r += dR
			c += dC
		}
	}

	// fmt.Println("ScoreDir", r, c, dR, dC, "=>", score)
	return score
}

func (f forest) scenicScoreAt(r, c int) int {
	fun := func(dR, dC int) int {
		return f.scenicScoreDir(r, c, dR, dC)
	}
	return fun(1, 0) *
		fun(-1, 0) *
		fun(0, 1) *
		fun(0, -1)
}

func (f forest) maxScenicScore() int {
	res := 0
	for r := 0; r < len(f); r++ {
		for c := 0; c < len(f[r]); c++ {
			score := f.scenicScoreAt(r, c)
			if score > res {
				res = score
			}
			if debugOut {
				fmt.Printf("%X", score)
			}
		}
		if debugOut {
			fmt.Println()
		}
	}
	return res
}

func main() {
	f := makeForest("input")
	i := f.countVisibleTrees()
	fmt.Println("There are", i, "visible trees")
	s := f.maxScenicScore()
	fmt.Println("Max scientic score:", s)
}
