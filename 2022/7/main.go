package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type file struct {
	name string
	size int
}

type dir struct {
	files  []file
	dirs   []*dir
	parent *dir
	size   int
	name   string
}

func (d *dir) addFile(name string, size int) {
	file := file{name, size}
	d.files = append(d.files, file)
}

func (d *dir) addDirectory(dir *dir) {
	d.dirs = append(d.dirs, dir)
}

func (d *dir) addSize(s int) {
	d.size = d.size + s
	if d.parent != nil {
		d.parent.addSize(s)
	}
}

func main() {
	inputs := readInput()
	problem1(inputs)
	problem2(inputs)
}

func problem1(inputs []string) {
	root := &dir{name: "/", size: 0}
	currentDir := root
	var dirs []*dir
	for i := 1; i < len(inputs); i++ {
		cmd := inputs[i]
		if strings.Contains(cmd, "$") {
			if strings.Contains(cmd, "cd") {
				dirCmd := strings.Split(cmd, " ")[2]
				if strings.Contains(dirCmd, "..") {
					currentDir = currentDir.parent
				} else {
					dirName := currentDir.name + "/" + dirCmd
					if currentDir.name == "/" {
						dirName = "/" + dirCmd
					}

					newDir := &dir{parent: currentDir, name: dirName, size: 0}
					currentDir.addDirectory(newDir)
					currentDir = newDir
					dirs = append(dirs, newDir)
				}
			}
		} else if !strings.Contains(cmd, "dir") {
			fileInfo := strings.Split(cmd, " ")
			size, _ := strconv.Atoi(fileInfo[0])
			name := fileInfo[1]
			currentDir.addFile(name, size)
			currentDir.addSize(size)
		}
	}

	totsize := 0
	for _, dir := range dirs {
		if dir.size <= 100000 {
			totsize = totsize + dir.size
		}
	}
	fmt.Println(totsize)
}

func problem2(inputs []string) {
	root := &dir{name: "/", size: 0}
	currentDir := root
	var dirs []*dir
	for i := 1; i < len(inputs); i++ {
		cmd := inputs[i]
		if strings.Contains(cmd, "$") {
			if strings.Contains(cmd, "cd") {
				dirCmd := strings.Split(cmd, " ")[2]
				if strings.Contains(dirCmd, "..") {
					currentDir = currentDir.parent
				} else {
					dirName := currentDir.name + "/" + dirCmd
					if currentDir.name == "/" {
						dirName = "/" + dirCmd
					}

					newDir := &dir{parent: currentDir, name: dirName, size: 0}
					currentDir.addDirectory(newDir)
					currentDir = newDir
					dirs = append(dirs, newDir)
				}
			}
		} else if !strings.Contains(cmd, "dir") {
			fileInfo := strings.Split(cmd, " ")
			size, _ := strconv.Atoi(fileInfo[0])
			name := fileInfo[1]
			currentDir.addFile(name, size)
			currentDir.addSize(size)
		}
	}

	unused := 70000000 - root.size
	needed := 30000000 - unused
	smallest := math.MaxInt
	for _, dir := range dirs {
		if dir.size >= needed {
			if dir.size < smallest {
				smallest = dir.size
			}
		}
	}

	fmt.Println(smallest)
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
