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
	p1string := strings.Split(v[0], " ")
	p1pos, _ := strconv.Atoi(p1string[len(p1string)-1])
	p2string := strings.Split(v[1], " ")
	p2pos, _ := strconv.Atoi(p2string[len(p2string)-1])
	dieRolls := 0
	pawns := [2]int{p1pos, p2pos}
	scores := [2]int{0, 0}
	currPlayer := 0
	for scores[0] < 1000 && scores[1] < 1000 {
		moves := 0
		for j := 0; j < 3; j++ {
			moves += dieRolls%100 + 1
			dieRolls++
		}
		pawns[currPlayer] = (pawns[currPlayer]-1+moves)%10 + 1
		scores[currPlayer] += pawns[currPlayer]
		if currPlayer == 0 {
			currPlayer = 1
		} else {
			currPlayer = 0
		}
	}

	res = dieRolls * scores[currPlayer]
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	p1string := strings.Split(v[0], " ")
	p1pos, _ := strconv.Atoi(p1string[len(p1string)-1])
	p2string := strings.Split(v[1], " ")
	p2pos, _ := strconv.Atoi(p2string[len(p2string)-1])
	universes := make(map[[4]int]int)
	universes[[4]int{p1pos, p2pos, 0, 0}] = 1

	var auxUniverses map[[4]int]int
	auxUniverses = make(map[[4]int]int)
	for rounds := 0; rounds < 8; rounds++ { // after 20 turns on each side someone must have won
		//auxUniverses = make(map[[4]int]int)
		for coords, universeCount := range universes { // coords: p1pos, p2pos, p1score, p2score
			if coords[2] < 21 && coords[3] < 21 {
				for roll := 1; roll <= 3; roll++ {
					pawn := (coords[0]-1+roll)%10 + 1
					_, exists := auxUniverses[[4]int{pawn, coords[1], coords[2] + pawn, coords[3]}]
					if !exists {
						auxUniverses[[4]int{pawn, coords[1], coords[2] + pawn, coords[3]}] = 0
					}
					auxUniverses[[4]int{pawn, coords[1], coords[2] + pawn, coords[3]}] += universeCount
				}
			} else {
				auxUniverses[coords] = universeCount
			}
		}
		//universes = auxUniverses
		//auxUniverses = make(map[[4]int]int)

		// garbage collector issues??
		for coords := range universes {
			universes[coords] = 0
		}
		for coords, universeCount := range auxUniverses {
			universes[coords] = universeCount
			auxUniverses[coords] = 0
		}

		totalUniverses := 0
		for coords, universeCount := range universes {
			if coords[2] < 21 && coords[3] < 21 {
				for roll := 1; roll <= 3; roll++ {
					pawn := (coords[1]-1+roll)%10 + 1
					_, exists := auxUniverses[[4]int{coords[0], pawn, coords[2], coords[3] + pawn}]
					if !exists {
						auxUniverses[[4]int{coords[0], pawn, coords[2], coords[3] + pawn}] = 0
					}
					auxUniverses[[4]int{coords[0], pawn, coords[2], coords[3] + pawn}] += universeCount
				}
			} else {
				auxUniverses[coords] = universeCount
			}
			totalUniverses += universeCount
		}
		println(totalUniverses)

		//universes = auxUniverses
		for coords := range universes {
			universes[coords] = 0
		}
		for coords, universeCount := range auxUniverses {
			universes[coords] = universeCount
			auxUniverses[coords] = 0
		}
	}

	p1wins := 0
	p2wins := 0
	for coords, universeCount := range universes {
		if coords[2] > coords[3] {
			p1wins += universeCount
		} else {
			p2wins += universeCount
		}
	}

	if p1wins > p2wins {
		res = p1wins
	} else {
		res = p2wins
	}
	fmt.Println(res)
}

func part2arraeys(v []string) {
	res := 0
	p1string := strings.Split(v[0], " ")
	p1pos, _ := strconv.Atoi(p1string[len(p1string)-1])
	p2string := strings.Split(v[1], " ")
	p2pos, _ := strconv.Atoi(p2string[len(p2string)-1])
	universes := make([][][][]int, 10)
	auxUniverses := make([][][][]int, 10)

	for i := 0; i < 10; i++ {
		universes[i] = make([][][]int, 10)
		auxUniverses[i] = make([][][]int, 10)
		for j := 0; j < 10; j++ {
			universes[i][j] = make([][]int, 21)
			auxUniverses[i][j] = make([][]int, 21)
			for k := 0; k < 21; k++ {
				universes[i][j][k] = make([]int, 21)
				auxUniverses[i][j][k] = make([]int, 21)
			}
		}
	}

	universes[p1pos-1][p2pos-1][0][0] = 1

	p1wins := 0
	p2wins := 0
	for rounds := 0; rounds < 8; rounds++ { // after sufficient turns one side must win
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 21; k++ {
					for l := 0; l < 21; l++ {
						for roll := 1; roll <= 3; roll++ {
							pawn := (i+roll)%10 + 1
							score := k + pawn
							if score < 21 {
								auxUniverses[pawn-1][j][score][l] = universes[i][j][k][l]
							} else {
								p1wins += universes[i][j][k][l] // and make sure not to simulate next step
							}
						}
					}
				}
			}
		}

		universes, auxUniverses = auxUniverses, universes
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 21; k++ {
					for l := 0; l < 21; l++ {
						auxUniverses[i][j][k][l] = 0
					}
				}
			}
		}

		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 21; k++ {
					for l := 0; l < 21; l++ {
						for roll := 1; roll <= 3; roll++ {
							pawn := (j+roll)%10 + 1
							score := l + pawn
							if score < 21 {
								auxUniverses[i][pawn-1][k][score] = universes[i][j][k][l]
							} else {
								p2wins += universes[i][j][k][l] // and make sure not to simulate next step
							}
						}
					}
				}
			}
		}

		universes, auxUniverses = auxUniverses, universes
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				for k := 0; k < 21; k++ {
					for l := 0; l < 21; l++ {
						auxUniverses[i][j][k][l] = 0
					}
				}
			}
		}
	}

	if p1wins > p2wins {
		res = p1wins
	} else {
		res = p2wins
	}
	fmt.Println(res)
}

