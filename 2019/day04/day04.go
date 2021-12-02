package main

import (
	"fmt"
	"strconv"
)

const (
	boundLower = 137683
	boundUpper = 596253
)

func splitNumber(n int) (string, []int) {
	strDigits := strconv.Itoa(n)
	intDigits := make([]int, len(strDigits))

	for i := 0; i < len(strDigits); i++ {
		dint, _ := strconv.Atoi(string(strDigits[i]))
		intDigits[i] = dint
	}

	return strDigits, intDigits
}

func tryRange(from, to int, valid func(int) bool) int {
	var counter int
	for i := from; i <= to; i++ {
		if valid(i) {
			counter++
		}
	}
	return counter
}

func part1() {
	valid := func(n int) bool {
		strDigits, intDigits := splitNumber(n)
		increasing := true
		double := false

		for i := 1; i < len(strDigits); i++ {
			if intDigits[i-1] > intDigits[i] {
				increasing = false
				break
			}
			if intDigits[i-1] == intDigits[i] {
				double = true
			}
		}

		return increasing && double
	}

	c := tryRange(boundLower, boundUpper, valid)
	fmt.Printf("Solution Day 4A: %v\n", c)
}

func part2() {
	valid := func(n int) bool {
		strDigits, intDigits := splitNumber(n)
		k := len(strDigits)
		increasing := true
		double := false
		subsecMatches := 0

		for i := 1; i < k; i++ {
			if intDigits[i-1] > intDigits[i] {
				increasing = false
				break
			}
			if intDigits[i-1] == intDigits[i] {
				subsecMatches++
				if i == k-1 && subsecMatches == 1 {
					double = true
				}
			} else {
				if subsecMatches == 1 {
					double = true
				}
				subsecMatches = 0
			}
		}

		return increasing && double
	}

	c := tryRange(boundLower, boundUpper, valid)
	fmt.Printf("Solution Day 4B: %v\n", c)
}

func main() {
	part1()
	part2()
}
