package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	problem1and2(inputs)
}

func problem1and2(inputs []string) {
	contained := 0
	overlap := 0
	for _, row := range inputs {
		assignments := strings.Split(row, ",")
		assign1 := strings.Split(assignments[0], "-")
		assign2 := strings.Split(assignments[1], "-")

		if contains(assign1, assign2) {
			contained++
		} else if overlaps(assign1, assign2) {
			overlap++
		}
	}
	fmt.Println(contained)
	fmt.Println(overlap + contained)
}

func contains(assign1, assign2 []string) bool {
	low1, high1, low2, high2 := highLow(assign1, assign2)
	return low1 >= low2 && high1 <= high2 || low2 >= low1 && high2 <= high1
}

func overlaps(assign1, assign2 []string) bool {
	low1, high1, low2, high2 := highLow(assign1, assign2)
	return (high1 >= low2 && high1 <= high2) || (high2 >= low1 && high2 <= high1)
}

func highLow(assign1, assign2 []string) (int, int, int, int) {
	low1, _ := strconv.Atoi(assign1[0])
	high1, _ := strconv.Atoi(assign1[1])
	low2, _ := strconv.Atoi(assign2[0])
	high2, _ := strconv.Atoi(assign2[1])
	return low1, high1, low2, high2
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
