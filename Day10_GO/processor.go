package main

type processor struct {
	commands  []command
	currCmd   int // pos in commands
	currCycle int
	allCycles int
	regX      int
}

func makeProcessor(cmds []command) processor {
	return processor{
		commands:  cmds,
		currCmd:   0,
		currCycle: -1, // off by one start error ???
		allCycles: 0,
		regX:      1,
	}
}

func (p processor) signal() int {
	return p.allCycles * p.regX
}

func (p *processor) cycle() {
	if p.currCmd >= len(p.commands) {
		return // we are done
	}

	p.currCycle++
	p.allCycles++
	cmd := &p.commands[p.currCmd]
	if p.currCycle >= cmd.cycles {
		// the > is actually not needed
		p.regX += cmd.deltaX
		p.currCmd++
		p.currCycle = 0
	}
}

func (p *processor) cycles(c int) {
	for i := 0; i < c; i++ {
		p.cycle()
	}
}
