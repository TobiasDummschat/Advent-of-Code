package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// interface to hold both types of instructions
type Instruction interface {
	isMask() bool
	isMem() bool
}

// The fields of the mask are 1 at the bits where that fields value is present and 0 otherwise
type Mask struct {
	Ones, Xs, Zeros uint64
}

type Mem struct {
	address, value uint64
}

func (_ Mask) isMask() bool { return true }
func (_ Mask) isMem() bool  { return false }
func (_ Mem) isMask() bool  { return false }
func (_ Mem) isMem() bool   { return true }

// apply mask to value for part 1
func (mask Mask) applyToValue(n uint64) uint64 {
	// force zeros by doing bitwise AND with the inverse of the places zeros are
	// force ones by doing bitwise OR with the places where ones are
	return (n & ^mask.Zeros) | mask.Ones
}

// apply mask to address for part 2
func (mask Mask) applyToAddress(n uint64) []uint64 {
	// force ones by doing bitwise OR with the places where ones are
	n = n | mask.Ones

	// replaces Xs with 0s so that we can easily add stuff later to get the possibilities
	n = n & ^(mask.Xs)

	// convert to string for easier processing in helper function
	xString := strconv.FormatUint(mask.Xs, 2)
	return applyToAddressHelper([]uint64{n}, xString)
}

func applyToAddressHelper(addresses []uint64, xString string) []uint64 {
	// recursively find the leftmost 1 in xString and add corresponding addresses
	// as we previously replaced the bits of Xs with 0s, we can just add the right power of 2 to get the other possibility
	forwardsIndex := strings.Index(xString, "1")
	if forwardsIndex == -1 {
		return addresses
	}

	backwardsIndex := len(xString) - forwardsIndex - 1
	for _, address := range addresses {
		addresses = append(addresses, address+uint64(math.Pow(2, float64(backwardsIndex))))
	}
	return applyToAddressHelper(addresses, xString[forwardsIndex+1:])
}

func emptyMask() Mask {
	// X everywhere so nothing is forced
	return Mask{uint64(0), uint64(math.Pow10(36) - 1), uint64(0)}
}

func main() {
	instructions := parseInput("2020\\Day 14\\day14_input")
	runProgram(instructions, make(map[uint64]uint64), false)
	runProgram(instructions, make(map[uint64]uint64), true)
}

func runProgram(instructions []Instruction, memory map[uint64]uint64, version2 bool) {
	mask := emptyMask()
	for _, instruction := range instructions {
		if instruction.isMask() { // update mask
			mask = instruction.(Mask)
		} else if instruction.isMem() { // update memory
			mem := instruction.(Mem)
			if !version2 { // part 1
				memory[mem.address] = mask.applyToValue(mem.value)
			} else { // part 2
				addresses := mask.applyToAddress(mem.address)
				for _, address := range addresses {
					memory[address] = mem.value
				}
			}
		} else {
			log.Panicf("Unknown instruction: %s.", instruction)
		}
	}
	sumValues(memory)
}

func sumValues(memory map[uint64]uint64) uint64 {
	sum := uint64(0)
	for _, value := range memory {
		sum += value
	}
	fmt.Printf("The values in memory sum to %d.\n", sum)
	return sum
}

func parseInput(path string) (instructions []Instruction) {
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	reMask := regexp.MustCompile("[X01]{36}")
	reMem := regexp.MustCompile("mem\\[(\\d+)] = (\\d+)")
	for _, line := range lines {
		var nextInstruction Instruction
		if line[0:4] == "mask" {
			match := reMask.FindString(line)

			// replace values in match to get an uint64 that shows where each of the properties Ones, Xs, Zeros is
			// by which bits of it are 1 and which are 0
			OnesString :=
				strings.ReplaceAll(match, "X", "0")
			XsString :=
				strings.ReplaceAll(
					strings.ReplaceAll(match, "1", "0"), "X", "1")
			ZerosString :=
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(match, "1", "X"), "0", "1"), "X", "0")
			Ones, _ := strconv.ParseInt(OnesString, 2, 64)
			Xs, _ := strconv.ParseInt(XsString, 2, 64)
			Zeros, _ := strconv.ParseInt(ZerosString, 2, 64)

			nextInstruction = Mask{Ones: uint64(Ones), Xs: uint64(Xs), Zeros: uint64(Zeros)}
		} else if line[0:3] == "mem" {
			matches := reMem.FindAllStringSubmatch(line, -1)[0]
			address, _ := strconv.Atoi(matches[1])
			value, _ := strconv.Atoi(matches[2])
			nextInstruction = Mem{address: uint64(address), value: uint64(value)}
		} else {
			log.Panicf("Cannot parse line %s.", line)
		}
		instructions = append(instructions, nextInstruction)
	}
	return instructions
}
