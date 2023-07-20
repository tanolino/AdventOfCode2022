package main

import (
	"fmt"
	"testing"
)

func Test_1(t *testing.T) {
	monkeys := readMonkeys("test1")
	monkeys.rounds(20)
	sumCorr := 10605
	sum := monkeys.getMonkeyBusiness()
	if sum != sumCorr {
		t.Error("Sum missmatch.", sum, "instead of", sumCorr)
	}
}

func Test_Part1(t *testing.T) {
	monkeys := readMonkeys("input")
	monkeys.rounds(20)
	sumCorr := 120756
	sum := monkeys.getMonkeyBusiness()
	if sum != sumCorr {
		t.Error("Sum missmatch.", sum, "instead of", sumCorr)
	}
}

func Test_2(t *testing.T) {
	monkeys := readMonkeys("test1")
	monkeys.maximumWorry = true

	for i := 0; i < 10000; i++ {
		if i == 1 || i == 20 || i == 1000 ||
			i == 2000 || i == 3000 || i == 4000 {
			fmt.Println("After", i, "rounds")
			monkeys.printList()
		}
		monkeys.oneRound()
	}

	fmt.Println("After", 10000, "rounds")
	monkeys.printList()

	sumCorr := 2713310158
	sum := monkeys.getMonkeyBusiness()
	if sum != sumCorr {
		t.Error("Sum missmatch.", sum, "instead of", sumCorr)
	}
}
