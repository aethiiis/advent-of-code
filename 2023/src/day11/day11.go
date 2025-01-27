package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

func processing(filename string) ([]utils.Pos, []int, []int) {
	lines := strings.SplitN(utils.ReadFile(filename), "\n", -1)
	rows, cols := len(lines), len(lines[0])
	galaxies := make([]utils.Pos, 0)
	emptyRows, emptyCols := make([]int, rows), make([]int, cols)
	for r, line := range lines {
		for c, ch := range line {
			pos := utils.Pos{R: r, C: c}
			if ch == '#' {
				galaxies = append(galaxies, pos)
				emptyRows[r] = 1
				emptyCols[c] = 1
			}
		}
	}
	return galaxies, emptyRows, emptyCols
}

func calculate(scale int, galaxies []utils.Pos, emptyRows, emptyCols []int) int {
	var distance int
	for i, pos := range galaxies {
		for _, other := range galaxies[:i] {
			for r := min(pos.R, other.R); r < max(pos.R, other.R); r++ {
				if emptyRows[r] == 0 {
					distance += scale
				} else {
					distance += 1
				}
			}
			for c := min(pos.C, other.C); c < max(pos.C, other.C); c++ {
				if emptyCols[c] == 0 {
					distance += scale
				} else {
					distance += 1
				}
			}
		}
	}
	return distance
}

func part1(filename string) int {
	galaxies, emptyRows, emptyCols := processing(filename)
	return calculate(2, galaxies, emptyRows, emptyCols)
}

func part2(filename string) int {
	galaxies, emptyRows, emptyCols := processing(filename)
	return calculate(1000000, galaxies, emptyRows, emptyCols)
}

func main() {
	filename := "day11/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
