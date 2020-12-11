package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Position string

const (
	Floor        Position = "."
	EmptySeat    Position = "L"
	OccupiedSeat Position = "#"
)

type Area [][]Position

func (thisArea Area) equals(area Area) bool {
	if len(thisArea) != len(area) {
		return false
	}
	for i := 0; i < len(thisArea); i++ {
		if len(thisArea[i]) != len(area[i]) {
			return false
		}
		for j := 0; j < len(thisArea[i]); j++ {
			if thisArea[i][j] != area[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	area := parseInput("2020\\Day 11\\day11_input")
	part1Areas := findEquilibrium(area, adjacentOccupiedCounter, 4)
	part2Areas := findEquilibrium(area, visibleOccupiedCounter, 5)

	writeGifToFile(areasToGIF(part1Areas), "2020\\Day 11\\part1.gif")
	writeGifToFile(areasToGIF(part2Areas), "2020\\Day 11\\part2.gif")
}

func findEquilibrium(area Area, nearbyOccupiedCounter func(Area, int, int) int, tooManyPeople int) (areas []Area) {
	iterations := 0
	areas = append(areas, area)

	oldArea, newArea := area, iterate(area, nearbyOccupiedCounter, tooManyPeople)
	for !oldArea.equals(newArea) {
		iterations++
		areas = append(areas, newArea)
		oldArea, newArea = newArea, iterate(newArea, nearbyOccupiedCounter, tooManyPeople)
	}

	fmt.Printf("\nArea constant after %d iterations with %d occupied seats in state:\n%v",
		iterations, countAllOccupied(oldArea), oldArea)
	return areas
}

func iterate(oldArea Area, nearbyOccupiedCounter func(Area, int, int) int, tooManyPeople int) Area {
	newArea := make([][]Position, len(oldArea))
	for i, row := range oldArea {
		newArea[i] = make([]Position, len(row))
		for j, pos := range row {
			count := nearbyOccupiedCounter(oldArea, i, j)
			if pos == EmptySeat && count == 0 {
				newArea[i][j] = OccupiedSeat
			} else if pos == OccupiedSeat && count >= tooManyPeople {
				newArea[i][j] = EmptySeat
			} else {
				newArea[i][j] = oldArea[i][j]
			}
		}
	}
	return newArea
}

func adjacentOccupiedCounter(area Area, row, col int) (count int) {
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			} else if i < 0 || j < 0 || i >= len(area) || j >= len(area[i]) {
				continue
			} else if area[i][j] == OccupiedSeat {
				count++
			}
		}
	}
	return count
}

func visibleOccupiedCounter(area Area, row, col int) (count int) {
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}
	for _, d := range directions {
		i, j := row+d[0], col+d[1]
		for 0 <= i && 0 <= j && i < len(area) && j < len(area[i]) {
			pos := area[i][j]
			if pos == OccupiedSeat {
				count++
				break
			} else if pos == EmptySeat {
				break
			} else {
				i, j = i+d[0], j+d[1]
			}
		}
	}
	return count
}

func countAllOccupied(area Area) (count int) {
	for i := range area {
		for _, pos := range area[i] {
			if pos == OccupiedSeat {
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
		for j, strPos := range positions {
			pos := Position(strPos)
			if pos != Floor && pos != EmptySeat && pos != OccupiedSeat {
				log.Panicf("Found position not valid: %s", strPos)
			}
			area[i][j] = pos
		}
	}
	return area
}

//goland:noinspection GoNilness
func areasToGIF(areas []Area) *gif.GIF {
	if len(areas) == 0 {
		log.Panic("Cannot convert empty area slice to gif.")
	}

	var images []*image.Paletted
	var delays []int
	delay := 10

	for _, area := range areas {
		img := areaToImage(area)
		images = append(images, img)
		delays = append(delays, delay)
	}

	delays[0] = delay * 10
	delays[len(delays)-1] = delay * 10

	return &gif.GIF{
		Image: images,
		Delay: delays,
	}
}

func areaToImage(area Area) *image.Paletted {
	white := color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff}
	gray := color.RGBA{R: 0x44, G: 0x44, B: 0x44, A: 0xff}
	black := color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff}

	palette := color.Palette{white, gray, black}
	height := len(area)
	width := len(area[0])
	img := image.NewPaletted(image.Rect(0, 0, height, width), palette)

	for x := range area {
		for y, pos := range area[x] {
			if pos == Floor {
				img.Set(x, y, white)
			} else if pos == EmptySeat {
				img.Set(x, y, gray)
			} else if pos == OccupiedSeat {
				img.Set(x, y, black)
			} else {
				log.Printf("Unkown type of position: %s", pos)
			}
		}
	}

	return img
}

func writeGifToFile(areaGif *gif.GIF, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		err1 := gif.EncodeAll(file, areaGif)
		err2 := file.Close()
		if err1 != nil {
			fmt.Println(err1)
		}
		if err2 != nil {
			fmt.Println(err1)
		}
	}
}
