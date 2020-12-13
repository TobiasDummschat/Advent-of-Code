package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Bus struct {
	index, period int
}

func main() {
	myArrival, busses := parseInput("2020\\Day 13\\day13_input")
	nextBus(myArrival, busses)
	challenge(busses)
}

func nextBus(myArrival int, busses []Bus) (Bus, int) {
	// As all busses start at timestamp 0 and their period is their period, we can find the time since the last departure
	// of a Bus by taking the current timestamp modulo the Bus period. Subtracting the result from the period then gives the
	// time until the busses next departure.

	// ^uint(0) is the inverse of uint(0) which is the max value of uint. int has one bit less, so shift once by >> 1
	maxInt := int(^uint(0) >> 1)
	bestBus, bestWait := Bus{}, maxInt

	for _, bus := range busses {
		period := bus.period
		if period <= 0 {
			continue
		}

		sinceLastDeparture := myArrival % period
		wait := period - sinceLastDeparture
		if wait < bestWait {
			bestBus, bestWait = bus, wait
		}
	}

	fmt.Printf("The next Bus to depart after timestamp %d is the Bus with id %d in %d minutes. The product of the Bus period and the necessary waiting time is %d.\n",
		myArrival, bestBus.period, bestWait, bestBus.period*bestWait)
	return bestBus, bestWait
}

func challenge(busses []Bus) int {
	// We want to find the smallest positive time t such that the bus in the timetable at index i departs at time t + i.
	// For a bus with period p this can be restated as the condition (t + i) % p == 0.
	// Equivalently there is some a such that t = a * p - i.
	// Finding the smallest t for one such condition is equivalent to finding the smallest a such that t is positive.
	// As each of the busses gives such a condition we get t = a0 * p0 - i0 = a1 * p1 - i1 = ...
	// Here aX, pX, and iX correspond to the bus at busses[X].
	// We recursively find the smallest valid start for aX and the smallest step it can be incremented by to stay valid:
	//     Lets say that we have four busses indexed from 0 to 3. We start with a3 = 0 and step3 = 1.
	//     The conditions for the two busses 2 and 3 can be combined from t = a2 * p2 - i2 = a3 * p3 - i3 to
	//     a2 * p2 = a3 * p3 - i3 + i2. As the left side is divisible by p2 we try to find the smallest a3 such that
	//     the right side is also divisible by p2. We do this by incrementing a3 by step until that is the case.
	//     Dividing the right by p2 then gives the smallest possible a2. Incrementing a3 by p2 and a2 by p3 gives
	//     another valid solution, so we set step2 = step3 * p2 taking into account the previous step size.
	// The smallest start for a0 yields t using t = a0 * p0 - i0.

	if len(busses) == 0 {
		return 0
	}

	a0, _ := challengeHelper(busses)
	bus := busses[0]
	t := a0*bus.period - bus.index
	fmt.Printf("The smallest timestamp for the challenge is %d.\n", t)
	return t
}

// return smallest possible starting value and increment. See challenge function for more explanation
func challengeHelper(busses []Bus) (int, int) {
	if len(busses) == 1 {
		return 0, 1
	}

	aRight, step := challengeHelper(busses[1:])
	busRight, busLeft := busses[0], busses[1]
	pLeft, iLeft, pRight, iRight := busRight.period, busRight.index, busLeft.period, busLeft.index

	for !(aRight*pRight-iRight >= 0 && (aRight*pRight-iRight+iLeft)%pLeft == 0) {
		aRight += step
	}

	var aLeft = (aRight*pRight - iRight + iLeft) / pLeft
	step = step * pRight
	return aLeft, step
}

// Parses the puzzle input and returns the earliest departure timestamp and an array of the departing busses
func parseInput(path string) (myArrival int, busses []Bus) {
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	myArrival, _ = strconv.Atoi(lines[0])
	re := regexp.MustCompile("\\d+|x")
	matches := re.FindAllString(lines[1], -1)
	for index, match := range matches {
		if match != "x" {
			id, _ := strconv.Atoi(match)
			busses = append(busses, Bus{index: index, period: id})
		}
	}
	return myArrival, busses
}

func gcd(a, b int) int {
	if a == 0 {
		return abs(b)
	}

	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
