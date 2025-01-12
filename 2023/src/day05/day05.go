package main

import (
	"2023/src/utils"
	"fmt"
	"golang.org/x/exp/slices"
	"strconv"
	"strings"
)

type IntervalMap struct {
	source, destination, length int
}

func processing(filename string) ([]int, [][]IntervalMap) {
	almanac := make([][]IntervalMap, 7)
	maps := strings.SplitN(utils.ReadFile(filename), "\n\n", -1)
	seeds := utils.Map2(strings.SplitN(maps[0][7:], " ", -1), func(str string) int {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return number
	})
	for i, MapText := range maps[1:] {
		almanac[i] = make([]IntervalMap, 0)
		lines := strings.SplitN(MapText, "\n", -1)
		for _, line := range lines[1:] {
			numbers := utils.Map2(strings.SplitN(line, " ", -1), func(str string) int {
				number, err := strconv.Atoi(str)
				if err != nil {
					panic(err)
				}
				return number
			})
			intervalMap := IntervalMap{
				source:      numbers[1],
				destination: numbers[0],
				length:      numbers[2],
			}
			almanac[i] = append(almanac[i], intervalMap)
		}
	}
	return seeds, almanac
}

func part1(filename string) int {
	seeds, almanac := processing(filename)
	result := make([]int, len(seeds))
	copy(result, seeds)
	for _, intervals := range almanac {
		next := make([]int, 0)
		for _, seed := range result {
			found := false
			for _, interval := range intervals {
				if interval.source <= seed && seed < interval.source+interval.length {
					next = append(next, seed-interval.source+interval.destination)
					found = true
					break
				}
			}
			if !found {
				next = append(next, seed)
			}
		}
		result = next
	}
	return slices.Min(result)
}

func part2(filename string) int {
	var result []int
	seeds, almanac := processing(filename)
	for i := 0; i < len(seeds); i += 2 {
		result = append(result, seeds[i], seeds[i]+seeds[i+1])
	}
	for _, intervals := range almanac {
		next := make([]int, 0)
		for len(result) > 0 {
			start, end := result[0], result[1]
			result = result[2:]
			found := false
			for _, interval := range intervals {
				overlapStart, overlapEnd := max(start, interval.source), min(end, interval.source+interval.length)
				if overlapStart < overlapEnd {
					next = append(next, overlapStart-interval.source+interval.destination, overlapEnd-interval.source+interval.destination)
					if overlapStart > start {
						result = append(result, start, overlapStart)
					}
					if end > overlapEnd {
						result = append(result, overlapEnd, end)
					}
					found = true
					break
				}
			}
			if !found {
				next = append(next, start, end)
			}
		}
		result = next
	}
	return slices.Min(result)
}

func main() {
	filename := "day05/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
