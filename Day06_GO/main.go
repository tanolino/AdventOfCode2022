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

func findStart(s string, l int) int {
	for i := 0; i+l <= len(s); i++ {
		if isStart(s[i : i+l]) {
			return i + l // start pos is after the sequence
		}
	}
	panic("No start found")
}

func findStartOfPackage(s string) int {
	return findStart(s, 4)
}

func findStartOfMessage(s string) int {
	return findStart(s, 14)
}

func main() {
	b, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	txt := string(b)
	fmt.Println("Start (Package): ", findStartOfPackage(txt))
	fmt.Println("Start (Message): ", findStartOfMessage(txt))
}
