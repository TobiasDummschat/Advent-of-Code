package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	trees := readInput("2020\\Day 3\\day3_input")
	part1(trees)
	part2(trees)
}

func part1(trees [][]bool) {
	fmt.Println("Part 1:")
	countTreesOnSlope(trees, 3, 1)
}

func part2(trees [][]bool) {
	fmt.Println("Part 2:")
	product := countTreesOnSlope(trees, 1, 1)
	product *= countTreesOnSlope(trees, 3, 1)
	product *= countTreesOnSlope(trees, 5, 1)
	product *= countTreesOnSlope(trees, 7, 1)
	product *= countTreesOnSlope(trees, 1, 2)
	fmt.Printf("Product of found trees: %d", product)
}

func countTreesOnSlope(trees [][]bool, right, down int) (count int) {
	i, j := down, right
	period := len(trees[0])
	for i < len(trees) {
		if trees[i][j%period] {
			count++
		}
		i += down
		j += right
	}
	fmt.Printf("Found %d trees on slope right %d, down %d.\n", count, right, down)
	return
}

func readInput(path string) [][]bool {
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")

	// setup trees array
	height, width := len(lines), len(lines[0])
	trees := make([][]bool, height)
	for i := range trees {
		trees[i] = make([]bool, width)
	}

	for i, line := range lines {
		for j := 0; j < width; j++ {
			char := string(line[j])
			trees[i][j] = char == "#"
		}
	}

	return trees
}
