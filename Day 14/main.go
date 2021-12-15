package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

func part1(v []string) {

	chain := []uint8(v[0])
	res := 0
	rules := strings.Split(v[1], "\n")
	for step := 1; step <= 10; step++ {
		newstring := []uint8{}
	NextStep:
		for i := 0; i < len(chain)-1; i++ {
			for _, rule := range rules {
				rulesplit := strings.Split(rule, " -> ")
				priori := rulesplit[0]
				posteriori := rulesplit[1]
				if chain[i] == priori[0] && chain[i+1] == priori[1] {
					newstring = append(newstring, priori[0], posteriori[0])
					continue NextStep
				}
			}
			newstring = append(newstring, chain[i])
		}
		newstring = append(newstring, chain[len(chain)-1])
		chain = newstring
	}

	frequencies := make(map[uint8]int)
	for i := 0; i < len(chain); i++ {
		_, exists := frequencies[chain[i]]
		if !exists {
			frequencies[chain[i]] = 0
		}
		frequencies[chain[i]]++
	}

	minfreq := len(chain) + 1
	maxfreq := 0
	for _, value := range frequencies {
		if value < minfreq {
			minfreq = value
		} else if value > maxfreq {
			maxfreq = value
		}
	}
	res = maxfreq - minfreq
	fmt.Println(res)
}

type SteppedPolymer struct {
	Element uint8
	Step    int
}

// I'm assuming here maps are passed by reference in Golang
func findreplacement(input []uint8, stepdelta int, dyntable map[[2]uint8][][]uint8) []uint8 {

	output := make([]uint8, len(input)*int(math.Pow(2, float64(stepdelta))))
	for i := 0; i < len(input)-1; i++ {
		pair := [2]uint8{input[i], input[i+1]}
		replacementarray := dyntable[pair]
		// find the calculated replacement that forwards time the most, then continue calculating

	}

}

func part2(v []string) {

	chain := []uint8(v[0])
	res := 0
	rules := strings.Split(v[1], "\n")

	// we face two challenges: 1) it is intractable to allot enough time to naively generate the whole string
	// 2) even if we are able to replace parts of the original string directly with their equivalents 40 steps later,
	// it costs > 1TB to just to hold it in memory.
	dynrules := make(map[[2]uint8][][]uint8, len(rules))
	for _, rule := range rules {
		rulesteps := make([][]uint8, 0, 40)
		rulesplit := strings.Split(rule, " -> ")
		priori := rulesplit[0]
		posteriori := rulesplit[1]
		rulesteps = append(rulesteps, []uint8{priori[0], posteriori[0], priori[1]})
		dynrules[[2]uint8{priori[0], priori[1]}] = rulesteps
	}
	// populate dynamic programming table
	chain = findreplacement(chain, 40, dynrules)

	// set up steps

	frequencies := make(map[uint8]int)
	for i := 0; i < len(chain); i++ {
		_, exists := frequencies[chain[i]]
		if !exists {
			frequencies[chain[i]] = 0
		}
		frequencies[chain[i]]++
	}

	minfreq := len(chain) + 1
	maxfreq := 0
	for _, value := range frequencies {
		if value < minfreq {
			minfreq = value
		} else if value > maxfreq {
			maxfreq = value
		}
	}
	res = maxfreq - minfreq
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 14/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n\n")
	part1(v)
	part2(v)
}
