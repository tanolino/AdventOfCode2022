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

func countVisibleTrees(f forest) int {
	res := 0
	debugOut := true
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
		fmt.Println()
	}
	return res
}

func main() {
	f := makeForest("input")
	i := countVisibleTrees(f)
	fmt.Println("There are", i, "visible trees")
}
