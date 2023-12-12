package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		groups := []int{}
		for i := 0; i < 5; i++ {
			for _, numStr := range strings.Split(parts[1], ",") {
				num, _ := strconv.Atoi(numStr)
				groups = append(groups, num)
			}
		}
		record := strings.Join([]string{parts[0], parts[0], parts[0], parts[0], parts[0]}, "?")
		rowPoss := permute(record, groups)
		total += rowPoss
	}

	fmt.Println(total)
}

var cache map[string]map[string]int = map[string]map[string]int{}

func cacheGet(record string, groups []int) int {
	gMap, rOk := cache[record]
	if !rOk {
		return -1
	}
	if val, ok := gMap[cacheKey(groups)]; ok {
		return val
	}
	return -1
}

func cacheSet(record string, groups []int, val int) int {
	gMap, ok := cache[record]
	if !ok {
		cache[record] = map[string]int{}
		gMap = cache[record]
	}
	gMap[cacheKey(groups)] = val
	return val
}

func cacheKey(groups []int) string {
	return fmt.Sprintf("%v", groups)
}

func permute(record string, groups []int) int {
	g := groups[0]
	cacheVal := cacheGet(record, groups)
	if cacheVal > -1 {
		return cacheVal
	}

	for i, char := range record {
		if char == '.' {
			continue
		}
		if char == '?' {
			r := permute("."+record[i+1:], groups) + permute("#"+record[i+1:], groups)
			return cacheSet(record, groups, r)
		}
		if i+g > len(record) {
			return cacheSet(record, groups, 0)
		}
		match := !strings.ContainsRune(record[i:i+g], '.')
		if match {
			if i+g < len(record) && record[i+g] == '#' {
				return cacheSet(record, groups, 0)
			}

			if len(groups) == 1 {
				if strings.ContainsRune(record[i+g:], '#') {
					return cacheSet(record, groups, 0)
				} else {
					return cacheSet(record, groups, 1)
				}
			}

			if i+g+1 > len(record)-1 {
				return cacheSet(record, groups, 0)
			}
			return cacheSet(record, groups, permute(record[i+g+1:], groups[1:]))
		} else {
			return cacheSet(record, groups, 0)
		}
	}

	return cacheSet(record, groups, 0)
}
