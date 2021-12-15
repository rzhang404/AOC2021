package main

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func inbounds(arr [][]int, y int, x int) bool {
	width := len(arr[0])
	height := len(arr)

	return x >= 0 && x < width && y >= 0 && y < height
}

// The following is with credit to https://pkg.go.dev/container/heap#example__priorityQueue
// An Item is something we manage in a priority queue.
type Item struct {
	value    []int // The value of the item; arbitrary.
	priority int   // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value []int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func part1(v []string) {

	res := 0
	cavemap := make([][]int, len(v))
	seenmap := make([][]bool, len(v))
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		seenline := make([]bool, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
			seenline[j] = false
		}
		cavemap[i] = line
		seenmap[i] = seenline
	}

	// make priorityqueue
	pq := make(PriorityQueue, 0)
	// add topleft corner to pq
	heap.Init(&pq)
	hi := 0
	heap.Push(&pq, &Item{[]int{0, 0}, 0, hi})
	hi++

	// while true:
	for {
		// pop item from priority queue, if bottomright corner break
		curr := heap.Pop(&pq).(*Item)
		hi--
		seenmap[curr.value[0]][curr.value[1]] = true
		if curr.value[0] == len(cavemap)-1 && curr.value[1] == len(cavemap[len(cavemap)-1])-1 {
			res = curr.priority
			break
		} else {
			// if not add all neighbours
			possneighbours := [][]int{{curr.value[0] - 1, curr.value[1]}, {curr.value[0] + 1, curr.value[1]},
				{curr.value[0], curr.value[1] - 1}, {curr.value[0], curr.value[1] + 1}}
			for _, neighbour := range possneighbours {
				if inbounds(cavemap, neighbour[0], neighbour[1]) && !seenmap[neighbour[0]][neighbour[1]] {
					heap.Push(&pq, &Item{neighbour, curr.priority + cavemap[neighbour[0]][neighbour[1]], hi})
					hi++
				}
			}
		}
	}

	fmt.Println(res)
}

func part2(v []string) {
	// expand cavemap before finding our way through, adjust priority queue to use heuristics
	// note that the heuristics of width-x + height-y guarantees the first path found is optimal,
	// since each movement down or right 1 unit always costs at least 1

	res := 0
	cavemap := make([][]int, len(v))
	seenmap := make([][]bool, len(v)*5)
	for i := 0; i < len(v); i++ {
		line := make([]int, len(v[i]))
		for j := 0; j < len(v[i]); j++ {
			point, _ := strconv.Atoi(string(v[i][j]))
			line[j] = point
		}
		cavemap[i] = line
	}

	for i := 0; i < len(seenmap); i++ {
		seenline := make([]bool, len(v[0])*5)
		for j := 0; j < len(v[0]); j++ {
			seenline[j] = false
		}
		seenmap[i] = seenline
	}

	for i := 0; i < len(v); i++ {
		riskline := cavemap[i]
		for j := 1; j < 5; j++ {
			newline := make([]int, len(v[i]))
			for k := 0; k < len(v[i]); k++ {
				newrisk := riskline[k] + j
				if newrisk > 9 {
					newrisk -= 9
				}
				newline[k] = newrisk
			}
			riskline = append(riskline, newline...)
		}
		cavemap[i] = riskline
	}

	for j := 1; j < 5; j++ {
		for i := 0; i < len(v); i++ {
			riskline := cavemap[i]
			newline := make([]int, len(v[i])*5)
			for k, risk := range riskline {
				newrisk := risk + j
				if newrisk > 9 {
					newrisk -= 9
				}
				newline[k] = newrisk
			}
			cavemap = append(cavemap, newline)
		}
	}

	// make priorityqueue
	pq := make(PriorityQueue, 0)
	// add topleft corner to pq
	heap.Init(&pq)
	hi := 0
	heap.Push(&pq, &Item{[]int{0, 0, 0}, 0, hi})
	hi++

	// while true:
	for {
		// pop item from priority queue, if bottomright corner break
		curr := heap.Pop(&pq).(*Item)
		hi--
		seenmap[curr.value[0]][curr.value[1]] = true
		if curr.value[0] == len(cavemap)-1 && curr.value[1] == len(cavemap[len(cavemap)-1])-1 {
			res = curr.value[2]
			break
		} else {
			// if not add all neighbours
			possneighbours := [][]int{{curr.value[0] - 1, curr.value[1], curr.value[2]}, {curr.value[0] + 1, curr.value[1], curr.value[2]},
				{curr.value[0], curr.value[1] - 1, curr.value[2]}, {curr.value[0], curr.value[1] + 1, curr.value[2]}}
			for _, neighbour := range possneighbours {
				if inbounds(cavemap, neighbour[0], neighbour[1]) && !seenmap[neighbour[0]][neighbour[1]] {
					neighbour[2] += cavemap[neighbour[0]][neighbour[1]]
					heap.Push(&pq, &Item{neighbour, neighbour[2] + len(cavemap) - neighbour[0] + len(cavemap) - neighbour[1], hi})
					hi++
				}
			}
		}
	}

	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 15/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n")
	part1(v)
	part2(v)
}
