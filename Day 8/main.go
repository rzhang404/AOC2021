package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func part1(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {
		line := strings.Split(v[i], " | ")
		outputs := strings.Split(line[1], " ")
		for j := 0; j < len(outputs); j++ {
			if (len(outputs[j]) > 1 && len(outputs[j]) < 5) || len(outputs[j]) == 7 {
				res++
			}
		}
	}
	fmt.Println(res)
}

type ByLen []string

func (a ByLen) Len() int {
	return len(a)
}

func (a ByLen) Less(i, j int) bool {
	return len(a[i]) < len(a[j])
}

func (a ByLen) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

//MakeSet initialize the set
func makeSet() *customSet {
	return &customSet{
		container: make(map[string]struct{}),
	}
}

type customSet struct {
	container map[string]struct{}
}

func (c *customSet) Exists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *customSet) Add(key string) {
	c.container[key] = struct{}{}
}

func (c *customSet) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Remove Error: Item doesn't exist in set")
	}
	delete(c.container, key)
	return nil
}

func (c *customSet) Size() int {
	return len(c.container)
}

func part2(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {
		line := strings.Split(v[i], " | ")
		outputs := strings.Split(line[1], " ")
		inputs := strings.Split(line[0], " ")
		wireslice := []string{"a", "b", "c", "d", "e", "f", "g"}
		wiremap := make(map[string]customSet)
		wiremap["a"] = *makeSet()
		wiremap["b"] = *makeSet()
		wiremap["c"] = *makeSet()
		wiremap["d"] = *makeSet()
		wiremap["e"] = *makeSet()
		wiremap["f"] = *makeSet()
		wiremap["g"] = *makeSet()

		// initialize
		for _, wireset := range wiremap {
			wireset.Add("a")
			wireset.Add("b")
			wireset.Add("c")
			wireset.Add("d")
			wireset.Add("e")
			wireset.Add("f")
			wireset.Add("g")
		}

		sort.Sort(ByLen(inputs))

		// 1
		wireset := makeSet()
		wireset.Add("a")
		wireset.Add("b")
		wireset.Add("c")
		wireset.Add("d")
		wireset.Add("e")
		wireset.Add("f")
		wireset.Add("g")
		for _, wire := range wireslice {
			signalset := wiremap[wire]
			if !strings.Contains(inputs[0], wire) {
				signalset.Remove("c")
				signalset.Remove("f")
			} else {
				signalset.Remove("a")
				signalset.Remove("b")
				signalset.Remove("d")
				signalset.Remove("e")
				signalset.Remove("g")
			}
		}

		// 7
		for _, wire := range wireslice {
			signalset := wiremap[wire]
			if !strings.Contains(inputs[1], wire) {
				signalset.Remove("a")
				signalset.Remove("c")
				signalset.Remove("f")
			} else {
				signalset.Remove("b")
				signalset.Remove("d")
				signalset.Remove("e")
				signalset.Remove("g")
			}
		}

		// 4
		for _, wire := range wireslice {
			signalset := wiremap[wire]
			if !strings.Contains(inputs[2], wire) {
				signalset.Remove("b")
				signalset.Remove("c")
				signalset.Remove("d")
				signalset.Remove("f")
			} else {
				signalset.Remove("a")
				signalset.Remove("e")
				signalset.Remove("g")
			}
		}

		// ignore 8

		//for j := 0; j < 7; j++ {
		//	var mappedwire string
		//	for _, wire := range wireslice {
		//		signalset := wiremap[wire]
		//		if signalset.Size() == 1 {
		//			for _, wire2 := range wireslice {
		//				if signalset.Exists(wire2) {
		//					mappedwire = wire2
		//				}
		//			}
		//		}
		//	}
		//}

		for j := 0; j < len(outputs); j++ {
			if (len(outputs[j]) > 1 && len(outputs[j]) < 5) || len(outputs[j]) == 7 {
				res++
			}
		}
	}
	fmt.Println(res)
}

