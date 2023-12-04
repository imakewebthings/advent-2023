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

	won := []int{}
	count := []int{}
	idx := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, ": ")[1]
		sides := strings.Split(nums, " | ")

		won = append(won, 0)
		count = append(count, 1)
		winners := map[string]bool{}
		for _, winner := range strings.Fields(sides[0]) {
			winners[winner] = true
		}

		for _, num := range strings.Fields(sides[1]) {
			if winners[num] {
				won[idx]++
			}
		}

		idx++
	}

	total := 0
	for i, n := range won {
		for c := 0; c < count[i]; c++ {
			for j := 1; j <= n; j++ {
				count[i+j]++
			}
		}
		total += count[i]
	}

	fmt.Println(total)
}
