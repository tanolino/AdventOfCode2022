package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputFile() string {
	test := false
	if test {
		return "test_input"
	} else {
		return "input"
	}
}

type stack []rune

func makeStack() stack {
	return make([]rune, 0)
}

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() rune {
	if len(*s) == 0 {
		panic("Stack already empty")
	}
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s stack) last() rune {
	if len(s) == 0 {
		panic("Stack empty")
	}
	return s[len(s)-1]
}

type stacks []stack

func makeStacks(s []string) stacks {
	// dat is the actual stack pieces
	dat := s[:len(s)-1]

	// cols is the position where to find
	// definitions for stack pieces
	cols := make([]int, 0)
	for i, ru := range s[len(s)-1] {
		if ru != ' ' {
			cols = append(cols, i)
		}
	}

	// Initialize result structures
	res := make([]stack, len(cols))
	for i := 0; i < len(res); i++ {
		res[i] = makeStack()
	}

	// fill said structures
	for inv := len(dat) - 1; inv >= 0; inv-- {
		line := dat[inv]
		for i, c := range cols {
			if len(line) <= c {
				continue
			}
			d := rune(line[c])
			if d != ' ' {
				res[i].push(d)
			}
		}
	}

	return res
}

func (s stacks) toString() string {
	res := ""

	for _, v := range s {
		for _, ru := range v {
			res += string(ru)
			res += " "
		}
		res += "\n"
	}

	return res
}

func (s stacks) moveLike9000(cmd command) {
	for i := 0; i < cmd.count; i++ {
		s[cmd.to].push(s[cmd.from].pop())
	}
}

func (s stacks) moveLike9001(cmd command) {
	tmp := makeStack()
	for i := 0; i < cmd.count; i++ {
		tmp.push(s[cmd.from].pop())
	}
	for len(tmp) > 0 {
		s[cmd.to].push(tmp.pop())
	}
}

func (s stacks) getTopCrates() string {
	res := ""
	for _, st := range s {
		res += string(st.last())
	}
	return res
}

type command struct {
	count int
	from  int
	to    int
}

func makeCommand(s string) command {
	sp := strings.Split(s, " ")
	var toI = func(s string) int {
		r, e := strconv.Atoi(s)
		if e != nil {
			panic(e)
		} else {
			return r
		}
	}
	cmd := command{
		count: toI(sp[1]),
		from:  toI(sp[3]) - 1, // correct index offset
		to:    toI(sp[5]) - 1, // correct index offset
	}

	// fmt.Println("Cmd: ", cmd, sp[1], sp[3], sp[5])
	return cmd
}

type inputData struct {
	stacks   stacks
	commands []command
}

func getInput() inputData {
	file, err := os.Open(getInputFile())
	if err != nil {
		panic(err)
	}

	state := make([]string, 0)
	commands := make([]command, 0)

	readState := func(s string) {
		state = append(state, s)
	}
	readCommands := func(s string) {
		commands = append(
			commands,
			makeCommand(s),
		)
	}

	var read = readState

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()

		if len(line) == 0 {
			read = readCommands
		} else {
			read(line)
		}
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}

	return inputData{
		stacks:   makeStacks(state),
		commands: commands,
	}
}

func simulateCrane9000() {
	input := getInput()
	// fmt.Println(input.stacks.toString())

	for _, cmd := range input.commands {

		// fmt.Println("Move ", cmd.from, " to ", cmd.to, " times ", cmd.count)
		input.stacks.moveLike9000(cmd)

		// fmt.Println(input.stacks.toString())
	}
	// fmt.Println(input.stacks.toString())
	fmt.Println("Top Crates (9000): ", input.stacks.getTopCrates())
}

func simulateCrane9001() {
	input := getInput()
	// fmt.Println(input.stacks.toString())

	for _, cmd := range input.commands {

		// fmt.Println("Move ", cmd.from, " to ", cmd.to, " times ", cmd.count)
		input.stacks.moveLike9001(cmd)

		// fmt.Println(input.stacks.toString())
	}
	// fmt.Println(input.stacks.toString())
	fmt.Println("Top Crates (9001): ", input.stacks.getTopCrates())
}

func main() {
	simulateCrane9000()
	simulateCrane9001()
}
