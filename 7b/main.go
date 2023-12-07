package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	HighCard int = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

var labels = map[byte]int{
	'J': -1,
	'2': 0,
	'3': 1,
	'4': 2,
	'5': 3,
	'6': 4,
	'7': 5,
	'8': 6,
	'9': 7,
	'T': 8,
	'Q': 10,
	'K': 11,
	'A': 12,
}

type HandBid struct {
	Hand     string
	Bid      int
	HandType int
}

type ByCamelRules []HandBid

func (r ByCamelRules) Len() int {
	return len(r)
}

func (r ByCamelRules) Less(a, b int) bool {
	if r[a].HandType == r[b].HandType {
		for i := 0; i < 5; i++ {
			aChar, bChar := r[a].Hand[i], r[b].Hand[i]
			if aChar != bChar {
				return labels[aChar] < labels[bChar]
			}
		}
		return false
	}
	return r[a].HandType < r[b].HandType
}

func (r ByCamelRules) Swap(a, b int) {
	r[a], r[b] = r[b], r[a]
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	hbs := []HandBid{}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		bid, _ := strconv.Atoi(parts[1])
		hbs = append(hbs, HandBid{
			Hand: parts[0],
			Bid:  bid,
		})
	}

	for i, hb := range hbs {
		cardMap := map[rune]int{}
		cardKeys := []rune{}
		for _, card := range hb.Hand {
			if card != 'J' && cardMap[card] == 0 {
				cardKeys = append(cardKeys, card)
			}
			cardMap[card]++
		}
		if cardMap['J'] == 5 {
			cardMap['A'] = 5
			cardKeys = append(cardKeys, 'A')
		} else if cardMap['J'] > 0 {
			moveTo := maxKey(cardMap)
			cardMap[moveTo] += cardMap['J']
		}
		if len(cardKeys) == 1 {
			hbs[i].HandType = FiveOfAKind
		} else if len(cardKeys) == 2 {
			if cardMap[cardKeys[0]] == 1 || cardMap[cardKeys[0]] == 4 {
				hbs[i].HandType = FourOfAKind
			} else {
				hbs[i].HandType = FullHouse
			}
		} else if len(cardKeys) == 3 {
			if cardMap[cardKeys[0]] == 3 || cardMap[cardKeys[1]] == 3 || cardMap[cardKeys[2]] == 3 {
				hbs[i].HandType = ThreeOfAKind
			} else {
				hbs[i].HandType = TwoPair
			}
		} else if len(cardKeys) == 4 {
			hbs[i].HandType = OnePair
		} else {
			hbs[i].HandType = HighCard
		}
	}

	sort.Sort(ByCamelRules(hbs))

	answer := 0
	for i, hb := range hbs {
		answer += hb.Bid * (i + 1)
	}
	fmt.Println(answer)
}

func maxKey(m map[rune]int) rune {
	curMax := 0
	curR := 'A'
	for r, count := range m {
		if r == 'J' {
			continue
		}
		if count > curMax {
			curMax = count
			curR = r
		}
	}
	return curR
}
