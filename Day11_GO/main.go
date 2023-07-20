package main

import "fmt"

func main() {
	monkeys := readMonkeys("input")
	monkeys.rounds(20)
	sum := monkeys.getMonkeyBusiness()
	fmt.Println("Monkey Business", sum)
}
