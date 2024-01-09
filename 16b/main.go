package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Point struct {
	X int
	Y int
}

type Beam struct {
	Pos Point
	Dir Direction
}

func (b *Beam) Move(board []string) (*Beam, bool) {
	switch b.Dir {
	case East:
		b.Pos.X++
	case West:
		b.Pos.X--
	case North:
		b.Pos.Y--
	case South:
		b.Pos.Y++
	}

	// Out of bounds, beam done
	if b.Pos.X < 0 || b.Pos.X >= len(board[0]) || b.Pos.Y < 0 || b.Pos.Y >= len(board) {
		return nil, true
	}

	space := board[b.Pos.Y][b.Pos.X]

	// Split em
	if (b.Dir == West || b.Dir == East) && space == '|' {
		split := &Beam{
			Pos: b.Pos,
			Dir: South,
		}
		b.Dir = North
		return split, false
	}
	if (b.Dir == North || b.Dir == South) && space == '-' {
		split := &Beam{
			Pos: b.Pos,
			Dir: West,
		}
		b.Dir = East
		return split, false
	}

	// Reflect em
	if b.Dir == West && space == '/' || b.Dir == East && space == '\\' {
		b.Dir = South
	} else if b.Dir == West && space == '\\' || b.Dir == East && space == '/' {
		b.Dir = North
	} else if b.Dir == North && space == '/' || b.Dir == South && space == '\\' {
		b.Dir = East
	} else if b.Dir == North && space == '\\' || b.Dir == South && space == '/' {
		b.Dir = West
	}

	return nil, false
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	board := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		board = append(board, line)
	}

	winner := 0
	for x := 0; x < len(board[0]); x++ {
		winner = max(winner, run(board, x, 0, South))
		winner = max(winner, run(board, x, len(board)-1, North))
	}
	for y := 0; y < len(board); y++ {
		winner = max(winner, run(board, 0, y, East))
		winner = max(winner, run(board, len(board[0])-1, y, West))
	}

	fmt.Println(winner)
}

func run(board []string, startX, startY int, startDir Direction) int {
	beams := []*Beam{
		&Beam{
			Pos: Point{
				X: startX,
				Y: startY,
			},
			Dir: startDir,
		},
	}

	switch startDir {
	case East:
		beams[0].Pos.X--
	case West:
		beams[0].Pos.X++
	case South:
		beams[0].Pos.Y--
	case North:
		beams[0].Pos.Y++
	}

	energized := map[Point]map[Direction]bool{}

	for len(beams) > 0 {
		for i := len(beams) - 1; i >= 0; i-- {
			beam := beams[i]
			split, done := beam.Move(board)
			if done || (energized[beam.Pos] != nil && energized[beam.Pos][beam.Dir]) {
				beams = append(beams[:i], beams[i+1:]...)
			} else {
				if energized[beam.Pos] == nil {
					energized[beam.Pos] = map[Direction]bool{}
				}
				energized[beam.Pos][beam.Dir] = true
			}
			if split != nil {
				beams = append(beams, split)
			}
		}
	}

	return len(energized)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
