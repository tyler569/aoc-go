package main

import (
	"testing"
)

const sample = `
A Y
B X
C Z
`

func TestPart1(t *testing.T) {
	if part1(sample) != 15 {
		t.Fail()
	}
}


func TestPart2(t *testing.T) {
	if part2(sample) != 12 {
		t.Fail()
	}
}
