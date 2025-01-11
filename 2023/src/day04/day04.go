package main

import (
	"2023/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	winning []int
	numbers []int
}

func processing(filename string) []Card {
	var list []Card
	lines := strings.SplitN(utils.ReadFile(filename), "\n", -1)
	for _, line := range lines {
		idRest := strings.Split(strings.Replace(strings.Replace(line, "  ", " ", -1), "  ", " ", -1), ": ")
		id, _ := strconv.Atoi(strings.Split(idRest[0], " ")[1])
		winningNumbers := strings.Split(idRest[1], " | ")
		winning := utils.Map2(strings.SplitN(winningNumbers[0], " ", -1), func(str string) int {
			number, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			return number
		})
		numbers := utils.Map2(strings.SplitN(winningNumbers[1], " ", -1), func(str string) int {
			number, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			return number
		})
		list = append(list, Card{id: id, winning: winning, numbers: numbers})
	}
	return list
}

func getNumber(card Card) int {
	var number int
	for _, win := range card.winning {
		if utils.Contains(card.numbers, win) {
			number++
		}
	}
	return number
}

func part1(filename string) int {
	var result int
	cards := processing(filename)
	for _, card := range cards {
		number := getNumber(card)
		if number > 0 {
			result += 1 << (number - 1)
		}
	}
	return result
}

func part2(filename string) int {
	var result int
	cards := processing(filename)
	counter := make(map[int]int)
	for _, card := range cards {
		if _, exist := counter[card.id]; !exist {
			counter[card.id] = 1
		} else {
			counter[card.id] += 1
		}
		result += counter[card.id]
		number := getNumber(card)
		for count := 0; count < counter[card.id]; count++ {
			for i := 1; i <= number; i++ {
				counter[card.id+i]++
			}
		}
	}
	return result
}

func main() {
	filename := "day04/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
