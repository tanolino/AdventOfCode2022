package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	name              string
	items             []int
	operation         func(int) int
	test              func(int) bool
	targetOnTestTrue  int
	targetOnTestFalse int
	inspections       int
}

type monkeyList []*monkey

func newMonkey(scan *bufio.Scanner) *monkey {
	m := new(monkey)
	if !scan.Scan() {
		return nil // No more monkey in the list
	}
	m.name = scan.Text()
	if !scan.Scan() {
		panic("Expected starting items, found EOF")
	}
	m.items = parseItems(scan.Text())
	if !scan.Scan() {
		panic("Expected operation, found EOF")
	}
	m.operation = parseOperation(scan.Text())
	if !scan.Scan() {
		panic("Expected test, found EOF")
	}
	m.test = parseTest(scan.Text())
	if !scan.Scan() {
		panic("Expected target monkey, found EOF")
	}
	parseTargetMonkey(m, scan.Text())
	if !scan.Scan() {
		panic("Expected target monkey, found EOF")
	}
	parseTargetMonkey(m, scan.Text())
	if scan.Scan() {
		e := scan.Text()
		if e != "" {
			panic("Expected empty line, received: " + e)
		}
	}
	m.inspections = 0 // redundant ?
	return m
}

func parseItems(txt string) []int {
	res := make([]int, 0)
	txt = strings.ReplaceAll(txt, ",", "")
	sp := strings.Fields(txt)
	if sp[0] != "Starting" {
		panic("Malformed (0) " + sp[0])
	} else if sp[1] != "items:" {
		panic("Malformed (1) " + sp[1])
	}
	for i := 2; i < len(sp); i++ {
		num, err := strconv.Atoi(sp[i])
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}

func parseOperand(o string) func(int) int {
	if o == "old" {
		return func(i int) int {
			return i
		}
	} else {
		num, err := strconv.Atoi(o)
		if err != nil {
			panic(err)
		}
		return func(_ int) int {
			return num
		}
	}
}

func parseOperationSign(o string) func(int, int) int {
	if o == "+" {
		return func(i1, i2 int) int {
			return i1 + i2
		}
	} else if o == "*" {
		return func(i1, i2 int) int {
			return i1 * i2
		}
	} else {
		panic("Unexpected operand: " + o)
	}
}

func parseOperation(o string) func(int) int {
	sp := strings.Fields(o)
	if sp[0] != "Operation:" ||
		sp[1] != "new" ||
		sp[2] != "=" {
		panic("Malformed operation")
	}

	o1f := parseOperand(sp[3])
	op := parseOperationSign(sp[4])
	o2f := parseOperand(sp[5])

	return func(i int) int {
		return op(o1f(i), o2f(i))
	}
}

func parseTest(t string) func(int) bool {
	sp := strings.Fields(t)
	if sp[0] != "Test:" ||
		sp[1] != "divisible" ||
		sp[2] != "by" {
		panic("Test line malformed")
	}
	num, err := strconv.Atoi(sp[3])
	if err != nil {
		panic(err)
	}

	return func(i int) bool {
		return i%num == 0
	}
}

func parseTargetMonkey(m *monkey, t string) {
	sp := strings.Fields(t)
	if sp[0] != "If" ||
		sp[2] != "throw" ||
		sp[3] != "to" ||
		sp[4] != "monkey" {
		panic("Throw to money malformed")
	}
	num, err := strconv.Atoi(sp[5])
	if err != nil {
		panic(err)
	}
	if sp[1] == "true:" {
		m.targetOnTestTrue = num
	} else if sp[1] == "false:" {
		m.targetOnTestFalse = num
	} else {
		panic("Malformed condition for throwing")
	}
}

func readMonkeys(filename string) monkeyList {
	file, err := os.Open("data/" + filename)
	if err != nil {
		panic(err)
	}
	res := make(monkeyList, 0)
	scan := bufio.NewScanner(file)
	for m := newMonkey(scan); m != nil; m = newMonkey(scan) {
		res = append(res, m)
	}
	return res
}

// (item, targetMonkey)
func (m *monkey) inspectAndThrow(monkeys *monkeyList) {
	for len(m.items) > 0 {
		m.inspections++
		item := m.items[0]
		m.items = m.items[1:]
		itemNew := m.operation(item) / 3
		if m.test(itemNew) {
			(*monkeys)[m.targetOnTestTrue].catchItem(itemNew)
		} else {
			(*monkeys)[m.targetOnTestFalse].catchItem(itemNew)
		}
	}
}

func (m *monkey) catchItem(item int) {
	m.items = append(m.items, item)
}

func (ml *monkeyList) oneRound() {
	for _, m := range *ml {
		m.inspectAndThrow(ml)
	}
}

func (ml *monkeyList) rounds(r int) {
	for i := 0; i < r; i++ {
		ml.oneRound()
	}
}

func (ml monkeyList) printList() {
	for _, m := range ml {
		fmt.Println(m)
	}
}

func (ml monkeyList) getMonkeyBusiness() int {
	ins := make([]int, len(ml))
	for i, m := range ml {
		ins[i] = m.inspections
	}
	sort.Ints(ins)
	return ins[len(ins)-1] * ins[len(ins)-2]
}
