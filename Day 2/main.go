package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(v []string) {
	h := 0
	d := 0
	for i := 0; i < len(v); i++ {
		cmd := strings.Split(v[i], " ")
		if cmd[0] == "forward" {
			v, _ := strconv.Atoi(cmd[1])
			h += v
		}
		if cmd[0] == "backward" {
			v, _ := strconv.Atoi(cmd[1])
			h -= v
		}
		if cmd[0] == "up" {
			v, _ := strconv.Atoi(cmd[1])
			d -= v
		}
		if cmd[0] == "down" {
			v, _ := strconv.Atoi(cmd[1])
			d += v
		}
	}
	fmt.Println(d * h)
}

func part2(v []string) {

	h := 0
	d := 0
	aim := 0
	for i := 0; i < len(v); i++ {
		cmd := strings.Split(v[i], " ")
		if cmd[0] == "forward" {
			v, _ := strconv.Atoi(cmd[1])
			h += v
			d += v * aim
		}
		if cmd[0] == "backward" {
			v, _ := strconv.Atoi(cmd[1])
			h -= v
			d -= v * aim
		}
		if cmd[0] == "up" {
			v, _ := strconv.Atoi(cmd[1])
			//d -= v
			aim -= v
		}
		if cmd[0] == "down" {
			v, _ := strconv.Atoi(cmd[1])
			//d += v
			aim += v
		}
	}
	fmt.Println(d * h)
}

func main() {

	content, err := ioutil.ReadFile("Day 2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)

	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
