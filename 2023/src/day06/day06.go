package main

import (
	"2023/src/utils"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func processing(filename string) ([][]int, []int) {
	tab1 := make([][]int, 2)
	tab2 := make([]int, 2)
	timeDistance := strings.Split(utils.ReadFile(filename), "\n")
	timeText := strings.Replace(strings.Replace(strings.Replace(strings.Trim(timeDistance[0][9:], " "), "  ", " ", -1), "  ", " ", -1), "  ", " ", -1)
	distanceText := strings.Replace(strings.Replace(strings.Replace(strings.Trim(timeDistance[1][9:], " "), "  ", " ", -1), "  ", " ", -1), "  ", " ", -1)
	tab1[0] = utils.Map2(strings.SplitN(timeText, " ", -1), func(str string) int {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return number
	})
	tab1[1] = utils.Map2(strings.SplitN(distanceText, " ", -1), func(str string) int {
		number, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		return number
	})
	tab2[0], _ = strconv.Atoi(strings.Replace(timeText, " ", "", -1))
	tab2[1], _ = strconv.Atoi(strings.Replace(distanceText, " ", "", -1))
	return tab1, tab2
}

func part1(filename string) int {
	result := 1
	tab, _ := processing(filename)
	races := utils.Zip(tab[0], tab[1])
	for _, race := range races {
		count := 0
		for hold := 0; hold < race.First; hold++ {
			if hold*(race.First-hold) > race.Second {
				count++
			}
		}
		result *= count
	}
	return result
}

func part2(filename string) int {
	_, tab := processing(filename)
	delta := tab[0]*tab[0] - 4*tab[1]
	if delta == 0 {
		return 1
	} else if delta > 0 {
		x1 := int((-float64(tab[0]) + math.Sqrt(float64(delta))) / 2)
		x2 := int((-float64(tab[0]) - math.Sqrt(float64(delta))) / 2)
		return x1 - x2
	}
	return 0
}

func main() {
	filename := "day06/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
