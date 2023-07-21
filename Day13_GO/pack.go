package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type data struct {
	value int
	list  []data
}

func makeDataValue(i int) data {
	return data{
		value: i,
		list:  nil,
	}
}

func (d data) isValue() bool {
	return d.value >= 0
}

func makeDataList(d []data) data {
	return data{
		value: -1,
		list:  d,
	}
}

func (d data) isList() bool {
	return d.list != nil
}

func (d data) String() string {
	if d.isValue() {
		return fmt.Sprint(d.value)
	} else if d.isList() {
		res := "["
		for i, v := range d.list {
			res += fmt.Sprint(v)
			if i+1 < len(d.list) {
				res += ","
			}
		}
		res += "]"
		return res
	} else {
		return "invalid"
	}
}

func (d data) len() int {
	if d.isList() {
		return len(d.list)
	} else {
		return 1
	}
}

func compareDataList(left, right data) int {
	end := left.len()
	if right.len() < end {
		end = right.len()
	}
	for i := 0; i < end; i++ {
		cmp := left.list[i].compare(
			right.list[i],
		)
		if cmp != 0 {
			return cmp
		}
	}
	if left.len() < right.len() {
		return -1
	} else if left.len() > right.len() {
		return 1
	} else {
		return 0
	}
}

// returns -1, 0, 1
func (left data) compare(right data) int {
	if left.isValue() && right.isValue() {
		if left.value < right.value {
			return -1
		} else if left.value > right.value {
			return 1
		} else {
			return 0
		}
	} else if left.isList() && right.isList() {
		return compareDataList(
			left,
			right,
		)
	} else if left.isValue() {
		tmp := make([]data, 1)
		tmp[0] = left
		return compareDataList(
			makeDataList(tmp),
			right,
		)
	} else if right.isValue() {
		tmp := make([]data, 1)
		tmp[0] = right
		return compareDataList(
			left,
			makeDataList(tmp),
		)
	} else {
		p := fmt.Sprintln("Can not compare")
		p += fmt.Sprintln(left)
		p += fmt.Sprint(right)
		panic(p)
	}
}

var parseDebug bool = false

// expects a string like "1,2,3]..."
// will return the "..." remainder
func parseDataList(str string) (data, string) {
	if parseDebug {
		fmt.Println("parseDataList  ", str)
	}

	res := make([]data, 0)
	// Check for empty List
	if str[0] == ']' {
		return makeDataList(res), str[1:]
	}

	// Check for data in List
	for len(str) > 0 {
		d, s := parseData(str)
		res = append(res, d)
		str = s
		if str[0] == ',' {
			str = str[1:]
			// and continue
		} else if str[0] == ']' {
			return makeDataList(res), str[1:]
		} else {
			panic("Unexpected rune (not ','/']'): " + str)
		}
	}
	panic("Unexpected end")
}

// expecs Data like "1,..." or "2]..."
// returns the ",..." or "]..." part
func parseDataValue(str string) (data, string) {
	if parseDebug {
		fmt.Println("parseDataValue ", str)
	}

	numStr := ""
	ru := (int)(str[0])
	for ru <= (int)('9') && ru >= (int)('0') {
		numStr += string(rune(str[0]))
		str = str[1:]
		ru = (int)(str[0])
	}
	num, err := strconv.Atoi(numStr)
	if err != nil {
		panic(err)
	}
	return makeDataValue(num), str
}

func parseData(str string) (data, string) {
	if parseDebug {
		fmt.Println("parseData      ", str)
	}

	if len(str) == 0 {
		panic("Unexpected end of package")
	}
	if str[0] == '[' {
		return parseDataList(str[1:])
	} else {
		return parseDataValue(str)
	}
}

type dataPair struct {
	d1, d2 data
}

func (pair dataPair) compare() int {
	return pair.d1.compare(pair.d2)
}

func (pair dataPair) String() string {
	return fmt.Sprintln(pair.d1) +
		fmt.Sprintln(pair.d2)
}

type pack []dataPair

func (p pack) String() string {
	res := ""
	for _, v := range p {
		res += fmt.Sprintln(v)
	}
	return res
}

func readData(filename string) pack {
	res := make(pack, 0)
	file, err := os.Open("data/" + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)

	moreData := true
	for moreData {
		var p1 data
		var p2 data
		if scan.Scan() {
			p1, _ = parseData(scan.Text())
		} else {
			panic("Early abort pack 1")
		}
		if scan.Scan() {
			p2, _ = parseData(scan.Text())
		} else {
			panic("Early abort pack 2")
		}
		moreData = scan.Scan()
		res = append(res, dataPair{
			d1: p1,
			d2: p2,
		})
	}
	return res
}
