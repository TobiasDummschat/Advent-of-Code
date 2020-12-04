package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr, iyr, eyr, hgt, hcl, ecl, pid, cid string
}

func main() {
	passports := readInput("2020\\Day 4\\day4_input")
	checkFields(passports)
	checkValues(passports)
}

func checkFields(passports []passport) {
	count := 0
	for _, passport := range passports {
		if passport.byr == "" {
			continue
		}
		if passport.iyr == "" {
			continue
		}
		if passport.eyr == "" {
			continue
		}
		if passport.hgt == "" {
			continue
		}
		if passport.hcl == "" {
			continue
		}
		if passport.ecl == "" {
			continue
		}
		if passport.pid == "" {
			continue
		}
		count++
	}
	fmt.Printf("There are %d valid passports by present fields.", count)
}

func checkValues(passports []passport) {
	count := 0
	for _, pp := range passports {
		valid := true
		byr, err := strconv.ParseInt(pp.byr, 10, 32)
		valid = valid && err == nil && 1920 <= byr && byr <= 2002

		iyr, err := strconv.ParseInt(pp.iyr, 10, 32)
		valid = valid && err == nil && 2010 <= iyr && iyr <= 2020

		eyr, err := strconv.ParseInt(pp.eyr, 10, 32)
		valid = valid && err == nil && 2020 <= eyr && eyr <= 2030

		hgtLen := len(pp.hgt)
		if hgtLen > 2 {
			unit := pp.hgt[hgtLen-2 : hgtLen]
			height, err := strconv.ParseInt(pp.hgt[:hgtLen-2], 10, 32)
			valid = valid && err == nil && ((unit == "cm" && 150 <= height && height <= 193) || (unit == "in" && 59 <= height && height <= 76))
		} else {
			valid = false
		}

		if len(pp.hcl) == 7 {
			_, err := strconv.ParseInt(pp.hcl[1:], 16, 32)
			valid = valid && err == nil && string(pp.hcl[0]) == "#"
		} else {
			valid = false
		}

		valid = valid && (pp.ecl == "amb" || pp.ecl == "blu" || pp.ecl == "brn" || pp.ecl == "gry" || pp.ecl == "grn" || pp.ecl == "hzl" || pp.ecl == "oth")

		_, err = strconv.ParseInt(pp.pid, 10, 32)
		valid = valid && len(pp.pid) == 9 && err == nil

		if valid {
			count++
		}
	}
	fmt.Printf("There are %d valid passports by present values.", count)
}

func readInput(path string) []passport {
	contents, _ := ioutil.ReadFile(path)
	rawPassports := strings.Split(string(contents), "\r\n\r\n")
	return parsePassports(rawPassports)
}

func parsePassports(rawPassports []string) (passports []passport) {
	tail := ":([\\w#]+)"
	reByr := regexp.MustCompile("byr" + tail)
	reIyr := regexp.MustCompile("iyr" + tail)
	reEyr := regexp.MustCompile("eyr" + tail)
	reHgt := regexp.MustCompile("hgt" + tail)
	reHcl := regexp.MustCompile("hcl" + tail)
	reEcl := regexp.MustCompile("ecl" + tail)
	rePid := regexp.MustCompile("pid" + tail)
	reCid := regexp.MustCompile("cid" + tail)
	for _, rawPassport := range rawPassports {
		newPassport := passport{}
		matches := reByr.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.byr = matches[1]
		}

		matches = reIyr.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.iyr = matches[1]
		}

		matches = reEyr.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.eyr = matches[1]
		}

		matches = reHgt.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.hgt = matches[1]
		}

		matches = reHcl.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.hcl = matches[1]
		}

		matches = reEcl.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.ecl = matches[1]
		}

		matches = rePid.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.pid = matches[1]
		}

		matches = reCid.FindStringSubmatch(rawPassport)
		if len(matches) >= 1 {
			newPassport.cid = matches[1]
		}
		passports = append(passports, newPassport)
	}
	return passports
}
