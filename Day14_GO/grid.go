package main

import "fmt"

type cellType int

const (
	air cellType = iota
	rock
	sand
	infinite
	sandspawner
)

var cellRunes = []string{
	".", // air
	"#", // rock
	"O", // sand
	"~", // infinite
	"+", // sandspawner
}

type row []cellType

func makeRow(width int) row {
	return make(row, width)
}

type grid struct {
	offset        coord
	width, height int
	rows          []row
	sands         int
}

func generateGrid(sc scan) grid {
	// First we build a rectangle with air
	min, max := sc.minmax()
	// we extend for sand overflow
	min.x -= 1
	max.x += 1
	width := max.x - min.x + 1
	height := max.y - min.y + 1
	rows := make([]row, 0)
	for i := 0; i < height; i++ {
		rows = append(rows, makeRow(width))
	}
	res := grid{
		offset: min,
		width:  width,
		height: height,
		rows:   rows,
	}
	// Second we fill in the rocks
	for _, sl := range sc {
		pts := sl.interpolate()
		for _, pt := range pts {
			pt = pt.sub(res.offset)
			rows[pt.y][pt.x] = rock
		}
	}
	rows[0][500-res.offset.x] = sandspawner

	return res
}

func (g grid) printGrid() {
	fmt.Println("Sands: ", g.sands)
	for _, row := range g.rows {
		for _, elem := range row {
			fmt.Print(cellRunes[(int)(elem)])
		}
		fmt.Println()
	}
}

// returns true if sand stick around
func (g *grid) spawn() bool {
	p := coord{500 - g.offset.x, 0}

	for {
		if p.y+1 >= g.height {
			return false
		}
		if g.rows[p.y+1][p.x] == air {
			p.y++
		} else if g.rows[p.y+1][p.x-1] == air {
			p.y++
			p.x--
		} else if g.rows[p.y+1][p.x+1] == air {
			p.y++
			p.x++
		} else {
			break
		}
	}
	g.rows[p.y][p.x] = sand
	g.sands++
	if p.y == 0 {
		newRow := makeRow(g.width)
		newRow[500-g.offset.x] = sandspawner
		g.rows = append([]row{newRow}, g.rows...)
		g.height++
	}
	return true
}

func (g *grid) spawnAll() int {
	for g.spawn() {
		// g.printGrid()
	}
	return g.sands
}
