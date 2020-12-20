package main

import (
	"github.com/muety/aoc2020/util"
	"log"
	"regexp"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *passport) Valid() bool {
	return p.ValidIgnoringCountry() && p.cid != ""
}

func (p *passport) ValidIgnoringCountry() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func readData(strict bool) []passport {
	data := util.MustRead("input.txt")
	reNormal := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile("(?m)^byr:([^\\s]+)$"),
		"iyr": regexp.MustCompile("(?m)^iyr:([^\\s]+)$"),
		"eyr": regexp.MustCompile("(?m)^eyr:([^\\s]+)$"),
		"hgt": regexp.MustCompile("(?m)^hgt:([^\\s]+)$"),
		"hcl": regexp.MustCompile("(?m)^hcl:([^\\s]+)$"),
		"ecl": regexp.MustCompile("(?m)^ecl:([^\\s]+)$"),
		"pid": regexp.MustCompile("(?m)^pid:([^\\s]+)$"),
		"cid": regexp.MustCompile("(?m)^cid:([^\\s]+)$"),
	}

	// https://www.richie-bendall.ml/ros-regex-numeric-range-generator/
	reStrict := map[string]*regexp.Regexp{
		"byr": regexp.MustCompile("(?m)^byr:(192[0-9]|19[3-9][0-9]|200[0-2])$"),
		"iyr": regexp.MustCompile("(?m)^iyr:(201[0-9]|2020)$"),
		"eyr": regexp.MustCompile("(?m)^eyr:(202[0-9]|2030)$"),
		"hgt": regexp.MustCompile("(?m)^hgt:((?:(?:15[0-9]|1[6-8][0-9]|19[0-3])cm)|(?:(?:59|6[0-9]|7[0-6])in))$"),
		"hcl": regexp.MustCompile("(?m)^hcl:(#[a-f0-9]{6})$"),
		"ecl": regexp.MustCompile("(?m)^ecl:((?:amb|blu|brn|gry|grn|hzl|oth))$"),
		"pid": regexp.MustCompile("(?m)^pid:([0-9]{9})$"),
		"cid": regexp.MustCompile("(?m)^cid:([^\\s]+)$"),
	}

	re := reNormal
	if strict {
		re = reStrict
	}

	matchOrEmpty := func(s string, re *regexp.Regexp) string {
		match := re.FindStringSubmatch(s)
		if len(match) == 2 {
			return match[1]
		}
		return ""
	}

	rawPassports := strings.Split(data, "\n\n")
	passports := make([]passport, len(rawPassports))

	for i, d := range rawPassports {
		d = strings.ReplaceAll(d, " ", "\n")
		passports[i] = passport{
			byr: matchOrEmpty(d, re["byr"]),
			iyr: matchOrEmpty(d, re["iyr"]),
			eyr: matchOrEmpty(d, re["eyr"]),
			hgt: matchOrEmpty(d, re["hgt"]),
			hcl: matchOrEmpty(d, re["hcl"]),
			ecl: matchOrEmpty(d, re["ecl"]),
			pid: matchOrEmpty(d, re["pid"]),
			cid: matchOrEmpty(d, re["cid"]),
		}
	}

	return passports
}

func solveFirst() {
	var count int
	for _, p := range readData(false) {
		if p.ValidIgnoringCountry() {
			count++
		}
	}
	log.Printf("Solution 4.1: %v\n", count)
}

func solveSecond() {
	var count int
	for _, p := range readData(true) {
		if p.ValidIgnoringCountry() {
			count++
		}
	}
	log.Printf("Solution 4.2: %v\n", count)
}

func main() {
	solveFirst()
	solveSecond()
}
