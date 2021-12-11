package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func inbounds(arr [][]int, x int, y int) bool {
	width := len(arr[0])
	height := len(arr)

	return x >= 0 && x < width && y >= 0 && y < height
}

func part1(v []string) {
	res := 0
	cavemap := make([][]int, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
		}
		cavemap[i] = line
	}

	for day := 1; day <= 100; day++ {
		for i := 0; i < len(cavemap); i++ {
			for j := 0; j < len(cavemap[i]); j++ {
				cavemap[i][j]++
			}
		}
		flashed := true
		for flashed {
			flashed = false
			for i := 0; i < len(cavemap); i++ {
				for j := 0; j < len(cavemap[i]); j++ {
					if cavemap[i][j] > 9 {
						flashed = true
						res++
						cavemap[i][j] = 0
						for k := -1; k <= 1; k++ {
							for l := -1; l <= 1; l++ {
								if inbounds(cavemap, j+k, i+l) && cavemap[i+l][j+k] > 0 {
									cavemap[i+l][j+k]++
								}
							}
						}
					}
				}
			}
		}

	}
	fmt.Println(res)
}

func part2(v []string) {
	res := 0
	cavemap := make([][]int, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
		}
		cavemap[i] = line
	}

	for day := 1; ; day++ {
		for i := 0; i < len(cavemap); i++ {
			for j := 0; j < len(cavemap[i]); j++ {
				cavemap[i][j]++
			}
		}
		flashed := true
		for flashed {
			flashed = false
			for i := 0; i < len(cavemap); i++ {
				for j := 0; j < len(cavemap[i]); j++ {
					if cavemap[i][j] > 9 {
						flashed = true
						cavemap[i][j] = 0
						for k := -1; k <= 1; k++ {
							for l := -1; l <= 1; l++ {
								if inbounds(cavemap, j+k, i+l) && cavemap[i+l][j+k] > 0 {
									cavemap[i+l][j+k]++
								}
							}
						}
					}
				}
			}
		}

		synced := true
		for i := 0; i < len(cavemap); i++ {
			for j := 0; j < len(cavemap[i]); j++ {
				if cavemap[i][j] > 0 {
					synced = false
				}
			}
		}
		if synced {
			res = day
			break
		}

	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 11/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
