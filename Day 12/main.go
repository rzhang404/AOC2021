package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func deepcopymap(mapsrc map[string]bool) map[string]bool {
	newmap := make(map[string]bool)
	for key, value := range mapsrc {
		newmap[key] = value
	}
	return newmap
}

func traversept1(cavemap map[string][]string, nextnode string, visited map[string]bool, pathsfound int) int { // add path for part 2 likely
	if nextnode == "end" {
		return pathsfound + 1
	}
	neighbours, exists := cavemap[nextnode]
	if !exists {
		// if no neighbours, no valid paths, return, but is impossible given input
		return pathsfound
	}

	uppercase := true
	for _, letter := range nextnode {
		if int(letter) >= 97 && int(letter) <= 122 {
			uppercase = false
		}
	}

	visitedmodified := deepcopymap(visited)

	// update visited
	if !uppercase {
		visitedmodified[nextnode] = true
	}
	visitedcopy := deepcopymap(visitedmodified)

	newpathsfound := pathsfound
	for _, neighbour := range neighbours {
		if !visitedcopy[neighbour] {
			newpathsfound = traversept1(cavemap, neighbour, visitedcopy, newpathsfound)
		}
		visitedcopy = deepcopymap(visitedmodified)
	}
	return newpathsfound
}

func traversept2(cavemap map[string][]string, nextnode string,
	visited map[string]bool, smallcavevisited bool, pathsfound int) int {
	if nextnode == "end" {
		return pathsfound + 1
	}
	neighbours, exists := cavemap[nextnode]
	if !exists {
		// if no neighbours, no valid paths, return, but is impossible given input
		return pathsfound
	}

	uppercase := true
	for _, letter := range nextnode {
		if int(letter) >= 97 && int(letter) <= 122 {
			uppercase = false
		}
	}

	visitedmodified := deepcopymap(visited)

	// update visited
	if !uppercase {
		visitedmodified[nextnode] = true
	}
	visitedcopy := deepcopymap(visitedmodified)

	newpathsfound := pathsfound
	for _, neighbour := range neighbours {
		if !visitedcopy[neighbour] {
			newpathsfound = traversept2(cavemap, neighbour, visitedcopy, smallcavevisited, newpathsfound)
		} else if visitedcopy[neighbour] && !smallcavevisited && neighbour != "start" && neighbour != "end" {
			newpathsfound = traversept2(cavemap, neighbour, visitedcopy, true, newpathsfound)
		}
		visitedcopy = deepcopymap(visitedmodified)
	}
	return newpathsfound
}

func part1(v []string) {

	res := 0
	cavemap := make(map[string][]string)
	for i := 0; i < len(v); i++ {
		edge := v[i]
		edgesplit := strings.Split(edge, "-")
		endpoint1 := edgesplit[0]
		endpoint2 := edgesplit[1]
		neighbours1, exists := cavemap[endpoint1]
		if !exists {
			neighbours1 = make([]string, 0, 10)
		}
		neighbours2, exists := cavemap[endpoint2]
		if !exists {
			neighbours2 = make([]string, 0, 10)
		}
		cavemap[endpoint1] = append(neighbours1, endpoint2)
		cavemap[endpoint2] = append(neighbours2, endpoint1)
	}
	visited := make(map[string]bool)
	for endpoint := range cavemap {
		visited[endpoint] = false
	}
	res = traversept1(cavemap, "start", visited, 0)
	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	cavemap := make(map[string][]string)
	for i := 0; i < len(v); i++ {
		edge := v[i]
		edgesplit := strings.Split(edge, "-")
		endpoint1 := edgesplit[0]
		endpoint2 := edgesplit[1]
		neighbours1, exists := cavemap[endpoint1]
		if !exists {
			neighbours1 = make([]string, 0, 10)
		}
		neighbours2, exists := cavemap[endpoint2]
		if !exists {
			neighbours2 = make([]string, 0, 10)
		}
		cavemap[endpoint1] = append(neighbours1, endpoint2)
		cavemap[endpoint2] = append(neighbours2, endpoint1)
	}
	visited := make(map[string]bool)
	for endpoint := range cavemap {
		visited[endpoint] = false
	}
	res = traversept2(cavemap, "start", visited, false, 0)
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 12/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
