package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func part1(v []string) {

	ones := 0
	zeroes := 0
	gamma := 0
	epsilon := 0
	for j := 0; j < len(v[0]); j++ {
		for i := 0; i < len(v)-1; i++ {
			str := v[i][j]
			if str == 49 {
				ones++
			} else {
				zeroes++
			}
		}
		gamma = gamma * 2
		epsilon = epsilon * 2
		if ones > zeroes {
			gamma++
		} else {
			epsilon++
		}
		ones = 0
		zeroes = 0
	}
	fmt.Println(gamma * epsilon)
}

func part2(v []string) {

	ones := 0
	zeroes := 0
	oxygen := 0
	co2 := 0
	most_commons := v[:len(v)-1]
	least_commons := v[:len(v)-1]
	for j := 0; j < len(most_commons[0]); j++ {
		for i := 0; i < len(most_commons); i++ {
			str := most_commons[i][j]
			if str == 49 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones >= zeroes {
			mcnext := []string{}
			for i := range most_commons {
				if most_commons[i][j] == 49 {
					mcnext = append(mcnext, most_commons[i])
				}
			}
			most_commons = mcnext
		} else {
			mcnext := []string{}
			for i := range most_commons {
				if most_commons[i][j] == 48 {
					mcnext = append(mcnext, most_commons[i])
				}
			}
			most_commons = mcnext
		}
		ones = 0
		zeroes = 0
		if len(most_commons) == 1 {
			break
		}
	}

	for j := 0; j < len(least_commons[0]); j++ {
		for i := 0; i < len(least_commons); i++ {
			str := least_commons[i][j]
			if str == 49 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones < zeroes {
			lcnext := []string{}
			for i := range least_commons {
				if least_commons[i][j] == 49 {
					lcnext = append(lcnext, least_commons[i])
				}
			}
			least_commons = lcnext
		} else {
			lcnext := []string{}
			for i := range least_commons {
				if least_commons[i][j] == 48 {
					lcnext = append(lcnext, least_commons[i])
				}
			}
			least_commons = lcnext
		}
		ones = 0
		zeroes = 0

		if len(least_commons) == 1 {
			break
		}
	}
	for i := 0; i < len(most_commons[0]); i++ {
		oxygen = oxygen * 2
		if most_commons[0][i] == 49 {
			oxygen++
		}
	}
	for i := 0; i < len(least_commons[0]); i++ {
		co2 = co2 * 2
		if least_commons[0][i] == 49 {
			co2++
		}
	}
	fmt.Println(oxygen * co2)
}

func main() {

	content, err := ioutil.ReadFile("Day 3/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
