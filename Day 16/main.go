package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// parses one packet, and sums all the contained version numbers, as well as gives the right edge
func sumversionnums(binpac string, left int) (sum int, right int) {
	i := left
	j := left
	sum = 0

	j = i + 3
	version, _ := strconv.ParseInt(binpac[i:j], 2, 8)
	i = j
	sum += int(version)
	j = i + 3
	pactype, _ := strconv.ParseInt(binpac[i:j], 2, 8)
	i = j
	if pactype == 4 { // literal value
		nextgrouptype := binpac[i]
		i += 1
		litval := make([]uint8, 0, 50)
		for string(nextgrouptype) == "1" {
			j = i + 4
			nextgroup := binpac[i:j]
			i = j
			litval = append(litval, nextgroup...)
			nextgrouptype = binpac[i]
			i += 1
		}
		j = i + 4
		// last group
		nextgroup := binpac[i:j]
		i = j
		litval = append(litval, nextgroup...)
		return sum, i
	} else { // operators
		lengthtypeid := binpac[i]
		i += 1
		if string(lengthtypeid) == "0" {
			j = i + 15
			length, _ := strconv.ParseInt(binpac[i:j], 2, 16)
			i = j
			j = i + int(length)
			for i < j {
				subsum, subright := sumversionnums(binpac, i)
				sum += subsum
				i = subright
			}
			return sum, i
		} else {
			j = i + 11
			containedsubs, _ := strconv.ParseInt(binpac[i:j], 2, 16)
			i = j
			for k := 0; k < int(containedsubs); k++ {
				subsum, subright := sumversionnums(binpac, i)
				sum += subsum
				i = subright
			}
			return sum, i
		}
	}
}

func opdispatch(opid int64, subvals []int64) int64 {
	if opid == 0 {
		sum := int64(0)
		for _, subval := range subvals {
			sum += subval
		}
		return sum
	} else if opid == 1 {
		product := int64(1)
		for _, subval := range subvals {
			product *= subval
		}
		return product
	} else if opid == 2 {
		min := int64(99999)
		for _, subval := range subvals {
			if subval < min {
				min = subval
			}
		}
		return min
	} else if opid == 3 {
		max := int64(0)
		for _, subval := range subvals {
			if subval > max {
				max = subval
			}
		}
		return max
	} else if opid == 5 {
		if subvals[0] > subvals[1] {
			return 1
		} else {
			return 0
		}
	} else if opid == 6 {
		if subvals[0] < subvals[1] {
			return 1
		} else {
			return 0
		}
	} else if opid == 7 {
		if subvals[0] == subvals[1] {
			return 1
		} else {
			return 0
		}
	} else {
		return -99999 // just for gofmt
	}
}
func interppacs(binpac string, left int) (value int64, right int) {
	i := left
	j := left
	value = 0

	j = i + 3
	//version, _ := strconv.ParseInt(binpac[i:j], 2, 8)
	i = j
	//sum += int(version)
	j = i + 3
	pactype, _ := strconv.ParseInt(binpac[i:j], 2, 8)
	i = j
	if pactype == 4 { // literal value
		nextgrouptype := binpac[i]
		i += 1
		litval := make([]uint8, 0, 50)
		for string(nextgrouptype) == "1" {
			j = i + 4
			nextgroup := binpac[i:j]
			i = j
			litval = append(litval, nextgroup...)
			nextgrouptype = binpac[i]
			i += 1
		}
		j = i + 4
		// last group
		nextgroup := binpac[i:j]
		i = j
		litval = append(litval, nextgroup...)
		value, _ = strconv.ParseInt(string(litval), 2, 64)
		return value, i
	} else { // operators
		lengthtypeid := binpac[i]
		i += 1
		subvalues := make([]int64, 0, 10)
		if string(lengthtypeid) == "0" {
			j = i + 15
			length, _ := strconv.ParseInt(binpac[i:j], 2, 16)
			i = j
			j = i + int(length)
			for i < j {
				subsum, subright := interppacs(binpac, i)
				subvalues = append(subvalues, subsum)
				i = subright
			}
			value = opdispatch(pactype, subvalues)
			return value, i
		} else {
			j = i + 11
			containedsubs, _ := strconv.ParseInt(binpac[i:j], 2, 16)
			i = j
			for k := 0; k < int(containedsubs); k++ {
				subsum, subright := interppacs(binpac, i)
				subvalues = append(subvalues, subsum)
				i = subright
			}
			value = opdispatch(pactype, subvalues)
			return value, i
		}
	}
}

func part1(v string) {

	binpac := ""
	for i := 0; i < len(v); i++ {
		hexint, _ := strconv.ParseInt(string(v[i]), 16, 8)
		binpac = binpac + fmt.Sprintf("%04b", hexint)
	}
	res := 0

	res, end := sumversionnums(binpac, 0)
	if end != len(binpac) {
		//fmt.Println("oops")
		fmt.Println("oh yeah, padding")
	}

	fmt.Println(res)
}

func part2(v string) {

	binpac := ""
	for i := 0; i < len(v); i++ {
		hexint, _ := strconv.ParseInt(string(v[i]), 16, 8)
		binpac = binpac + fmt.Sprintf("%04b", hexint)
	}

	res, end := interppacs(binpac, 0)
	if end != len(binpac)-1 {
		//fmt.Println("oops")
		fmt.Println("oh yeah, padding")
	}

	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 16/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	part1(strcontent)
	part2(strcontent)
}
