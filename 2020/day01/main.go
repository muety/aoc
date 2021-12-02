package main

import (
	"log"
	"math"
	"strconv"

	"github.com/muety/aoc2020/util"
)

func readData() []int {
	lines := util.MustReadLines("input.txt")
	digits := make([]int, len(lines), len(lines))
	for i, l := range lines {
		d, _ := strconv.Atoi(l)
		digits[i] = d
	}
	return digits
}

func solveFirst() {
	digits := readData()
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

func solveSecond() {
	digits := readData()
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

func main() {
	solveFirst()
	solveSecond()
}
