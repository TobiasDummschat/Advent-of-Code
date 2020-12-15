package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	startingNumbers := parseInput("2020\\Day 15\\day15_input")
	play(2020, startingNumbers)
	play(30000000, startingNumbers)
}

func play(upToTurn int, startingNumbers []int) (spokenNumbers []int, memory map[int][]int) {
	memory = make(map[int][]int)
	previous := startingNumbers[0]
	// process starting numbers
	for i, number := range startingNumbers {
		turn := i + 1
		previous = number
		say(number, turn, &spokenNumbers, &memory)
	}

	for turn := len(spokenNumbers) + 1; turn <= upToTurn; turn++ {
		next := nextNumber(previous, memory)
		say(next, turn, &spokenNumbers, &memory)
		previous = next
	}

	fmt.Printf("The spoken number at turn %d is %d.\n", upToTurn, previous)
	return spokenNumbers, memory
}

func say(number, turn int, spokenNumbers *[]int, memory *map[int][]int) {
	_, ok := (*memory)[number]
	if !ok {
		(*memory)[number] = []int{turn}
	} else {
		(*memory)[number] = append((*memory)[number], turn)
	}
	*spokenNumbers = append(*spokenNumbers, number)
}

func nextNumber(previous int, memory map[int][]int) int {
	turnsSaid, ok := memory[previous]
	if !ok {
		return 0
	} else {
		timesSaid := len(turnsSaid)
		if timesSaid == 0 || timesSaid == 1 {
			return 0
		} else {
			return turnsSaid[timesSaid-1] - turnsSaid[timesSaid-2]
		}

	}
}

func parseInput(path string) (startingNumbers []int) {
	contents, _ := ioutil.ReadFile(path)
	numStrings := strings.Split(string(contents), ",")
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		startingNumbers = append(startingNumbers, num)
	}
	return startingNumbers
}
