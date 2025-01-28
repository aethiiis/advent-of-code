package main

import (
	"2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Record struct {
	springs    []byte
	constraint []int
}

type State struct {
	constraintIndex, countConsecutive, dotExpected int
}

func processing(filename string) []Record {
	records := make([]Record, 0)
	for _, line := range strings.SplitN(utils.ReadFile(filename), "\n", -1) {
		springs, constraint, _ := strings.Cut(line, " ")
		records = append(records, Record{springs: []byte(springs), constraint: utils.Map2(strings.SplitN(constraint, ",", -1), func(str string) int {
			num, _ := strconv.Atoi(str)
			return num
		})})
	}
	return records
}

func count(record Record) int {
	var res int
	current := map[State]int{State{0, 0, 0}: 1}
	next := map[State]int{}
	for _, springStatus := range record.springs {
		for state, num := range current {
			constraintIndex, countConsecutive, dotExpected := state.constraintIndex, state.countConsecutive, state.dotExpected
			switch {
			case (springStatus == '#' || springStatus == '?') && constraintIndex < len(record.constraint) && dotExpected == 0:
				if springStatus == '?' && countConsecutive == 0 {
					next[State{constraintIndex, countConsecutive, dotExpected}] += num
				}
				countConsecutive++
				if countConsecutive == record.constraint[constraintIndex] {
					constraintIndex, countConsecutive, dotExpected = constraintIndex+1, 0, 1
				}
				next[State{constraintIndex, countConsecutive, dotExpected}] += num
			case (springStatus == '.' || springStatus == '?') && countConsecutive == 0:
				dotExpected = 0
				next[State{constraintIndex, countConsecutive, dotExpected}] += num
			}
		}
		current, next = next, current
		clear(next)
	}
	for s, v := range current {
		if s.constraintIndex == len(record.constraint) {
			res += v
		}
	}
	return res
}

func part1(filename string) int {
	return utils.Reduce(processing(filename), count, 0)
}

func part2(filename string) int {
	return utils.Reduce(utils.Map2(processing(filename), func(record Record) Record {
		newSprings := make([]byte, 0)
		newConstraint := make([]int, 0)
		for i := 0; i < 5; i++ {
			newSprings = append(newSprings, record.springs...)
			newSprings = append(newSprings, byte('?'))
			newConstraint = append(newConstraint, record.constraint...)
		}
		return Record{
			springs:    newSprings[:len(newSprings)-1],
			constraint: newConstraint,
		}
	}), count, 0)
}

func main() {
	filename := "day12/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
