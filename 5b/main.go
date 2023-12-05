package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	seedLine := scanner.Text()
	seeds := []Range{}
	start := 0
	for i, s := range strings.Split(strings.Split(seedLine, ": ")[1], " ") {
		num, _ := strconv.Atoi(s)
		if i%2 == 0 {
			start = num
		} else {
			seeds = append(seeds, Range{
				Start: start,
				End:   start + num - 1,
			})
		}
	}
	scanner.Scan()
	scanner.Scan()

	diffMap := map[Range]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if strings.Contains(line, "map:") {
			seeds = applyMap(seeds, diffMap)
			diffMap = map[Range]int{}
			continue
		}
		parts := strings.Split(line, " ")
		dest, _ := strconv.Atoi(parts[0])
		start, _ := strconv.Atoi(parts[1])
		length, _ := strconv.Atoi(parts[2])
		diff := dest - start
		diffMap[Range{Start: start, End: start + length - 1}] = diff
	}
	seeds = applyMap(seeds, diffMap)

	min := seeds[0].Start
	for _, seed := range seeds {
		if seed.Start < min {
			min = seed.Start
		}
	}
	fmt.Println(min)
}

func applyMap(seeds []Range, diffMap map[Range]int) []Range {
	for i := 0; i < len(seeds); i++ {
		seed := seeds[i]
		for r, diff := range diffMap {
			// No intersection
			if seed.End < r.Start || seed.Start > r.End {
				continue
			}
			// Seed range completely within
			if seed.Start >= r.Start && seed.End <= r.End {
				seeds[i].Start += diff
				seeds[i].End += diff
				break
			}
			// Partial overlap with start of map range
			if seed.Start < r.Start && seed.End <= r.End {
				overhang := Range{
					Start: seed.Start,
					End:   r.Start - 1,
				}
				seeds = append(seeds, overhang)
				seeds[i].Start = r.Start
				seeds[i].Start += diff
				seeds[i].End += diff
				break
			}
			// Partial overlap with end of map range
			if seed.Start >= r.Start && seed.End > r.End {
				overhang := Range{
					Start: r.End + 1,
					End:   seed.End,
				}
				seeds = append(seeds, overhang)
				seeds[i].End = r.End
				seeds[i].Start += diff
				seeds[i].End += diff
				break
			}
			// Map range completely within seed range, overlap both sides
			seeds = append(seeds, Range{
				Start: seed.Start,
				End:   r.Start - 1,
			})
			seeds = append(seeds, Range{
				Start: r.End + 1,
				End:   seed.End,
			})
			seeds[i].Start = r.Start
			seeds[i].End = r.End
			seeds[i].Start += diff
			seeds[i].End += diff
			break
		}
	}
	return seeds
}
