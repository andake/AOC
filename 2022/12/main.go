package main

import (
	"bufio"
	"log"
	"math"
	"os"
)

type node struct {
	dist    int
	val     rune
	x       int
	y       int
	visited bool
	dest    bool
}

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func findNextNode(nodes [][]*node) *node {
	var node *node

	for _, nr := range nodes {
		for _, n := range nr {
			if !n.visited {
				if node == nil {
					node = n
				} else if n.dist < node.dist {
					node = n
				}
			}
		}
	}

	if node.dist == math.MaxInt {
		return nil
	}

	return node
}

func dijstras(nodes [][]*node, current *node, w int, h int) {
	for current != nil {
		x := current.x
		y := current.y

		if y != 0 {
			n := nodes[y-1][x]
			if !n.visited {
				if n.val-current.val <= 1 {
					d := current.dist + 1
					if d < n.dist {
						n.dist = d
					}
				}
			}
		}
		if x != 0 {
			n := nodes[y][x-1]
			if !n.visited {
				if n.val-current.val <= 1 {
					d := current.dist + 1
					if d < n.dist {
						n.dist = d
					}
				}
			}
		}
		if y != h {
			n := nodes[y+1][x]
			if !n.visited {
				if n.val-current.val <= 1 {
					d := current.dist + 1
					if d < n.dist {
						n.dist = d
					}
				}
			}
		}
		if x != w {
			n := nodes[y][x+1]
			if !n.visited {
				if n.val-current.val <= 1 {
					d := current.dist + 1
					if d < n.dist {
						n.dist = d
					}
				}
			}
		}
		current.visited = true
		if current.dest {
			break
		}

		current = findNextNode(nodes)
	}
}

func problem1(inputs []string) {
	var nodes [][]*node
	var current *node
	var goal *node
	h := 0
	w := 0
	for y, row := range inputs {
		nodes = append(nodes, make([]*node, 0))
		for x, c := range row {
			node := &node{math.MaxInt, c, x, y, false, false}
			if c == 'S' {
				node.dist = 0
				node.val = 'a'
				current = node
			} else if c == 'E' {
				node.val = 'z'
				node.dest = true
				goal = node
			}
			nodes[y] = append(nodes[y], node)
		}
		h = y
		w = len(row) - 1
	}

	dijstras(nodes, current, w, h)
	println(goal.dist)
}

func reset(nodes [][]*node) {
	for _, nr := range nodes {
		for _, n := range nr {
			n.visited = false
			n.dist = math.MaxInt
		}
	}
}

func problem2(inputs []string) {
	var nodes [][]*node
	var current *node
	var goal *node
	aNodes := make([]*node, 0)
	h := 0
	w := 0
	for y, row := range inputs {
		nodes = append(nodes, make([]*node, 0))
		for x, c := range row {
			node := &node{math.MaxInt, c, x, y, false, false}
			if c == 'S' {
				node.val = 'a'
				aNodes = append(aNodes, node)
			} else if c == 'E' {
				node.val = 'z'
				node.dest = true
				goal = node
			} else if c == 'a' {
				aNodes = append(aNodes, node)
			}
			nodes[y] = append(nodes[y], node)
		}
		h = y
		w = len(row) - 1
	}

	dists := make([]int, 0)
	for _, a := range aNodes {
		current = a
		current.dist = 0

		dijstras(nodes, current, w, h)
		dists = append(dists, goal.dist)
		reset(nodes)
	}

	dist := math.MaxInt
	for _, d := range dists {
		if d < dist {
			dist = d
		}
	}
	println(dist)
}

func readInput() []string {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var inputs []string

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		inputs = append(inputs, scan.Text())
	}

	return inputs
}
