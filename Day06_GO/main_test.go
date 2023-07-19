package main

import "testing"

func Test_Package_1(t *testing.T) {
	pos := findStartOfPackage("bvwbjplbgvbhsrlpgdmjqwftvncz")
	posRight := 5
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Package_2(t *testing.T) {
	pos := findStartOfPackage("nppdvjthqldpwncqszvftbrmjlhg")
	posRight := 6
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
func Test_Package_3(t *testing.T) {
	pos := findStartOfPackage("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	posRight := 10
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
func Test_Package_4(t *testing.T) {
	pos := findStartOfPackage("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	posRight := 11
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Message_1(t *testing.T) {
	pos := findStartOfMessage("mjqjpqmgbljsphdztnvjfqwrcgsmlb")
	posRight := 19
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Message_2(t *testing.T) {
	pos := findStartOfMessage("bvwbjplbgvbhsrlpgdmjqwftvncz")
	posRight := 23
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Message_3(t *testing.T) {
	pos := findStartOfMessage("nppdvjthqldpwncqszvftbrmjlhg")
	posRight := 23
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Message_4(t *testing.T) {
	pos := findStartOfMessage("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg")
	posRight := 29
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}

func Test_Message_5(t *testing.T) {
	pos := findStartOfMessage("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw")
	posRight := 26
	if pos != posRight {
		t.Error("Failed. Position was ", pos, " instead of ", posRight)
	}
}
