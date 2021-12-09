package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func part1(v []string) {

	res := 0
	cavemap := make([][]int, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
		}
		cavemap[i] = line
	}

	for i := 0; i < len(cavemap); i++ {
		for j := 0; j < len(cavemap[i]); j++ {
			if (i == 0 || cavemap[i][j] < cavemap[i-1][j]) && (i == len(cavemap[i])-1 || cavemap[i][j] < cavemap[i+1][j]) &&
				(j == 0 || cavemap[i][j] < cavemap[i][j-1]) && (j == len(cavemap)-1 || cavemap[i][j] < cavemap[i][j+1]) {
				res += cavemap[i][j] + 1
			}
		}
	}
	fmt.Println(res)
}

type Coords struct {
	Y int
	X int
}

func part2(v []string) {

	res := 0
	cavemap := make([][]int, len(v))
	seenmap := make([][]bool, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		seenline := make([]bool, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
			seenline[j] = false
		}
		cavemap[i] = line
		seenmap[i] = seenline
	}

	var basinsizes []int
	// find all the basins
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			if seenmap[i][j] {
				continue
			}

			// else start a new basin
			visitqueue := make([]Coords, 1, 100)
			visitqueue[0] = Coords{i, j}
			basinsize := 0
			for len(visitqueue) > 0 {
				curr := visitqueue[0]
				visitqueue = visitqueue[1:]
				if curr.Y < 0 || curr.Y >= len(cavemap) || curr.X < 0 || curr.X >= len(cavemap[curr.Y]) {
					continue
				}
				if seenmap[curr.Y][curr.X] {
					continue
				}
				seenmap[curr.Y][curr.X] = true
				if cavemap[curr.Y][curr.X] == 9 {
					continue
				}
				basinsize++
				visitqueue = append(visitqueue, Coords{curr.Y - 1, curr.X}, Coords{curr.Y + 1, curr.X},
					Coords{curr.Y, curr.X - 1}, Coords{curr.Y, curr.X + 1})

			}
			basinsizes = append(basinsizes, basinsize)
		}
	}

	sort.Ints(basinsizes)
	// multiply the biggest three
	res = basinsizes[len(basinsizes)-1] * basinsizes[len(basinsizes)-2] * basinsizes[len(basinsizes)-3]
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 9/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
