package main

import (
	"math"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	sol1(lines)
	sol2(lines)
}

func sol2(lines []string) {
	cardCount := make(map[int]int)
	for i := 1; i <= len(lines); i++ {
		cardCount[i] = 1
	}
	for i, line := range lines {
		_, card, _ := strings.Cut(line, ": ")
		winningNums, guesses := parseCard(card)
		winners := 0
		for k := range guesses {
			if winningNums[k] {
				winners++
			}
		}
		factor := cardCount[i+1]
		for j := 1; j <= winners; j++ {
			cardCount[i+1+j] += factor
		}
	}
	sum := 0
	for _, v := range cardCount {
		sum += v
	}
	println(sum)
}

func sol1(lines []string) {
	winningSum := 0
	for _, line := range lines {
		_, card, _ := strings.Cut(line, ": ")
		winningNums, guesses := parseCard(card)
		count := -1
		for k := range guesses {
			if winningNums[k] {
				count++
			}
		}
		if count > -1 {
			winningSum += int(math.Pow(2, float64(count)))
		}
	}
	println(winningSum)
}

func parseCard(card string) (winningNums, guesses map[string]bool) {
	winningNumList, guessList, _ := strings.Cut(card, " | ")
	winningNums = make(map[string]bool)
	for _, n := range strings.Split(winningNumList, " ") {
		if len(n) > 0 {
			winningNums[n] = true
		}
	}
	guesses = make(map[string]bool)
	for _, n := range strings.Split(guessList, " ") {
		if len(n) > 0 {
			guesses[n] = true
		}
	}
	return
}
