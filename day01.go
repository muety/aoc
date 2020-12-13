package main

import (
	"log"
	"math"
	"strconv"

	"github.com/muety/aoc2020/util"
)

type SolverDay01 struct {
}

func (s SolverDay01) readData() []int {
	lines, err := util.ReadLines("data/input1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	digits := make([]int, len(lines), len(lines))
	for i, l := range lines {
		d, _ := strconv.Atoi(l)
		digits[i] = d
	}
	return digits
}

func (s SolverDay01) SolveFirst() {
	digits := s.readData()
	total := len(digits)

	// Brute force approach with O(n²)
	for i, num1 := range digits {
		for _, num2 := range digits[i+1 : total-1] {
			if num1+num2 == 2020 {
				log.Printf("Solution 1.1: %v\n", num1*num2)
				return
			}
		}
	}

	log.Println("No solution found for 1.1")
}

func (s SolverDay01) SolveSecond() {
	digits := s.readData()
	total := len(digits)

	// Even worse brute force with O(n³)
	for i, num1 := range digits {
		for j, num2 := range digits[i+1 : total-1] {
			if num1+num2 > 2020 {
				continue
			}

			for _, num3 := range digits[int(math.Max(float64(i), float64(j)))+1 : total-1] {
				if num1+num2+num3 == 2020 {
					log.Printf("Solution 1.2: %v\n", num1*num2*num3)
					return
				}
			}
		}
	}

	log.Println("No solution found for 1.2")
}
