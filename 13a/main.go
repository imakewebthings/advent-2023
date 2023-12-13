package main

import (
	"fmt"
	"os"
	"strings"
)

var debug bool = false

func main() {
	bInput, _ := os.ReadFile("input.txt")
	input := string(bInput)

	horizontals := [][]string{}
	patterns := strings.Split(input, "\n\n")
	for _, pattern := range patterns {
		h := strings.Split(pattern, "\n")
		if h[len(h)-1] == "" {
			h = h[:len(h)-1]
		}
		horizontals = append(horizontals, h)
	}

	verticals := make([][]string, len(horizontals))
	for i, horizontal := range horizontals {
		verticals[i] = make([]string, len(horizontal[0]))
		for _, row := range horizontal {
			for j, char := range row {
				verticals[i][j] = string([]rune{char}) + verticals[i][j]
			}
		}
	}

	total := 0
	for i, _ := range patterns {
		total += reflection(horizontals[i]) * 100
		total += reflection(verticals[i])
	}

	fmt.Println(total)
}

func reflection(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		j := i - 1
		k := i
		matching := true
		for matching && j >= 0 && k < len(pattern) {
			matching = pattern[j] == pattern[k]
			j--
			k++
		}
		if matching {
			return i
		}
	}
	return 0
}
