package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type IntCodeMachine struct {
	Seq []int
	idx int
}

func (m *IntCodeMachine) Run() {
	for !m.Exec() {
	}
}

func (m *IntCodeMachine) Exec() bool {
	maxIdx := len(m.Seq)

	if m.Seq[m.idx] == 99 || m.idx > maxIdx {
		return true
	}

	reg1 := m.Seq[m.idx+1]
	reg2 := m.Seq[m.idx+2]
	regRes := m.Seq[m.idx+3]

	if reg1 > maxIdx || reg2 > maxIdx || regRes > maxIdx {
		return true
	}

	if m.Seq[m.idx] == 1 {
		m.Seq[regRes] = m.Seq[reg1] + m.Seq[reg2]
	} else if m.Seq[m.idx] == 2 {
		m.Seq[regRes] = m.Seq[reg1] * m.Seq[reg2]
	}

	m.idx += 4

	return false
}

func part1(codes []int) {
	m := IntCodeMachine{Seq: codes}
	m.Run()

	fmt.Printf("Solution Day 2A: %v\n", m.Seq[0])
}

func part2(codes []int) {
	target := 19690720
	found := false
	tryCodes := make([]int, len(codes))

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			copy(tryCodes, codes)
			tryCodes[1] = i
			tryCodes[2] = j
			m := IntCodeMachine{Seq: tryCodes}
			m.Run()
			if m.Seq[0] == target {
				found = true
				break
			}
		}

		if found {
			break
		}
	}

	fmt.Printf("Solution Day 2B: %v%v\n", tryCodes[1], tryCodes[2])
}

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

	part1(codes1)
	part2(codes2)
}
