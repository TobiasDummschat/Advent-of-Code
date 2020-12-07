package main

import (
	"io/ioutil"
	"regexp"
	"strings"
)

// directed weighted graph with optional color field for each vertex
type graph struct {
	// TODO
}

func main() {

}

func readInput(path string) {
	contents, _ := ioutil.ReadFile(path)
	rules := strings.Split(string(contents), "\r\n")
	reOuterBag := regexp.MustCompile("(\\w+(?: \\w+)) bags contain")
	reInnerBag := regexp.MustCompile("(\\d+) (\\w+(?: \\w+)) bags?")
	for _, rule := range rules {
		containingBag := reOuterBag.FindStringSubmatch(rule)
		containedBags := reInnerBag.FindAllStringSubmatch(rule, -1)
		// TODO continue
	}
}
