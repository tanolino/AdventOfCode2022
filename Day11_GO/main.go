package main

import "fmt"

func part1() {
	monkeys := readMonkeys("input")
	monkeys.rounds(20)
	sum := monkeys.getMonkeyBusiness()
	fmt.Println("Monkey Business", sum)
}

func part2() {
	monkeys := readMonkeys("input")
	monkeys.maximumWorry = true
	monkeys.rounds(10000)
	sum := monkeys.getMonkeyBusiness()
	fmt.Println("Monkey Business", sum)
}

func main() {
	part1()
	part2()
}
