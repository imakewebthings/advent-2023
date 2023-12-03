package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		rows = append(rows, line)
	}

	answer := 0

	for y, row := range rows {
		for x := 0; x < len(row); x++ {
			if isDigit(row[x]) {
				bNum := []byte{}
				start := x - 1
				for x < len(row) && isDigit(row[x]) {
					bNum = append(bNum, row[x])
					x++
				}
				if symbolAdjacent(rows, y, start, x) {
					num, _ := strconv.Atoi(string(bNum))
					answer += num
				}
			}
		}
	}

	fmt.Println(answer)
}

func isDigit(b byte) bool {
	return b-'0' < 10 && b-'0' >= 0
}

func symbolAdjacent(rows []string, y, start, end int) bool {
	checkLeft, checkRight := true, true

	if start < 0 {
		start = 0
		checkLeft = false
	}
	if end >= len(rows[0]) {
		end = len(rows[0]) - 1
		checkRight = false
	}

	if y > 0 {
		for x := start; x <= end; x++ {
			if rows[y-1][x] != '.' {
				return true
			}
		}
	}

	if checkLeft && rows[y][start] != '.' {
		return true
	}
	if checkRight && rows[y][end] != '.' {
		return true
	}

	if y < len(rows)-1 {
		for x := start; x <= end; x++ {
			if rows[y+1][x] != '.' {
				return true
			}
		}
	}

	return false
}
