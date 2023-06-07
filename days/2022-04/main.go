package main

import (
	"fmt"
	"strings"

	"github.com/tyler569/aoc-go"
)

func main() {
	input := strings.Trim(aoc.Input(2022, 4), "\n")
	lines := strings.Split(input, "\n")

	containsCount := 0
	overlapsCount := 0
	for _, line := range lines {
		var x1, x2, y1, y2 int
		fmt.Sscanf(line, "%v-%v,%v-%v", &x1, &x2, &y1, &y2)

		if x1 <= y1 && x2 >= y2 || y1 <= x1 && y2 >= x2 {
			containsCount += 1
		}

		if x1 <= y2 && y1 <= x2 {
			overlapsCount += 1
		}
	}

	fmt.Println("1:", containsCount)
	fmt.Println("2:", overlapsCount)
}
