package utils

import "strings"

type Pos struct {
	X, Y int
}

type Grid[T comparable] map[Pos]T

func GetIntGridFromList(list [][]int) Grid[int] {
	grid := make(Grid[int])
	for r := range list {
		for c := range list[r] {
			grid[Pos{X: r, Y: c}] = list[r][c]
		}
	}
	return grid
}

func GetIntGridFromString(text string) Grid[int] {
	grid := make(Grid[int])
	lines := strings.SplitN(text, "\n", -1)
	for r, line := range lines {
		for c, char := range line {
			grid[Pos{r, c}] = int(char)
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
