package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	w := len(inputs[0]) - 1
	h := 0

	var trees [][]int

	for i := 0; i < len(inputs); i++ {
		row := inputs[i]
		trees = append(trees, make([]int, 0))
		trees[i] = make([]int, 0)
		for y := 0; y < len(row); y++ {
			r := row[y]
			height, _ := strconv.Atoi(string(r))
			trees[i] = append(trees[i], height)
		}
		h = i
	}

	seen := 0
	for i := 0; i < len(trees); i++ {
		treeRow := trees[i]
		for y := 0; y < len(treeRow); y++ {
			tree := treeRow[y]
			if i == 0 || y == 0 || i == h || y == w {
				seen = seen + 1
			} else {
				up, down := getUpDown(trees, i, y)
				if walk(tree, treeRow[:y]) || walk(tree, treeRow[y+1:]) || walk(tree, up) || walk(tree, down) {
					seen = seen + 1
				}
			}

		}
	}
	println(seen)
}

func walk(h int, hs []int) bool {
	for _, i := range hs {
		if h <= i {
			return false
		}
	}
	return true
}

func getUpDown(hs [][]int, coord1, coord2 int) ([]int, []int) {
	up := make([]int, 0)
	down := make([]int, 0)
	for i := 0; i < len(hs); i++ {
		row := hs[i]
		for y := 0; y < len(row); y++ {
			if y == coord2 && i < coord1 {
				up = append(up, row[y])
			} else if y == coord2 && i > coord1 {
				down = append(down, row[y])
			}
		}
	}
	return up, down
}

func problem2(inputs []string) {

	var trees [][]int

	for i := 0; i < len(inputs); i++ {
		row := inputs[i]
		trees = append(trees, make([]int, 0))
		trees[i] = make([]int, 0)
		for y := 0; y < len(row); y++ {
			r := row[y]
			height, _ := strconv.Atoi(string(r))
			trees[i] = append(trees[i], height)
		}
	}

	scenicScores := make([]int, 0)
	for i := 0; i < len(trees); i++ {
		treeRow := trees[i]
		for y := 0; y < len(treeRow); y++ {
			tree := treeRow[y]
			up, down := getUpDown(trees, i, y)
			scenicScores = append(scenicScores, score2(tree, treeRow[:y])*score1(tree, treeRow[y+1:])*score2(tree, up)*score1(tree, down))
		}
	}

	maxScenicScore := 0
	for _, s := range scenicScores {
		if s > maxScenicScore {
			maxScenicScore = s
		}
	}
	println(maxScenicScore)
}

func score1(h int, hs []int) int {
	score := 0
	for _, i := range hs {
		if h <= i {
			score = score + 1
			return score
		} else {
			score = score + 1
		}
	}
	return score
}

func score2(h int, hs []int) int {
	score := 0
	if len(hs) > 0 {
		for y := len(hs) - 1; y >= 0; y-- {
			i := hs[y]
			if h <= i {
				score = score + 1
				return score
			} else {
				score = score + 1
			}
		}
	} else {
		return 0
	}
	return score
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
