package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point struct {
	x, y, z int
}

type Grid map[Point]bool // grid of active points. a map and not a list for easier lookup, if a point is inactive

func cycle(grid Grid, times int) Grid {
	for c := 0; c < times; c++ {
		activeNeighborCount := countActiveNeighbors(grid)
		grid = nextGrid(grid, activeNeighborCount)
		fmt.Printf("After %d cycles the grid contains %d active cubes.\n", c+1, len(grid))
	}
	return grid
}

// go through points with counted neighbors and make the next map
func nextGrid(grid Grid, activeNeighborCount map[Point]int) map[Point]bool {
	newGrid := make(map[Point]bool)
	for point, count := range activeNeighborCount {
		if grid[point] { // active
			// remain active if count is 2 or 3
			if count == 2 || count == 3 {
				newGrid[point] = true
			}
		} else { // inactive
			// activate if count is 3
			if count == 3 {
				newGrid[point] = true
			}
		}
	}
	return newGrid
}

// counts the active neighbors of each point that has active neighbors or is active itself
func countActiveNeighbors(grid Grid) map[Point]int {
	activeNeighborCount := make(map[Point]int)
	for point, active := range grid {
		if active {
			// put active point into neighbor count so that it might get deactivated later
			_, ok := activeNeighborCount[point]
			if !ok {
				activeNeighborCount[point] = 0
			}

			// increment counts for neighbors
			for _, neighbor := range point.getNeighbors() {
				activeNeighborCount[neighbor]++
			}
		}
	}
	return activeNeighborCount
}

func (p Point) getNeighbors() (neighbors []Point) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if !(dx == 0 && dy == 0 && dz == 0) {
					neighbors = append(neighbors, Point{p.x + dx, p.y + dy, p.z + dz})
				}
			}
		}
	}
	return neighbors
}

func main() {
	grid := parseInput("2020\\Day 17\\day17_input")
	cycle(grid, 6)
}

func parseInput(path string) (grid Grid) {
	grid = make(map[Point]bool)
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	z := 0
	for x, line := range lines {
		chars := strings.Split(line, "")
		for y, char := range chars {
			if char == "#" {
				grid[Point{x, y, z}] = true
			} else if char != "." {
				log.Panicf("Unknown character '%s' in line %d at position %d.", char, x, y)
			}
		}
	}
	return grid
}
