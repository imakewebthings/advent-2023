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
			if row[x] == '*' {
				answer += gearRatio(rows, y, x)
			}
		}
	}

	fmt.Println(answer)
}

func isDigit(b byte) bool {
	return b-'0' < 10 && b-'0' >= 0
}

func gearRatio(rows []string, y, x int) int {
	nums := []int{}

	if x > 0 && isDigit(rows[y][x-1]) {
		nums = append(nums, extractNum(rows, y, x-1))
	}
	if x < len(rows[0])-1 && isDigit(rows[y][x+1]) {
		nums = append(nums, extractNum(rows, y, x+1))
	}

	if y > 0 {
		if isDigit(rows[y-1][x]) {
			nums = append(nums, extractNum(rows, y-1, x))
		} else {
			if x > 0 && isDigit(rows[y-1][x-1]) {
				nums = append(nums, extractNum(rows, y-1, x-1))
			}
			if x < len(rows[0])-1 && isDigit(rows[y-1][x+1]) {
				nums = append(nums, extractNum(rows, y-1, x+1))
			}
		}
	}

	if y < len(rows)-1 {
		if isDigit(rows[y+1][x]) {
			nums = append(nums, extractNum(rows, y+1, x))
		} else {
			if x > 0 && isDigit(rows[y+1][x-1]) {
				nums = append(nums, extractNum(rows, y+1, x-1))
			}
			if x < len(rows[0])-1 && isDigit(rows[y+1][x+1]) {
				nums = append(nums, extractNum(rows, y+1, x+1))
			}
		}
	}

	if len(nums) < 2 {
		return 0
	}
	return nums[0] * nums[1]
}

func extractNum(rows []string, y, x int) int {
	for x >= 0 && isDigit(rows[y][x]) {
		x--
	}
	x++
	bNum := []byte{}
	for x < len(rows[y]) && isDigit(rows[y][x]) {
		bNum = append(bNum, rows[y][x])
		x++
	}
	num, _ := strconv.Atoi(string(bNum))
	return num
}
