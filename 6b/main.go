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
	timeStr := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")
	scanner.Scan()
	line = scanner.Text()
	distanceStr := strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")

	time, _ := strconv.Atoi(timeStr)
	distance, _ := strconv.Atoi(distanceStr)
	ways := 0

	for down := 1; down < time; down++ {
		if down*(time-down) > distance {
			ways++
		}
	}

	fmt.Println(ways)
}
