package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
)

type instruction struct {
	operation string
	argument  int
}

func main() {
	program := parseInput("2020\\Day 8\\day8_input")
	run(program, true)
	fixProgram(program)
}

// runs the program until loop or termination. returns the global value from that point and if the program terminated
func run(program []instruction, showFailure bool) (int, bool) {
	accumulator, line, count := 0, 0, 0
	hasExecuted := make([]bool, len(program))

	for !hasExecuted[line] {

		hasExecuted[line] = true
		operation := program[line].operation
		argument := program[line].argument

		if operation == "acc" {
			accumulator += argument
			line++
		} else if operation == "jmp" {
			line += argument
		} else if operation == "nop" {
			line++
		} else {
			log.Panicf("Unknown operation: %s", operation)
		}
		count++

		if line == len(program) {
			fmt.Printf("After %d executions the program terminated. The accumulator value was %d.\n",
				count, accumulator)
			return accumulator, true
		}
	}

	if showFailure {
		fmt.Printf("After %d executions the program attempted to execute line %d again. The accumulator value was %d.\n",
			count, line, accumulator)
	}
	return accumulator, false
}

// tries to fix the program by changing one jmp to nop or nop to jmp.
func fixProgram(program []instruction) {
	for line, oldInstruction := range program {
		var newInstruction instruction
		if oldInstruction.operation == "jmp" {
			newInstruction = instruction{
				operation: "nop",
				argument:  oldInstruction.argument,
			}
		} else if oldInstruction.operation == "nop" {
			newInstruction = instruction{
				operation: "jmp",
				argument:  oldInstruction.argument,
			}
		} else {
			continue
		}
		program[line] = newInstruction
		_, terminated := run(program, false)
		if terminated {
			fmt.Printf("The program was fixed by swapping the instruction on line %d.\n", line)
			break
		} else {
			program[line] = oldInstruction
		}

	}
}

func parseInput(path string) (program []instruction) {
	contents, _ := ioutil.ReadFile(path)
	re := regexp.MustCompile("(\\w{3}) ([+-]\\d+)")
	matches := re.FindAllStringSubmatch(string(contents), -1)
	for _, match := range matches {
		operation := match[1]
		argument, err := strconv.ParseInt(match[2], 10, 32)
		if err != nil {
			fmt.Println(err)
		} else {
			instruction := instruction{operation, int(argument)}
			program = append(program, instruction)
		}
	}
	return program
}
