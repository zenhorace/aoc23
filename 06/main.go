package main

import (
	"strconv"
	"strings"
	"time"
)

type record struct {
	totalTime, minDist, minTime, maxTime int
}

// TODO: add the constant time solution using the quadratic equation

func main() {
	sol1()
	st := time.Now()
	sol2()
	println("time for pt2", time.Since(st).Milliseconds(), "ms")
}

func sol1() {
	var records []record
	times := strings.Split(timeStr, " ")
	distances := strings.Split(distanceStr, " ")
	for i := 0; i < len(times); i++ {
		t, _ := strconv.Atoi(times[i])
		d, _ := strconv.Atoi(distances[i])
		records = append(records, record{totalTime: t, minDist: d})
	}
	count := 1
	for _, r := range records {
		r.min()
		r.max()
		count *= r.ways()
	}
	println("pt1", count)
}

func sol2() {
	r := record{totalTime: 46828479, minDist: 347152214061471}
	r.min()
	r.max()
	println("pt2", r.ways())
}

func (r *record) min() {
	for i := 1; i < r.totalTime; i++ {
		if i*(r.totalTime-i) > r.minDist {
			r.minTime = i
			return
		}
	}
}

func (r *record) max() {
	for i := r.totalTime - 1; i > 1; i-- {
		if i*(r.totalTime-i) > r.minDist {
			r.maxTime = i
			return
		}
	}
}

func (r *record) ways() int {
	return r.maxTime - (r.minTime - 1)
}
