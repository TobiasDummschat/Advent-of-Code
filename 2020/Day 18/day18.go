package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Value interface {
	evaluate() int
}

type Integer int

type Operation struct {
	left, right Value
	operator    string
}

// in parenthesis
type Block struct {
	value Value
}

func (i Integer) evaluate() int {
	return int(i)
}

func (o Operation) evaluate() int {
	if o.operator == "+" {
		return o.left.evaluate() + o.right.evaluate()
	} else if o.operator == "*" {
		return o.left.evaluate() * o.right.evaluate()
	}
	panic(fmt.Errorf("unknown operation '%s' in '%q %s %q'", o.operator, o.left, o.operator, o.right))
}

func (b Block) evaluate() int {
	return b.value.evaluate()
}

func main() {
	calculations := parseInput("2020\\Day 18\\day18_input", false)
	sumResults(calculations)
	calculations = parseInput("2020\\Day 18\\day18_input", true)
	sumResults(calculations)
}

func sumResults(calculations []Value) (sum int) {
	for _, calculation := range calculations {
		sum += calculation.evaluate()
	}
	fmt.Printf("The results of the calculations sum to %d.\n", sum)
	return sum
}

// parse a list of calculations separated by new lines
func parseInput(path string, plusFirst bool) (calculations []Value) {
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	for _, line := range lines {
		calculation := parseCalculation(line, plusFirst)
		calculations = append(calculations, calculation)
	}
	return calculations
}

// takes in a string calculation and parses it either using the same precedence for '*' and '+' or evaluating '+' first.
func parseCalculation(s string, plusFirst bool) Value {
	s = strings.ReplaceAll(s, " ", "")
	var left, right Value
	left, s = nextValue(s, plusFirst)
	for len(s) != 0 {
		operator := string(s[0])
		right, s = nextValue(s[1:], plusFirst)
		operation, leftIsOperation := left.(Operation)

		// if plusFirst and the current operation is addition and the previous operation was multiplication,
		// then we need replace the previous operations right side with the current addition as it takes precedent
		if plusFirst && operator == "+" && leftIsOperation && operation.operator == "*" {
			operation.right = Operation{
				left:     operation.right,
				right:    right,
				operator: "+",
			}
			left = operation
		} else { // else left-associativity, so just append the current operation to the previous one
			left = Operation{left: left, operator: operator, right: right}
		}
	}
	return left
}

// returns the Integer or Block Value starting at the start of s and the rest of s after that value
func nextValue(s string, plusFirst bool) (value Value, rest string) {
	firstChar := s[0]
	if firstChar >= '0' && firstChar <= '9' { // case: value is Integer
		// find maximally long number and return it
		for i := 1; i <= len(s); i++ {
			if i == len(s) || s[i] < '0' || s[i] > '9' {
				value, err := strconv.Atoi(s[:i])
				if err != nil {
					panic(err)
				}
				return Integer(value), s[i:]
			}
		}
	} else if firstChar == '(' { // case: value is Block
		// count currently open parentheses until the inital '(' is closed
		count, i := 1, 1
		for ; i < len(s) && count > 0; i++ {
			switch s[i] {
			case '(':
				count++
			case ')':
				count--
			}
		}
		if count == 0 {
			// parse inner calculation and return with rest the rest of s
			return Block{parseCalculation(s[1:i-1], plusFirst)}, s[i:]
		} else {
			panic(fmt.Errorf("unmatched parentheses in '%s'", s))
		}
	}
	panic(fmt.Errorf("couldn't find next value in string '%s'", s))
}
