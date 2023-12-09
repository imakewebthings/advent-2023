package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

var nodeRe = regexp.MustCompile(`^([A-Z|0-9]{3}) = \(([A-Z|0-9]{3}), ([A-Z|0-9]{3})\)$`)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	directions := scanner.Text()
	scanner.Scan()

	lefts := map[string]string{}
	rights := map[string]string{}
	currents := []string{}
	endMoves := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := nodeRe.FindStringSubmatch(line)
		source, left, right := parts[1], parts[2], parts[3]
		lefts[source] = left
		rights[source] = right
		if source[2] == 'A' {
			currents = append(currents, source)
		}
	}

	moves := 0
	i := 0
	for len(endMoves) != len(currents) {
		for j, current := range currents {
			if directions[i] == 'L' {
				currents[j] = lefts[current]
			} else {
				currents[j] = rights[current]
			}
		}
		moves++
		i++
		if i == len(directions) {
			i = 0
		}
		for _, current := range currents {
			if current[2] == 'Z' {
				endMoves = append(endMoves, moves)
			}
		}
	}

	multiples := make([]int, len(endMoves))
	copy(multiples, endMoves)

	for !equal(multiples) {
		x := smallest(multiples)
		multiples[x] += endMoves[x]
	}

	fmt.Println(multiples[0])
}

func noEmpties(slices [][]int) bool {
	for _, slices := range slices {
		if len(slices) == 0 {
			return false
		}
	}
	return true
}

func equal(nums []int) bool {
	num := nums[0]
	for _, n := range nums[1:] {
		if num != n {
			return false
		}
	}
	return true
}

func smallest(nums []int) int {
	min := math.MaxInt
	i := -1
	for j, num := range nums {
		if num < min {
			min = num
			i = j
		}
	}
	return i
}
