package utils

import (
	"strconv"
	"strings"
)

type Pos struct {
	R, C int
}

type Dims struct {
	Rows, Cols int
}

type Grid[T any] map[Pos]T

func GetGridFromList[T any](list [][]T) Grid[T] {
	grid := make(Grid[T])
	for r := range list {
		for c := range list[r] {
			grid[Pos{R: r, C: c}] = list[r][c]
		}
	}
	return grid
}

func GetDefaultGrid[T any](defaultValue T, dims Dims) Grid[T] {
	grid := make(Grid[T])
	for r := 0; r < dims.Rows; r++ {
		for c := 0; c < dims.Cols; c++ {
			grid[Pos{R: r, C: c}] = defaultValue
		}
	}
	return grid
}

func (grid Grid[T]) GetListFromGrid() [][]T {
	dims := grid.GetDimsFromGrid()
	list := make([][]T, dims.Rows)
	for i := range list {
		list[i] = make([]T, dims.Cols)
	}
	for r := 0; r < dims.Rows; r++ {
		for c := 0; c < dims.Cols; c++ {
			list[r][c] = grid[Pos{R: r, C: c}]
		}
	}
	return list
}

func (grid Grid[T]) GetDimsFromGrid() Dims {
	rows := 0
	cols := 0
	for pos := range grid {
		if pos.R > rows {
			cols = pos.R
		}
		if pos.C > cols {
			rows = pos.C
		}
	}
	return Dims{Rows: rows, Cols: cols}
}

func GetGridFromString[T any](text, sep1, sep2 string) Grid[any] {
	grid := make(Grid[any])
	var test T
	lines := strings.SplitN(text, sep1, -1)
	switch any(test).(type) {
	case int:
		for r, line := range lines {
			values := strings.SplitN(line, sep2, -1)
			for c, char := range values {
				grid[Pos{R: r, C: c}], _ = strconv.Atoi(char)
			}
		}
	case rune:
		for r, line := range lines {
			values := strings.SplitN(line, sep2, -1)
			for c, char := range values {
				grid[Pos{R: r, C: c}] = char[0]
			}
		}
	case bool:
		for r, line := range lines {
			values := strings.SplitN(line, sep2, -1)
			for c, char := range values {
				if char[0] == '0' {
					grid[Pos{R: r, C: c}] = false
				} else {
					grid[Pos{R: r, C: c}] = true
				}
			}
		}
	case float64:
		if sep2 == "" {
			panic("Can't have empty secondary separator when working with floats")
		}
		for r, line := range lines {
			values := strings.SplitN(line, sep2, -1)
			for c, char := range values {
				grid[Pos{R: r, C: c}], _ = strconv.ParseFloat(char, 64)
			}
		}
	}
	return grid
}

func (grid Grid[T]) GetNeighbours(pos Pos, diag bool) map[Pos]T {
	neighbours := make(map[Pos]T)
	if value, ok := grid[Pos{pos.R + 1, pos.C}]; ok {
		neighbours[Pos{pos.R + 1, pos.C}] = value
	}
	if value, ok := grid[Pos{pos.R, pos.C - 1}]; ok {
		neighbours[Pos{pos.R, pos.C - 1}] = value
	}
	if value, ok := grid[Pos{pos.R, pos.C + 1}]; ok {
		neighbours[Pos{pos.R, pos.C + 1}] = value
	}
	if value, ok := grid[Pos{pos.R - 1, pos.C}]; ok {
		neighbours[Pos{pos.R - 1, pos.C}] = value
	}
	if diag {
		if value, ok := grid[Pos{pos.R + 1, pos.C + 1}]; ok {
			neighbours[Pos{pos.R + 1, pos.C + 1}] = value
		}
		if value, ok := grid[Pos{pos.R - 1, pos.C - 1}]; ok {
			neighbours[Pos{pos.R - 1, pos.C - 1}] = value
		}
		if value, ok := grid[Pos{pos.R - 1, pos.C + 1}]; ok {
			neighbours[Pos{pos.R - 1, pos.C + 1}] = value
		}
		if value, ok := grid[Pos{pos.R + 1, pos.C - 1}]; ok {
			neighbours[Pos{pos.R + 1, pos.C - 1}] = value
		}
	}
	return neighbours
}
