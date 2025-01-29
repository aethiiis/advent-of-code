package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Pos struct {
	R, C int
}

type Dims struct {
	Rows, Cols int
}

type Direction struct {
	Dr, Dc int
}

type Grid[T any] map[Pos]T

type BoundedGrid[T any] struct {
	Grid Grid[T]
	Dims Dims
}

func (p *Pos) Add(other Pos) Pos {
	return Pos{R: p.R + other.R, C: p.C + other.C}
}

func (p *Pos) Move(dir Direction) Pos {
	return Pos{R: p.R + dir.Dr, C: p.C + dir.Dc}
}

func GetGridFromList[T any](list [][]T) Grid[T] {
	grid := make(Grid[T])
	for r := range list {
		for c := range list[r] {
			grid[Pos{R: r, C: c}] = list[r][c]
		}
	}
	return grid
}

func GetBoundedGridFromList[T any](list [][]T) BoundedGrid[T] {
	return BoundedGrid[T]{Grid: GetGridFromList(list), Dims: Dims{Rows: len(list), Cols: len(list[0])}}
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

func GetDefaultBoundedGrid[T any](defaultValue T, dims Dims) BoundedGrid[T] {
	return BoundedGrid[T]{Grid: GetDefaultGrid(defaultValue, dims), Dims: dims}
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

func GetBoundedGridFromString[T any](text, sep1, sep2 string) BoundedGrid[any] {
	return BoundedGrid[any]{Grid: GetGridFromString[any](text, sep1, sep2), Dims: Dims{Rows: len(text), Cols: len(strings.SplitN(text, "\n", 1)[0])}}
}

func (grid *Grid[T]) Transpose() Grid[T] {
	transposed := make(Grid[T])
	for pos, value := range *grid {
		transposed[Pos{R: pos.C, C: pos.R}] = value
	}
	return transposed
}

func (grid *BoundedGrid[T]) Transpose() BoundedGrid[T] {
	return BoundedGrid[T]{Grid: grid.Grid.Transpose(), Dims: Dims{Rows: grid.Dims.Cols, Cols: grid.Dims.Rows}}
}

func (grid *Grid[T]) GetListFromGrid() [][]T {
	dims := grid.GetDimsFromGrid()
	list := make([][]T, dims.Rows)
	for i := range list {
		list[i] = make([]T, dims.Cols)
	}
	for r := 0; r < dims.Rows; r++ {
		for c := 0; c < dims.Cols; c++ {
			list[r][c] = (*grid)[Pos{R: r, C: c}]
		}
	}
	return list
}

func (grid *BoundedGrid[T]) GetListFromBoundedGrid() [][]T {
	dims := grid.Dims
	list := make([][]T, dims.Rows)
	for i := range list {
		list[i] = make([]T, dims.Cols)
	}
	for r := 0; r < dims.Rows; r++ {
		for c := 0; c < dims.Cols; c++ {
			list[r][c] = grid.Grid[Pos{R: r, C: c}]
		}
	}
	return list
}

func (grid *Grid[T]) GetDimsFromGrid() Dims {
	rows := 0
	cols := 0
	for pos := range *grid {
		if pos.R > rows {
			rows = pos.R
		}
		if pos.C > cols {
			cols = pos.C
		}
	}
	return Dims{Rows: rows + 1, Cols: cols + 1}
}

func (grid *Grid[T]) GetNeighbours(pos Pos, diag bool) map[Pos]T {
	neighbours := make(map[Pos]T)
	if value, ok := (*grid)[Pos{pos.R + 1, pos.C}]; ok {
		neighbours[Pos{pos.R + 1, pos.C}] = value
	}
	if value, ok := (*grid)[Pos{pos.R, pos.C - 1}]; ok {
		neighbours[Pos{pos.R, pos.C - 1}] = value
	}
	if value, ok := (*grid)[Pos{pos.R, pos.C + 1}]; ok {
		neighbours[Pos{pos.R, pos.C + 1}] = value
	}
	if value, ok := (*grid)[Pos{pos.R - 1, pos.C}]; ok {
		neighbours[Pos{pos.R - 1, pos.C}] = value
	}
	if diag {
		if value, ok := (*grid)[Pos{pos.R + 1, pos.C + 1}]; ok {
			neighbours[Pos{pos.R + 1, pos.C + 1}] = value
		}
		if value, ok := (*grid)[Pos{pos.R - 1, pos.C - 1}]; ok {
			neighbours[Pos{pos.R - 1, pos.C - 1}] = value
		}
		if value, ok := (*grid)[Pos{pos.R - 1, pos.C + 1}]; ok {
			neighbours[Pos{pos.R - 1, pos.C + 1}] = value
		}
		if value, ok := (*grid)[Pos{pos.R + 1, pos.C - 1}]; ok {
			neighbours[Pos{pos.R + 1, pos.C - 1}] = value
		}
	}
	return neighbours
}

func (grid *BoundedGrid[T]) GetNeighbours(pos Pos, diag bool) map[Pos]T {
	return grid.Grid.GetNeighbours(pos, diag)
}

func (grid *Grid[T]) PrintGrid(tab string) {
	list := grid.GetListFromGrid()
	if len(list) == 0 {
		return
	}
	for _, line := range list {
		for _, item := range line {
			fmt.Print(item)
			fmt.Print(tab)
		}
		fmt.Println()
	}
}

func (grid *BoundedGrid[T]) PrintBoundedGrid(tab string) {
	list := grid.GetListFromBoundedGrid()
	if len(list) == 0 {
		return
	}
	for _, line := range list {
		for _, item := range line {
			fmt.Print(item)
			fmt.Print(tab)
		}
		fmt.Println()
	}
}

func (grid *Grid[T]) Len() int {
	return len(*grid)
}

func (grid *BoundedGrid[T]) Len() int {
	return len(grid.Grid)
}

func (grid *Grid[T]) GetRowFromGrid(r int) []T {
	dims := grid.GetDimsFromGrid()
	row := make([]T, dims.Cols)
	for c := 0; c < dims.Cols; c++ {
		row[c] = (*grid)[Pos{R: r, C: c}]
	}
	return row
}

func (grid *BoundedGrid[T]) GetRowFromBoundedGrid(r int) []T {
	row := make([]T, grid.Dims.Cols)
	for c := 0; c < grid.Dims.Cols; c++ {
		row[c] = grid.Grid[Pos{R: r, C: c}]
	}
	return row
}

func (grid *Grid[T]) GetColFromGrid(c int) []T {
	dims := grid.GetDimsFromGrid()
	col := make([]T, dims.Rows)
	for r := 0; r < dims.Cols; r++ {
		col[r] = (*grid)[Pos{R: r, C: c}]
	}
	return col
}

func (grid *BoundedGrid[T]) GetColFromBoundedGrid(c int) []T {
	col := make([]T, grid.Dims.Cols)
	for r := 0; r < grid.Dims.Rows; r++ {
		col[r] = grid.Grid[Pos{R: r, C: c}]
	}
	return col
}

func (grid *Grid[T]) SetRow(r int, row []T) {
	dims := grid.GetDimsFromGrid()
	if len(row) != dims.Cols {
		panic("The new row needs to have the correct size")
	}
	for c := 0; c < dims.Cols; c++ {
		(*grid)[Pos{R: r, C: c}] = row[c]
	}
}

func (grid *BoundedGrid[T]) SetRow(r int, row []T) {
	if len(row) != grid.Dims.Cols {
		panic("The new row needs to have the correct size")
	}
	for c := 0; c < grid.Dims.Cols; c++ {
		grid.Grid[Pos{R: r, C: c}] = row[c]
	}
}

func (grid *Grid[T]) SetCol(c int, col []T) {
	dims := grid.GetDimsFromGrid()
	if len(col) != dims.Rows {
		panic("The new column needs to have the correct size")
	}
	for r := 0; r < dims.Rows; r++ {
		(*grid)[Pos{R: r, C: c}] = col[r]
	}
}

func (grid *BoundedGrid[T]) SetCol(c int, col []T) {
	if len(col) != grid.Dims.Rows {
		panic("The new column needs to have the correct size")
	}
	for r := 0; r < grid.Dims.Rows; r++ {
		grid.Grid[Pos{R: r, C: c}] = col[r]
	}
}

func ReduceGrid[T any, U int | float32 | float64 | string](grid Grid[T], f func(T) U, initial U) U {
	acc := initial
	for _, element := range grid {
		acc += f(element)
	}
	return acc
}

func ReduceBoundedGrid[T any, U int | float32 | float64 | string](grid BoundedGrid[T], f func(Pos, T) U, initial U) U {
	acc := initial
	for pos, element := range grid.Grid {
		acc += f(pos, element)
	}
	return acc
}
