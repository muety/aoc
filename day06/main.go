package main

import (
	"github.com/muety/aoc2020/util"
	"log"
	"strings"
)

func readData() [][]string {
	data := util.MustRead("input.txt")
	groups := strings.Split(data, "\n\n")
	answers := make([][]string, len(groups))
	for i, group := range groups {
		answers[i] = strings.Split(strings.TrimSpace(group), "\n")
	}
	return answers
}

func SolveFirst() {
	var total int
	for _, group := range readData() {
		answers := make(map[rune]bool)
		for _, person := range group {
			for _, answer := range person {
				answers[answer] = true
			}
		}
		total += len(answers)
	}
	log.Printf("Solution 6.1: %v\n", total)
}

func SolveSecond() {
	var total int
	for _, group := range readData() {
		answers := make(map[rune]int)
		for _, person := range group {
			for _, answer := range person {
				answers[answer] += 1
			}
		}
		for _, count := range answers {
			if count == len(group) {
				total++
			}
		}
	}
	log.Printf("Solution 6.2: %v\n", total)
}

func main() {
	SolveFirst()
	SolveSecond()
}
