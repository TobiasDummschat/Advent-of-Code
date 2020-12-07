package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	graph := readInput("2020\\Day 7\\day7_input")
	countColorsAbleToContains("shiny gold", graph)
	countBagsRequiredInside("shiny gold", graph)
}

func countColorsAbleToContains(name string, graph DirectedGraph) (count int) {
	// breadth search via incoming edges starting with the given vertex
	// color the mainBag and bags able to contain the mainBag green
	// stop searching whenever reaching green
	// count will be number of bags colored green except for the mainBag

	mainBag := graph.GetOrCreateVertex(name)
	green := "green"
	mainBag.color = green

	// edges to go through starting from the mainBag
	queue := mainBag.incoming

	for len(queue) > 0 {
		edge := queue[0]
		queue = queue[1:]
		vertex := edge.tail

		if vertex.color != green {
			count++
			vertex.color = green
			queue = append(queue, vertex.incoming...)
		}
	}

	fmt.Printf("%d bags are able to contain the %s bag trasitively.\n", count, name)
	return count
}

func countBagsRequiredInside(name string, graph DirectedGraph) (count int) {
	mainBag := graph.GetOrCreateVertex(name)

	count = requiredInHelper(mainBag)

	fmt.Printf("%d bags are required inside a %s bag.\n", count, name)
	return count
}

// returns number of bags required in the given bag
func requiredInHelper(bag *Vertex) int {
	if len(bag.outgoing) == 0 {
		return 0
	}

	count := 0
	// go through all the outgoing edges and count the bags and their contents as often as they appear
	for _, edge := range bag.outgoing {
		count += edge.weight * (1 + requiredInHelper(edge.head))
	}
	return count
}

func readInput(path string) DirectedGraph {
	contents, _ := ioutil.ReadFile(path)
	rules := strings.Split(string(contents), "\r\n")

	reOuterBag := regexp.MustCompile("(\\w+(?: \\w+)) bags contain")
	reInnerBag := regexp.MustCompile("(\\d+) (\\w+(?: \\w+)) bags?")

	graph := NewGraph()

	// create all vertices
	for _, rule := range rules {
		outerBagName := reOuterBag.FindStringSubmatch(rule)[1]
		innerMatches := reInnerBag.FindAllStringSubmatch(rule, -1)

		outerBag := graph.GetOrCreateVertex(outerBagName)

		for _, innerMatch := range innerMatches {
			weight, _ := strconv.ParseInt(innerMatch[1], 10, 32)
			innerBagName := innerMatch[2]
			innerBag := graph.GetOrCreateVertex(innerBagName)
			graph.AddEdge(outerBag, innerBag, int(weight))
		}
	}
	return graph
}
