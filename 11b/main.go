package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct{ X, Y int }

const ExpansionRate = 1000000

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	galaxies := []Point{}
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, Point{X: x, Y: y})
			}
		}
		y++
	}

	sort.Slice(galaxies, func(a, b int) bool {
		return galaxies[a].Y < galaxies[b].Y
	})
	expansion := 0
	for i := 1; i < len(galaxies); i++ {
		galaxies[i].Y += expansion
		diff := (galaxies[i].Y-galaxies[i-1].Y-1)*ExpansionRate - 1
		if diff < 0 {
			diff = 0
		}
		expansion += diff
		galaxies[i].Y += diff
	}

	sort.Slice(galaxies, func(a, b int) bool {
		return galaxies[a].X < galaxies[b].X
	})
	expansion = 0
	for i := 1; i < len(galaxies); i++ {
		galaxies[i].X += expansion
		diff := (galaxies[i].X-galaxies[i-1].X-1)*ExpansionRate - 1
		if diff < 0 {
			diff = 0
		}
		expansion += diff
		galaxies[i].X += diff
	}

	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			xDiff := galaxies[j].X - galaxies[i].X
			yDiff := galaxies[j].Y - galaxies[i].Y
			if yDiff < 0 {
				yDiff = yDiff * -1
			}
			sum += xDiff + yDiff
		}
	}

	fmt.Println(sum)
}
