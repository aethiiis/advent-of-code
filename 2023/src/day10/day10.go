package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

func processing(filename string) (utils.Grid[rune], utils.Pos) {
	var start utils.Pos
	grid := make(utils.Grid[rune])
	for r, line := range strings.SplitN(utils.ReadFile(filename), "\n", -1) {
		for c, char := range line {
			if char == 83 {
				start = utils.Pos{R: r, C: c}
			}
			grid[utils.Pos{R: r, C: c}] = char
		}
	}
	return grid, start
}

func loop(grid utils.Grid[rune], start utils.Pos) []utils.Pos {
	path := make([]utils.Pos, 0)
	dims := grid.GetDimsFromGrid()
	seen := utils.NewSet[utils.Pos]()
	current, first := start, true
	for current != start || first {
		if first {
			first = false
		}
		r, c, ch := current.R, current.C, grid[current]
		to, bo, le, ri := utils.Pos{R: r - 1, C: c}, utils.Pos{R: r + 1, C: c}, utils.Pos{R: r, C: c - 1}, utils.Pos{R: r, C: c + 1}
		seen.Add(current)
		if top := grid[to]; r > 0 && (ch == 'S' || ch == '|' || ch == 'J' || ch == 'L') &&
			(top == '|' || top == '7' || top == 'F') && !seen.Contains(to) {
			current = to
		} else if bottom := grid[bo]; r < dims.Rows-1 && (ch == 'S' || ch == '|' || ch == '7' || ch == 'F') &&
			(bottom == '|' || bottom == 'J' || bottom == 'L') && !seen.Contains(bo) {
			current = bo
		} else if left := grid[le]; c > 0 && (ch == 'S' || ch == '-' || ch == 'J' || ch == '7') &&
			(left == '-' || left == 'L' || left == 'F') && !seen.Contains(le) {
			current = le
		} else if right := grid[ri]; c < dims.Cols-1 && (ch == 'S' || ch == '-' || ch == 'L' || ch == 'F') &&
			(right == '-' || right == 'J' || right == '7') && !seen.Contains(ri) {
			current = ri
		} else {
			break
		}
		path = append(path, current)
	}
	return append(path, start)
}

func part1(filename string) int {
	return len(loop(processing(filename))) / 2
}

func part2(filename string) int {
	path := loop(processing(filename))
	return utils.Shoelace(path) - (len(path) / 2) + 1
}

func main() {
	filename := "day10/test.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
