package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

type Ray struct {
	r, c, dr, dc int
}

func processing(filename string) utils.BoundedGrid[byte] {
	lines := strings.Split(utils.ReadFile(filename), "\n")
	grid, dims := make(utils.Grid[byte]), utils.Dims{Rows: len(lines), Cols: len(lines[0])}
	for r, line := range lines {
		for c, ch := range line {
			grid[utils.Pos{R: r, C: c}] = byte(ch)
		}
	}
	return utils.BoundedGrid[byte]{Grid: grid, Dims: dims}
}

func add(seen *utils.Set[Ray], queue *utils.Queue[Ray], next Ray) {
	if !seen.Contains(next) {
		seen.Add(next)
		queue.Put(next)
	}
}

func propagate(grid utils.BoundedGrid[byte], start Ray) utils.Set[Ray] {
	seen := utils.NewSet[Ray]()
	queue := utils.NewQueue[Ray](0)
	queue.Put(start)
	for !queue.Empty() {
		current := queue.Get()
		r, c, dr, dc := current.r, current.c, current.dr, current.dc
		r, c = r+dr, c+dc
		if r < 0 || r >= grid.Dims.Rows || c < 0 || c >= grid.Dims.Cols {
			continue
		}
		ch := grid.Grid[utils.Pos{R: r, C: c}]
		switch {
		case ch == 47:
			add(&seen, &queue, Ray{r: r, c: c, dr: -dc, dc: -dr})
		case ch == 92:
			add(&seen, &queue, Ray{r: r, c: c, dr: dc, dc: dr})
		case ch == 124 && dr == 0:
			add(&seen, &queue, Ray{r: r, c: c, dr: -1, dc: 0})
			add(&seen, &queue, Ray{r: r, c: c, dr: 1, dc: 0})
		case ch == 45 && dc == 0:
			add(&seen, &queue, Ray{r: r, c: c, dr: 0, dc: -1})
			add(&seen, &queue, Ray{r: r, c: c, dr: 0, dc: 1})
		default:
			add(&seen, &queue, Ray{r: r, c: c, dr: dr, dc: dc})
		}
	}
	return seen
}

func calculate(seen utils.Set[Ray]) int {
	energy := utils.NewSet[utils.Pos]()
	for _, ray := range seen.List() {
		energy.Add(utils.Pos{R: ray.r, C: ray.c})
	}
	return energy.Size()
}

func part1(filename string) int {
	return calculate(propagate(processing(filename), Ray{r: 0, c: -1, dr: 0, dc: 1}))
}

func part2(filename string) int {
	grid := processing(filename)
	starts := make([]Ray, 0)
	for r := 0; r < grid.Dims.Rows; r++ {
		starts = append(starts, Ray{r: r, c: -1, dr: 0, dc: 1}, Ray{r: r, c: grid.Dims.Cols, dr: 0, dc: -1})
	}
	for c := 0; c < grid.Dims.Cols; c++ {
		starts = append(starts, Ray{r: -1, c: c, dr: 1, dc: 0}, Ray{r: grid.Dims.Rows, c: c, dr: -1, dc: 0})
	}
	return utils.Max(utils.Map2(starts, func(start Ray) int {
		return calculate(propagate(grid, start))
	}), nil)
}

func main() {
	filename := "day16/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
