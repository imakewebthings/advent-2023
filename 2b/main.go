package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	gameId := 1
	answer := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		rounds := strings.Split(parts[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0

		for _, round := range rounds {
			pulls := strings.Split(round, ", ")
			for _, pull := range pulls {
				pullParts := strings.Split(pull, " ")
				num, _ := strconv.Atoi(pullParts[0])
				color := pullParts[1]
				if color == "red" {
					maxRed = max(maxRed, num)
				} else if color == "green" {
					maxGreen = max(maxGreen, num)
				} else if color == "blue" {
					maxBlue = max(maxBlue, num)
				}
			}
		}

		answer = answer + maxRed*maxGreen*maxBlue
		gameId++
	}

	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
