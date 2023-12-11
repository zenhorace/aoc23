package main

import (
	"fmt"
	"log"
	"strings"
)

type node struct {
	id   string
	l, r string // could be *node
}

const goal = "ZZZ"
const startNode = "AAA"

func main() {
	lines := strings.Split(networkStr, "\n")
	nodes := make(map[string]*node)
	var startNodes []string
	for _, line := range lines {
		id, edgeStr, _ := strings.Cut(line, " = ")
		edgeStr = strings.Trim(edgeStr, "()")
		edges := strings.Split(edgeStr, ", ")
		if len(edges) != 2 {
			log.Fatalf("invalid edge count: %d", len(edges))
		}
		nodes[id] = &node{id: id, l: edges[0], r: edges[1]}
		if id[2] == 'A' {
			startNodes = append(startNodes, id)
		}
	}
	fmt.Println("pt1", sol1(nodes))
	fmt.Println("pt2", sol2(nodes, startNodes))
}

func sol2(nodes map[string]*node, startNodes []string) int {
	var vals []int
	for _, startNode := range startNodes {
		currNode := nodes[startNode]
		steps := 0
		for currNode.id[2] != 'Z' {
			inst := instructions[(steps % len(instructions))]
			if inst == 'L' {
				currNode = nodes[currNode.l]
			} else {
				currNode = nodes[currNode.r]
			}
			steps++
		}
		vals = append(vals, steps)
	}
	return LCM(vals[0], vals[1], vals[2:]...)
}

func sol1(nodes map[string]*node) int {
	currNode := nodes[startNode]
	steps := 0
	for currNode.id != goal {
		inst := instructions[(steps % len(instructions))]
		if inst == 'L' {
			currNode = nodes[currNode.l]
		} else {
			currNode = nodes[currNode.r]
		}
		steps++
	}
	return steps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
