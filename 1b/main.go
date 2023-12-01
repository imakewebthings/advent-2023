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

	/*
		While not fastest, it would be simplest to replace the words with their digits
		and continue using the solution from Part 1. Doing this straight up creates an
		issue where multiple words share characters. Ex: "twone", which if evaluating
		from the left becomes "2ne" and if from the right becomes "tw1". Still, being
		stubborn, rather than change my approach, keeping the word form on both sides
		of the digit works around stealing characters from other words in
		either direction.
	*/
	words := map[string]string{
		"one":   "one1one",
		"two":   "two2two",
		"three": "three3three",
		"four":  "four4four",
		"five":  "five5five",
		"six":   "six6six",
		"seven": "seven7seven",
		"eight": "eight8eight",
		"nine":  "nine9nine",
	}
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		for word, num := range words {
			line = strings.ReplaceAll(line, word, num)
		}
		fmt.Println(line)

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
