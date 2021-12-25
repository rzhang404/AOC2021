package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func part1(v []string) {

	res := 0
	seamap := make([][]int, len(v))
	auxmap := make([][]int, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			if string(v[i][j]) == ">" {
				line[j] = 1 // east
			} else if string(v[i][j]) == "v" {
				line[j] = 2 // south
			} else if string(v[i][j]) == "." {
				line[j] = 0 // empty
			} else {
				log.Fatal("misinput")
			}
		}
		seamap[i] = line
		auxmap[i] = make([]int, len(v[i]))
	}

	changed := true
	for changed {
		changed = false
		res++

		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v[i]); j++ {
				jprime := j + 1
				if jprime == len(v[i]) {
					jprime = 0
				}
				if seamap[i][j] == 1 && seamap[i][jprime] == 0 {
					changed = true
					auxmap[i][j] = 0
					auxmap[i][jprime] = 1
					j++
				} else {
					auxmap[i][j] = seamap[i][j]
				}
			}
		}

		seamap, auxmap = auxmap, seamap

		for j := 0; j < len(v[0]); j++ {
			for i := 0; i < len(v); i++ {
				iprime := i + 1
				if iprime == len(v) {
					iprime = 0
				}
				if seamap[i][j] == 2 && seamap[iprime][j] == 0 {
					changed = true
					auxmap[i][j] = 0
					auxmap[iprime][j] = 2
					i++
				} else {
					auxmap[i][j] = seamap[i][j]
				}
			}
		}

		seamap, auxmap = auxmap, seamap
	}
	fmt.Println(res)
}

func part2(v []string) {

	//res := 0
	//for i := 0; i < len(v); i++ {
	//
	//}
	//fmt.Println(res)

	// yay done !
}

func main() {

	content, err := ioutil.ReadFile("Day 25/input.txt")
	//content, err := ioutil.ReadFile("Day 25/smallinput.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
