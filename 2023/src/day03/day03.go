package main

import (
	"2023/src/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Number struct {
	value     int
	positions []utils.Pos
}

func processing(filename string) (utils.Grid[any], []Number, []utils.Pos) {
	text := utils.ReadFile(filename)
	grid := utils.GetGridFromString[rune](text, "\n", "")
	lines := strings.SplitN(text, "\n", -1)
	numbers := getNumbers(lines)
	gears := getGears(lines)
	return grid, numbers, gears
}

func getNumbers(lines []string) []Number {
	var list []Number
	expression, _ := regexp.Compile("\\d+")
	for r, line := range lines {
		positions := expression.FindAllStringIndex(line, -1)
		numbers := expression.FindAllString(line, -1)
		for i := range numbers {
			number, _ := strconv.Atoi(numbers[i])
			var pos []utils.Pos
			for _, p := range positions[i] {
				pos = append(pos, utils.Pos{R: r, C: p})
			}
			pos[1].C = pos[1].C - 1
			list = append(list, Number{number, pos})
		}
	}
	return list
}

func getGears(lines []string) []utils.Pos {
	var list []utils.Pos
	expression, _ := regexp.Compile("\\*")
	for r, line := range lines {
		positions := expression.FindAllStringIndex(line, -1)
		for _, pos := range positions {
			list = append(list, utils.Pos{R: r, C: pos[0]})
		}
	}
	return list
}

func part1(filename string) int {
	var result int
	grid, numbers, _ := processing(filename)
	for _, number := range numbers {
		part := false
		for _, position := range number.positions {
			neighbours := grid.GetNeighbours(position, true)
			for _, neighbour := range neighbours {
				if uintNeighbour, ok := neighbour.(uint8); ok && uintNeighbour != 46 && uintNeighbour != 13 && (uintNeighbour > 57 || uintNeighbour < 48) {
					part = true
					break
				}
			}
		}
		if part {
			result += number.value
		}
	}
	return result
}

func part2(filename string) int {
	var result int
	grid, numbers, gears := processing(filename)
GearLoop:
	for _, gear := range gears {
		var first Number
		var second Number
		neighbours := grid.GetNeighbours(gear, true)
	NeighboursLoop:
		for neighbourPosition := range neighbours {
			for _, number := range numbers {
				for _, numberPositions := range number.positions {
					if neighbourPosition == numberPositions {
						if first.value == 0 {
							first = number
						} else if number.value == first.value {
							continue NeighboursLoop
						} else if second.value == 0 {
							second = number
						} else if number.value == second.value {
							continue NeighboursLoop
						} else {
							continue GearLoop
						}
					}
				}
			}
		}
		result += first.value * second.value
	}
	return result
}

func main() {
	filename := "src/day03/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
