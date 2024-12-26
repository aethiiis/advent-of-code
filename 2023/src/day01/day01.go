package main

import (
	"2023/src/utils"
	"fmt"
	"strings"
)

func processing(filename string) []string {
	text := utils.ReadFile(filename)
	lines := strings.SplitN(text, "\n", -1)
	return lines
}

func calc(line string) int {
	return 10*int(line[strings.IndexAny(line, "0123456789")]-'0') + int(line[strings.LastIndexAny(line, "0123456789")]-'0')
}

func part1(filename string) int {
	var result int
	lines := processing(filename)
	for _, line := range lines {
		result += calc(line)
	}
	return result
}

func part2(filename string) int {
	var result int
	lines := processing(filename)
	replacer := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four", "f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e", "zero", "z0o")
	for _, line := range lines {
		result += calc(replacer.Replace(replacer.Replace(line)))
	}
	return result
}

func main() {
	filename := "src/day01/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
