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
	part1Areas := iterateUntilConstant(area)

	writeGifToFile(areasToGIF(part1Areas), "2020\\Day 11\\part1.gif")
}

func iterateUntilConstant(area Area) (areas []Area) {
	iterations := 0
	areas = append(areas, area)

	oldArea, newArea := area, iterate(area)
	for !oldArea.equals(newArea) {
		iterations++
		areas = append(areas, newArea)
		oldArea, newArea = newArea, iterate(newArea)
	}

	fmt.Printf("Area constant after %d iterations with %d occupied seats in state:\n%v",
		iterations, countAllOccupied(oldArea), oldArea)
	return areas
}

func writeGifToFile(areaGif *gif.GIF, path string) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	} else {
		gif.EncodeAll(file, areaGif)
		file.Close()
	}
}

func iterate(oldArea Area) Area {
	newArea := make([][]Position, len(oldArea))
	for i, row := range oldArea {
		newArea[i] = make([]Position, len(row))
		for j, pos := range row {
			count := nearbyOccupied(oldArea, i, j)
			if pos == EmptySeat && count == 0 {
				newArea[i][j] = OccupiedSeat
			} else if pos == OccupiedSeat && count >= 4 {
				newArea[i][j] = EmptySeat
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
			} else if i < 0 || j < 0 || i >= len(area) || j >= len(area[i]) {
				continue
			} else if area[i][j] == OccupiedSeat {
				count++
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
