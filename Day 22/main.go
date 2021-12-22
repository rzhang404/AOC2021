package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(v []string) {

	res := 0
	cubes := make([][][]bool, 101)
	for i := 0; i < 101; i++ {
		cubes[i] = make([][]bool, 101)
		for j := 0; j < 101; j++ {
			cubes[i][j] = make([]bool, 101)
		}
	}

	for _, strinput := range v {
		line := strings.Split(strinput, " ")
		setting := false
		if line[0] == "on" {
			setting = true
		} else {
			setting = false
		}
		coords := strings.Split(line[1], ",")
		x := strings.Split(coords[0][2:], "..")
		y := strings.Split(coords[1][2:], "..")
		z := strings.Split(coords[2][2:], "..")
		xl, _ := strconv.Atoi(x[0])
		xh, _ := strconv.Atoi(x[1])
		yl, _ := strconv.Atoi(y[0])
		yh, _ := strconv.Atoi(y[1])
		zl, _ := strconv.Atoi(z[0])
		zh, _ := strconv.Atoi(z[1])

		if (xl < -50 && xh < -50) || (xl > 50 && xh > 50) {
			continue
		}
		if (yl < -50 && yh < -50) || (yl > 50 && yh > 50) {
			continue
		}
		if (zl < -50 && zh < -50) || (zl > 50 && zh > 50) {
			continue
		}

		if xl < -50 {
			xl = -50
		}
		if xh > 50 {
			xh = 50
		}
		if yl < -50 {
			yl = -50
		}
		if yh > 50 {
			yh = 50
		}
		if zl < -50 {
			zl = -50
		}
		if zh > 50 {
			zh = 50
		}
		for i := xl; i <= xh; i++ {
			for j := yl; j <= yh; j++ {
				for k := zl; k <= zh; k++ {
					cubes[i+50][j+50][k+50] = setting
				}
			}
		}

	}

	for i := -50; i <= 50; i++ {
		for j := -50; j <= 50; j++ {
			for k := -50; k <= 50; k++ {
				if cubes[i+50][j+50][k+50] {
					res++
				}
			}
		}
	}

	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {

	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 22/input.txt")
	//content, err := ioutil.ReadFile("Day 22/smallinput.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
