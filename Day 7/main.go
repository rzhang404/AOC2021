package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func part1(v []string) {

	res := float64(0)
	horiz := make([]int, len(v))
	total := 0
	for i := 0; i < len(v); i++ {
		horiz[i], _ = strconv.Atoi(v[i])
		total += horiz[i]
	}

	sort.Ints(horiz)
	median := horiz[len(v)/2]

	for i := 0; i < len(horiz); i++ {
		res += math.Abs(float64(horiz[i] - median))
	}

	fmt.Println(res)
}

func part2(v []string) {

	horiz := make([]int, len(v))
	total := 0
	for i := 0; i < len(v); i++ {
		horiz[i], _ = strconv.Atoi(v[i])
		total += horiz[i]
	}

	trinums := make([]int, 2000)
	trinums[0] = 0
	for i := 1; i < 2000; i++ {
		trinums[i] = trinums[i-1] + i
	}

	sort.Ints(horiz)

	minfuel := len(v) * trinums[len(trinums)-1]
	for i := 0; i < 2000; i++ {
		currfuel := 0
		for j := 0; j < len(horiz); j++ {
			currfuel += trinums[int(math.Abs(float64(horiz[j]-i)))]
		}
		if currfuel < minfuel {
			minfuel = currfuel
		}
	}

	fmt.Println(minfuel)
}

func main() {

	content, err := ioutil.ReadFile("Day 7/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, ",")
	part1(v)
	part2(v)
}
