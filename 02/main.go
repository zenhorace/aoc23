package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	limit = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
)

func main() {
	sol1()
	sol2()
}

func sol2() {
	lines := strings.Split(input1, "\n")
	var imp []int
	for _, l := range lines {
		line := strings.TrimLeft(l, "Game ")
		_, setStr, ok := strings.Cut(line, ":")
		if !ok {
			fmt.Println("Error parsing line: ", line)
			os.Exit(1)
		}

		imp = append(imp, powerSet(setStr))
	}
	sum := 0
	for _, i := range imp {
		sum += i
	}
	fmt.Println(sum)
}

func powerSet(setStr string) int {
	hwm := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}
	sets := strings.Split(setStr, ";")
	for _, s := range sets {
		for _, p := range strings.Split(s, ",") {
			countStr, color, _ := strings.Cut(strings.Trim(p, " "), " ")
			count, _ := strconv.Atoi(countStr)
			if count > hwm[color] {
				hwm[color] = count
			}
		}
	}
	return hwm["red"] * hwm["green"] * hwm["blue"]
}

func sol1() {
	lines := strings.Split(input1, "\n")
	var imp []int
	for _, l := range lines {
		line := strings.TrimLeft(l, "Game ")
		gameStr, setStr, ok := strings.Cut(line, ":")
		if !ok {
			fmt.Println("Error parsing line: ", line)
			os.Exit(1)
		}
		game, err := strconv.Atoi(gameStr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if possible(setStr) {
			imp = append(imp, game)
		}
	}
	sum := 0
	for _, i := range imp {
		sum += i
	}
	fmt.Println(sum)
}

func possible(setStr string) bool {
	sets := strings.Split(setStr, ";")
	for _, s := range sets {
		for _, p := range strings.Split(s, ",") {
			countStr, color, _ := strings.Cut(strings.Trim(p, " "), " ")
			count, _ := strconv.Atoi(countStr)
			if count > limit[color] {
				return false
			}
		}
	}
	return true
}
