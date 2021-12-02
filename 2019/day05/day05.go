package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const MaxValueLength = 4

type IntCodeMachine struct {
	Seq    []int
	Input  func() int
	Output func(int)
	idx    int
}

func (m *IntCodeMachine) Run() {
	for !m.Exec() {
	}
}

func (m *IntCodeMachine) Exec() bool {
	var noUpdate bool
	maxIdx := len(m.Seq)

	opcode, modes := m.parseHeader()
	if opcode == 99 || m.idx > maxIdx {
		return true
	}

	il := m.il(opcode)
	params := m.readParams(il)
	vals, err := m.resolveParams(params, modes)
	if err != nil {
		return true
	}

	if opcode == 1 {
		m.Seq[params[2]] = vals[0] + vals[1]
	} else if opcode == 2 {
		m.Seq[params[2]] = vals[0] * vals[1]
	} else if opcode == 3 {
		m.Seq[params[0]] = m.Input()
	} else if opcode == 4 {
		m.Output(vals[0])
	} else if opcode == 5 {
		if vals[0] != 0 {
			m.idx = vals[1]
			noUpdate = true
		}
	} else if opcode == 6 {
		if vals[0] == 0 {
			m.idx = vals[1]
			noUpdate = true
		}
	} else if opcode == 7 {
		if vals[0] < vals[1] {
			m.Seq[params[2]] = 1
		} else {
			m.Seq[params[2]] = 0
		}
	} else if opcode == 8 {
		if vals[0] == vals[1] {
			m.Seq[params[2]] = 1
		} else {
			m.Seq[params[2]] = 0
		}
	} else {
		panic("Unknown opcode")
	}

	if !noUpdate {
		m.idx += (il + 1)
	}

	return false
}

func (m *IntCodeMachine) il(opcode int) int {
	if opcode == 1 {
		return 3
	} else if opcode == 2 {
		return 3
	} else if opcode == 3 {
		return 1
	} else if opcode == 4 {
		return 1
	} else if opcode == 5 {
		return 2
	} else if opcode == 6 {
		return 2
	} else if opcode == 7 {
		return 3
	} else if opcode == 8 {
		return 3
	}
	return 0
}

func (m *IntCodeMachine) readParams(n int) []int {
	params := make([]int, n)
	for i := 0; i < len(params); i++ {
		params[i] = m.Seq[m.idx+i+1]
	}
	return params
}

func (m *IntCodeMachine) resolveParams(params, modes []int) ([]int, error) {
	resolved := make([]int, len(params))
	for i := 0; i < len(params); i++ {
		if modes[i] == 1 {
			resolved[i] = params[i]
		} else {
			if params[i] > len(m.Seq) {
				return nil, errors.New("Could not read param value")
			}
			resolved[i] = m.Seq[params[i]]
		}
	}
	return resolved, nil
}

func (m *IntCodeMachine) parseHeader() (int, []int) {
	digits := m.splitDigits(m.Seq[m.idx])
	n := len(digits)
	opcode := digits[n-2]*10 + digits[n-1]
	il := m.il(opcode)
	modes := make([]int, il)

	for i := 0; i < len(modes) && i <= n-3; i++ {
		modes[i] = digits[n-i-3]
	}

	return opcode, modes
}

func (m *IntCodeMachine) splitDigits(n int) []int {
	digits := make([]int, MaxValueLength)
	for i := MaxValueLength - 1; i >= 0; i-- {
		digits[i] = (n / int(math.Pow(10, float64(MaxValueLength-1-i)))) % 10
	}
	return digits
}

func part1(codes []int) {
	inputResolver := func() int {
		return 1
	}
	outputResolver := func(data int) {
		fmt.Printf("%v ", data)
	}

	m := IntCodeMachine{Seq: codes, Input: inputResolver, Output: outputResolver}
	m.Run()
}

func part2(codes []int) {
	inputResolver := func() int {
		return 5
	}
	outputResolver := func(data int) {
		fmt.Printf("%v ", data)
	}

	m := IntCodeMachine{Seq: codes, Input: inputResolver, Output: outputResolver}
	m.Run()
}

func readPuzzleInput(filename string) {}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	characters := strings.Split(string(inputBytes), ",")

	codes := make([]int, len(characters))
	for i, s := range characters {
		mass, _ := strconv.Atoi(s)
		codes[i] = mass
	}

	codes1 := make([]int, len(codes))
	codes2 := make([]int, len(codes))

	copy(codes1, codes)
	copy(codes2, codes)

	fmt.Printf("Part 1: ")
	part1(codes1)
	fmt.Println()

	fmt.Printf("Part 2: ")
	part2(codes2)
	fmt.Println()
}
