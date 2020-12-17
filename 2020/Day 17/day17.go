package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Point struct {
	x, y, z, w int
}

type Grid map[Point]bool // grid of active points. a map and not a list for easier lookup, if a point is inactive

func cycle(grid Grid, times, dim int) Grid {
	for c := 0; c < times; c++ {
		activeNeighborCount := countActiveNeighbors(grid, dim)
		grid = nextGrid(grid, activeNeighborCount)
		fmt.Printf("After %d cycles using %d dimensions the grid contains %d active cubes.\n", c+1, dim, len(grid))
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
func countActiveNeighbors(grid Grid, dim int) map[Point]int {
	activeNeighborCount := make(map[Point]int)
	for point, active := range grid {
		if active {
			// put active point into neighbor count so that it might get deactivated later
			_, ok := activeNeighborCount[point]
			if !ok {
				activeNeighborCount[point] = 0
			}

			// increment counts for neighbors
			for _, neighbor := range point.getNeighbors(dim) {
				activeNeighborCount[neighbor]++
			}
		}
	}
	return activeNeighborCount
}

func (p Point) getNeighbors(dim int) (neighbors []Point) {
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if dim == 3 && !(dx == 0 && dy == 0 && dz == 0) {
					neighbors = append(neighbors, Point{x: p.x + dx, y: p.y + dy, z: p.z + dz})
				} else if dim == 4 {
					for dw := -1; dw <= 1; dw++ {
						if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
							neighbors = append(neighbors, Point{x: p.x + dx, y: p.y + dy, z: p.z + dz, w: p.w + dw})
						}
					}
				}
			}
		}
	}
	return neighbors
}

func main() {
	grid := parseInput("2020\\Day 17\\day17_input")
	cycle(grid, 6, 3)
	cycle(grid, 6, 4)
}

func parseInput(path string) (grid Grid) {
	grid = make(map[Point]bool)
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	for x, line := range lines {
		chars := strings.Split(line, "")
		for y, char := range chars {
			if char == "#" {
				grid[Point{x: x, y: y}] = true // z and w default to 0
			} else if char != "." {
				log.Panicf("Unknown character '%s' in line %d at position %d.", char, x, y)
			}
		}
	}
	return grid
}
