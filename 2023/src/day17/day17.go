package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

type Crucible struct {
	r, c, dr, dc, n int
}

func processing(filename string) utils.BoundedGrid[int] {
	lines := strings.Split(utils.ReadFile(filename), "\n")
	grid, dims := make(utils.Grid[int]), utils.Dims{Rows: len(lines), Cols: len(lines[0])}
	for r, line := range lines {
		for c, ch := range line {
			grid[utils.Pos{R: r, C: c}] = int(ch) - 48
		}
	}
	return utils.BoundedGrid[int]{Grid: grid, Dims: dims}
}

func dijkstra(grid utils.BoundedGrid[int], max, min int) int {
	seen := utils.NewSet[Crucible]()
	queue := utils.PriorityQueue[Crucible]{}
	queue.Put(Crucible{r: 0, c: 0, dr: 0}, 0)
	for !queue.Empty() {
		current, priority := queue.Get()
		r, c, dr, dc, n := current.r, current.c, current.dr, current.dc, current.n
		if r == grid.Dims.Rows-1 && c == grid.Dims.Cols-1 && n >= min {
			return priority
		}
		if seen.Contains(current) {
			continue
		}
		seen.Add(current)
		if n < max && (dr != 0 || dc != 0) {
			nr, nc := r+dr, c+dc
			if 0 <= nr && nr < grid.Dims.Rows && 0 <= nc && nc < grid.Dims.Cols {
				queue.Put(Crucible{r: nr, c: nc, dr: dr, dc: dc, n: n + 1}, priority+grid.Grid[utils.Pos{R: nr, C: nc}])
			}
		}
		if n >= min || (dr == 0 && dc == 0) {
			directions := []utils.Direction{{Dr: 0, Dc: 1}, {Dr: 0, Dc: -1}, {Dr: -1, Dc: 0}, {Dr: 1, Dc: 0}}
			same, opposite := utils.Direction{Dr: dr, Dc: dc}, utils.Direction{Dr: -dr, Dc: -dc}
			for _, dir := range directions {
				if dir != same && dir != opposite {
					nr, nc := r+dir.Dr, c+dir.Dc
					if 0 <= nr && nr < grid.Dims.Rows && 0 <= nc && nc < grid.Dims.Cols {
						queue.Put(Crucible{r: nr, c: nc, dr: dir.Dr, dc: dir.Dc, n: 1}, priority+grid.Grid[utils.Pos{R: nr, C: nc}])
					}
				}
			}
		}
	}
	return 0
}

func part1(filename string) int {
	return dijkstra(processing(filename), 3, 0)
}

func part2(filename string) int {
	return dijkstra(processing(filename), 10, 4)
}

func main() {
	filename := "day17/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
