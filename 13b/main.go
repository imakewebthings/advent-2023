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
		hVal := reflection(horizontals[i]) * 100
		total += hVal
		if hVal == 0 {
			total += reflection(verticals[i])
		}
	}

	fmt.Println(total)
}

func reflection(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		j := i - 1
		k := i
		matching := true
		usedSmudge := false
		for matching && j >= 0 && k < len(pattern) {
			hd := hamming(pattern[j], pattern[k])
			if hd > 1 || (hd == 1 && usedSmudge) {
				matching = false
			} else if hd == 1 && !usedSmudge {
				usedSmudge = true
			}
			j--
			k++
		}
		if !usedSmudge {
			continue
		}
		if matching {
			return i
		}
	}
	return 0
}

func hamming(s1, s2 string) int {
	d := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			d++
		}
	}
	return d
}
