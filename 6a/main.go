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

	scanner.Scan()
	line := scanner.Text()
	timeStrs := strings.Fields(strings.Split(line, ":")[1])
	scanner.Scan()
	line = scanner.Text()
	distanceStrs := strings.Fields(strings.Split(line, ":")[1])
	times := []int{}
	for _, str := range timeStrs {
		num, _ := strconv.Atoi(str)
		times = append(times, num)
	}
	distances := []int{}
	for _, str := range distanceStrs {
		num, _ := strconv.Atoi(str)
		distances = append(distances, num)
	}

	answer := 1
	for i, time := range times {
		ways := 0
		for down := 1; down < time; down++ {
			if down*(time-down) > distances[i] {
				ways++
			}
		}
		answer *= ways
	}

	fmt.Println(answer)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
