package main

import (
	"2023/src/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func processing(filename string) utils.BoundedGrid[rune] {
	grid := make(utils.Grid[rune])
	lines := strings.Split(utils.ReadFile(filename), "\n")
	for r, line := range lines {
		for c, ch := range line {
			grid[utils.Pos{R: r, C: c}] = ch // '#' = 35 ; '.' = 46 ; 'O' = 79
		}
	}
	return utils.BoundedGrid[rune]{Grid: grid, Dims: utils.Dims{Rows: len(lines), Cols: len(lines[0])}}
}

func slide(grid *utils.BoundedGrid[rune]) {
	for c := 0; c < grid.Dims.Cols; c++ {
		col := grid.GetColFromBoundedGrid(c)
		groups := strings.Split(string(col), "#")
		for i, group := range groups {
			groupList := []rune(group)
			slices.Sort(groupList)
			slices.Reverse(groupList)
			groups[i] = string(groupList)
		}
		newCol := strings.Join(groups, "#")
		grid.SetCol(c, []rune(newCol))
	}
}

func convert(grid utils.BoundedGrid[rune]) string {
	var str string
	for r := 0; r < grid.Dims.Rows; r++ {
		str += string(grid.GetRowFromBoundedGrid(r))
	}
	return str
}

func reconvert(str string, dims utils.Dims) utils.BoundedGrid[rune] {
	grid := make(utils.Grid[rune])
	i := 0
	for r := 0; r < dims.Rows; r++ {
		for c := 0; c < dims.Cols; c++ {
			grid[utils.Pos{R: r, C: c}] = rune(str[i])
			i++
		}
	}
	return utils.BoundedGrid[rune]{Grid: grid, Dims: dims}
}

func cycle(grid *utils.BoundedGrid[rune]) {
	slide(grid)
	for i := 0; i < 3; i++ {
		*grid = grid.Transpose()
		reverse(grid)
		slide(grid)
	}
	*grid = grid.Transpose()
	reverse(grid)
}

func reverse(grid *utils.BoundedGrid[rune]) {
	for r := 0; r < grid.Dims.Rows; r++ {
		newRow := grid.GetRowFromBoundedGrid(r)
		slices.Reverse(newRow)
		grid.SetRow(r, newRow)
	}
}

func part1(filename string) int {
	grid := processing(filename)
	slide(&grid)
	return utils.ReduceBoundedGrid(grid, func(pos utils.Pos, element rune) int {
		if element == 'O' {
			return grid.Dims.Rows - pos.R
		}
		return 0
	}, 0)
}

func part2(filename string) int {
	grid := processing(filename)
	dims := grid.Dims
	gridStr := convert(grid)
	seen := utils.NewSet[string]()
	seen.Add(gridStr)
	list := []string{gridStr}
	i := 0
	for {
		i++
		cycle(&grid)
		gridStr = convert(grid)
		if seen.Contains(gridStr) {
			break
		}
		seen.Add(gridStr)
		list = append(list, gridStr)
	}
	first := slices.Index(list, gridStr)
	grid = reconvert(list[(1000000000-first)%(i-first)+first], dims)
	return utils.ReduceBoundedGrid(grid, func(pos utils.Pos, element rune) int {
		if element == 'O' {
			return grid.Dims.Rows - pos.R
		}
		return 0
	}, 0)
}

func main() {
	filename := "day14/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
