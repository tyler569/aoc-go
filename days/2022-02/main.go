package main

import (
	"fmt"
	"strings"

	"github.com/tyler569/aoc-go"
)

func main() {
	input := aoc.Input(2022, 2)

	fmt.Println("1:", part1(input))
	fmt.Println("2:", part2(input))
}

func part1(input string) int {
	totalScore := 0
	lines := strings.Split(strings.Trim(input, " \n"), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)

		their := toPlay(parts[0])
		my := toPlay(parts[1])

		totalScore += score(their, my)
	}

	return totalScore
}

func part2(input string) int {
	totalScore := 0
	lines := strings.Split(strings.Trim(input, " \n"), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)

		their, my := toPlay2(parts[0], parts[1])

		totalScore += score(their, my)
	}

	return totalScore
}

//go:generate stringer -type=play
type play int

const (
	rock play = iota
	scissors
	paper
)

func (p play) result(other play) int {
	if p == other {
		return 0
	}

	if p == rock && other == scissors ||
		p == paper && other == rock ||
		p == scissors && other == paper {

		return 1
	} else {
		return -1
	}
}

func (p play) score() int {
	switch p {
	case rock:
		return 1
	case paper:
		return 2
	case scissors:
		return 3
	}
	panic("invalid play")
}

func toPlay(p string) play {
	switch p {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	case "X":
		return rock
	case "Y":
		return paper
	case "Z":
		return scissors
	}
	panic("invalid play")
}

func toPlay2(p, o string) (play, play) {
	var their play
	switch p {
	case "A":
		their = rock
	case "B":
		their = paper
	case "C":
		their = scissors
	default:
		panic("invalid play")
	}

	var result int
	switch o {
	case "X":
		result = -1
	case "Y":
		result = 0
	case "Z":
		result = 1
	default:
		panic("invalid play")
	}

	for p := rock; p <= paper; p += 1 {
		if p.result(their) == result {
			return their, p
		}
	}

	panic("unreachable")
}

func score(their, my play) int {
	total := my.score()
	switch my.result(their) {
	case 1:
		total += 6
	case 0:
		total += 3
	}
	return total
}
