package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := []string{}
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	cols := make([][]rune, len(rows[0]))
	for i := 0; i < len(rows[0]); i++ {
		cols[i] = []rune{}
	}
	for _, row := range rows {
		for j, char := range row {
			cols[j] = append(cols[j], char)
		}
	}

	total := 0
	for _, col := range cols {
		cur := 0
		colSum := 0
		for j, char := range col {
			if char == '#' {
				cur = j + 1
			} else if char == 'O' {
				colSum += len(col) - cur
				cur++
			}
		}
		total += colSum
	}

	fmt.Println(total)
}
