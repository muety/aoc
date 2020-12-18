package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/muety/aoc2020/util"
)

type SolverDay02 struct {
}

type day02InputElement struct {
	a1 int
	a2 int
	r  rune
	s  string
}

func (s SolverDay02) readData() []day02InputElement {
	lines, err := util.ReadLines("data/input2.txt")
	if err != nil {
		log.Fatalln(err)
	}

	elements := make([]day02InputElement, len(lines))
	reg := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")

	for i, l := range lines {
		match := reg.FindStringSubmatch(l)
		if len(match) != 5 {
			log.Fatalf("failed to parse line '%s'\n", l)
		}

		from, _ := strconv.Atoi(match[1])
		to, _ := strconv.Atoi(match[2])
		elements[i] = day02InputElement{
			a1: from,
			a2: to,
			r:  []rune(match[3])[0],
			s:  match[4],
		}
	}

	return elements
}

func (s SolverDay02) SolveFirst() {
	var count int
	checkValid := func(e day02InputElement) bool {
		count := strings.Count(e.s, string(e.r))
		return count >= e.a1 && count <= e.a2
	}
	for _, e := range s.readData() {
		if checkValid(e) {
			count++
		}
	}
	log.Printf("Solution 2.1: %v\n", count)
}

func (s SolverDay02) SolveSecond() {
	var count int
	checkValid := func(e day02InputElement) bool {
		return (rune(e.s[e.a1-1]) == e.r) != (rune(e.s[e.a2-1]) == e.r) // logical xor
	}
	for _, e := range s.readData() {
		if checkValid(e) {
			count++
		}
	}
	log.Printf("Solution 2.2: %v\n", count)
}
