package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getFuel(mass int) int {
	if mass < 2 {
		return 0
	}
	return mass/3 - 2
}

func getFuelRecursive(mass int) int {
	baseFuel := getFuel(mass)
	additionalFuel := getFuel(baseFuel)
	if additionalFuel > 0 {
		return baseFuel + getFuelRecursive(baseFuel)
	}
	return baseFuel
}

func part1(masses []int) {
	var totalFuel int
	for _, m := range masses {
		totalFuel += getFuel(m)
	}
	fmt.Printf("Solution Day 1A: %v\n", totalFuel)
}

func part2(masses []int) {
	var totalFuel int
	for _, m := range masses {
		totalFuel += getFuelRecursive(m)
	}
	fmt.Printf("Solution Day 1B: %v\n", totalFuel)
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputBytes), "\n")

	masses := make([]int, len(lines))
	for i, s := range lines {
		mass, _ := strconv.Atoi(s)
		masses[i] = mass
	}

	part1(masses)
	part2(masses)
}
