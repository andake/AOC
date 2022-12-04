package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {

	points := map[string]int{"X": 1, "Y": 2, "Z": 3}
	ms := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	os := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}

	score := 0
	for _, s := range inputs {
		split := strings.Split(s, " ")
		o := split[0]
		m := split[1]

		score = score + points[m]
		if ms[m] == os[o] {
			score = score + 3
		} else if win(ms[m], os[o]) {
			score = score + 6
		}
	}

	fmt.Println(score)
}

func problem2(inputs []string) {

	points := map[string]int{"A": 1, "B": 2, "C": 3}
	lose := map[string]string{"A": "C", "B": "A", "C": "B"}
	win := map[string]string{"A": "B", "B": "C", "C": "A"}

	score := 0
	for _, s := range inputs {
		split := strings.Split(s, " ")
		o := split[0]
		m := split[1]

		if m == "X" {
			//lose
			score = score + points[lose[o]]
		} else if m == "Y" {
			//draw
			score = score + 3 + points[o]
		} else if m == "Z" {
			//win
			score = score + 6 + points[win[o]]
		}
	}

	fmt.Println(score)
}

func win(m, o string) bool {
	if m == "rock" && o == "scissors" {
		return true
	} else if m == "scissors" && o == "paper" {
		return true
	} else if m == "paper" && o == "rock" {
		return true
	}
	return false
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
