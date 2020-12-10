package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	adapters := parseAndSortInput("2020\\Day 10\\day10_input")
	adapterDifferenceDistribution(adapters)
	cache := make(map[int]int)
	arrangements := countArrangements(adapters, 0, 0, &cache)
	fmt.Printf("Found %d possible arrangements.\n", arrangements)
}

// Returns the number of 1, 2, and 3 jolt differences when using all adapters
func adapterDifferenceDistribution(adapters []int) (int, int, int) {
	if adapters[0] < 0 || adapters[0] > 3 {
		fmt.Printf("Smallest adapter with %d jolts not valid to plug into the charging outlet.", adapters[0])
	}

	diff1Count, diff2Count, diff3Count := 0, 0, 1 // 1 for diff to device
	lastAdapter := 0
	for _, adapter := range adapters {
		diff := adapter - lastAdapter
		switch diff {
		case 1:
			diff1Count++
		case 2:
			diff2Count++
		case 3:
			diff3Count++
		default:
			fmt.Printf("No adapters to bridge the jolt gap from %d to %d.\n", lastAdapter, adapter)
		}
		lastAdapter = adapter
	}
	fmt.Printf("Found %d differences of size 1, %d differences of size 2, and %d differences of size 3.\n",
		diff1Count, diff2Count, diff3Count)
	return diff1Count, diff2Count, diff3Count
}

func countArrangements(adapters []int, nextIndex, currentJoltage int, cache *map[int]int) (arrangements int) {
	if nextIndex >= len(adapters)-1 {
		return 1
	}

	cachedResult, cachePresent := (*cache)[nextIndex]
	if cachePresent {
		return cachedResult
	}

	if adapters[nextIndex]-currentJoltage < 0 || adapters[nextIndex]-currentJoltage > 3 {
		log.Panicf("Cannot jump from current joltage %d to smalles adapter %d.\n", currentJoltage, adapters[nextIndex])
	}

	var validIndices []int
	for i := nextIndex; i < len(adapters); i++ {
		adapter := adapters[i]
		if adapter-currentJoltage <= 3 {
			validIndices = append(validIndices, i)
		} else {
			break
		}
	}

	for _, index := range validIndices {
		arrangements += countArrangements(adapters, index+1, adapters[index], cache)
	}

	(*cache)[nextIndex] = arrangements
	return arrangements
}

func parseAndSortInput(path string) (result []int) {
	contents, _ := ioutil.ReadFile(path)
	numStrings := strings.Split(string(contents), "\r\n")
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		result = append(result, num)
	}
	sort.Ints(result)
	return result
}
