package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func part1(v []string) {
	draw_order := strings.Split(v[0], ",")
	cards := v[1:]
	var boards = make([][]int, len(cards))
	for i := 0; i < len(cards); i++ {
		boards[i] = make([]int, 25)
		rows := strings.Split(cards[i], "\n")
		for j := 0; j < 5; j++ {
			row := strings.Fields(rows[j])
			for k := 0; k < 5; k++ {
				boards[i][j*5+k], _ = strconv.Atoi(row[k])
			}
		}
	}

	res := 0
	var winning_board []int
	var winning_call int
GameLoop:
	for i := 0; i < len(draw_order); i++ {
		draw, _ := strconv.Atoi(draw_order[i])

		for j := 0; j < len(boards); j++ {
			for k := 0; k < 25; k++ {
				if boards[j][k] == draw {
					// mark new number
					boards[j][k] = 100

					// check wins
					// check row
					winner := true
					rownum := k / 5
				RowLoop:
					for l := 0; l < 5; l++ {
						if boards[j][rownum*5+l] != 100 {
							winner = false
							break RowLoop
						}
					}

					if winner {
						winning_board = boards[j]
						winning_call = draw
						break GameLoop
					}

					// check column
					winner = true
					colnum := k % 5
				ColumnLoop:
					for l := 0; l < 5; l++ {
						if boards[j][l*5+colnum] != 100 {
							winner = false
							break ColumnLoop
						}
					}

					if winner {
						winning_board = boards[j]
						winning_call = draw
						break GameLoop
					}
				}
			}
		}
	}

	for i := 0; i < 25; i++ {
		if winning_board[i] != 100 {
			res += winning_board[i]
		}
	}

	res *= winning_call

	fmt.Println(res)
}

func part2(v []string) {

	draw_order := strings.Split(v[0], ",")
	cards := v[1:]
	var boards = make([][]int, len(cards))
	for i := 0; i < len(cards); i++ {
		boards[i] = make([]int, 25)
		rows := strings.Split(cards[i], "\n")
		for j := 0; j < 5; j++ {
			row := strings.Fields(rows[j])
			for k := 0; k < 5; k++ {
				boards[i][j*5+k], _ = strconv.Atoi(row[k])
			}
		}
	}

	res := 0
	var win_order [][]int
	var last_winning_call int
GameLoop:
	for i := 0; i < len(draw_order); i++ {
		draw, _ := strconv.Atoi(draw_order[i])

	BoardLoop:
		for j := 0; j < len(boards); j++ {
			for k := 0; k < 25; k++ {
				if boards[j][k] == draw {
					// mark new number
					boards[j][k] = 100

					// check wins
					// check row
					winner := true
					rownum := k / 5
				RowLoop:
					for l := 0; l < 5; l++ {
						if boards[j][rownum*5+l] != 100 {
							winner = false
							break RowLoop
						}
					}

					if winner {
						win_order = append(win_order, boards[j])
						boards = append(boards[:j], boards[j+1:]...)
						j--
						if len(boards) == 0 {
							last_winning_call = draw
							break GameLoop
						} else {
							continue BoardLoop
						}
					}

					// check column
					winner = true
					colnum := k % 5
				ColumnLoop:
					for l := 0; l < 5; l++ {
						if boards[j][l*5+colnum] != 100 {
							winner = false
							break ColumnLoop
						}
					}

					if winner {
						win_order = append(win_order, boards[j])
						boards = append(boards[:j], boards[j+1:]...)
						j--
						if len(boards) == 0 {
							last_winning_call = draw
							break GameLoop
						} else {
							continue BoardLoop
						}
					}
				}
			}
		}
	}

	for i := 0; i < 25; i++ {
		if win_order[len(win_order)-1][i] != 100 {
			res += win_order[len(win_order)-1][i]
		}
	}

	res *= last_winning_call

	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 4/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n\n")
	part1(v)
	part2(v)
}
