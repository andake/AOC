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
	var s [][]string
	for y := 0; y < 200; y++ {
		s = append(s, make([]string, 200))
		for i := 0; i < len(s[y]); i++ {
			s[y][i] = "."
		}
	}
	s[0][500-450] = "+"

	for _, row := range inputs {
		cs := strings.Split(row, " -> ")
		for i := 0; i < len(cs)-1; i++ {
			var x1, y1, x2, y2 int
			fmt.Sscanf(cs[i], "%d,%d", &x1, &y1)
			fmt.Sscanf(cs[i+1], "%d,%d", &x2, &y2)
			if x1 == x2 {
				start := y1
				stop := y2
				if y2 < y1 {
					start = y2
					stop = y1
				}
				for y := start; y <= stop; y++ {
					s[y][x1-450] = "#"
				}
			} else {
				start := x1
				stop := x2
				if x2 < x1 {
					start = x2
					stop = x1
				}
				for x := start; x <= stop; x++ {
					s[y1][x-450] = "#"
				}
			}

		}
	}

	sands := 0
Loop:
	for true {
		sy := 0
		sx := 500 - 450
	Loop1:
		for true {
			if sy >= len(s)-1 {
				break Loop
			}
			if sx > len(s[sy])-1 {
				break Loop
			}

			if s[sy+1][sx] == "." {
				sy = sy + 1
			} else if s[sy+1][sx-1] == "." {
				sy = sy + 1
				sx = sx - 1
			} else if s[sy+1][sx+1] == "." {
				sy = sy + 1
				sx = sx + 1
			} else {
				s[sy][sx] = "o"
				sands = sands + 1
				break Loop1
			}
		}
	}

	println(sands)
}

func problem2(inputs []string) {
	var s [][]string
	for y := 0; y < 200; y++ {
		s = append(s, make([]string, 1000))
		for i := 0; i < len(s[y]); i++ {
			s[y][i] = "."
		}
	}
	s[0][500] = "+"

	maxy := 0
	for _, row := range inputs {
		cs := strings.Split(row, " -> ")
		for i := 0; i < len(cs)-1; i++ {
			var x1, y1, x2, y2 int
			fmt.Sscanf(cs[i], "%d,%d", &x1, &y1)
			fmt.Sscanf(cs[i+1], "%d,%d", &x2, &y2)
			if y1 > maxy {
				maxy = y1
			} else if y2 > maxy {
				maxy = y2
			}
			if x1 == x2 {
				start := y1
				stop := y2
				if y2 < y1 {
					start = y2
					stop = y1
				}
				for y := start; y <= stop; y++ {
					s[y][x1] = "#"
				}
			} else {
				start := x1
				stop := x2
				if x2 < x1 {
					start = x2
					stop = x1
				}
				for x := start; x <= stop; x++ {
					s[y1][x] = "#"
				}
			}

		}
	}

	for i := 0; i < 1000; i++ {
		s[maxy+2][i] = "#"
	}

	sands := 0
Loop:
	for true {
		sy := 0
		sx := 500
	Loop1:
		for true {

			if s[sy+1][sx] == "." {
				sy = sy + 1
			} else if s[sy+1][sx-1] == "." {
				sy = sy + 1
				sx = sx - 1
			} else if s[sy+1][sx+1] == "." {
				sy = sy + 1
				sx = sx + 1
			} else {
				s[sy][sx] = "o"
				sands = sands + 1
				if sy == 0 && sx == 500 {
					break Loop
				} else {
					break Loop1
				}
			}
		}
	}

	println(sands)
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
