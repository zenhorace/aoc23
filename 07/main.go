package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

// To get the answer to part 1, set J's strength back to 11 and remove the call to complexCardType.

var (
	strength = map[byte]int{
		'A': 14,
		'K': 13,
		'Q': 12,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
)

type hand struct {
	bid   int
	cards []byte
	cType int
}

type byHands []hand

func (a byHands) Len() int      { return len(a) }
func (a byHands) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byHands) Less(i, j int) bool {
	if a[i].cType < a[j].cType {
		return true
	}
	if a[i].cType > a[j].cType {
		return false
	}
	for k := 0; k < 5; k++ {
		if strength[a[i].cards[k]] < strength[a[j].cards[k]] {
			return true
		}
		if strength[a[i].cards[k]] > strength[a[j].cards[k]] {
			return false
		}
	}
	return false
}

func main() {
	lines := strings.Split(input, "\n")
	hands := make([]hand, len(lines))
	for i, line := range lines {
		hands[i] = parseHand(line)
	}
	sort.Sort(byHands(hands))
	sum := 0
	for i, h := range hands {
		sum += (i + 1) * h.bid
	}
	println(sum)
}

func parseHand(line string) hand {
	var h hand
	var err error
	f := strings.Fields(line)
	h.bid, err = strconv.Atoi(f[1])
	if err != nil {
		log.Fatal(err)
	}
	h.cards = []byte(f[0])
	h.cType = complexCardType(cardType(h.cards))
	return h
}

// upgrades type based on Jokers
func complexCardType(t int, m map[byte]int) int {
	jCount := m['J']
	if jCount == 0 {
		return t
	}
	if jCount >= 4 || len(m) == 2 {
		return fiveOfAKind
	}
	if jCount == 3 {
		return fourOfAKind
	}
	if jCount == 2 {
		if len(m) == 3 { // two pair
			return fourOfAKind
		}
		return threeOfAKind // there were 4 different cards
	}

	// only 1 J
	switch t {
	case highCard:
		return onePair
	case onePair:
		return threeOfAKind
	case twoPair:
		return fullHouse
	case threeOfAKind:
		return fourOfAKind
	}
	return t // should never happen
}

func cardType(cards []byte) (int, map[byte]int) {
	m := make(map[byte]int)
	for _, c := range cards {
		m[c]++
	}
	if len(m) == 5 {
		return highCard, m
	}
	if len(m) == 4 {
		return onePair, m
	}
	if len(m) == 3 {
		for _, v := range m {
			if v == 3 {
				return threeOfAKind, m
			}
		}
		return twoPair, m
	}
	if len(m) == 2 {
		for _, v := range m {
			if v == 4 {
				return fourOfAKind, m
			}
		}
		return fullHouse, m
	}
	return fiveOfAKind, m
}
