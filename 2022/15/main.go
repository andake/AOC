package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

type Coord struct {
	x int
	y int
}

func problem1(inputs []string) {
	coverage := make(map[Coord]bool)

	yt := 2000000
	for _, row := range inputs {
		var sx int
		var sy int
		var bx int
		var by int

		fmt.Sscanf(row, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)

		coverage[Coord{sx, sy}] = true
		coverage[Coord{bx, by}] = false
		md := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))

		if sy+md > yt && sy < yt {
			o := yt - sy
			for x := sx - md + o; x <= sx+md-o; x++ {
				c := Coord{x, yt}
				if _, found := coverage[c]; !found {
					coverage[c] = true
				}
			}
		} else if sy-md < yt && sy >= yt {
			o := sy - yt
			for x := sx - md + o; x <= sx+md-o; x++ {
				c := Coord{x, yt}
				if _, found := coverage[c]; !found {
					coverage[c] = true
				}
			}
		}
	}

	c := 0
	for k, v := range coverage {
		if k.y == yt {
			if v {
				c = c + 1
			}
		}
	}
	println(c)
}

type coordmd struct {
	x  int
	y  int
	md int
}

func problem2(inputs []string) {
	coords := make([]coordmd, 0)
	for _, row := range inputs {
		var sx int
		var sy int
		var bx int
		var by int

		fmt.Sscanf(row, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)

		md := int(math.Abs(float64(sx-bx)) + math.Abs(float64(sy-by)))
		c := coordmd{sx, sy, md}
		coords = append(coords, c)
	}

	xg := 0
	yg := 0
	limit := 4000000
Loop:
	for _, c := range coords {
		for x, y := c.x-c.md-1, c.y; x <= c.x; x, y = x+1, y-1 {
			if x >= 0 && x <= limit && y >= 0 && y <= limit {
				for pos, c2 := range coords {
					if !included(x, y, c2) {
						//println("res", x, y)
						if pos == len(coords)-1 {
							println("res", x, y)
							xg = x
							yg = y
							break Loop
						}
					} else {
						break
					}
				}
			}
		}
		for x, y := c.x+c.md+1, c.y; x >= c.x; x, y = x-1, y-1 {
			if x >= 0 && x <= limit && y >= 0 && y <= limit {
				for pos, c2 := range coords {
					if !included(x, y, c2) {
						if pos == len(coords)-1 {
							println("blaa3")
							println("res", x, y)
							xg = x
							yg = y
							break Loop
						}
					} else {
						break
					}
				}
			}
		}
		for x, y := c.x+c.md+1, c.y; x >= c.x; x, y = x-1, y+1 {
			if x >= 0 && x <= limit && y >= 0 && y <= limit {
				for pos, c2 := range coords {
					if !included(x, y, c2) {
						if pos == len(coords)-1 {
							println("res", x, y)
							xg = x
							yg = y
							break Loop
						}
					} else {
						break
					}
				}
			}
		}
		for x, y := c.x-c.md-1, c.y; x <= c.x; x, y = x+1, y+1 {
			if x >= 0 && x <= limit && y >= 0 && y <= limit {
				for pos, c2 := range coords {
					if !included(x, y, c2) {
						if pos == len(coords)-1 {
							println("res", x, y)
							xg = x
							yg = y
							break Loop
						}
					} else {
						break
					}
				}
			}
		}
	}

	println(xg*4000000 + yg)
}

func included(x, y int, c coordmd) bool {
	//println("c.md: ", c.md)
	// println("xy md: ", int(math.Abs(float64(x-c.x))+math.Abs(float64(y-c.y))))
	return c.md >= int(math.Abs(float64(x-c.x))+math.Abs(float64(y-c.y)))
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
