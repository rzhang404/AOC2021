package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func part1(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {

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

	content, err := ioutil.ReadFile("Day 24/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
