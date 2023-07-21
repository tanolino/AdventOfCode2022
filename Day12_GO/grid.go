package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type status int

const (
	hot status = iota
	end
	none
)

type cell struct {
	height, steps int
	status        status
}

func makeCell(ru rune) cell {
	c := cell{
		steps:  -1,
		status: none,
	}
	if ru == 'S' {
		c.status = hot
		c.height = 0
		c.steps = 0
	} else if ru == 'E' {
		c.status = end
		c.height = (int)('z') - (int)('a')
	} else {
		c.height = (int)(ru) - (int)('a')
	}
	return c
}

func (this *cell) discover(other *cell) bool {
	if other == nil {
		return false // cant march outside
	}
	if other.height > this.height+1 {
		return false
	}

	stepsToOther := this.steps + 1
	if other.steps >= 0 && stepsToOther > other.steps {
		return false
	}

	other.steps = stepsToOther
	if other.status != end {
		other.status = hot
	}
	return true
}

type grid struct {
	x, y  int
	cells []cell
}

func makeGrid(filename string) grid {
	g := grid{}

	file, err := os.Open("grid/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		if g.x <= 0 {
			g.x = len(scan.Text())
		}
		g.y++
	}
	g.cells = make([]cell, g.x*g.y)

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		panic(err)
	}
	scan = bufio.NewScanner(file)
	for line := 0; scan.Scan(); line++ {
		txt := scan.Text()

		for i, ru := range txt {
			g.cells[line*g.x+i] = makeCell(ru)
		}
	}
	return g
}

func (g *grid) getCell(x, y int) *cell {
	if x < 0 || y < 0 || x >= g.x || y >= g.y {
		return nil
	} else {
		return &g.cells[y*g.x+x]
	}
}

func (g *grid) printGridHeight() {
	fmt.Println()
	for y := 0; y < g.y; y++ {
		for x := 0; x < g.x; x++ {
			c := g.getCell(x, y)
			if c.status == hot {
				fmt.Print("H")
			} else if c.status == end {
				fmt.Print("E")
			} else if c.status == none {
				s := (rune)((int)('a') + c.height)
				fmt.Print(string(s))
			}
		}
		fmt.Println()
	}
}

func (g *grid) printGridSteps() {
	fmt.Println()
	for y := 0; y < g.y; y++ {
		for x := 0; x < g.x; x++ {
			c := g.getCell(x, y)
			if c.steps < 0 {
				fmt.Print(".")
			} else if c.steps < 10 {
				fmt.Print(c.steps)
			} else if c.steps < 36 {
				s := (rune)((int)('a') + c.steps - 10)
				fmt.Print(string(s))
			} else {
				s := (rune)((int)('A') + c.steps - 36)
				fmt.Print(string(s))
			}
		}
		fmt.Println()
	}
}

func (g *grid) microStepAt(x, y int) bool {
	// c is never nil
	c := g.getCell(x, y)
	if c.status != hot {
		return false // no progress here
	}

	progress := false
	check := func(x, y int) {
		if c.discover(g.getCell(x, y)) {
			progress = true
		}
	}
	check(x-1, y)
	check(x+1, y)
	check(x, y-1)
	check(x, y+1)
	c.status = none

	return progress
}

func (g *grid) step() bool {
	progress := false

	for y := 0; y < g.y; y++ {
		for x := 0; x < g.x; x++ {
			if g.microStepAt(x, y) {
				progress = true
			}
		}
	}

	return progress
}

func (g *grid) marchToEnd() int {
	for g.step() {
		// g.printGridSteps()
	}

	for y := 0; y < g.y; y++ {
		for x := 0; x < g.x; x++ {
			c := g.getCell(x, y)
			if c.status == end {
				return c.steps
			}
		}
	}
	return -1
}

func (g *grid) resetStartAtAllLowPoints() {
	for y := 0; y < g.y; y++ {
		for x := 0; x < g.x; x++ {
			c := g.getCell(x, y)
			if c.height == 0 {
				c.steps = 0
				c.status = hot
			} else {
				c.steps = -1
				if c.status != end {
					c.status = none
				}
			}
		}
	}
}
