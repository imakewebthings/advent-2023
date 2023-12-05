package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type RangeMap struct {
	DestStart   int
	SourceStart int
	Length      int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seedLine := scanner.Text()
	seeds := []int{}
	for _, s := range strings.Split(strings.Split(seedLine, ": ")[1], " ") {
		seedNum, _ := strconv.Atoi(s)
		seeds = append(seeds, seedNum)
	}
	scanner.Scan()
	scanner.Scan()

	maps := [][]RangeMap{}
	var curMap []RangeMap
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map:") {
			maps = append(maps, curMap)
			curMap = []RangeMap{}
			continue
		}
		parts := strings.Split(line, " ")
		m := RangeMap{}
		m.DestStart, _ = strconv.Atoi(parts[0])
		m.SourceStart, _ = strconv.Atoi(parts[1])
		m.Length, _ = strconv.Atoi(parts[2])
		curMap = append(curMap, m)
	}
	maps = append(maps, curMap)

	cur := 0
	min := math.MaxInt
	for _, seed := range seeds {
		cur = seed
		for _, rangeMaps := range maps {
			for _, rm := range rangeMaps {
				if rm.SourceStart <= cur && cur <= rm.SourceStart+rm.Length {
					cur = cur - rm.SourceStart + rm.DestStart
					break
				}
			}
		}
		if cur < min {
			min = cur
		}
	}

	fmt.Println(min)
}
