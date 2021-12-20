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
	enhancementAlgorithm := ""
	for i := 0; i < len(v[0]); i++ {
		if string(v[0][i]) == "#" {
			enhancementAlgorithm = enhancementAlgorithm + "1"
		} else {
			enhancementAlgorithm = enhancementAlgorithm + "0"
		}
	}
	imageInput := v[2:]

	leftEdge := 0
	rightEdge := len(imageInput[0])
	topEdge := 0
	bottomEdge := len(imageInput)
	defaultVal := "0" // pixels outside of input's influence changes between 0 and 1 due to enhancement every iteration
	image := make(map[[2]int]string)
	for i := 0; i < bottomEdge; i++ {
		for j := 0; j < rightEdge; j++ {
			if string(imageInput[i][j]) == "#" {
				image[[2]int{i, j}] = "1"
			} else {
				image[[2]int{i, j}] = "0"
			}
		}
	}

	auxImage := make(map[[2]int]string)
	for iter := 0; iter < 50; iter++ {
		leftEdge--
		rightEdge++
		topEdge--
		bottomEdge++
		for i := topEdge; i < bottomEdge; i++ {
			for j := leftEdge; j < rightEdge; j++ {
				binString := ""
				for k := -1; k <= 1; k++ {
					for l := -1; l <= 1; l++ {
						val, exists := image[[2]int{i + k, j + l}]
						if !exists {
							val = defaultVal
						}
						binString = binString + val
					}
				}

				// locate enhancement
				enhanceIndex, _ := strconv.ParseInt(binString, 2, 16)
				auxImage[[2]int{i, j}] = string(enhancementAlgorithm[enhanceIndex])
			}
		}
		image, auxImage = auxImage, image
		if defaultVal == "0" {
			defaultVal = "1"
		} else {
			defaultVal = "0"
		}
	}

	for i := leftEdge; i < rightEdge; i++ {
		for j := topEdge; j < bottomEdge; j++ {
			if image[[2]int{i, j}] == "1" {
				res++
			}
		}
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

	content, err := ioutil.ReadFile("Day 20/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
