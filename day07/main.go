package main

import (
	"github.com/muety/aoc2020/util"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type rule struct {
	subject string
	objects []string
	counts  []int
}

func (r *rule) sum() (total int) {
	for _, c := range r.counts {
		total += c
	}
	return total
}

func readData() []rule {
	lines := util.MustReadLines("input.txt")

	outerRe := regexp.MustCompile("(\\w+ \\w+) bags contain (.+\\.)")
	innerRe := regexp.MustCompile("(\\d)+ (\\w+ \\w+) bags?")

	rules := make([]rule, len(lines))

	for i, l := range lines {
		match1 := outerRe.FindStringSubmatch(l)

		r := rule{
			subject: match1[1],
			objects: make([]string, 0),
			counts:  make([]int, 0),
		}

		for _, o := range strings.Split(match1[2], ", ") {
			match2 := innerRe.FindStringSubmatch(o)

			if len(match2) != 3 {
				continue
			}

			c, err := strconv.Atoi(match2[1])
			if err != nil {
				log.Fatalln(err)
			}

			r.objects = append(r.objects, match2[2])
			r.counts = append(r.counts, c)
		}

		rules[i] = r
	}

	return rules
}

// bag -> all rules, which the bag requires
func mapRules(rules []rule) map[string]rule {
	m := make(map[string]rule)
	for _, r := range rules {
		m[r.subject] = r
	}
	return m
}

// bag -> all rules, which the bag is contained in
func mapRulesInverse(rules []rule) map[string][]rule {
	m := make(map[string][]rule)
	for _, r := range rules {
		for _, o := range r.objects {
			if _, ok := m[o]; !ok {
				m[o] = make([]rule, 0)
			}
			m[o] = append(m[o], r)
		}
	}
	return m
}

// recursively get all bags that lead to the target bag
func resolveOptions(m map[string][]rule, target string) []string {
	lookup := make(map[string]bool)
	rules, ok := m[target]
	if !ok {
		return []string{}
	}

	for _, r := range rules {
		lookup[r.subject] = true
		for _, rr := range resolveOptions(m, r.subject) {
			lookup[rr] = true
		}
	}

	uniqueBags := make([]string, 0, len(lookup))
	for k := range lookup {
		uniqueBags = append(uniqueBags, k)
	}
	return uniqueBags
}

// recursively count the number of bags required for the target bag
func resolveCount(m map[string]rule, target string) (count int) {
	r := m[target]
	for i, o := range r.objects {
		count += r.counts[i] * resolveCount(m, o)
	}
	return count + 1
}

func SolveFirst() {
	rules := resolveOptions(mapRulesInverse(readData()), "shiny gold")
	log.Printf("Solution 7.1: %v\n", len(rules))
}

func SolveSecond() {
	count := resolveCount(mapRules(readData()), "shiny gold")
	log.Printf("Solution 7.2: %v\n", count-1)
}

func main() {
	SolveFirst()
	SolveSecond()
}
