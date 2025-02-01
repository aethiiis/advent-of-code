package main

import (
	"2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	dir utils.Direction
	num int
}

var directions = map[string]utils.Direction{"U": {-1, 0}, "D": {1, 0}, "L": {0, -1}, "R": {0, 1}}
var hex = map[string]string{"0": "R", "1": "D", "2": "L", "3": "U"}

func processing(filename string) ([]Instruction, []Instruction) {
	instructions, rgb := make([]Instruction, 0), make([]Instruction, 0)
	for _, line := range strings.Split(utils.ReadFile(filename), "\n") {
		parts := strings.Split(line, " ")
		dir := directions[parts[0]]
		num, _ := strconv.Atoi(parts[1])
		rgbDir := directions[hex[string(parts[2][len(parts[2])-2])]]
		rgbNum, _ := strconv.ParseInt(parts[2][2:len(parts[2])-2], 16, 64)
		instructions, rgb = append(instructions, Instruction{dir: dir, num: num}), append(rgb, Instruction{dir: rgbDir, num: int(rgbNum)})
	}
	return instructions, rgb
}

func calculate(instructions []Instruction) int {
	points, boundary := []utils.Pos{{0, 0}}, 0
	for _, instruction := range instructions {
		boundary += instruction.num
		points = append(points, utils.Pos{
			R: points[len(points)-1].R + instruction.dir.Dr*instruction.num,
			C: points[len(points)-1].C + instruction.dir.Dc*instruction.num,
		})
	}
	return utils.Shoelace(points) - boundary/2 + 1 + boundary
}

func part1(filename string) int {
	instructions, _ := processing(filename)
	return calculate(instructions)
}

func part2(filename string) int {
	_, instructions := processing(filename)
	return calculate(instructions)
}

func main() {
	filename := "day18/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
