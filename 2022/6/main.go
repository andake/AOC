package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var prev1 rune = 0
	var prev2 rune = 0
	var prev3 rune = 0
	chars := 0
	for x, r := range inputs[0] {

		if prev1 != prev2 && prev1 != prev3 && prev1 != r && prev2 != prev3 && prev2 != r && prev3 != r {
			if prev1 != 0 && prev2 != 0 && prev3 != 0 {
				chars = x + 1
				break
			}
		}

		prev1 = prev2
		prev2 = prev3
		prev3 = r

	}
	fmt.Println(chars)
}

func problem2(inputs []string) {
	comp := make([]rune, 14)
	line := inputs[0]
	for i := range comp {
		comp[i] = rune(line[i])
	}
	chars := 0
	for x, r := range inputs[0] {
		same := false
		for b, p1 := range comp {
			for c, p2 := range comp {
				if b != c && p1 == p2 {
					same = true
					break
				}
			}
		}
		for i := 0; i < len(comp)-1; i++ {
			comp[i] = comp[i+1]
		}
		comp[len(comp)-1] = r
		if !same {
			chars = x
			break
		}
	}
	fmt.Println(chars)
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
