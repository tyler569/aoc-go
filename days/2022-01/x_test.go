package main

import "testing"

const sample1 = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

var elvesCalories = elves(sample1)

func TestPart1(t *testing.T) {
	if part1(elvesCalories) != 24000 {
		t.Fatal("Top elf should be carrying 24000 calories")
	}
}

func TestPart2(t *testing.T) {
	if part2(elvesCalories) != 45000 {
		t.Fatal("Top 3 elves should be carrying 45000 calories")
	}
}
