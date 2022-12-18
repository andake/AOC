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

type cube struct {
	x int
	y int
	z int
}

func problem1(inputs []string) {
	cubes := make(map[cube]int, 0)
	for _, r := range inputs {
		c := cube{}
		fmt.Sscanf(r, "%d,%d,%d", &c.x, &c.y, &c.z)
		cubes[c] = 0
	}

	for k := range cubes {
		covered := 0
		if _, f := cubes[cube{k.x, k.y, k.z - 1}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y - 1, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x - 1, k.y, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x + 1, k.y, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y + 1, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y, k.z + 1}]; f {
			covered = covered + 1
		}
		cubes[k] = covered
	}

	t := 0
	for _, c := range cubes {
		t = t + c
	}
	println(len(cubes)*6 - t)
}

func problem2(inputs []string) {
	cubes := make(map[cube]int, 0)
	mx := 0
	my := 0
	mz := 0
	for _, r := range inputs {
		c := cube{}
		fmt.Sscanf(r, "%d,%d,%d", &c.x, &c.y, &c.z)
		cubes[c] = 0
		if mx < c.x {
			mx = c.x
		}
		if my < c.y {
			my = c.y
		}
		if mz < c.z {
			mz = c.z
		}
	}

	air := make(map[cube]bool, 0)
	for x := 0; x <= mx; x++ {
		for y := 0; y <= my; y++ {
			for z := 0; z <= mz; z++ {
				if _, f := cubes[cube{x, y, z}]; !f {
					air[cube{x, y, z}] = false
				}
			}
		}
	}

	air[cube{0, 0, 0}] = true
	aq := make([]cube, 0)
	aq = append(aq, cube{0, 0, 0})
	for len(aq) != 0 {
		c := aq[len(aq)-1]
		aq = aq[:len(aq)-1]

		if v, f := air[cube{c.x, c.y, c.z - 1}]; f {
			if !v {
				air[cube{c.x, c.y, c.z - 1}] = true
				aq = append(aq, cube{c.x, c.y, c.z - 1})
			}
		}
		if v, f := air[cube{c.x, c.y, c.z + 1}]; f {
			if !v {
				air[cube{c.x, c.y, c.z + 1}] = true
				aq = append(aq, cube{c.x, c.y, c.z + 1})
			}
		}
		if v, f := air[cube{c.x, c.y - 1, c.z}]; f {
			if !v {
				air[cube{c.x, c.y - 1, c.z}] = true
				aq = append(aq, cube{c.x, c.y - 1, c.z})
			}
		}
		if v, f := air[cube{c.x, c.y + 1, c.z}]; f {
			if !v {
				air[cube{c.x, c.y + 1, c.z}] = true
				aq = append(aq, cube{c.x, c.y + 1, c.z})
			}
		}
		if v, f := air[cube{c.x - 1, c.y, c.z}]; f {
			if !v {
				air[cube{c.x - 1, c.y, c.z}] = true
				aq = append(aq, cube{c.x - 1, c.y, c.z})
			}
		}
		if v, f := air[cube{c.x + 1, c.y, c.z}]; f {
			if !v {
				air[cube{c.x + 1, c.y, c.z}] = true
				aq = append(aq, cube{c.x + 1, c.y, c.z})
			}
		}
	}

	for k := range cubes {
		covered := 0
		if _, f := cubes[cube{k.x, k.y, k.z - 1}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y - 1, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x - 1, k.y, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x + 1, k.y, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y + 1, k.z}]; f {
			covered = covered + 1
		}
		if _, f := cubes[cube{k.x, k.y, k.z + 1}]; f {
			covered = covered + 1
		}
		if r, f := air[cube{k.x, k.y, k.z + 1}]; f {
			if !r {
				covered = covered + 1
			}
		}
		if r, f := air[cube{k.x, k.y, k.z - 1}]; f {
			if !r {
				covered = covered + 1
			}
		}
		if r, f := air[cube{k.x, k.y - 1, k.z}]; f {
			if !r {
				covered = covered + 1
			}
		}
		if r, f := air[cube{k.x, k.y + 1, k.z}]; f {
			if !r {
				covered = covered + 1
			}
		}
		if r, f := air[cube{k.x - 1, k.y, k.z}]; f {
			if !r {
				covered = covered + 1
			}
		}
		if r, f := air[cube{k.x + 1, k.y, k.z}]; f {
			if !r {
				covered = covered + 1
			}
		}

		cubes[k] = covered
	}

	t := 0
	for _, c := range cubes {
		t = t + c
	}

	println(len(cubes)*6 - t)
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
