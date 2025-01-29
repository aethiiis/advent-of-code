package main

import (
	"2023/src/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"strings"
)

func processing(filename string) [][]string {
	text := utils.ReadFile(filename)
	blocks := strings.SplitN(text, "\n\n", -1)
	mirrors := make([][]string, len(blocks))
	for i := range mirrors {
		mirrors[i] = strings.SplitN(blocks[i], "\n", -1)
	}
	return mirrors
}

func symmetry(mirror []string, diff int) int {
	for i := 1; i < len(mirror); i++ {
		above, below := append([]string{}, mirror[:i]...), append([]string{}, mirror[i:]...)
		slices.Reverse(above)
		sum := 0
	Symmetry:
		for j := 0; j < min(len(below), len(above)); j++ {
			if above[j] != below[j] {
				for k := range above[j] {
					if above[j][k] != below[j][k] {
						sum++
					}
					if sum > diff {
						break Symmetry
					}
				}
			}
		}
		if sum == diff {
			return i
		}
	}
	return 0
}

func transpose(mirror []string) []string {
	rows, cols := len(mirror), len(mirror[0])
	transposed, tmp := make([]string, cols), make([][]uint8, cols)
	for c := 0; c < cols; c++ {
		transposed[c], tmp[c] = strings.Repeat(" ", rows), make([]uint8, rows)
		for r := 0; r < rows; r++ {
			tmp[c][r] = mirror[r][c]
		}
		transposed[c] = string(tmp[c])
	}
	return transposed
}

func part1(filename string) int {
	return utils.Sum(utils.Map2(processing(filename), func(mirror []string) int {
		return symmetry(mirror, 0)*100 + symmetry(transpose(mirror), 0)
	}))
}

func part2(filename string) int {
	return utils.Sum(utils.Map2(processing(filename), func(mirror []string) int {
		return symmetry(mirror, 1)*100 + symmetry(transpose(mirror), 1)
	}))
}

func main() {
	filename := "day13/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
