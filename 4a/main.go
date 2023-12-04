package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ": ")[1]
		sides := strings.Split(nums, " | ")

		winners := map[string]bool{}
		for _, winner := range strings.Fields(sides[0]) {
			winners[winner] = true
		}

		value := 0
		for _, num := range strings.Fields(sides[1]) {
			if winners[num] {
				if value == 0 {
					value = 1
				} else {
					value = value * 2
				}
			}
		}

		total += value
	}

	fmt.Println(total)
}
