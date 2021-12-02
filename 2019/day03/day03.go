package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

const (
	left  = iota
	right = iota
	up    = iota
	down  = iota
)

type Waypoint struct {
	X    int
	Y    int
	Cost int
}

func (c *Waypoint) Manhattan(other *Waypoint) int {
	return int(math.Abs(float64(c.X)-float64(other.X)) + math.Abs(float64(c.Y)-float64(other.Y)))
}

func (c *Waypoint) String() string {
	return strconv.Itoa(c.X) + "," + strconv.Itoa(c.Y)
}

type Path struct {
	Commands  []WalkCommand
	Waypoints map[string]int
	hDiff     int
	vDiff     int
	idx       int
	count     int
}

func (p *Path) StepAndCheckIntersectWith(otherPath *Path) []*Waypoint {
	intersections := make([]*Waypoint, 0)

	if p.idx == len(p.Commands) {
		return intersections
	}

	var hFactor, vFactor int

	if p.Commands[p.idx].Dir == left {
		hFactor = -1
	} else if p.Commands[p.idx].Dir == right {
		hFactor = 1
	} else if p.Commands[p.idx].Dir == up {
		vFactor = -1
	} else if p.Commands[p.idx].Dir == down {
		vFactor = 1
	}

	for i := 0; i < p.Commands[p.idx].N; i++ {
		p.hDiff += hFactor
		p.vDiff += vFactor
		p.count++
		wp := &Waypoint{X: p.hDiff, Y: p.vDiff, Cost: p.count}
		p.Waypoints[wp.String()] = p.count

		if otherPath != nil && otherPath.CheckWaypoint(wp) {
			costOther := otherPath.GetWaypointCost(wp)
			intersections = append(intersections, &Waypoint{wp.X, wp.Y, p.count + costOther})
		}
	}

	p.idx++

	return intersections
}

func (p *Path) CheckWaypoint(needle *Waypoint) bool {
	if _, ok := p.Waypoints[needle.String()]; ok {
		return true
	}
	return false
}

func (p *Path) GetWaypointCost(needle *Waypoint) int {
	if cost, ok := p.Waypoints[needle.String()]; ok {
		return cost
	}
	return -1
}

func (p *Path) HasNext() bool {
	return p.idx < len(p.Commands)
}

func NewPath(commands []WalkCommand) *Path {
	return &Path{
		Commands:  commands,
		Waypoints: make(map[string]int),
	}
}

type WalkCommand struct {
	Dir int
	N   int
}

func parseWalkCommand(str string) WalkCommand {
	var dir int
	var n int

	if str[0] == []byte("L")[0] {
		dir = left
	} else if str[0] == []byte("R")[0] {
		dir = right
	} else if str[0] == []byte("U")[0] {
		dir = up
	} else if str[0] == []byte("D")[0] {
		dir = down
	}

	n, _ = strconv.Atoi(str[1:])

	return WalkCommand{Dir: dir, N: n}
}

func part1(path1, path2 *Path) {
	var origin, intersection Waypoint
	intersection.X = math.MaxInt32

	for path1.HasNext() {
		path1.StepAndCheckIntersectWith(nil)
	}

	for path2.HasNext() {
		intersections := path2.StepAndCheckIntersectWith(path1)

		for _, is := range intersections {
			if origin.Manhattan(is) < origin.Manhattan(&intersection) {
				intersection = *is
			}
		}
	}

	fmt.Printf("Solution Day 3A: %v\n", origin.Manhattan(&intersection))
}

func part2(path1, path2 *Path) {
	var intersection Waypoint
	intersection.Cost = math.MaxInt32

	for path1.HasNext() {
		path1.StepAndCheckIntersectWith(nil)
	}

	for path2.HasNext() {
		intersections := path2.StepAndCheckIntersectWith(path1)

		for _, is := range intersections {
			if is.Cost < intersection.Cost {
				intersection = *is
			}
		}
	}

	fmt.Printf("Solution Day 3B: %v\n", intersection.Cost)
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputBytes), "\n")

	strCommands1 := strings.Split(lines[0], ",")
	strCommands2 := strings.Split(lines[1], ",")

	commands1 := make([]WalkCommand, len(strCommands1))
	commands2 := make([]WalkCommand, len(strCommands2))

	for i, str := range strCommands1 {
		commands1[i] = parseWalkCommand(str)
	}
	for i, str := range strCommands2 {
		commands2[i] = parseWalkCommand(str)
	}

	part1(NewPath(commands1), NewPath(commands2))
	part2(NewPath(commands1), NewPath(commands2))
}
