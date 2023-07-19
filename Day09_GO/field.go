package main

type pos struct {
	x, y int
}

var mapPosDir = map[direction]pos{
	left:  {-1, 0},
	right: {1, 0},
	up:    {0, 1},
	down:  {0, -1},
}

func (p1 pos) diff(p2 pos) pos {
	return pos{
		x: p1.x - p2.x,
		y: p1.y - p2.y,
	}
}

func (p1 pos) equal(p2 pos) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func (p1 *pos) moveBy(p2 pos) {
	p1.x += p2.x
	p1.y += p2.y
}

func (p1 *pos) follow(p2 pos) {
	diff := p2.diff(*p1)
	if intAbs(diff.x) <= 1 &&
		intAbs(diff.y) <= 1 {
		return
	}

	// we need to move the tail
	move := diff.clamp()
	p1.moveBy(move)
}

// allow max distance of 1
func (p pos) clamp() pos {
	return pos{
		x: intClamp(p.x, -1, 1),
		y: intClamp(p.y, -1, 1),
	}
}

// start is implicit 0, 0
type field struct {
	head       pos
	tail       pos
	tailResult []pos
	knots      []pos // points inbetween
}

func makeField() field {
	tailResult := make([]pos, 1)
	// tailResult[0] = pos{0, 0}
	return field{
		// head:       pos{0, 0},
		// tail:       pos{0, 0},
		tailResult: tailResult,
		knots:      make([]pos, 0),
	}
}

func (f *field) addKnots(num int) {
	f.knots = append(f.knots, make([]pos, num)...)
}

func intAbs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}

func intClamp(i, min, max int) int {
	if i <= min {
		return min
	} else if i >= max {
		return max
	} else {
		return i
	}
}

func (f *field) moveHeadOneStep(move pos) {
	f.head.moveBy(move)
	var lastKnot *pos = &f.head
	for i := range f.knots {
		f.knots[i].follow(*lastKnot)
		lastKnot = &f.knots[i]
	}
	f.tail.follow(*lastKnot)
	f.addTailResult(f.tail)
	// fmt.Println("Head:", f.head, "Tail:", f.tail)
}

func (f *field) moveHeadByCmd(cmd cmd) {
	move, avail := mapPosDir[cmd.dir]
	if !avail {
		panic("cmd.dir not in pos map")
	}

	for i := 0; i < cmd.steps; i++ {
		f.moveHeadOneStep(move)
	}
}

func (f *field) addTailResult(p pos) {
	for _, p2 := range f.tailResult {
		if p.equal(p2) {
			return
		}
	}
	f.tailResult = append(f.tailResult, p)
}
