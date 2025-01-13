package main

import (
	"2023/src/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards []int
	class int
	bid   int
}

var symbol2Strength = map[int32]int{
	50: 1,  // 2
	51: 2,  // 3
	52: 3,  // 4
	53: 4,  // 5
	54: 5,  // 6
	55: 6,  // 7
	56: 7,  // 8
	57: 8,  // 9
	84: 9,  // T
	74: 10, // J
	81: 11, // Q
	75: 12, // K
	65: 13, // A
}

func processing(filename string) []Hand {
	hands := make([]Hand, 0)
	lines := strings.SplitN(utils.ReadFile(filename), "\n", -1)
	for _, line := range lines {
		var hand Hand
		handBid := strings.Split(line, " ")
		cards := handBid[0]
		hand.bid, _ = strconv.Atoi(handBid[1])
		for _, card := range cards {
			hand.cards = append(hand.cards, symbol2Strength[card])
		}
		hands = append(hands, hand)
	}
	return hands
}

func class(hand Hand) int {
	counts := utils.Map2(hand.cards, func(card int) int {
		return utils.Count(hand.cards, card)
	})
	if utils.Contains(counts, 5) {
		return 6
	} else if utils.Contains(counts, 4) {
		return 5
	} else if utils.Contains(counts, 3) {
		if utils.Contains(counts, 2) {
			return 4
		} else {
			return 3
		}
	} else if utils.Count(counts, 2) == 4 {
		return 2
	} else if utils.Contains(counts, 2) {
		return 1
	} else {
		return 0
	}
}

func winnings(hands []Hand) int {
	var count int
	for rank, hand := range hands {
		count += (rank + 1) * hand.bid
	}
	return count
}

func part1(filename string) int {
	hands := processing(filename)
	for i, hand := range hands {
		hands[i].class = class(hand)
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].class != hands[j].class {
			return hands[i].class < hands[j].class
		} else {
			for k := range hands[i].cards {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
			return false
		}
	})
	return winnings(hands)
}

func part2(filename string) int {
	hands := processing(filename)
	for i := range hands {
		hands[i].class = class(hands[i])
		indexes := utils.FindAll(hands[i].cards, 10)
		if len(indexes) != 0 {
			if hands[i].class == 6 {
			} else if hands[i].class == 5 {
				hands[i].class += 1
			} else if hands[i].class == 4 {
				hands[i].class += 2
			} else if hands[i].class == 3 {
				hands[i].class += 2
			} else if hands[i].class == 2 {
				if len(indexes) == 2 {
					hands[i].class += 3
				} else {
					hands[i].class += 2
				}
			} else if hands[i].class == 1 {
				hands[i].class += 2
			} else {
				hands[i].class += 1
			}
		}
		for _, index := range indexes {
			hands[i].cards[index] = 0
		}
	}
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].class != hands[j].class {
			return hands[i].class < hands[j].class
		} else {
			for k := range hands[i].cards {
				if hands[i].cards[k] != hands[j].cards[k] {
					return hands[i].cards[k] < hands[j].cards[k]
				}
			}
			return false
		}
	})
	return winnings(hands)
}

func main() {
	filename := "day07/input.txt"
	fmt.Println(part1(filename))
	fmt.Println(part2(filename))
}