func permutestrings(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func permutateints(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func sortString(input string) string {
	runeArray := []rune(input)
	sort.Sort(sortRuneString(runeArray))
	return string(runeArray)
}

type sortRuneString []rune

func (s sortRuneString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRuneString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRuneString) Len() int {
	return len(s)
}

func part2try2(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {
		line := strings.Split(v[i], " | ")
		outputs := strings.Split(line[1], " ")
		inputs := strings.Split(line[0], " ")

		letters := []string{"a", "b", "c", "d", "e", "f", "g"}
		// correct numbers from 0 to 9
		correctinputs := []string{"abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdefg"}
		sort.Strings(correctinputs)
		var sortedcorrectinputs [10]string
		for j, correctinput := range correctinputs {
			sortedcorrectinputs[j] = correctinput
		}

		perms := permutestrings(letters)
		var mappedinputs, mappedoutputs []string
		for _, perm := range perms {
			testinputs := inputs
			testoutputs := outputs
			for j, letter := range letters {
				// replace with mapping
				for k, testinput := range testinputs {
					testinputs[k] = strings.Replace(testinput, letter, perm[j], -1) // ooohh... this doesn't work
				}
				for k, testoutput := range testoutputs {
					testoutputs[k] = strings.Replace(testoutput, letter, perm[j], -1)
				}
			}
			var sortedinputs [10]string
			for j, input := range inputs {
				sortedinputs[j] = sortString(input)
			}
			if sortedinputs == sortedcorrectinputs {
				mappedinputs = testinputs
				mappedoutputs = testoutputs
				break
			}
		}

		println(mappedinputs)
		var sortedoutputs [4]string

		for j, mappedoutput := range mappedoutputs {
			sortedoutputs[j] = sortString(mappedoutput)
		}

		for _, sortedoutput := range sortedoutputs {
			for j, correctinput := range correctinputs {
				if sortedoutput == correctinput {
					res += j
				}
			}
		}
	}
	fmt.Println(res)
}

func part2try3(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {
		line := strings.Split(v[i], " | ")
		outputs := strings.Split(line[1], " ")
		inputs := strings.Split(line[0], " ")

		letters := []string{"a", "b", "c", "d", "e", "f", "g"}
		powers := []int{1, 2, 4, 8, 16, 32, 64}
		digitmap := []int{1 + 2 + 4 + 16 + 32 + 64, 4 + 32, 1 + 4 + 8 + 16 + 64, 1 + 4 + 8 + 32 + 64, 2 + 4 + 8 + 32,
			1 + 2 + 8 + 32 + 64, 1 + 2 + 8 + 16 + 32 + 64, 1 + 4 + 32, 1 + 2 + 4 + 8 + 16 + 32 + 64, 1 + 2 + 4 + 8 + 32 + 64}
		sorteddigitmap := make([]int, 10)
		copy(sorteddigitmap, digitmap)
		sort.Ints(sorteddigitmap)
		var truncsortdigimap [10]int
		for j, mapnum := range sorteddigitmap {
			truncsortdigimap[j] = mapnum
		}

		// permute
		perms := permutateints(powers)
		mappedinputs := make([]int, 10)
		mappedoutputs := make([]int, 4)
		var truncmap [10]int
		// for each permutation, replace input values
		for _, perm := range perms {
			for j, input := range inputs {
				mappednum := 0
				for k, letter := range letters {
					if strings.Contains(input, letter) {
						mappednum += perm[k]
					}
				}
				mappedinputs[j] = mappednum
			}
			sort.Ints(mappedinputs)

			for j := range truncmap {
				truncmap[j] = mappedinputs[j]
			}

			// find the correct one, then replace output values
			if truncmap == truncsortdigimap {
				outval := 0
				for j, output := range outputs {
					outval *= 10
					mappednum := 0
					for k, letter := range letters {
						if strings.Contains(output, letter) {
							mappednum += perm[k]
						}
					}
					mappedoutputs[j] = mappednum
					for digit, digimapnum := range digitmap {
						if digimapnum == mappednum {
							outval += digit
						}
					}
				}
				res += outval
				break
			}
		}

		// add output values to res

	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 8/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2try3(v)
}
