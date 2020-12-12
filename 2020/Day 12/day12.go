package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
)

type Instruction struct {
	Action string
	Value  int
}

type Ship struct {
	Position, Direction, Waypoint [2]int
}

func execute(ship *Ship, instructions []Instruction, execute func(*Ship, Instruction)) {
	oldPos := ship.Position
	for _, instruction := range instructions {
		execute(ship, instruction)
	}
	newPos := ship.Position
	distance := distance(oldPos, newPos)
	fmt.Printf("Executed %d instructions and ended up at position %v which is %d units from previous position.\n", len(instructions), ship.Position, distance)
}

func withoutWaypoint(ship *Ship, instruction Instruction) {
	switch instruction.Action {
	case "N":
		ship.Position[1] += instruction.Value
	case "S":
		ship.Position[1] -= instruction.Value
	case "E":
		ship.Position[0] += instruction.Value
	case "W":
		ship.Position[0] -= instruction.Value
	case "L":
		turnLeft(ship, instruction.Value)
	case "R":
		turnLeft(ship, 360-instruction.Value)
	case "F":
		ship.Position[0] += instruction.Value * ship.Direction[0]
		ship.Position[1] += instruction.Value * ship.Direction[1]
	}
}

func withWaypoint(ship *Ship, instruction Instruction) {
	switch instruction.Action {
	case "N":
		ship.Waypoint[1] += instruction.Value
	case "S":
		ship.Waypoint[1] -= instruction.Value
	case "E":
		ship.Waypoint[0] += instruction.Value
	case "W":
		ship.Waypoint[0] -= instruction.Value
	case "L":
		turnLeft(ship, instruction.Value)
	case "R":
		turnLeft(ship, 360-instruction.Value)
	case "F":
		ship.Position[0] += instruction.Value * ship.Waypoint[0]
		ship.Position[1] += instruction.Value * ship.Waypoint[1]
	}
}

// turn both the ships direction and the ships waypoint left
// degrees must be a multiple of 90
func turnLeft(ship *Ship, degrees int) {
	steps := degrees / 90
	for i := 0; i < steps; i++ {
		// turn 90 degrees
		ship.Direction[0], ship.Direction[1] = -ship.Direction[1], ship.Direction[0]
		ship.Waypoint[0], ship.Waypoint[1] = -ship.Waypoint[1], ship.Waypoint[0]
	}
}

func main() {
	instructions := parseInput("2020\\Day 12\\day12_input")
	ship := newShip()
	execute(&ship, instructions, withoutWaypoint)
	ship = newShip()
	execute(&ship, instructions, withWaypoint)
}

func newShip() Ship {
	return Ship{
		Position:  [2]int{0, 0},
		Direction: [2]int{1, 0},
		Waypoint:  [2]int{10, 1},
	}
}

func parseInput(path string) (instructions []Instruction) {
	contents, _ := ioutil.ReadFile(path)
	re := regexp.MustCompile("(N|S|E|W|L|R|F)(\\d+)")
	matches := re.FindAllStringSubmatch(string(contents), -1)
	for _, match := range matches {
		action := match[1]
		value, _ := strconv.Atoi(match[2])
		instructions = append(instructions, Instruction{Action: action, Value: value})
	}
	return instructions
}

// Manhattan distance
func distance(x, y [2]int) int {
	return Abs(x[0]-y[0]) + Abs(x[1]-y[1])
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
