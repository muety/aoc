package main

import (
	"github.com/muety/aoc2020/util"
	"log"
	"math"
)

const (
	nRows    = 128
	nCols    = 8
	nBitsRow = 7
	nBitsCol = 3
)

func readData() []string {
	return util.MustReadLines("input.txt")
}

func parseMask(s string) ([]int, []int) {
	row, col := make([]int, nBitsRow), make([]int, nBitsCol)
	for i := 0; i < nBitsRow; i++ {
		if s[i] == 'F' {
			row[i] = -1
		}
	}
	for i := nBitsRow; i < nBitsRow+nBitsCol; i++ {
		if s[i] == 'L' {
			col[i-nBitsRow] = -1
		}
	}
	return row, col
}

func seatId(row, col int) int {
	return row*8 + col
}

func solveSeatIds(lines []string) []int {
	ids := make([]int, len(lines))

	for k, l := range lines {
		rowMask, colMask := parseMask(l)

		var row int
		for i, m := range rowMask {
			x := int(nRows / math.Pow(2, float64(i+1)))
			row += x + x*m
		}

		var col int
		for i, m := range colMask {
			x := int(nCols / math.Pow(2, float64(i+1)))
			col += x + x*m
		}

		ids[k] = seatId(row, col)
	}

	return ids
}

func SolveFirst() {
	var maxSeatId int
	for _, id := range solveSeatIds(readData()) {
		if id > maxSeatId {
			maxSeatId = id
		}
	}
	log.Printf("Solution 5.1: %v\n", maxSeatId)
}

func SolveSecond() {
	occupied := make(map[int]bool)
	for _, id := range solveSeatIds(readData()) {
		occupied[id] = true
	}

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			id := seatId(i, j)
			if _, ok := occupied[id]; ok {
				continue
			}
			if _, ok1 := occupied[id-1]; ok1 {
				if _, ok2 := occupied[id+1]; ok2 {
					log.Printf("Solution 5.2: %v\n", id)
					return
				}
			}
		}
	}
}

func main() {
	SolveFirst()
	SolveSecond()
}
