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
	prev, _ := strconv.Atoi(v[0])
	for i := 1; i < len(v); i++ {
		curr, _ := strconv.Atoi(v[i])
		if curr > prev {
			res++
		}
		prev = curr
	}
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	prev, _ := strconv.Atoi(v[2])
	prev2, _ := strconv.Atoi(v[1])
	prev3, _ := strconv.Atoi(v[0])
	for i := 3; i < len(v); i++ {
		curr, _ := strconv.Atoi(v[i])
		if curr+prev+prev2 > prev+prev2+prev3 {
			res++
		}
		prev3 = prev2
		prev2 = prev
		prev = curr
	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 1/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
