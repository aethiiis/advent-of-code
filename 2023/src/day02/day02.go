package main

import (
	"2023/src/utils"
	"strconv"
	"strings"
)

func processing(filename string) map[int][][]int {
	games := make(map[int][][]int)
	lines := strings.SplitN(utils.ReadFile(filename), "\n", -1)
	for _, line := range lines {
		game := strings.Split(line, ": ")
		id, _ := strconv.Atoi(game[0][5:])
		sets := strings.SplitN(game[1], "; ", -1)
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			count := []int{0, 0, 0}
			for _, cube := range cubes {
				single := strings.Split(cube, " ")
				number, _ := strconv.Atoi(single[0])
				color := single[1]
				if color == "red" {
					count[0] += number
				} else if color == "green" {
					count[1] += number
				} else {
					count[2] += number
				}
			}
			games[id] = append(games[id], count)
		}
	}
	return games
}

func part1(filename string) int {
	var result int
	games := processing(filename)
	for id, game := range games {
		enough := true
		for _, set := range game {
			if set[0] > 12 {
				enough = false
			}
			if set[1] > 13 {
				enough = false
			}
			if set[2] > 14 {
				enough = false
			}
		}
		if enough {
			result += id
		}
	}
	return result
}

func part2(filename string) int {
	var result int
	games := processing(filename)
	for _, game := range games {
		red := 0
		green := 0
		blue := 0
		for _, set := range game {
			if set[0] > red {
				red = set[0]
			}
			if set[1] > green {
				green = set[1]
			}
			if set[2] > blue {
				blue = set[2]
			}
		}
		result += red * green * blue
	}
	return result
}

func main() {
	filename := "src/day02/input.txt"
	println(part1(filename))
	println(part2(filename))
}
