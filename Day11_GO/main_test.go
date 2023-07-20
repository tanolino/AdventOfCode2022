package main

import (
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
