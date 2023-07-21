package main

import "fmt"

func main() {
	d := readData("input")
	sum := 0
	for i, v := range d {
		cmp := v.compare()
		fmt.Println(v, "=====>", cmp)
		if cmp < 0 {
			sum += i + 1
		} else if cmp == 0 {
			err := fmt.Sprintln("undefined", i, "pair")
			err += fmt.Sprintln(v)
			panic(err)
		}
	}
	// 5362 too low
	// 5832 too height
	fmt.Println("Sum", sum)
}
