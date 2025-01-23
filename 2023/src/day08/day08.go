package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

func processing(filename string) ([]int, map[string][]string, []string) {
	instructionsNodes := strings.SplitN(utils.ReadFile(filename), "\n\n", -1)
	instructionsText := instructionsNodes[0]
	nodesLines := strings.SplitN(instructionsNodes[1], "\n", -1)
	instructions := make([]int, len(instructionsText))
	nodes := make(map[string][]string)
	starts := make([]string, 0)
	ends := make([]string, 0)
	for i, char := range instructionsText {
		if char == 76 {
			instructions[i] = 0
		} else if char == 82 {
			instructions[i] = 1
		}
	}
	for _, line := range nodesLines {
		nodeNext := strings.Split(line, " = ")
		leftRight := strings.Split(strings.Trim(nodeNext[1], "()"), ", ")
		nodes[nodeNext[0]] = []string{leftRight[0], leftRight[1]}
		if strings.HasSuffix(nodeNext[0], "A") {
			starts = append(starts, nodeNext[0])
		}
		if strings.HasSuffix(nodeNext[0], "Z") {
			ends = append(ends, nodeNext[0])
		}
	}
	return instructions, nodes, starts
}

func part1(filename string) int {
	instructions, nodes, _ := processing(filename)
	start, end := "AAA", "ZZZ"
	current := start
	n := len(instructions)
	i := 0
	for current != end {
		current = nodes[current][instructions[i%n]]
		i++
	}
	return i
}

func part2(filename string) int {
	var result int
	instructions, nodes, starts := processing(filename)
	for _, start := range starts {
		current := start
		n := len(instructions)
		i := 0
		for !strings.HasSuffix(current, "Z") {
			current = nodes[current][instructions[i%n]]
			i++
		}
		if result == 0 {
			result = i
		} else {
			result = utils.Lcm(result, i)
		}
	}
	return result
}

func main() {
	filename := "day08/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
