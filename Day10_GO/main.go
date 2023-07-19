package main

import "fmt"

func part1() {
	p := makeProcessor(readCommands("part1"))
	sum := 0
	p.cycles(20)
	for i := 0; i < 6; i++ {
		sig := p.signal()
		// fmt.Println("Signal:", sig)
		sum += sig
		p.cycles(40)
	}
	fmt.Println("Signal Sum:", sum)
}

func part2() {
	progName := "part1"
	//progName := "test2"
	p := makeProcessor(readCommands(progName))

	for row := 0; row < 6; row++ {
		for pos := 0; pos < 40; pos++ {
			p.cycle()
			if closeEnough(p.regX, pos) {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func closeEnough(i1, i2 int) bool {
	res := i1 - i2
	return res >= -1 && res <= 1
}

func main() {
	part1()
	part2()
}
