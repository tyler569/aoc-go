package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tyler569/aoc-go"

	"golang.org/x/exp/slices"
)

func main() {
	input := aoc.Input(2022, 1)

	elves := elves(input)

	fmt.Println("1:", part1(elves))
	fmt.Println("2:", part2(elves))
}

func elves(input string) (elves []int) {
	packs := strings.Split(input, "\n\n")
	for _, pack := range packs {
		items := strings.Fields(pack)

		calories := 0

		for _, item := range items {
			itemCalories, err := strconv.Atoi(item)
			if err != nil {
				panic(err)
			}
			calories += itemCalories
		}

		elves = append(elves, calories)
	}

	slices.SortFunc(elves, func(a int, b int) bool {
		return a > b
	})

	return
}

func part1(elves []int) int {
	return elves[0]
}

func part2(elves []int) int {
	return elves[0] + elves[1] + elves[2]
}
