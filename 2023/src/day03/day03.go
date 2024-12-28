package main

import (
	"2023/src/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func processing(filename string) (utils.Grid[rune], map[int][]utils.Pos) {
	text := utils.ReadFile(filename)
	grid := utils.GetRuneGridFromString(text)
	numbers := getNumbers(text)
	return grid, numbers
}

func getNumbers(text string) map[int][]utils.Pos {
	list := make(map[int][]utils.Pos)
	expression, _ := regexp.Compile("\\d+")
	lines := strings.SplitN(text, "\n", -1)
	for r, line := range lines {
		positions := expression.FindAllStringIndex(line, -1)
		numbers := expression.FindAllString(line, -1)
		for i := range len(numbers) {
			number, _ := strconv.Atoi(numbers[i])
			var pos []utils.Pos
			for _, p := range positions[i] {
				pos = append(pos, utils.Pos{X: r, Y: p})
			}
			list[number] = pos
		}
	}
	return list
}

func part1(filename string) int {
	var result int
	grid, numbers := processing(filename)
	for number, positions := range numbers {
		part := true
		for _, position := range positions {
			neighbours := grid.GetRuneNeighbours(position, true)
			for pos, neighbour := range neighbours {
				if neighbour != rune('.') {
					part = false
					break
				}
			}
		}
	}
	return result
}

func part2(filename string) int {
	var result int
	return result
}

func main() {
	filename := "src/day03/test.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
