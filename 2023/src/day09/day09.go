package main

import (
	"2023/src/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

func processing(filename string) [][]int {
	return utils.Map2(strings.SplitN(utils.ReadFile(filename), "\n", -1), func(str string) []int {
		return utils.Map2(strings.SplitN(str, " ", -1), func(num string) int {
			number, _ := strconv.Atoi(num)
			return number
		})
	})
}

func extrapolate(sequence []int) int {
	n := len(sequence) - 1
	if n > -1 {
		differences := make([]int, n)
		for i := 0; i < n; i++ {
			differences[i] = sequence[i+1] - sequence[i]
		}
		return sequence[n] + extrapolate(differences)
	} else {
		return 0
	}
}

func part1(filename string) int {
	return utils.Sum(utils.Map2(processing(filename), extrapolate))
}

func part2(filename string) int {
	return utils.Sum(utils.Map2(processing(filename), func(sequence []int) int {
		slices.Reverse(sequence)
		return extrapolate(sequence)
	}))
}

func main() {
	filename := "day09/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
