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

func problem2(inputs []string) {

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
