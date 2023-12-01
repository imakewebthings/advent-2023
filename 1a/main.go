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

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		num := 0

		for i := 0; i < len(line); i++ {
			if line[i]-'0' >= 0 && line[i]-'0' <= 9 {
				num = int(line[i]-'0') * 10
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			if line[i]-'0' >= 0 && line[i]-'0' <= 9 {
				num += int(line[i] - '0')
				break
			}
		}

		total += num
	}

	fmt.Println(total)
}
