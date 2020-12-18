package main

import (
	"log"
	"os"
)

func main() {
	dayInput := os.Args[1]

	var solver Solver
	if dayInput == "day01" {
		solver = SolverDay01{}
	} else if dayInput == "day02" {
		solver = SolverDay02{}
	} else if dayInput == "day03" {
		solver = SolverDay03{}
	} else {
		log.Fatalln("Missing or unknown day arguments.")
	}

	solver.SolveFirst()
	solver.SolveSecond()
}
