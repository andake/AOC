package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	sum := 0
	for _, comps := range inputs {
		counted := make(map[rune]bool)
		len := len(comps)
		comp1 := comps[:len/2]
		comp2 := comps[len/2:]
		for _, item1 := range comp1 {
			for _, item2 := range comp2 {
				if item1 == item2 {
					if ok := counted[item1]; !ok {
						counted[item1] = true
						sum = sum + getPrio(item1)
						break
					}

				}
			}
		}
	}
	fmt.Println(sum)
}

func problem2(inputs []string) {
	sum := 0
	for i := 0; i < len(inputs); i = i + 3 {
		comp1 := inputs[i]
		comp2 := inputs[i+1]
		comp3 := inputs[i+2]

		counted := make(map[rune]bool)
		var common []rune
		for _, item1 := range comp1 {
			for _, item2 := range comp2 {
				if item2 == item1 {
					if ok := counted[item1]; !ok {
						counted[item1] = true
						common = append(common, item1)
					}
				}
			}
		}

		counted = make(map[rune]bool)
		for _, itemCom := range common {
			for _, item3 := range comp3 {
				if item3 == itemCom {
					if ok := counted[itemCom]; !ok {
						counted[itemCom] = true
						sum = sum + getPrio(itemCom)
					}
				}
			}
		}
	}
	fmt.Println(sum)
}

func getPrio(r rune) int {
	if unicode.IsLower(r) {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
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
