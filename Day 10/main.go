package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"sort"
	"strings"
)

func part1(v []string) {

	res := 0
OuterLoop:
	for i := 0; i < len(v); i++ {
		line := make([]string, 0, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			char := string(v[i][j])
			if string(v[i][j]) == "(" || string(v[i][j]) == "[" || string(v[i][j]) == "{" || string(v[i][j]) == "<" {
				line = append(line, string(v[i][j]))
			} else if char == ")" {
				if line[len(line)-1] == "(" {
					line = line[:len(line)-1]
				} else {
					res += 3
					continue OuterLoop
				}
			} else if char == "]" {
				if line[len(line)-1] == "[" {
					line = line[:len(line)-1]
				} else {
					res += 57
					continue OuterLoop
				}
			} else if char == "}" {
				if line[len(line)-1] == "{" {
					line = line[:len(line)-1]
				} else {
					res += 1197
					continue OuterLoop
				}
			} else if char == ">" {
				if line[len(line)-1] == "<" {
					line = line[:len(line)-1]
				} else {
					res += 25137
					continue OuterLoop
				}
			}
		}
	}
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	verifieds := make([][]string, 0, len(v))
OuterLoop:
	for i := 0; i < len(v); i++ {
		line := make([]string, 0, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			char := string(v[i][j])
			if string(v[i][j]) == "(" || string(v[i][j]) == "[" || string(v[i][j]) == "{" || string(v[i][j]) == "<" {
				line = append(line, string(v[i][j]))
			} else if char == ")" {
				if line[len(line)-1] == "(" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == "]" {
				if line[len(line)-1] == "[" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == "}" {
				if line[len(line)-1] == "{" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == ">" {
				if line[len(line)-1] == "<" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			}
		}
		verifieds = append(verifieds, line)
	}

	scores := make([]int, 0, len(verifieds))
	for i := 0; i < len(verifieds); i++ {
		var score int
		score = 0
		for j := len(verifieds[i]) - 1; j >= 0; j-- {
			score *= 5
			char := string(verifieds[i][j])
			if char == "(" {
				score += 1
			} else if char == "[" {
				score += 2
			} else if char == "{" {
				score += 3
			} else if char == "<" {
				score += 4
			}
		}
		scores = append(scores, score)
	}
	sort.Ints(scores)
	res = scores[len(scores)/2]
	fmt.Println(res)
}

/** Wrong implementation assuming big ints are required
func part2(v []string) {

	verifieds := make([]string, 0, len(v))
OuterLoop:
	for i := 0; i < len(v); i++ {
		line := make([]string, 0, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			char := string(v[i][j])
			if string(v[i][j]) == "(" || string(v[i][j]) == "[" || string(v[i][j]) == "{" || string(v[i][j]) == "<" {
				line = append(line, string(v[i][j]))
			} else if char == ")" {
				if line[len(line)-1] == "(" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == "]" {
				if line[len(line)-1] == "[" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == "}" {
				if line[len(line)-1] == "{" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			} else if char == ">" {
				if line[len(line)-1] == "<" {
					line = line[:len(line)-1]
				} else {
					continue OuterLoop
				}
			}
		}
		verifieds = append(verifieds, v[i])
	}

	var scores BI
	for i := 0; i < len(verifieds); i++ {
		var score *big.Int
		var aux *big.Int
		score = big.NewInt(0)
		for j := len(verifieds[i])-1; j >= 0; j-- {
			aux = big.NewInt(0)
			aux.Mul(score, big.NewInt(5))
			score = aux
			char := string(verifieds[i][j])
			if char == "(" {
				aux = big.NewInt(0)
				aux.Add(score, big.NewInt(1))
				score = aux
			} else if char == "[" {
				aux = big.NewInt(0)
				aux.Add(score, big.NewInt(2))
				score = aux
			} else if char == "{" {
				aux = big.NewInt(0)
				aux.Add(score, big.NewInt(3))
				score = aux
			} else if char == "<" {
				aux = big.NewInt(0)
				aux.Add(score, big.NewInt(4))
				score = aux
			}
		}
		scores = append(scores, score)
	}
	sort.Sort(scores)
	fmt.Println(scores[len(scores)/2])
}
**/

type BI []*big.Int

func (bi BI) Len() int {
	return len(bi)
}

func (bi BI) Less(i, j int) bool {
	checkval := big.NewInt(0)
	checkval = checkval.Sub(bi[i], bi[j])
	if checkval.Sign() == -1 {
		return true
	} else {
		return false
	}
}

func (bi BI) Swap(i, j int) {
	bi[i], bi[j] = bi[j], bi[i]
}

func main() {

	content, err := ioutil.ReadFile("Day 10/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