func coordToString(c [4]int) string {
	strcoords := make([]string, 4)
	for i := 0; i < 4; i++ {
		strcoords[i] = strconv.Itoa(c[i])
	}
	return strings.Join(strcoords, " ")
}

func stringToCoord(s string) [4]int {
	pieces := strings.Split(s, " ")
	coords := [4]int{}
	for i := 0; i < 4; i++ {
		coords[i], _ = strconv.Atoi(pieces[i])
	}
	return coords
}

func part2stringmaps(v []string) { // maybe golang just doesn't like [4]int as keys?

	res := 0
	p1string := strings.Split(v[0], " ")
	p1pos, _ := strconv.Atoi(p1string[len(p1string)-1])
	p2string := strings.Split(v[1], " ")
	p2pos, _ := strconv.Atoi(p2string[len(p2string)-1])
	p1wins := 0
	p2wins := 0
	universes := make(map[string]int)
	universes[coordToString([4]int{p1pos, p2pos, 0, 0})] = 1

	var auxUniverses map[string]int
	auxUniverses = make(map[string]int)
	for rounds := 0; rounds < 8; rounds++ { // after 20 turns on each side someone must have won
		auxUniverses = make(map[string]int)
		//strcoords: p1pos, p2pos, p1score, p2score
		for strcoords, universeCount := range universes {
			coords := stringToCoord(strcoords)
			if coords[2] < 21 && coords[3] < 21 {
				for roll := 1; roll <= 3; roll++ {
					pawn := (coords[0]-1+roll)%10 + 1
					_, exists := auxUniverses[coordToString([4]int{pawn, coords[1], coords[2] + pawn, coords[3]})]
					if !exists {
						auxUniverses[coordToString([4]int{pawn, coords[1], coords[2] + pawn, coords[3]})] = 0
					}
					auxUniverses[coordToString([4]int{pawn, coords[1], coords[2] + pawn, coords[3]})] += universeCount
				}
			} else {
				if coords[2] > coords[3] {
					p1wins += universeCount
				} else {
					p2wins += universeCount
				}
			}
		}
		universes = auxUniverses
		auxUniverses = make(map[string]int)

		//// garbage collector issues??
		//for strcoords := range universes {
		//	universes[strcoords] = 0
		//}
		//for strcoords, universeCount := range auxUniverses {
		//	universes[strcoords] = universeCount
		//	auxUniverses[strcoords] = 0
		//}

		totalUniverses := 0
		for strcoords, universeCount := range universes {
			coords := stringToCoord(strcoords)
			if coords[2] < 21 && coords[3] < 21 {
				for roll := 1; roll <= 3; roll++ {
					pawn := (coords[1]-1+roll)%10 + 1
					_, exists := auxUniverses[coordToString([4]int{coords[0], pawn, coords[2], coords[3] + pawn})]
					if !exists {
						auxUniverses[coordToString([4]int{coords[0], pawn, coords[2], coords[3] + pawn})] = 0
					}
					auxUniverses[coordToString([4]int{coords[0], pawn, coords[2], coords[3] + pawn})] += universeCount
				}
			} else {
				if coords[2] > coords[3] {
					p1wins += universeCount
				} else {
					p2wins += universeCount
				}
			}
			totalUniverses += universeCount
		}
		println(totalUniverses)

		universes = auxUniverses
		//for coords := range universes {
		//	universes[coords] = 0
		//}
		//for coords, universeCount := range auxUniverses {
		//	universes[coords] = universeCount
		//	auxUniverses[coords] = 0
		//}
	}

	//p1wins := 0
	//p2wins := 0
	//for coords, universeCount := range universes {
	//	if coords[2] > coords[3] {
	//		p1wins += universeCount
	//	} else {
	//		p2wins += universeCount
	//	}
	//}

	if p1wins > p2wins {
		res = p1wins
	} else {
		res = p2wins
	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 21/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	//part2(v)
	//part2arraeys(v)
	part2stringmaps(v)
}
