package main

import "testing"

func Test1(t *testing.T) {
	pos := findStart("bvwbjplbgvbhsrlpgdmjqwftvncz")
	posRight := 5
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test2(t *testing.T) {
	pos := findStart("nppdvjthqldpwncqszvftbrmjlhg")
	posRight := 6
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
func Test3(t *testing.T) {
	pos := findStart("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	posRight := 10
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
func Test4(t *testing.T) {
	pos := findStart("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	posRight := 11
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
