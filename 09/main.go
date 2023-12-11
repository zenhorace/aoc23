package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input, "\n")
	endSum := 0
	beginSum := 0
	for _, line := range lines {
		fields := strings.Fields(line)
		seq := toInt(fields)
		var track [][]int
		track = append(track, seq)
		for !isZero(track[len(track)-1]) {
			diff := make([]int, len(seq)-1)
			for i := 1; i < len(seq); i++ {
				diff[i-1] = seq[i] - seq[i-1]
			}
			track = append(track, diff)
			seq = diff
		}
		nextVal := 0
		prevVal := 0
		for i := len(track) - 2; i >= 0; i-- {
			nextVal += track[i][len(track[i])-1]
			prevVal = track[i][0] - prevVal
		}
		endSum += nextVal
		beginSum += prevVal
	}
	fmt.Println(beginSum, "|", endSum)
}

func isZero(seq []int) bool {
	for _, v := range seq {
		if v != 0 {
			return false
		}
	}
	return true
}

func toInt(s []string) []int {
	ints := make([]int, len(s))
	var err error
	for i, v := range s {
		ints[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal("umm", err)
		}
	}
	return ints
}
