package main

import (
	"fmt"
	"os"
	"strings"
)

// Start means 4 contiguous but differnt following runes
func isStart(s string) bool {
	for len(s) > 1 {
		ru := s[0]
		rem := s[1:]
		if strings.Contains(rem, string(ru)) {
			return false
		}
		s = rem
	}
	return true
}

func findStart(s string) int {
	for i := 0; i+3 < len(s); i++ {
		if isStart(s[i : i+4]) {
			return i + 4 // start pos is after the sequence
		}
	}
	panic("No start found")
}

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Start: ", findStart(string(b)))
}
