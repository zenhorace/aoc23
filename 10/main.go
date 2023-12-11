package main

import (
	"fmt"
	"log"
	"strings"
)

var (
	maze [][]byte

	north = delta{0, -1}
	south = delta{0, 1}
	east  = delta{1, 0}
	west  = delta{-1, 0}

	movements = map[byte][]delta{
		'|': {north, south},
		'-': {east, west},
		'L': {north, east},
		'J': {north, west},
		'7': {south, west},
		'F': {east, south},
	}

	inverse = map[delta]delta{
		north: south,
		south: north,
		east:  west,
		west:  east,
	}
)

type delta struct {
	x, y int
}

type point struct {
	x, y int
}

func (p point) add(d delta) point {
	return point{p.x + d.x, p.y + d.y}
}

func main() {
	var start point
	lines := strings.Split(mazeStr, "\n")
	for i, line := range lines {
		maze = append(maze, []byte(line))
		for j, c := range line {
			if c == 'S' {
				start = point{j, i}
			}
		}
	}
	swapS(start) // hard coded by manaual inspection but could be done programmatically
	hwm := bfs(start)
	fmt.Println("pt 1", hwm)
	path := walk(start, south) // approaching from the south means we walk counter clockwise
	fmt.Println("pt 2", shoelace(path))
}

// swap S for its correct pipe
func swapS(s point) {
	var sDirs []delta
	for _, mv := range []delta{north, east, south, west} {
		pt := s.add(mv)
		if t, ok := movements[maze[pt.y][pt.x]]; ok {
			if t[0] == inverse[mv] || t[1] == inverse[mv] {
				sDirs = append(sDirs, mv)
			}
		}
	}
	if len(sDirs) != 2 {
		log.Fatal("bad s", len(sDirs))
	}
	for k, v := range movements { // I can do this because I list directions clockwise
		if v[0] == sDirs[0] && v[1] == sDirs[1] {
			maze[s.y][s.x] = k
			return
		}
	}
}

func bfs(start point) int {
	hwm := 0
	dist := make(map[point]int)
	dist[start] = 0
	queue := []point{start}
	for len(queue) > 0 {
		pt := queue[0]
		queue = queue[1:]
		for _, b := range branches(pt) {
			if _, ok := dist[b]; !ok {
				dist[b] = dist[pt] + 1
				queue = append(queue, b)
				if dist[b] > hwm {
					hwm = dist[b]
				}
			}
		}
	}
	return hwm
}

func branches(start point) []point {
	var res []point
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			x, y := start.x+i, start.y+j
			if x < 0 || y < 0 || x >= len(maze[0]) || y >= len(maze) {
				continue
			}
			pt := point{x, y}
			mv := movements[maze[y][x]]
			for _, d := range mv {
				if start == pt.add(d) {
					res = append(res, pt)
				}
			}
		}
	}
	return res
}

func shoelace(path []point) int { // https://en.wikipedia.org/wiki/Shoelace_formula
	var sum int
	for i := len(path) - 1; i > 0; i-- {
		sum += path[i].x*path[i-1].y - path[i-1].x*path[i].y
	}
	area := sum / 2
	// by default, area includes the tiles occupied by the path. We only want the interior,
	// so subtract the path from the area.
	return area - (len(path)/2 - 1) // subtract one because start is duplicated to complete the loop
}

func walk(start point, dir delta) []point {
	res := []point{start}
	pt, dir := next(start, dir)
	for pt != start {
		res = append(res, pt)
		pt, dir = next(pt, dir)
	}
	res = append(res, start)
	return res
}

func next(pt point, from delta) (point, delta) {
	for _, d := range movements[maze[pt.y][pt.x]] {
		if d != from {
			return pt.add(d), inverse[d]
		}
	}
	log.Fatal("no next")
	// never happens
	return pt.add(from), from
}
