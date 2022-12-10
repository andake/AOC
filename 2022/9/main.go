package main

import (
	"bufio"
	"log"
	"math"
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
	var visited [][]bool
	visited = make([][]bool, 1000)
	for x := 0; x < len(visited); x++ {
		visited[x] = make([]bool, 1000)
		for y := 0; y < len(visited[x]); y++ {
			visited[x][y] = false
		}
	}
	posxh := 500
	posyh := 500
	posxt := 500
	posyt := 500

	numVisit := 0
	for _, action := range inputs {
		as := strings.Split(action, " ")
		dir := as[0]
		num, _ := strconv.Atoi(as[1])

		for i := 0; i < num; i++ {
			if dir == "U" {
				posyh = posyh - 1
			} else if dir == "D" {
				posyh = posyh + 1
			} else if dir == "L" {
				posxh = posxh - 1
			} else if dir == "R" {
				posxh = posxh + 1
			}

			if posxt != posxh || posyt != posyh {
				if math.Abs(float64(posxt)-float64(posxh)) > 1 && posyt == posyh {
					if posxt > posxh {
						posxt = posxt - 1
					} else {
						posxt = posxt + 1
					}
				} else if math.Abs(float64(posyt)-float64(posyh)) > 1 && posxt == posxh {
					if posyt > posyh {
						posyt = posyt - 1
					} else {
						posyt = posyt + 1
					}
				} else if math.Abs(float64(posyt)-float64(posyh)) > 1 || math.Abs(float64(posxt)-float64(posxh)) > 1 {
					if dir == "U" || dir == "D" {
						posxt = posxh
						if posyt > posyh {
							posyt = posyt - 1
						} else if posyt < posyh {
							posyt = posyt + 1
						}
					} else {
						posyt = posyh
						if posxt > posxh {
							posxt = posxt - 1
						} else if posxt < posxh {
							posxt = posxt + 1
						}
					}
				}
			}

			if !visited[posxt][posyt] {
				numVisit = numVisit + 1
				visited[posxt][posyt] = true
			}
		}
	}

	println(numVisit)
}

type position struct {
	x int
	y int
}

func (p *position) minusY() {
	p.y = p.y - 1
}

func (p *position) minusX() {
	p.x = p.x - 1
}

func (p *position) plusY() {
	p.y = p.y + 1
}

func (p *position) plusX() {
	p.x = p.x + 1
}

func problem2(inputs []string) {

	visited := make(map[position]bool)
	positions := make([]position, 10)
	for i := 0; i < len(positions); i++ {
		positions[i] = position{500, 500}
	}

	for _, action := range inputs {
		as := strings.Split(action, " ")
		dir := as[0]
		num, _ := strconv.Atoi(as[1])
		// println(action)
		for i := 0; i < num; i++ {
			if dir == "U" {
				positions[0].minusY()
			} else if dir == "D" {
				positions[0].plusY()
			} else if dir == "L" {
				positions[0].minusX()
			} else if dir == "R" {
				positions[0].plusX()
			}

			for n := 1; n < len(positions); n++ {
				updatePos(&positions[n], &positions[n-1], dir)
			}
			visited[positions[9]] = true
		}
	}

	println(len(visited))
}

func updatePos(p, pprev *position, dir string) {
	posxt := p.x
	posyt := p.y
	posxh := pprev.x
	posyh := pprev.y

	if posxt != posxh || posyt != posyh {
		if math.Abs(float64(posxt)-float64(posxh)) > 1 && posyt == posyh {
			if posxt > posxh {
				p.minusX()
			} else {
				p.plusX()
			}
		} else if math.Abs(float64(posyt)-float64(posyh)) > 1 && posxt == posxh {
			if posyt > posyh {
				p.minusY()
			} else {
				p.plusY()
			}
		} else if math.Abs(float64(posyt)-float64(posyh)) > 1 || math.Abs(float64(posxt)-float64(posxh)) > 1 {
			if posxt < posxh && posyt > posyh {
				p.plusX()
				p.minusY()
			} else if posxt > posxh && posyt > posyh {
				p.minusY()
				p.minusX()
			} else if posxt > posxh && posyt < posyh {
				p.minusX()
				p.plusY()
			} else {
				p.plusY()
				p.plusX()
			}
		}
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
