package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const Cycles int = 1_000_000_000

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := [][]byte{}
	for scanner.Scan() {
		rows = append(rows, []byte(scanner.Text()))
	}

	platMap := map[string]int{}
	platArr := []string{}
	i := 0
	answer := 0
	for answer == 0 {
		cycle(rows)
		key := platStr(rows)
		platArr = append(platArr, key)
		if platMap[key] != 0 {
			answerIndex := ((Cycles - platMap[key] - 1) % (i - platMap[key])) + platMap[key]
			answer = northLoad(platArr[answerIndex])
		}
		platMap[key] = i
		i++
	}

	fmt.Println(answer)
}

func tiltNorth(rows [][]byte) {
	for x := 0; x < len(rows[0]); x++ {
		cur := 0
		for y := 0; y < len(rows); y++ {
			char := rows[y][x]
			if char == '#' {
				cur = y + 1
			} else if char == 'O' {
				rows[cur][x] = 'O'
				if cur != y {
					rows[y][x] = '.'
				}
				cur++
			}
		}
	}
}

func tiltWest(rows [][]byte) {
	for y := 0; y < len(rows); y++ {
		cur := 0
		for x := 0; x < len(rows[0]); x++ {
			char := rows[y][x]
			if char == '#' {
				cur = x + 1
			} else if char == 'O' {
				rows[y][cur] = 'O'
				if cur != x {
					rows[y][x] = '.'
				}
				cur++
			}
		}
	}
}

func tiltSouth(rows [][]byte) {
	for x := 0; x < len(rows[0]); x++ {
		cur := len(rows) - 1
		for y := len(rows) - 1; y >= 0; y-- {
			char := rows[y][x]
			if char == '#' {
				cur = y - 1
			} else if char == 'O' {
				rows[cur][x] = 'O'
				if cur != y {
					rows[y][x] = '.'
				}
				cur--
			}
		}
	}
}

func tiltEast(rows [][]byte) {
	for y := 0; y < len(rows); y++ {
		cur := len(rows[0]) - 1
		for x := len(rows[0]) - 1; x >= 0; x-- {
			char := rows[y][x]
			if char == '#' {
				cur = x - 1
			} else if char == 'O' {
				rows[y][cur] = 'O'
				if cur != x {
					rows[y][x] = '.'
				}
				cur--
			}
		}
	}
}

func northLoad(platKey string) int {
	rows := strings.Split(platKey, "\n")
	total := 0
	for i, row := range rows {
		for _, char := range row {
			if char == 'O' {
				total += len(rows) - i
			}
		}
	}
	return total
}

func cycle(rows [][]byte) {
	tiltNorth(rows)
	tiltWest(rows)
	tiltSouth(rows)
	tiltEast(rows)
}

func platStr(rows [][]byte) string {
	strs := make([]string, len(rows))
	for i, row := range rows {
		strs[i] = string(row)
	}
	return strings.Join(strs, "\n")
}
