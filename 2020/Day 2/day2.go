package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type entry struct {
	min, max            int
	character, password string
}

func main() {
	input := readInput("2020\\Day 2\\day2_input")
	oldRules(input)
	newRules(input)
}

func oldRules(input []entry) {
	var validEntries []entry
	for _, e := range input {
		re := regexp.MustCompile(e.character)
		count := len(re.FindAll([]byte(e.password), -1))
		if e.min <= count && e.max >= count {
			validEntries = append(validEntries, e)
		}
	}
	fmt.Printf("%d of the %d password entries are valid under the old rules.\n", len(validEntries), len(input))
}

func newRules(input []entry) {
	var validEntries []entry
	for _, e := range input {
		pos1 := []byte(e.password)[e.min-1]
		pos2 := []byte(e.password)[e.max-1]
		char := []byte(e.character)[0]
		if (char == pos1) != (char == pos2) {
			validEntries = append(validEntries, e)
		}
	}
	fmt.Printf("%d of the %d password entries are valid under the old rules.\n", len(validEntries), len(input))
}

func readInput(path string) (result []entry) {
	contents, _ := ioutil.ReadFile(path)
	re := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
	matches := re.FindAllSubmatch(contents, -1)
	for _, match := range matches {
		min, err1 := strconv.ParseInt(string(match[1]), 10, 32)
		max, err2 := strconv.ParseInt(string(match[2]), 10, 32)
		if err1 != nil {
			fmt.Printf("Couldn't parse min value from '%s'. %v\n", match[0], err1)
		} else if err2 != nil {
			fmt.Println("Couldn't parse max value.", err2)
		} else {
			char := string(match[3])
			password := string(match[4])
			newEntry := entry{int(min), int(max), char, password}
			result = append(result, newEntry)
		}
	}
	return
}
