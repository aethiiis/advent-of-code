package main

import (
	"2023/src/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Part struct {
	x, m, a, s int
}

func processing(filename string) (map[string]string, []Part) {
	text := utils.ReadFile(filename)
	wf, pa, _ := strings.Cut(text, "\n\n")
	parts := make([]Part, 0)
	workflows := make(map[string]string)
	re := regexp.MustCompile(`\d+`)
	for _, part := range strings.Split(pa, "\n") {
		numbers := utils.Map2(re.FindAllString(part, -1), func(s string) int {
			i, _ := strconv.Atoi(s)
			return i
		})
		parts = append(parts, Part{numbers[0], numbers[1], numbers[2], numbers[3]})
	}
	for _, workflow := range strings.Split(wf, "\n") {
		name, rest, _ := strings.Cut(workflow, "{")
		workflows[name] = rest[:len(rest)-1]
	}
	return workflows, parts
}

func eval(workflows map[string]string, part Part, name string) bool {
	possibilities := strings.Split(workflows[name], ",")
	for _, possibility := range possibilities {
		if possibility == "R" {
			return false
		}
		if possibility == "A" {
			return true
		}
		condition, destination, _ := strings.Cut(possibility, ":")
		variable, op := condition[0], condition[1]
		number, _ := strconv.Atoi(condition[2:])
		verified := false
		if destination == "" {
			return eval(workflows, part, condition)
		}
		switch op {
		case '>':
			switch variable {
			case 'x':
				if part.x > number {
					verified = true
				}
			case 'm':
				if part.m > number {
					verified = true
				}
			case 'a':
				if part.a > number {
					verified = true
				}
			case 's':
				if part.s > number {
					verified = true
				}
			}
		case '<':
			switch variable {
			case 'x':
				if part.x < number {
					verified = true
				}
			case 'm':
				if part.m < number {
					verified = true
				}
			case 'a':
				if part.a < number {
					verified = true
				}
			case 's':
				if part.s < number {
					verified = true
				}
			}
		}
		if verified {
			if destination == "R" {
				return false
			}
			if destination == "A" {
				return true
			}
			return eval(workflows, part, destination)
		}
	}
	panic("No possibility found")
}

func both(variable uint8, gt bool, number int, ranges [][][]int) [][][]int {
	i := strings.Index("xmas", string(variable))
	r := make([][][]int, 0)
	for _, rng := range ranges {
		lo, hi := rng[i][0], rng[i][1]
		if gt {
			lo = max(lo, number+1)
		} else {
			hi = min(hi, number-1)
		}
		if lo > hi {
			continue
		}
		rng[i] = []int{lo, hi}
		r = append(r, rng)
	}
	return r
}

func outer(workflows map[string]string, name string) [][][]int {
	return inner(workflows, strings.Split(workflows[name], ","))
}

func inner(workflows map[string]string, possibilities []string) [][][]int {
	possibility := possibilities[0]
	if possibility == "R" {
		return [][][]int{}
	}
	if possibility == "A" {
		return [][][]int{{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}}
	}
	condition, destination, _ := strings.Cut(possibility, ":")
	if destination == "" {
		return outer(workflows, possibility)
	}
	variable, gt := condition[0], strings.Contains(condition, ">")
	number, _ := strconv.Atoi(condition[2:])
	var valInverted int
	if gt {
		valInverted = number + 1
	} else {
		valInverted = number - 1
	}
	verified := both(variable, gt, number, inner(workflows, []string{destination}))
	notVerified := both(variable, !gt, valInverted, inner(workflows, possibilities[1:]))
	return append(verified, notVerified...)
}

func part1(filename string) int {
	var result int
	workflows, parts := processing(filename)
	for _, part := range parts {
		if eval(workflows, part, "in") {
			result += part.x + part.m + part.a + part.s
		}
	}
	return result
}

func part2(filename string) int {
	var result int
	workflows, _ := processing(filename)
	ranges := outer(workflows, "in")
	for _, rng := range ranges {
		product := 1
		for _, r := range rng {
			product *= r[1] - r[0] + 1
		}
		result += product
	}
	return result
}

func main() {
	filename := "day19/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
