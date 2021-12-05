package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(v []string) {

	seamap := make([][]int, 1000)
	for i := 0; i < len(seamap); i++ {
		seamap[i] = make([]int, 1000)
	}
	res := 0
	for i := 0; i < len(v); i++ {
		split := strings.Split(v[i], " -> ")
		start := strings.Split(split[0], ",")
		end := strings.Split(split[1], ",")
		startx, _ := strconv.Atoi(start[0])
		starty, _ := strconv.Atoi(start[1])
		endx, _ := strconv.Atoi(end[0])
		endy, _ := strconv.Atoi(end[1])

		if startx == endx && starty <= endy {
			for j := starty; j <= endy; j++ {
				seamap[startx][j]++
			}
		} else if startx == endx && starty >= endy {
			for j := starty; j >= endy; j-- {
				seamap[startx][j]++
			}
		} else if starty == endy && startx <= endx {
			for j := startx; j <= endx; j++ {
				seamap[j][starty]++
			}
		} else if starty == endy && startx >= endx {
			for j := startx; j >= endx; j-- {
				seamap[j][starty]++
			}
		}
	}

	for i := 0; i < len(seamap); i++ {
		for j := 0; j < len(seamap[0]); j++ {
			if seamap[i][j] >= 2 {
				res++
			}
		}
	}
	fmt.Println(res)
}

func part2(v []string) {

	seamap := make([][]int, 1000)
	for i := 0; i < len(seamap); i++ {
		seamap[i] = make([]int, 1000)
	}
	res := 0
	for i := 0; i < len(v); i++ {
		split := strings.Split(v[i], " -> ")
		start := strings.Split(split[0], ",")
		end := strings.Split(split[1], ",")
		startx, _ := strconv.Atoi(start[0])
		starty, _ := strconv.Atoi(start[1])
		endx, _ := strconv.Atoi(end[0])
		endy, _ := strconv.Atoi(end[1])

		if startx == endx && starty <= endy {
			for j := starty; j <= endy; j++ {
				seamap[startx][j]++
			}
		} else if startx == endx && starty >= endy {
			for j := starty; j >= endy; j-- {
				seamap[startx][j]++
			}
		} else if starty == endy && startx <= endx {
			for j := startx; j <= endx; j++ {
				seamap[j][starty]++
			}
		} else if starty == endy && startx >= endx {
			for j := startx; j >= endx; j-- {
				seamap[j][starty]++
			}
		} else if startx < endx && starty < endy {
			for j := 0; j <= endy-starty; j++ {
				seamap[startx+j][starty+j]++
			}
		} else if startx < endx && starty > endy {
			for j := 0; j <= starty-endy; j++ {
				seamap[startx+j][starty-j]++
			}
		} else if startx > endx && starty < endy {
			for j := 0; j <= endy-starty; j++ {
				seamap[startx-j][starty+j]++
			}
		} else if startx > endx && starty > endy {
			for j := 0; j <= starty-endy; j++ {
				seamap[startx-j][starty-j]++
			}
		}
	}

	for i := 0; i < len(seamap); i++ {
		for j := 0; j < len(seamap[0]); j++ {
			if seamap[i][j] >= 2 {
				res++
			}
		}
	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 5/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
