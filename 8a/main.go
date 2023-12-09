package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

var nodeRe = regexp.MustCompile(`^([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)$`)

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

	for scanner.Scan() {
		line := scanner.Text()
		parts := nodeRe.FindStringSubmatch(line)
		source, left, right := parts[1], parts[2], parts[3]
		lefts[source] = left
		rights[source] = right
	}

	moves := 0
	i := 0
	current := "AAA"
	for current != "ZZZ" {
		if directions[i] == 'L' {
			current = lefts[current]
		} else {
			current = rights[current]
		}
		moves++
		i++
		if i == len(directions) {
			i = 0
		}
	}

	fmt.Println(moves)
}
