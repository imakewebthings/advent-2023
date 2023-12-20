package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Lens struct {
	Label    string
	Strength int
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	line := scanner.Text()
	steps := strings.Split(line, ",")

	boxes := make([][]*Lens, 256)
	for _, step := range steps {
		if strings.ContainsRune(step, '-') {
			label := step[:len(step)-1]
			remove(boxes, label)
		} else {
			ls := strings.Split(step, "=")
			strength, _ := strconv.Atoi(ls[1])
			set(boxes, ls[0], strength)
		}
	}

	fmt.Println(focalPower(boxes))
}

func hash(step string) int {
	val := 0
	for _, char := range step {
		val += int(char)
		val *= 17
		val = val % 256
	}
	return val
}

func remove(boxes [][]*Lens, label string) {
	key := hash(label)
	box := boxes[key]
	for i, lens := range box {
		if lens.Label == label {
			boxes[key] = append(boxes[key][:i], boxes[key][i+1:]...)
			return
		}
	}
}

func set(boxes [][]*Lens, label string, strength int) {
	key := hash(label)
	box := boxes[key]
	for _, lens := range box {
		if lens.Label == label {
			lens.Strength = strength
			return
		}
	}
	boxes[key] = append(boxes[key], &Lens{
		Label:    label,
		Strength: strength,
	})
}

func focalPower(boxes [][]*Lens) int {
	val := 0
	for i, box := range boxes {
		for j, lens := range box {
			val += (i + 1) * (j + 1) * lens.Strength
		}
	}
	return val
}
