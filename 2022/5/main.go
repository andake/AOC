package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type stack struct {
	crates []string
}

func (st *stack) add(s string) {
	st.crates = append(st.crates, s)
}

func (st *stack) addLast(s string) {
	st.crates = append([]string{s}, st.crates...)
}

func (st *stack) addMulti(s []string) {
	st.crates = append(st.crates, s...)
}

func (st *stack) get(n int) []string {
	ret := st.crates[len(st.crates)-n:]
	st.crates = st.crates[:len(st.crates)-n]
	return ret
}

func (st *stack) readFirst() string {
	return st.crates[len(st.crates)-1]
}

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	var stacks []stack
	for x, row := range inputs {
		y := 0
		for i := 0; i < len(row); {
			s := string(row[i])
			if x == 0 {
				stack := stack{make([]string, 0)}
				stacks = append(stacks, stack)
			}
			if s == "[" {
				stacks[y].addLast(string(row[i+1]))
			}
			i = i + 4
			y++
		}
		if strings.Contains(row, "1") {
			break
		}
	}

	for _, row := range inputs {

		if strings.Contains(row, "move") {
			action := strings.Split(row, " ")
			total, _ := strconv.Atoi(action[1])
			fromStack, _ := strconv.Atoi(action[3])
			toStack, _ := strconv.Atoi(action[5])

			for i := 0; i < total; i++ {
				stacks[toStack-1].addMulti(stacks[fromStack-1].get(1))
			}
		}
	}
	for _, s := range stacks {
		fmt.Print(s.readFirst())
	}
	fmt.Println()
}

func problem2(inputs []string) {
	var stacks []stack
	for x, row := range inputs {
		y := 0
		for i := 0; i < len(row); {
			s := string(row[i])
			if x == 0 {
				stack := stack{make([]string, 0)}
				stacks = append(stacks, stack)
			}
			if s == "[" {
				stacks[y].addLast(string(row[i+1]))
			}
			i = i + 4
			y++
		}
		if strings.Contains(row, "1") {
			break
		}
	}

	for _, row := range inputs {

		if strings.Contains(row, "move") {
			action := strings.Split(row, " ")
			total, _ := strconv.Atoi(action[1])
			fromStack, _ := strconv.Atoi(action[3])
			toStack, _ := strconv.Atoi(action[5])

			stacks[toStack-1].addMulti(stacks[fromStack-1].get(total))
		}
	}
	for _, s := range stacks {
		fmt.Print(s.readFirst())
	}
	fmt.Println()
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
