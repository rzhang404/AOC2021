package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(v []string) {

	dots := strings.Split(v[0], "\n")
	folds := strings.Split(v[1], "\n")
	dotmap := make([][]bool, 2000)
	for i := 0; i < 2000; i++ {
		line := make([]bool, 2000)
		for j := 0; j < 2000; j++ {
			line[j] = false
		}
		dotmap[i] = line
	}

	maxx := 0
	maxy := 0
	res := 0
	for i := 0; i < len(dots); i++ {
		dot := strings.Split(dots[i], ",")
		x, _ := strconv.Atoi(dot[0])
		y, _ := strconv.Atoi(dot[1])
		dotmap[y][x] = true
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	fold := strings.Split(folds[0][11:], "=") // remove "fold along "
	foldaxis := fold[0]
	foldline, _ := strconv.Atoi(fold[1])

	if foldaxis == "x" {
		for i := 1; i <= (maxx-foldline) && i <= foldline; i++ {
			for j := 0; j <= maxy; j++ {
				dotmap[j][foldline-i] = dotmap[j][foldline+i] || dotmap[j][foldline-i]
			}
		}
		maxx = foldline - 1
	}

	for i := 0; i <= maxx; i++ {
		for j := 0; j <= maxy; j++ {
			if dotmap[j][i] {
				res++
			}
		}
	}
	fmt.Println(res)
}

func part2(v []string) {

	dots := strings.Split(v[0], "\n")
	folds := strings.Split(v[1], "\n")
	dotmap := make([][]bool, 2000)
	for i := 0; i < 2000; i++ {
		line := make([]bool, 2000)
		for j := 0; j < 2000; j++ {
			line[j] = false
		}
		dotmap[i] = line
	}

	maxx := 0
	maxy := 0
	for i := 0; i < len(dots); i++ {
		dot := strings.Split(dots[i], ",")
		x, _ := strconv.Atoi(dot[0])
		y, _ := strconv.Atoi(dot[1])
		dotmap[y][x] = true
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
	}

	for _, instr := range folds {
		fold := strings.Split(instr[11:], "=") // remove "fold along "
		foldaxis := fold[0]
		foldline, _ := strconv.Atoi(fold[1])
		if foldaxis == "x" {
			for i := 1; i <= (maxx-foldline) && i <= foldline; i++ {
				for j := 0; j <= maxy; j++ {
					dotmap[j][foldline-i] = dotmap[j][foldline+i] || dotmap[j][foldline-i]
				}
			}
			maxx = foldline - 1
		} else {
			for i := 0; i <= maxx; i++ {
				for j := 1; j <= (maxy-foldline) && j <= foldline; j++ {
					dotmap[foldline-j][i] = dotmap[foldline+j][i] || dotmap[foldline-j][i]
				}
			}
			maxy = foldline - 1
		}
	}

	for j := 0; j <= maxy; j++ {
		for i := 0; i <= maxx; i++ {
			if dotmap[j][i] {
				print("#")
			} else {
				print(".")
			}
		}
		print("\n")
	}
}

func main() {

	content, err := ioutil.ReadFile("Day 13/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n\n")
	part1(v)
	part2(v)
}
