package main

import (
	"github.com/muety/aoc2020/util"
	"log"
	"strconv"
)

const preambleLen = 25

type sequence struct {
	numbers []int
	win     int
	pos     int
	max     int
	min     int
}

func newSequence(preamble []int) *sequence {
	seq := &sequence{
		numbers: preamble,
		win:     len(preamble),
		pos:     len(preamble),
	}
	return seq
}

func (s *sequence) next(num int) bool {
	if !s.check(num) {
		return false
	}
	s.numbers = append(s.numbers, num)
	s.pos++
	return true
}

func (s *sequence) check(num int) bool {
	if len(s.numbers) < 2 {
		return true
	}

	for i := s.pos - 1; i < s.pos && i >= 0 && i >= s.pos-s.win; i-- {
		for j := s.pos - 1; j < s.pos && j >= 0 && j >= s.pos-s.win; j-- {
			if s.numbers[i]+s.numbers[j] == num && i != j {
				return true
			}
		}
	}

	return false
}

func readData() []int {
	lines := util.MustReadLines("input.txt")
	numbers := make([]int, len(lines))

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		numbers[i] = num
	}

	return numbers
}

func solveFirst() {
	numbers := readData()
	seq := newSequence(numbers[0:preambleLen])

	for _, num := range numbers[preambleLen : len(numbers)-1] {
		if !seq.next(num) {
			log.Printf("Solution 9.1: %v\n", num)
			return
		}
	}
	log.Println("Solution 9.1 not found")
}

func solveSecond() {
	// TODO
}

func main() {
	solveFirst()
	solveSecond()
}
