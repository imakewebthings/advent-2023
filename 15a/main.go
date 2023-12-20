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
	scanner.Scan()
	line := scanner.Text()
	steps := strings.Split(line, ",")

	total := 0
	for _, step := range steps {
		total += hash(step)
	}

	fmt.Println(total)
}

func hash(step string) int {
	val := 0
	for _, char := range step {
		val += int(char)
		val *= 17
		val = val % 256
	}
	return val
}
