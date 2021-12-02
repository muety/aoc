package main

/*
 Representing the orbited objects as a directed, asynclic graph
 and using known graph traversal methods on that probably would have
 been easier ;-)
*/

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func buildOrbitMap(orbitDefs []string) map[string]*list.List {
	orbitMap := make(map[string]*list.List)

	for _, def := range orbitDefs {
		objects := strings.Split(def, ")")
		obj1 := objects[0]
		obj2 := objects[1]

		if _, ok := orbitMap[obj1]; !ok {
			orbitMap[obj1] = list.New()
		}
		if _, ok := orbitMap[obj2]; !ok {
			orbitMap[obj2] = list.New()
		}

		orbitMap[obj2].PushBack(obj1)
	}

	return orbitMap
}

func countObject(obj string, orbitMap *map[string]*list.List) int {
	l, _ := (*orbitMap)[obj]
	c := l.Len()
	if c > 0 {
		elem := l.Front()
		for elem != nil {
			c += countObject(elem.Value.(string), orbitMap)
			elem = elem.Next()
		}
	}
	return c
}

func countTotal(orbitMap *map[string]*list.List) int {
	var c int

	for k := range *orbitMap {
		c += countObject(k, orbitMap)
	}
	return c
}

func countShortestPath(from, to string, orbitMap *map[string]*list.List, checked *map[string]bool) int {
	if checked == nil {
		tmpChecked := make(map[string]bool)
		checked = &tmpChecked
	}

	(*checked)[from+to] = true

	if from == to {
		return 0
	}

	l1, ok := (*orbitMap)[from]
	if !ok || l1.Len() == 0 {
		return -1
	}
	l2, ok := (*orbitMap)[to]
	if !ok || l2.Len() == 0 {
		return -1
	}

	if l1.Len() == 0 || l2.Len() == 0 {
		return -1
	}

	if contains(to, l1) || contains(from, l2) {
		return 1
	}

	minLength := math.MaxInt32

	elem := l1.Front()
	for elem != nil {
		val := elem.Value.(string)
		if _, ok := (*checked)[val+to]; !ok {
			length := countShortestPath(val, to, orbitMap, checked)
			if length > -1 && length < minLength {
				minLength = length + 1
			}
		}

		elem = elem.Next()
	}

	elem = l2.Front()
	for elem != nil {
		val := elem.Value.(string)
		if _, ok := (*checked)[from+val]; !ok {
			length := countShortestPath(from, val, orbitMap, checked)
			if length > -1 && length < minLength {
				minLength = length + 1
			}
		}
		elem = elem.Next()
	}

	if minLength == math.MaxInt32 {
		return -1
	}

	return minLength
}

func contains(needle string, haystack *list.List) bool {
	if haystack.Len() == 0 {
		return false
	}

	elem := haystack.Front()
	for elem != nil {
		if elem.Value.(string) == needle {
			return true
		}
		elem = elem.Next()
	}

	return false
}

func part1(lines []string) {
	orbitMap := buildOrbitMap(lines)
	count := countTotal(&orbitMap)
	fmt.Printf("Solution Day 6A: %v\n", count)
}

func part2(lines []string) {
	orbitMap := buildOrbitMap(lines)

	you := orbitMap["YOU"].Front().Value.(string)
	san := orbitMap["SAN"].Front().Value.(string)

	count := countShortestPath(you, san, &orbitMap, nil)
	fmt.Printf("Solution Day 6B: %v\n", count)
}

func main() {
	inputBytes, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(inputBytes), "\n")

	part1(lines)
	part2(lines)
}
