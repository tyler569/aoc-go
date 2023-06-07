package main

import (
	"fmt"
	"strings"

	"github.com/tyler569/aoc-go"
)

func main() {
	input := strings.Trim(aoc.Input(2022, 5), "\n")
	parts := strings.Split(input, "\n\n")
	pilesString := parts[0]
	instructionsString := parts[1]

	piles := parsePiles(pilesString)
	instructions := strings.Split(instructionsString, "\n")
	moveCrates(piles, instructions)
	fmt.Println("1:", getTops(piles))

	piles = parsePiles(pilesString)
	moveCrates2(piles, instructions)
	fmt.Println("2:", getTops(piles))
}

func parsePiles(pilesString string) [][]byte {
	lines := strings.Split(pilesString, "\n")
	lines = lines[:len(lines)-1]
	nPiles := (len(lines[0]) + 1) / 4

	scratch := [][]byte{}

	for _, line := range lines {
		scratchN := []byte{}
		for i := 0; i < nPiles; i++ {
			scratchN = append(scratchN, line[i*4+1])
		}
		scratch = append(scratch, scratchN)
	}

	piles := [][]byte{}

	for p := 0; p < len(scratch[0]); p += 1 {
		pile := []byte{}
		for c := len(scratch) - 1; c >= 0; c -= 1 {
			if scratch[c][p] != ' ' {
				pile = append(pile, scratch[c][p])
			}
		}
		piles = append(piles, pile)
	}

	return piles
}

func removeLast(pile []byte) ([]byte, byte) {
	last := pile[len(pile)-1]
	pile = pile[:len(pile)-1]

	return pile, last
}

func moveCrates(piles [][]byte, instructions []string) {
	var last byte
	for _, instr := range instructions {
		var count, from, to int
		fmt.Sscanf(instr, "move %v from %v to %v", &count, &from, &to)
		from -= 1
		to -= 1

		for i := 0; i < count; i += 1 {
			piles[from], last = removeLast(piles[from])
			piles[to] = append(piles[to], last)
		}
	}
}

func removeLastN(pile []byte, n int) ([]byte, []byte) {
	end := pile[len(pile)-n:]
	pile = pile[:len(pile)-n]
	return pile, end
}

func moveCrates2(piles [][]byte, instructions []string) {
	var last []byte
	for _, instr := range instructions {
		var count, from, to int
		fmt.Sscanf(instr, "move %v from %v to %v", &count, &from, &to)
		from -= 1
		to -= 1

		piles[from], last = removeLastN(piles[from], count)
		piles[to] = append(piles[to], last...)
	}
}

func getTops(piles [][]byte) string {
	var result []byte

	for _, pile := range piles {
		last := pile[len(pile)-1]
		result = append(result, last)
	}

	return string(result)
}
