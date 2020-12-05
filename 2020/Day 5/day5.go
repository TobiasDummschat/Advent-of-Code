package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ids := readInput("2020\\Day 5\\day5_input")
	highestID(ids)
	freeMiddleID(ids)
}

func highestID(ids []int) {
	_, max := minMax(ids)
	fmt.Printf("The highest seat id is %d.\n", max)
}

func freeMiddleID(ids []int) {
	sort.Ints(ids)
	lastID := ids[0]
	for _, id := range ids {
		if id == lastID {
			// skip first iteration
			continue
		}
		if id == lastID+2 {
			// if step size is 2, the id between these is the free one we want
			lastID += 1
			break
		}
		lastID = id
	}
	fmt.Printf("The free middle seat ID is %d.", lastID)
}

func minMax(ints []int) (int, int) {
	min, max := ints[0], ints[0]
	for _, n := range ints {
		if min > n {
			min = n
		}
		if max < n {
			max = n
		}
	}
	return min, max
}

func readInput(path string) (ids []int) {
	contents, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(contents), "\r\n")
	for _, line := range lines {
		// The id is computed in such a way that one can just convert the binary partitioning into a binary number
		line = strings.ReplaceAll(line, "F", "0")
		line = strings.ReplaceAll(line, "B", "1")
		line = strings.ReplaceAll(line, "L", "0")
		line = strings.ReplaceAll(line, "R", "1")
		id, err := strconv.ParseInt(line, 2, 16)
		if err != nil {
			fmt.Println(err)
		} else {
			ids = append(ids, int(id))
		}
	}
	return ids
}
