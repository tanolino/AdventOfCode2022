package main

import "fmt"

func part1() {
	d := readDataAsPairs("input")
	sum := 0
	for i, v := range d {
		cmp := v.compare()
		// fmt.Println(v, "=====>", cmp)
		if cmp < 0 {
			sum += i + 1
		} else if cmp == 0 {
			err := fmt.Sprintln("undefined", i, "pair")
			err += fmt.Sprintln(v)
			panic(err)
		}
	}
	fmt.Println("Sum", sum)
}

func part2() {
	key := readDecoderKey("input")
	fmt.Println("Key: ", key)
}

func main() {
	part1()
	part2()
}
