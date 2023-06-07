package main

import (
	"fmt"
	"strings"

	"github.com/tyler569/aoc-go"
)

func main() {
	input := strings.Trim(aoc.Input(2022, 3), "\n")

	sacks := strings.Split(input, "\n")

	totalPriority := 0
	for _, line := range sacks {
		item := mismatchedItem(line)
		totalPriority += priority(item)
	}

	badgePriority := 0
	for i := 0; i < len(sacks)/3; i += 1 {
		group := sacks[3*i : 3*i+3]
		item := groupBadge(group)
		badgePriority += priority(item)
	}

	fmt.Println("1:", totalPriority)
	fmt.Println("2:", badgePriority)
}

func priority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 'a' + 1
	}
	if item >= 'A' && item <= 'Z' {
		return int(item) - 'A' + 27
	}
	panic("invalid item")
}

func mismatchedItem(sack string) rune {
	l := len(sack)
	c1 := sack[:l/2]
	c2 := sack[l/2:]

	for _, c := range c1 {
		if strings.ContainsRune(c2, c) {
			return c
		}
	}
	panic("no mismatched item")
}

func groupBadge(group []string) rune {
	for _, item := range group[0] {
		if strings.ContainsRune(group[1], item) && strings.ContainsRune(group[2], item) {
			return item
		}
	}
	panic("no group badge item")
}
