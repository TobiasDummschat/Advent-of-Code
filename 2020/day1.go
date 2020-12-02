package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input := readAndSortInput("2020\\day1_input")
	twoNumbers(2020, input)
	threeNumbers(2020, input)
}

func twoNumbers(n int, input []int) {
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
			fmt.Printf("\nFound pair summing to %d in %d tries: %d, %d with product %d",
				n, tries, input[row], input[col], input[row]*input[col])
			return
		case !search && lastTooSmall != nowTooSmall:
			row++
			search = true
		case nowTooSmall:
			col++
			search = false
		case !nowTooSmall:
			col--
			search = false
		}
		lastTooSmall = nowTooSmall
	}
	fmt.Printf("No pair summing to %d found in %d tries.", n, tries)
}

func threeNumbers(n int, input []int) {
	// Same approach as in the two number case, but now with a 3D-table and a plane of color transition.
	// Search one layer similar to two number case, then go down one layer. Keep x < y < z
	// The runtime including sorting for input size m is at most O(m) * T * S where T is the runtime of the two number
	// case and S is the time the sorting takes, so O(m^2 log m). It might be even lower.

	x := 0
	y := 1
	z := 2

	lastTooSmall := true
	search := true

	tries := 0
	for x < y && y < z && z < len(input) {
		for y < z && z < len(input) {
			tries++

			sum := input[x] + input[y] + input[z]
			nowTooSmall := sum < n

			switch {
			case sum == n:
				fmt.Printf("\nFound triple summing to %d in %d tries: %d, %d, %d with product %d",
					n, tries, input[x], input[y], input[z], input[x]*input[y]*input[z])
				return
			case !search && lastTooSmall != nowTooSmall:
				y++
				search = true
			case nowTooSmall:
				z++
				search = false
			case !nowTooSmall:
				z--
				search = false
			}

			lastTooSmall = nowTooSmall
		}
		x++
		y, z = x+1, x+2
		search = true
	}
	fmt.Printf("No triple summing to %d found in %d tries.", n, tries)
}

func readAndSortInput(path string) (result []int) {
	contents, _ := ioutil.ReadFile(path)
	text := string(contents)
	lines := strings.Split(text, "\r\n")
	for _, s := range lines {
		if i, err := strconv.ParseInt(s, 10, 16); err == nil {
			result = append(result, int(i))
		} else {
			fmt.Println(err)
		}
	}
	sort.Ints(result)
	return
}
