package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type myFile struct {
	name string
	size int
}

func newMyFile(sizeStr string, name string) *myFile {
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		panic(err)
	}
	res := new(myFile)
	res.name = name
	res.size = size
	return res
}

type myDir struct {
	name   string
	dirs   []*myDir
	files  []*myFile
	parent *myDir
}

func newMyDir(name string, parent *myDir) *myDir {
	res := new(myDir)
	res.name = name
	res.dirs = make([]*myDir, 0)
	res.files = make([]*myFile, 0)
	res.parent = parent
	return res
}

func (d *myDir) addEntry(s string) {
	sp := strings.Split(s, " ")
	if len(sp) != 2 {
		panic("Invalid Entry: " + s)
	}
	if sp[0] == "dir" {
		d.dirs = append(
			d.dirs,
			newMyDir(sp[1], d),
		)
	} else {
		d.files = append(
			d.files,
			newMyFile(sp[0], sp[1]),
		)
	}
}

func (d myDir) getDir(s string) *myDir {
	for _, d := range d.dirs {
		if d.name == s {
			return d
		}
	}
	if s == ".." {
		return d.parent
	}
	panic("Dir " + s + " not found in " + d.name)
}

func (md *myDir) getSize() int {
	sum := 0
	for _, d := range md.dirs {
		sum += d.getSize()
	}
	for _, f := range md.files {
		sum += f.size
	}
	return sum
}

func (md *myDir) printOut(prefix string) {
	fmt.Println(prefix, md.name, md.getSize())
	indent := prefix + "    "
	for _, f := range md.files {
		fmt.Println(indent, f.name, f.size)
	}
	for _, d := range md.dirs {
		d.printOut(indent)
	}
}

type processor struct {
	current *myDir
	root    *myDir
	lsCmd   bool
}

const deviceSpace = 70000000
const deviceSpaceForUpdate = 30000000

func (p *processor) procLine(s string) {
	// fmt.Println("Processing ", s, " in ", p.current.name)
	if s[0] == '$' {
		// new command
		p.lsCmd = false
		sp := strings.Split(s, " ")
		if len(sp) <= 1 {
			panic("What should I do with: " + s)
		}

		// Command: cd
		if sp[1] == "cd" && len(sp) == 3 {
			if sp[2] == "/" {
				p.current = p.root
			} else {
				p.current = p.current.getDir(sp[2])
			}
			// Command: ls
		} else if sp[1] == "ls" {
			p.lsCmd = true
		} else {
			panic("Unknown command: " + sp[1])
		}
	} else if p.lsCmd {
		p.current.addEntry(s)
	} else {
		panic("What should I do with " + s)
	}
}

func readInput(filename string) *myDir {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	rootDir := newMyDir("<root>", nil)
	proc := processor{
		current: rootDir,
		root:    rootDir,
		lsCmd:   false,
	}

	for scan.Scan() {
		line := scan.Text()

		proc.procLine(line)
	}
	if err := scan.Err(); err != nil {
		panic(err)
	}

	return rootDir
}

func findDirsRecursive(root *myDir, filter func(*myDir) bool) []*myDir {
	res := make([]*myDir, 0)
	if filter(root) {
		res = append(res, root)
	}
	// no else, because files can count double
	for _, d := range root.dirs {
		res = append(res, findDirsRecursive(d, filter)...)
	}
	return res
}

func filterBelow100000(d *myDir) bool {
	return d.getSize() < 100000
}

func solve_part1(rootDir *myDir) {
	sum := 0
	for _, d := range findDirsRecursive(
		rootDir, filterBelow100000,
	) {
		sum += d.getSize()
	}
	fmt.Println("Sum of dirs with total size of max 100000 ", sum)
}

func solve_part2(rootDir *myDir) {
	freeSpace := deviceSpace - rootDir.getSize()
	spaceNeeded := deviceSpaceForUpdate - freeSpace
	// We are looking for the smalles folder with
	// more size then "spaceNeeded"

	filter := func(d *myDir) bool {
		return d.getSize() >= spaceNeeded
	}

	var delDir *myDir
	var delDirSize int
	for _, d := range findDirsRecursive(
		rootDir, filter,
	) {
		dSize := d.getSize()
		if delDir == nil || delDirSize > dSize {
			delDir = d
			delDirSize = dSize
		}
	}
	fmt.Println("Smallest dir to delete for Update: ",
		delDir.name, " Size: ", delDirSize)
}

func main() {
	rootDir := readInput("input")
	// rootDir.printOut("")

	solve_part1(rootDir)
	solve_part2(rootDir)
}
