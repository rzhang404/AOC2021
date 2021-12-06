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
	//fish_at_timer := make(map[int]int)
	fish_at_timer := make([]int, 9)
	for i := 0; i < len(v); i++ {
		fish, _ := strconv.Atoi(v[i])
		fish_at_timer[fish]++
	}

	for i := 0; i < 80; i++ {
		newfish := fish_at_timer[0]
		for j := 0; j < 8; j++ {
			fish_at_timer[j] = fish_at_timer[j+1]
		}
		fish_at_timer[8] = newfish
		fish_at_timer[6] += newfish
	}

	for i := 0; i < 9; i++ {
		res += fish_at_timer[i]
	}
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	//fish_at_timer := make(map[int]int)
	fish_at_timer := make([]int, 9)
	for i := 0; i < len(v); i++ {
		fish, _ := strconv.Atoi(v[i])
		fish_at_timer[fish]++
	}

	for i := 0; i < 256; i++ {
		newfish := fish_at_timer[0]
		for j := 0; j < 8; j++ {
			fish_at_timer[j] = fish_at_timer[j+1]
		}
		fish_at_timer[8] = newfish
		fish_at_timer[6] += newfish
	}

	for i := 0; i < 9; i++ {
		res += fish_at_timer[i]
	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 6/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, ",")
	part1(v)
	part2(v)
}
