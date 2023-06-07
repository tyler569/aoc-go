package main

import (
	"fmt"
	"strings"
	"testing"
)

const sample = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

var sampleParts = strings.Split(strings.Trim(sample, "\n"), "\n\n")

var samplePiles = sampleParts[0]
var sampleInstrs = sampleParts[1]

func printPiles(t *testing.T, piles [][]byte) {
	for _, pile := range piles {
		var s string
		for _, c := range pile {
			s += fmt.Sprintf("[%c] ", c)
		}
		t.Log(s)
	}
}

func TestPart1(t *testing.T) {
	piles := parsePiles(samplePiles)
	printPiles(t, piles)

	for _, instr := range strings.Split(sampleInstrs, "\n") {
		t.Log(instr)
		moveCrates(piles, []string{instr})
		printPiles(t, piles)
	}

	result := getTops(piles)

	t.Log(result)

	if result != "CMZ" {
		t.Fail()
	}
}

func TestPart2(t *testing.T) {
	piles := parsePiles(samplePiles)
	printPiles(t, piles)

	for _, instr := range strings.Split(sampleInstrs, "\n") {
		t.Log(instr)
		moveCrates2(piles, []string{instr})
		printPiles(t, piles)
	}

	result := getTops(piles)

	t.Log(result)

	if result != "MCD" {
		t.Fail()
	}
}
