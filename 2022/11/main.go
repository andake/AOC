package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

type item struct {
	wl uint64
	o  int
}

type monkey struct {
	items    []item
	oper     string
	test     uint64
	tm       int
	fm       int
	monkeys  *[]*monkey
	inspects int
}

func (m *monkey) addItem(i item) {
	m.items = append(m.items, i)
}

func (m *monkey) inspect() {
	m.inspects = m.inspects + 1
	over := new(big.Int).SetUint64(m.items[0].wl)
	if strings.Contains(m.oper, "+") {
		pv, _ := strconv.ParseUint(strings.TrimSpace(strings.Split(m.oper, "+")[1]), 10, 64)
		over = over.Add(new(big.Int).SetUint64(m.items[0].wl), new(big.Int).SetUint64(pv))
		m.items[0].wl = over.Uint64()
	} else {
		o := strings.TrimSpace(strings.Split(m.oper, "*")[1])
		if o == "old" {
			over = over.Mul(new(big.Int).SetUint64(m.items[0].wl), new(big.Int).SetUint64(m.items[0].wl))
			m.items[0].wl = over.Uint64()
		} else {
			val, _ := strconv.ParseUint(o, 10, 64)
			over = over.Mul(new(big.Int).SetUint64(m.items[0].wl), new(big.Int).SetUint64(val))
			m.items[0].wl = over.Uint64()
		}
	}
}

func (m *monkey) bored() {
	m.items[0].wl = m.items[0].wl / 3
}

func (m *monkey) throw(bored bool) {
	i := m.items[0]
	m.items = m.items[1:]
	s := *m.monkeys
	if !bored {
		i.wl = i.wl % 9699690
	}
	if i.wl%m.test == 0 {
		s[m.tm].addItem(i)
	} else {
		s[m.fm].addItem(i)
	}
}

func problem1(inputs []string) {
	monkeys := make([]*monkey, 0)
	mi := 0
	for i := 0; i < len(inputs); i++ {
		row := inputs[i]
		if strings.Contains(row, "Monkey") {
			mi, _ = strconv.Atoi(strings.TrimSuffix(strings.Split(row, " ")[1], ":"))
			m := monkey{monkeys: &monkeys, inspects: 0}
			monkeys = append(monkeys, &m)
		} else if strings.Contains(row, "Starting items") {
			itemString := strings.Split(row, ":")[1]
			for _, i := range strings.Split(itemString, ",") {
				val, _ := strconv.ParseUint(strings.TrimSpace(i), 10, 64)
				item := item{wl: val, o: 0}
				monkeys[mi].addItem(item)
			}
		} else if strings.Contains(row, "Operation") {
			oper := strings.TrimSpace(strings.Split(row, "=")[1])
			monkeys[mi].oper = oper
		} else if strings.Contains(row, "Test") {
			test, _ := strconv.ParseUint(strings.TrimSpace(strings.Split(row, "by")[1]), 10, 64)
			monkeys[mi].test = test
		} else if strings.Contains(row, "true") {
			tm, _ := strconv.Atoi(strings.TrimSpace(strings.Split(row, "monkey")[1]))
			monkeys[mi].tm = tm
		} else if strings.Contains(row, "false") {
			fm, _ := strconv.Atoi(strings.TrimSpace(strings.Split(row, "monkey")[1]))
			monkeys[mi].fm = fm
		}
	}

	for i := 0; i < 20; i++ {
		for _, m := range monkeys {
			length := len(m.items)
			for y := 0; y < length; y++ {
				m.inspect()
				m.bored()
				m.throw(true)
			}
		}
	}

	inspects := make([]int, 0)
	for i, m := range monkeys {
		fmt.Printf("Monkey %d:", i)
		fmt.Printf(" %d,", m.inspects)
		println()
		inspects = append(inspects, m.inspects)
	}

	sort.Ints(inspects)
	multi := inspects[len(inspects)-1] * inspects[len(inspects)-2]
	println(multi)

}

func problem2(inputs []string) {
	monkeys := make([]*monkey, 0)
	mi := 0
	for i := 0; i < len(inputs); i++ {
		row := inputs[i]
		if strings.Contains(row, "Monkey") {
			mi, _ = strconv.Atoi(strings.TrimSuffix(strings.Split(row, " ")[1], ":"))
			m := monkey{monkeys: &monkeys, inspects: 0}
			monkeys = append(monkeys, &m)
		} else if strings.Contains(row, "Starting items") {
			itemString := strings.Split(row, ":")[1]
			for _, i := range strings.Split(itemString, ",") {
				val, _ := strconv.ParseUint(strings.TrimSpace(i), 10, 64)
				item := item{wl: val, o: 0}
				monkeys[mi].addItem(item)
			}
		} else if strings.Contains(row, "Operation") {
			oper := strings.TrimSpace(strings.Split(row, "=")[1])
			monkeys[mi].oper = oper
		} else if strings.Contains(row, "Test") {
			test, _ := strconv.ParseUint(strings.TrimSpace(strings.Split(row, "by")[1]), 10, 64)
			monkeys[mi].test = test
		} else if strings.Contains(row, "true") {
			tm, _ := strconv.Atoi(strings.TrimSpace(strings.Split(row, "monkey")[1]))
			monkeys[mi].tm = tm
		} else if strings.Contains(row, "false") {
			fm, _ := strconv.Atoi(strings.TrimSpace(strings.Split(row, "monkey")[1]))
			monkeys[mi].fm = fm
		}
	}

	for i := 0; i < 10000; i++ {
		for _, m := range monkeys {
			length := len(m.items)
			for y := 0; y < length; y++ {
				m.inspect()
				m.throw(false)
			}
		}
	}

	inspects := make([]int, 0)
	for i, m := range monkeys {
		fmt.Printf("Monkey %d:", i)
		fmt.Printf(" %d,", m.inspects)
		println()
		inspects = append(inspects, m.inspects)
	}

	sort.Ints(inspects)
	multi := inspects[len(inspects)-1] * inspects[len(inspects)-2]
	println(multi)
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
