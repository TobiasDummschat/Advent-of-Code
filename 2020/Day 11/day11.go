package main

import (
	"io/ioutil"
	"strings"
)

type Position string

const (
	Floor    Position = "."
	Empty    Position = "L"
	Occupied Position = "#"
)

type Area [][]Position

func main() {
	parseInput("2020\\Day 11\\day11_input")
}

func doIteration(oldArea Area) Area {
	newArea := make([][]Position, len(oldArea))
	for i, row := range oldArea {
		newArea[i] = make([]Position, len(row))
		for j, pos := range row {
			count := nearbyOccupied(oldArea, i, j)
			if pos == Empty && count == 0 {
				newArea[i][j] = Occupied
			} else if pos == Occupied && count >= 4 {
				newArea[i][j] = Empty
			} else {
				newArea[i][j] = oldArea[i][j]
			}
		}
	}
	return newArea
}

func nearbyOccupied(area Area, row, col int) (count int) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			} else if 0 < i || 0 < j || len(area) <= i || len(area[i]) <= j {
				continue
			} else if area[i][j] == Occupied {
				count++
			}
		}
	}
	return count
}

func parseInput(path string) Area {
	contents, _ := ioutil.ReadFile(path)
	rows := strings.Split(string(contents), "\r\n")
	area := make([][]Position, len(rows))
	for i, row := range rows {
		area[i] = make([]Position, len(row))
		positions := strings.Split(row, "")
		for j, pos := range positions {
			area[i][j] = Position(pos)
		}
	}
	return area
}
