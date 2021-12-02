package main

import (
	"github.com/muety/aoc2020/util"
	"log"
)

// true <-> tree
func readData() [][]bool {
	lines := util.MustReadLines("input.txt")
	h, w := len(lines), len(lines[0])
	grid := make([][]bool, h) // col by row <-> 1st = row, 2nd = col
	for i, l := range lines {
		grid[i] = make([]bool, w)
		for j, c := range l {
			grid[i][j] = c == '#'
		}
	}

	return grid
}

// slope: [horizontal, vertical]
func traverse(grid [][]bool, slope []int) (count int) {
	h, w := len(grid), len(grid[0])
	for y, x := 0, 0; y < h; {
		if grid[y][x%w] {
			count++
		}
		y, x = y+slope[1], x+slope[0]
	}
	return count
}

func solveFirst() {
	count := traverse(readData(), []int{3, 1})
	log.Printf("Solution 3.1: %v\n", count)
}

func solveSecond() {
	grid := readData()
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	traverseMp := func(grid [][]bool, slope []int, out chan int) {
		out <- traverse(grid, slope)
	}

	c := make(chan int)
	for _, s := range slopes {
		go traverseMp(grid, s, c)
	}

	n, solution := 0, 1
	for count := range c {
		solution *= count
		n++
		if n == len(slopes) {
			close(c)
		}
	}

	log.Printf("Solution 3.2: %v\n", solution)
}

func main() {
	solveFirst()
	solveSecond()
}
