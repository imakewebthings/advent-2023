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

	rows := [][]byte{}
	start := [2]int{-1, -1}

	for scanner.Scan() {
		line := scanner.Text()
		sX := strings.IndexRune(line, 'S')
		if sX > -1 {
			start[0] = len(rows)
			start[1] = sX
		}
		rows = append(rows, []byte(line))
	}

	cur, dir := first(rows, start)
	moves := 1

	for cur != start {
		cur, dir = next(rows, cur, dir)
		moves++
	}

	fmt.Println(moves / 2)
}

func next(rows [][]byte, cur [2]int, dir byte) ([2]int, byte) {
	switch rows[cur[0]][cur[1]] {
	case '|':
		if dir == 'U' {
			return [2]int{cur[0] - 1, cur[1]}, 'U'
		}
		return [2]int{cur[0] + 1, cur[1]}, 'D'
	case '-':
		if dir == 'L' {
			return [2]int{cur[0], cur[1] - 1}, 'L'
		}
		return [2]int{cur[0], cur[1] + 1}, 'R'
	case 'J':
		if dir == 'R' {
			return [2]int{cur[0] - 1, cur[1]}, 'U'
		}
		return [2]int{cur[0], cur[1] - 1}, 'L'
	case '7':
		if dir == 'R' {
			return [2]int{cur[0] + 1, cur[1]}, 'D'
		}
		return [2]int{cur[0], cur[1] - 1}, 'L'
	case 'L':
		if dir == 'L' {
			return [2]int{cur[0] - 1, cur[1]}, 'U'
		}
		return [2]int{cur[0], cur[1] + 1}, 'R'
	case 'F':
		if dir == 'L' {
			return [2]int{cur[0] + 1, cur[1]}, 'D'
		}
		return [2]int{cur[0], cur[1] + 1}, 'R'
	default:
		return [2]int{-1, -1}, 'X'
	}
}

func first(rows [][]byte, start [2]int) ([2]int, byte) {
	look := rows[start[0]-1][start[1]]
	if look == 'F' || look == '7' || look == '|' {
		return [2]int{start[0] - 1, start[1]}, 'U'
	}
	look = rows[start[0]+1][start[1]]
	if look == 'L' || look == 'J' || look == '|' {
		return [2]int{start[0] + 1, start[1]}, 'D'
	}
	look = rows[start[0]][start[1]-1]
	if look == 'L' || look == 'F' || look == '-' {
		return [2]int{start[0], start[1] - 1}, 'L'
	}
	look = rows[start[0]][start[1]+1]
	if look == '7' || look == 'J' || look == '-' {
		return [2]int{start[0], start[1] + 1}, 'R'
	}
	return [2]int{-1, -1}, 'X'
}
