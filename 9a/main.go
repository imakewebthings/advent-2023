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

	histories := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		numStrs := strings.Fields(line)
		history := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			history[i], _ = strconv.Atoi(numStr)
		}
		histories = append(histories, history)
	}

	answer := 0

	for _, history := range histories {
		ds := [][]int{history}

		for !samesies(ds[len(ds)-1]) {
			cur := ds[len(ds)-1]
			next := []int{}
			for i := 1; i < len(cur); i++ {
				next = append(next, cur[i]-cur[i-1])
			}
			ds = append(ds, next)
		}

		for i := len(ds) - 2; i >= 0; i-- {
			proj := ds[i][len(ds[i])-1] + ds[i+1][len(ds[i+1])-1]
			ds[i] = append(ds[i], proj)
		}

		answer += ds[0][len(ds[0])-1]
	}

	fmt.Println(answer)
}

func samesies(nums []int) bool {
	num := nums[0]
	for _, n := range nums {
		if num != n {
			return false
		}
	}
	return true
}
