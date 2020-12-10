package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	numbers := parseInput("2020\\Day 9\\day9_input")
	success, _, n := findInvalidNumber(numbers, 25)
	if success {
		segment := continuousNumbers(n, numbers)
		min, max := MinMax(segment)
		fmt.Printf("Found segment summing to %d:\n%v\nSum of the smallest and the largest number in the segment is %d.\n",
			n, segment, min+max)
	}
}

// Finds the first number in numbers that is not the sum of any of the preambleLength previous numbers.
// Returns if an invalid number was found, the index of the found number and the number itself
func findInvalidNumber(numbers []int, preambleLength int) (bool, int, int) {
	// concurrently run through all numbers starting with index preambleLength and see, if a pair in their preamble sums
	// to that number.

	invalidIndexCh := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(len(numbers) - preambleLength)
	go closeChannelWhenDone(invalidIndexCh, &wg)

	for i := preambleLength; i < len(numbers); i++ {
		preamble := make([]int, 25)
		copy(preamble, numbers[i-25:i])
		go twoNumbers(numbers[i], preamble, i, invalidIndexCh, &wg)
	}

	var invalidIndices []int
	for index := range invalidIndexCh {
		invalidIndices = append(invalidIndices, index)
	}
	sort.Ints(invalidIndices)

	if len(invalidIndices) == 0 {
		fmt.Printf("No invalid number found.\n")
		return false, -1, -1
	} else {
		firstInvalidIndex := invalidIndices[0]
		firstInvalidNumber := numbers[firstInvalidIndex]
		fmt.Printf("The first invalid number is at index %d with value %d.\n", firstInvalidIndex, firstInvalidNumber)
		return true, firstInvalidIndex, firstInvalidNumber
	}
}

func closeChannelWhenDone(ch chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

// adapted from day 1
func twoNumbers(n int, input []int, index int, ch chan int, wg *sync.WaitGroup) {
	/*
	 * 1. Imagine a table where the rows and columns are indexed by the sorted input values and each cell contains the sum of
	 *    row + column.
	 * 2. Color a cell blue, if its sum is < n, red if it is > n, and green, if it is = n.
	 * 3. As the rows and columns are sorted, the colors in each row row to col and each column top to bottom always
	 *    appear in the same order: blue -> green -> red. Note that not all colors have to be present.
	 * 4. The transition from blue to red or green forms a line starting somewhere in the upper col that only moves
	 *    down and to the row from there.
	 * 5. Any green cell touches this line.
	 * 6. This algorithm uses this by going through the rows in the direction of the line (e.g. left, if cell < n)
	 *    and moving a row down after switching colors.
	 * 7. Given a sorted array of size m, this runs in O(m). Combined with the sorting it runs in O(m log m).
	 */

	defer wg.Done()
	sort.Ints(input)

	row, col := 0, 1
	lastTooSmall := true
	search := true
	tries := 0

	for row < col {
		tries++

		sum := input[row] + input[col]
		nowTooSmall := sum < n

		switch {
		case sum == n:
			return
		case !search && lastTooSmall != nowTooSmall:
			row++
			if !nowTooSmall {
				col--
			}
			search = true
		case nowTooSmall:
			col++
			search = false
		case !nowTooSmall:
			col--
			search = false
		}
		lastTooSmall = nowTooSmall
		if col >= len(input) {
			row++
			col = len(input) - 1
			search = true
		}
	}
	ch <- index
}

func continuousNumbers(n int, input []int) []int {
	/*
	 * 1. We have a left and a right index with left < right.
	 * 2. We look at the sum of all numbers from left to right (left inclusive, right exclusive)
	 * 3. If the sum is too small, move right one index up, if it is too big, move left one index up.
	 * 4. Repeat until desired continuous segment summing to n is found.
	 */

	left, right := 0, 1
	tries := 0

	for left < right && right <= len(input) {
		tries++

		sum := 0
		for i := left; i < right; i++ {
			sum += input[i]
		}

		switch {
		case sum == n:
			return input[left:right]
		case sum < n:
			right++
		case sum > n:
			left++
		}
	}
	fmt.Printf("Couldn't find continuous segment summing to %d.\n", n)
	return nil
}

func parseInput(path string) (result []int) {
	contents, _ := ioutil.ReadFile(path)
	numStrings := strings.Split(string(contents), "\r\n")
	for _, numString := range numStrings {
		n, _ := strconv.ParseInt(numString, 10, 64)
		result = append(result, int(n))
	}
	return result
}

func MinMax(input []int) (int, int) {
	min := input[0]
	max := input[0]
	for _, n := range input {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}
