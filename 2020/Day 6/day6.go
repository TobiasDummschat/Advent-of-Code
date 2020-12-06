package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// contains a bool array of size 26 to store which of the 26 letters is present in the set
type alphabetSet struct {
	letterIsPresent [26]bool
}

func (set *alphabetSet) add(letter int32) {
	set.letterIsPresent[letter-int32("a"[0])] = true
}

func (set *alphabetSet) contains(letter int32) bool {
	return set.letterIsPresent[letter-int32("a"[0])]
}

// contains a slice of alphabetSets representing which questions the members answered "yes" or "no" to
type group struct {
	members []alphabetSet
}

func (g *group) add(member alphabetSet) {
	g.members = append(g.members, member)
}

func main() {
	groups := readInput("2020\\Day 6\\day6_input")
	countAnswers(groups, "anyone")
	countAnswers(groups, "everyone")
}

func countAnswers(groups []group, quantor string) (sum int) {
	if quantor != "anyone" && quantor != "everyone" {
		log.Panicf("Unknown quantor: %s", quantor)
	}

	// loop through all members of all groups to count how many questions anyone/everyone in the group answered "yes"
	var containsLetter bool
	for _, g := range groups {
		alphabet := "abcdefghijklmnopqrstuvwxyz"
		for _, letter := range alphabet {
			// reset containsLetter to empty OR (anyone) or AND (everyone)
			if quantor == "anyone" {
				containsLetter = false
			} else if quantor == "everyone" {
				containsLetter = true
			}

			for _, member := range g.members {
				if quantor == "anyone" {
					containsLetter = containsLetter || member.contains(letter)
				} else if quantor == "everyone" {
					containsLetter = containsLetter && member.contains(letter)
				}
			}
			if containsLetter {
				sum++
			}
		}
	}
	fmt.Printf("The number of questions answered 'yes' by %s in a group summed over all groups is %d.\n", quantor, sum)
	return sum
}

// parse input to slice of groups
func readInput(path string) (groups []group) {
	contents, _ := ioutil.ReadFile(path)
	groupStrings := strings.Split(string(contents), "\r\n\r\n")
	for _, groupString := range groupStrings {
		newGroup := group{}
		memberStrings := strings.Split(groupString, "\r\n")
		for _, memberString := range memberStrings {
			newMember := alphabetSet{}
			for _, letter := range memberString {
				newMember.add(letter)
			}
			newGroup.add(newMember)
		}
		groups = append(groups, newGroup)
	}
	return groups
}
