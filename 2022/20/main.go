package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

type elem struct {
	prev *elem
	next *elem
	val  int
}

func problem1(inputs []string) {

	elems := make([]*elem, 0)
	var head *elem
	for r, i := range inputs {
		v, _ := strconv.Atoi(i)
		elem := &elem{nil, nil, v}
		elems = append(elems, elem)
		if r == 0 {
			head = elem
		} else if r == len(inputs)-1 {
			elem.prev = elems[r-1]
			elems[r-1].next = elem
			head.prev = elem
			elem.next = head
		} else {
			elem.prev = elems[r-1]
			elems[r-1].next = elem
		}
	}

	for e := head; ; e = e.next {
		print(e.val, ", ")
		if e.next == head {
			break
		}
	}
	println()

	ep := 0
	cnt := 0
	sum := 0
	ca := false
	for true {
		e := elems[ep]
		if ca {
			cnt = cnt + 1
		}
		println("cnt ", e.val, cnt)
		if e.val > 0 {
			eg := e
			for i := 0; i < e.val; i++ {
				//println("loop eg", eg.val)
				eg = eg.next
			}
			//fmt.Printf("ep: %d, eg: %d\n", elems[ep].val, eg.val)
			//println("bb ", elems[ep].prev.val, elems[ep].next.val)
			// if e == head {
			// 	head = elems[ep].next
			// 	println("head", head.val)
			// }
			elems[ep].prev.next = elems[ep].next
			elems[ep].next.prev = elems[ep].prev
			elems[ep].prev = eg
			elems[ep].next = eg.next
			eg.next.prev = elems[ep]
			eg.next = elems[ep]
		} else if e.val < 0 {
			eg := e
			for i := 0; i >= e.val; i-- {
				eg = eg.prev
			}
			//fmt.Printf("ep %d, pos %d\n", ep, pos)
			if e == head {
				head = elems[ep].next
			}
			elems[ep].prev.next = elems[ep].next
			elems[ep].next.prev = elems[ep].prev
			elems[ep].prev = eg
			elems[ep].next = eg.next
			eg.next.prev = elems[ep]
			eg.next = elems[ep]

		} else {
			// if !ca {
			// 	cnt = cnt + 1
			// }
			ca = true
		}
		// for e := head; ; e = e.next {
		// 	print(e.val, ", ")
		// 	if e.next == head {
		// 		break
		// 	}
		// }
		// println()
		if cnt == 1000 || cnt == 2000 || cnt == 3000 {
			sum = sum + e.val
			if cnt == 3000 {
				break
			}
		}
		ep = ep + 1
		if ep > len(elems)-1 {
			ep = 0
			// println("bal")
			elems = make([]*elem, 0)
			for et := head; ; et = et.next {
				// println(et.val)
				elems = append(elems, et)
				if et.next == head {
					break
				}
			}
			// println("woot: ", elems[0].val)
		}
	}

	println(sum)

	// for e := head; ; e = e.next {
	// 	print(e.val, ", ")
	// 	if e.next == head {
	// 		break
	// 	}
	// }
	// println()
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
