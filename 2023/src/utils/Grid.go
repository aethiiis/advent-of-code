package utils

import "strings"

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

func GetGridFromString[T any](text string, sep string) Grid[T] {
	grid := make(Grid[T])
	lines := strings.SplitN(text, sep, -1)
	for r, line := range lines {
		for c, char := range line {
			grid[Pos{R: r, C: c}] = int(char)
		}
	}
	return grid
}

func (grid Grid[int]) GetIntNeighbours(pos Pos, diag bool) map[Pos]int {
	neighbours := make(map[Pos]int)
	if value, ok := grid[Pos{pos.X + 1, pos.Y}]; ok {
		neighbours[Pos{pos.X + 1, pos.Y}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		neighbours[Pos{pos.X, pos.Y - 1}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y + 1}]; ok {
		neighbours[Pos{pos.X, pos.Y + 1}] = value
	}
	if value, ok := grid[Pos{pos.X - 1, pos.Y}]; ok {
		neighbours[Pos{pos.X - 1, pos.Y}] = value
	}
	if diag {
		if value, ok := grid[Pos{pos.X + 1, pos.Y + 1}]; ok {
			neighbours[Pos{pos.X + 1, pos.Y + 1}] = value
		}
		if value, ok := grid[Pos{pos.X - 1, pos.Y - 1}]; ok {
			neighbours[Pos{pos.X - 1, pos.Y - 1}] = value
		}
		if value, ok := grid[Pos{pos.X - 1, pos.Y + 1}]; ok {
			neighbours[Pos{pos.X - 1, pos.Y + 1}] = value
		}
		if value, ok := grid[Pos{pos.X + 1, pos.Y - 1}]; ok {
			neighbours[Pos{pos.X + 1, pos.Y - 1}] = value
		}
	}
	return neighbours
}

func GetRuneGridFromList(list [][]rune) Grid[rune] {
	grid := make(Grid[rune])
	for r := range list {
		for c := range list[r] {
			grid[Pos{X: r, Y: c}] = list[r][c]
		}
	}
	return grid
}

func GetRuneGridFromString(text string) Grid[rune] {
	grid := make(Grid[rune])
	lines := strings.SplitN(text, "\n", -1)
	for r, line := range lines {
		for c, char := range line {
			grid[Pos{r, c}] = char
		}
	}
	return grid
}

func (grid Grid[rune]) GetRuneNeighbours(pos Pos, diag bool) map[Pos]rune {
	neighbours := make(map[Pos]rune)
	if value, ok := grid[Pos{pos.X + 1, pos.Y}]; ok {
		neighbours[Pos{pos.X + 1, pos.Y}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		neighbours[Pos{pos.X, pos.Y - 1}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y + 1}]; ok {
		neighbours[Pos{pos.X, pos.Y + 1}] = value
	}
	if value, ok := grid[Pos{pos.X - 1, pos.Y}]; ok {
		neighbours[Pos{pos.X - 1, pos.Y}] = value
	}
	if diag {
		if value, ok := grid[Pos{pos.X + 1, pos.Y + 1}]; ok {
			neighbours[Pos{pos.X + 1, pos.Y + 1}] = value
		}
		if value, ok := grid[Pos{pos.X - 1, pos.Y - 1}]; ok {
			neighbours[Pos{pos.X - 1, pos.Y - 1}] = value
		}
		if value, ok := grid[Pos{pos.X - 1, pos.Y + 1}]; ok {
			neighbours[Pos{pos.X - 1, pos.Y + 1}] = value
		}
		if value, ok := grid[Pos{pos.X + 1, pos.Y - 1}]; ok {
			neighbours[Pos{pos.X + 1, pos.Y - 1}] = value
		}
	}
	return neighbours
}

func GetBoolGridFromList(list [][]bool) Grid[bool] {
	grid := make(Grid[bool])
	for r := range list {
		for c := range list[r] {
			grid[Pos{X: r, Y: c}] = list[r][c]
		}
	}
	return grid
}

func (grid Grid[bool]) GetBoolNeighbours(pos Pos) map[Pos]bool {
	neighbours := make(map[Pos]bool)
	if value, ok := grid[Pos{pos.X + 1, pos.Y}]; ok {
		neighbours[Pos{pos.X + 1, pos.Y}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y - 1}]; ok {
		neighbours[Pos{pos.X, pos.Y - 1}] = value
	}
	if value, ok := grid[Pos{pos.X, pos.Y + 1}]; ok {
		neighbours[Pos{pos.X, pos.Y + 1}] = value
	}
	if value, ok := grid[Pos{pos.X - 1, pos.Y}]; ok {
		neighbours[Pos{pos.X - 1, pos.Y}] = value
	}
	return neighbours
}
