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
		possible := true
		line := scanner.Text()
		parts := strings.Split(line, ": ")
		rounds := strings.Split(parts[1], "; ")

		for _, round := range rounds {
			pulls := strings.Split(round, ", ")
			for _, pull := range pulls {
				pullParts := strings.Split(pull, " ")
				num, _ := strconv.Atoi(pullParts[0])
				color := pullParts[1]
				if color == "red" && num > 12 ||
					color == "green" && num > 13 ||
					color == "blue" && num > 14 {
					possible = false
				}
			}
		}

		if possible {
			answer += gameId
		}
		gameId++
	}

	fmt.Println(answer)
}
