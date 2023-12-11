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

	path := map[[2]int]bool{}
	cur, dir, startPiece := first(rows, start)
	for cur != start {
		cur, dir = next(rows, path, cur, dir)
	}

	path[start] = true
	rows[start[0]][start[1]] = startPiece
	normalize(rows, path)

	count := 0
	for _, row := range rows {
		borders := 0
		for _, char := range row {
			if char == '|' || char == 'J' || char == 'L' {
				borders++
				continue
			}
			if char == '.' && borders%2 == 1 {
				count++
			}
		}
	}

	fmt.Println(count)
}

func next(rows [][]byte, path map[[2]int]bool, cur [2]int, dir byte) ([2]int, byte) {
	defer func() {
		path[cur] = true
	}()
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

func first(rows [][]byte, start [2]int) ([2]int, byte, byte) {
	var next [2]int
	var dir byte

	look := rows[start[0]-1][start[1]]
	connections := map[byte]bool{}
	if look == 'F' || look == '7' || look == '|' {
		next, dir = [2]int{start[0] - 1, start[1]}, 'U'
		connections['U'] = true
	}
	look = rows[start[0]+1][start[1]]
	if look == 'L' || look == 'J' || look == '|' {
		next, dir = [2]int{start[0] + 1, start[1]}, 'D'
		connections['D'] = true
	}
	look = rows[start[0]][start[1]-1]
	if look == 'L' || look == 'F' || look == '-' {
		next, dir = [2]int{start[0], start[1] - 1}, 'L'
		connections['L'] = true
	}
	look = rows[start[0]][start[1]+1]
	if look == '7' || look == 'J' || look == '-' {
		next, dir = [2]int{start[0], start[1] + 1}, 'R'
		connections['R'] = true
	}

	var startPiece byte
	if connections['U'] && connections['D'] {
		startPiece = '|'
	} else if connections['L'] && connections['R'] {
		startPiece = '-'
	} else if connections['L'] && connections['D'] {
		startPiece = '7'
	} else if connections['L'] && connections['U'] {
		startPiece = 'J'
	} else if connections['R'] && connections['D'] {
		startPiece = 'F'
	} else if connections['R'] && connections['U'] {
		startPiece = 'L'
	}
	return next, dir, startPiece
}

func normalize(rows [][]byte, path map[[2]int]bool) {
	for y, row := range rows {
		for x, _ := range row {
			if !path[[2]int{y, x}] {
				rows[y][x] = '.'
			}
		}
	}
}
