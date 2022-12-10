package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	cycles := 0
	x := 1
	strength := 0
	for _, in := range inputs {
		cycles = cycles + 1
		strength = strength + getStrength(cycles, x)

		if in != "noop" {
			cycles = cycles + 1
			strength = strength + getStrength(cycles, x)
			num, _ := strconv.Atoi(strings.Split(in, " ")[1])
			x = x + num
		}
	}
	println(strength)
}

func getStrength(cycles, x int) int {
	switch cycles {
	case 20:
		fallthrough
	case 60:
		fallthrough
	case 100:
		fallthrough
	case 140:
		fallthrough
	case 180:
		fallthrough
	case 220:
		return cycles * x
	}

	return 0
}

func problem2(inputs []string) {
	crt := make([][]string, 6)
	for i := 0; i < len(crt); i++ {
		crt[i] = make([]string, 40)
		for n := 0; n < len(crt[i]); n++ {
			crt[i][n] = "."
		}
	}

	cycles := 0
	x := 1
	crtWritePos := 0
	for _, in := range inputs {
		cycles = cycles + 1

		if cycles == 1 || cycles == 41 || cycles == 81 || cycles == 121 || cycles == 161 || cycles == 201 {
			crtWritePos = 0
		}

		if crtWritePos <= x+1 && crtWritePos >= x-1 {
			crt[getRow(cycles)][crtWritePos] = "#"
		}
		crtWritePos = crtWritePos + 1
		if in != "noop" {
			cycles = cycles + 1
			if cycles == 1 || cycles == 41 || cycles == 81 || cycles == 121 || cycles == 161 || cycles == 201 {
				crtWritePos = 0
			}
			if crtWritePos <= x+1 && crtWritePos >= x-1 {
				crt[getRow(cycles)][crtWritePos] = "#"
			}
			num, _ := strconv.Atoi(strings.Split(in, " ")[1])
			x = x + num
			crtWritePos = crtWritePos + 1
		}
	}

	for i := 0; i < len(crt); i++ {
		for n := 0; n < len(crt[i]); n++ {
			print(crt[i][n])
		}
		println()
	}
}

func getRow(c int) int {
	if c >= 1 && c <= 40 {
		return 0
	} else if c >= 41 && c <= 80 {
		return 1
	} else if c >= 81 && c <= 120 {
		return 2
	} else if c >= 121 && c <= 160 {
		return 3
	} else if c >= 161 && c <= 200 {
		return 4
	} else {
		return 5
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
