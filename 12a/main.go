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

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		groups := []int{}
		for _, numStr := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(numStr)
			groups = append(groups, num)
		}
		rowPoss := permute(parts[0], groups)
		total += rowPoss
	}

	fmt.Println(total)
}

func permute(record string, groups []int) int {
	g := groups[0]

	for i, char := range record {
		if char == '.' {
			continue
		}
		if char == '?' {
			return permute("."+record[i+1:], groups) + permute("#"+record[i+1:], groups)
		}
		if i+g > len(record) {
			return 0
		}
		match := !strings.ContainsRune(record[i:i+g], '.')
		if match {
			if i+g < len(record) && record[i+g] == '#' {
				return 0
			}

			if len(groups) == 1 {
				if strings.ContainsRune(record[i+g:], '#') {
					return 0
				} else {
					return 1
				}
			}

			if i+g+1 > len(record)-1 {
				return 0
			}
			return permute(record[i+g+1:], groups[1:])
		} else {
			return 0
		}
	}

	return 0
}
