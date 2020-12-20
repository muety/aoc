package main

import (
	"fmt"
	"github.com/muety/aoc2020/util"
	"log"
	"regexp"
	"strconv"
)

const (
	acc = iota
	jmp
	nop
)

const (
	run = iota
	exitOk
	exitFail
)

type instruction struct {
	cmd  int
	arg1 int
	cc   int // call count
	chk  bool
}

type program struct {
	seq  []instruction
	acc  int
	pc   int // program counter
	cc   int // call counter
	hist []*instruction
}

func (p *program) run() int {
	var code int
	for code = run; code == run; code = p.next() {
	}
	return code
}

func (p *program) runSelfHeal() *program {
	for {
		code := p.next()
		if code == run {
			continue
		}
		if code == exitOk {
			break
		}

		// try heal
		for p.prev() == run {
			if !p.cur().chk && (p.cur().cmd == jmp || p.cur().cmd == nop) {
				p.cur().chk = true

				log.Println("[DEBUG] instantiating new sub program.")
				p2 := p.clone()

				if p.cur().cmd == jmp {
					log.Printf("[DEBUG] swapping jmp for nop at index %d\n", p.pc)
					p2.cur().cmd = nop
				} else if p.cur().cmd == nop {
					log.Printf("[DEBUG] swapping nop for jmp at index %d\n", p.pc)
					p2.cur().cmd = jmp
				}

				log.Println("[DEBUG] running new sub program.")
				if p2.run() == exitOk {
					log.Println("[DEBUG] self-healing successful.")
					return p2
				}
			}
		}
	}
	return nil
}

func (p *program) clone() *program {
	return &program{
		seq: p.seq,
		acc: p.acc,
		pc:  p.pc,
	}
}

func (p *program) next() int {
	if p.pc >= len(p.seq) {
		return exitOk
	}

	if p.detectLoop() {
		log.Println("[DEBUG] loop detected. aborting.")
		return exitFail
	}
	p.record()

	in := p.cur()
	if in.cmd == nop {
		p.pc++
	} else if in.cmd == jmp {
		p.pc += in.arg1
	} else if in.cmd == acc {
		p.acc += in.arg1
		p.pc++
	}
	in.cc++
	p.cc++

	return run
}

func (p *program) prev() int {
	if p.pc <= 0 {
		return exitOk
	}
	in := p.hist[p.cc-1]
	if in.cmd == nop {
		p.pc--
	} else if in.cmd == jmp {
		p.pc -= in.arg1
	} else if in.cmd == acc {
		p.acc -= in.arg1
		p.pc--
	}
	in.cc--
	p.cc--
	return run
}

func (p *program) cur() *instruction {
	return &p.seq[p.pc]
}

func (p *program) record() {
	if p.hist == nil {
		p.hist = make([]*instruction, 0, 0)
	}
	p.hist = append(p.hist, p.cur())
}

func (p *program) detectLoop() bool {
	return p.cur().cc > 0
}

func (p *program) string() string {
	return fmt.Sprintf("Program{seq: [%d], acc: %d, pc: %d}", len(p.seq), p.acc, p.pc)
}

func readData() *program {
	lines := util.MustReadLines("input.txt")
	instructionRe := regexp.MustCompile("(nop|acc|jmp) ((?:\\+|-)\\d+)")

	parse := func(s string) instruction {
		match := instructionRe.FindStringSubmatch(s)
		if len(match) != 3 {
			log.Fatalln("failed to parse instruction")
		}
		var cmd int
		if match[1] == "acc" {
			cmd = acc
		} else if match[1] == "jmp" {
			cmd = jmp
		} else if match[1] == "nop" {
			cmd = nop
		}
		arg1, _ := strconv.Atoi(match[2])
		return instruction{
			cmd:  cmd,
			arg1: arg1,
		}
	}

	instructions := make([]instruction, len(lines))
	for i, l := range lines {
		instructions[i] = parse(l)
	}

	return &program{
		seq: instructions,
	}
}

func solveFirst() {
	p := readData()
	p.run()
	log.Printf("Solution 8.1: %v\n", p.acc)
}

func solveSecond() {
	p := readData().runSelfHeal()
	log.Printf("Solution 8.2: %v\n", p.acc)
}

func main() {
	solveFirst()
	solveSecond()
}
