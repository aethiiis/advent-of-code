package main

import (
	"2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Lens struct {
	label string
	focal int
}

type Box []Lens

var registry = map[string]int{}

func processing(filename string) []string {
	return strings.Split(utils.ReadFile(filename), ",")
}

func hash(word string) int {
	var result int
	for _, ch := range word {
		result = (result + int(ch)) * 17 % 256
	}
	return result
}

func update(label string) {
	if _, ok := registry[label]; !ok {
		registry[label] = hash(label)
	}
}

func dash(boxes *map[int]Box, label string) {
	number := registry[label]
	box, ok := (*boxes)[number]
	if !ok {
		return
	}
	for i, lens := range box {
		if lens.label == label {
			box = append(box[:i], box[i+1:]...)
			(*boxes)[number] = box
			break
		}
	}
}

func equal(boxes *map[int]Box, label string, focal int) {
	number := registry[label]
	present := false
	box, ok := (*boxes)[number]
	if !ok {
		(*boxes)[number] = Box{Lens{label: label, focal: focal}}
		return
	}
	for i, lens := range box {
		if lens.label == label {
			box[i].focal = focal
			present = true
			break
		}
	}
	if !present {
		box = append(box, Lens{label: label, focal: focal})
	}
	(*boxes)[number] = box
}

func power(boxes map[int]Box) int {
	pow := 0
	for number, box := range boxes {
		for pos, lens := range box {
			pow += (number + 1) * (pos + 1) * lens.focal
		}
	}
	return pow
}

func part1(filename string) int {
	return utils.Sum(utils.Map2(processing(filename), hash))
}

func part2(filename string) int {
	boxes := make(map[int]Box)
	for _, word := range processing(filename) {
		if strings.Contains(word, "=") {
			label, digits, _ := strings.Cut(word, "=")
			update(label)
			focal, _ := strconv.Atoi(digits)
			equal(&boxes, label, focal)
		} else {
			label, _, _ := strings.Cut(word, "-")
			update(label)
			dash(&boxes, label)
		}
	}
	return power(boxes)
}

func main() {
	filename := "day15/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
