package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	inputs := readInput()
	cals := problem1(inputs)
	problem2(cals)
}
func problem2(inputs []int) {
	sort.Ints(inputs)
	fmt.Println(inputs[len(inputs)-1] + inputs[len(inputs)-3] + inputs[len(inputs)-2])
}

func problem1(inputs []string) []int {
	var calories []int
	total := 0
	for _, calorie := range inputs {
		c, _ := strconv.Atoi(calorie)
		total = total + c
		if calorie == "" {
			calories = append(calories, total)
			total = 0
		}
	}
	max := 0
	for _, cal := range calories {
		if cal > max {
			max = cal
		}
	}
	fmt.Println(max)
	return calories
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
