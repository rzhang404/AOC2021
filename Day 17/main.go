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
	xstr := strings.Split(v[2][2:], "..")
	ystr := strings.Split(v[3][2:], "..")
	xl, _ := strconv.Atoi(xstr[0])
	xr, _ := strconv.Atoi(xstr[1][:len(xstr[1])-1]) // removes the ","
	yb, _ := strconv.Atoi(ystr[0])
	yt, _ := strconv.Atoi(ystr[1])
	for i := 0; i <= xr; i++ {
		for j := 0; j <= 2000; j++ {
			xv := i
			yv := j
			x := 0
			y := 0
			peak := 0
			// simulate
			for y >= yb {
				x += xv
				y += yv
				if y > peak {
					peak = y
				}
				yv--
				if xv > 0 {
					xv--
				} else if xv < 0 {
					xv++
				}

				if x >= xl && x <= xr && y >= yb && y <= yt {
					if peak > res {
						res = peak
					}
					break
				}
			}
		}
	}
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	xstr := strings.Split(v[2][2:], "..")
	ystr := strings.Split(v[3][2:], "..")
	xl, _ := strconv.Atoi(xstr[0])
	xr, _ := strconv.Atoi(xstr[1][:len(xstr[1])-1]) // removes the ","
	yb, _ := strconv.Atoi(ystr[0])
	yt, _ := strconv.Atoi(ystr[1])
	for i := 0; i <= xr; i++ {
		for j := yb; j <= 5000; j++ {
			xv := i
			yv := j
			x := 0
			y := 0
			// simulate
			for y >= yb {
				x += xv
				y += yv
				yv--
				if xv > 0 {
					xv--
				} else if xv < 0 {
					xv++
				}

				if x >= xl && x <= xr && y >= yb && y <= yt {
					res++
					break
				}
			}
		}
	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 17/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, " ")
	part1(v)
	part2(v)
}
