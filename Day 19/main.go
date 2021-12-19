package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func generateOrientations(beacon [3]int) [24][3]int {
	orientations := [24][3]int{}

	// screw it I'm not making this fancy with matrix algebra
	// orient around +x and rotate
	orientations[0] = [3]int{beacon[0], beacon[1], beacon[2]}
	orientations[1] = [3]int{beacon[0], beacon[2], -beacon[1]}
	orientations[2] = [3]int{beacon[0], -beacon[1], -beacon[2]}
	orientations[3] = [3]int{beacon[0], -beacon[2], beacon[1]}
	// orient around -x and rotate
	orientations[4] = [3]int{-beacon[0], -beacon[1], beacon[2]}
	orientations[5] = [3]int{-beacon[0], beacon[2], beacon[1]}
	orientations[6] = [3]int{-beacon[0], beacon[1], -beacon[2]}
	orientations[7] = [3]int{-beacon[0], -beacon[2], -beacon[1]}

	// orient around +y and rotate
	orientations[8+0] = [3]int{orientations[0][2], orientations[0][0], orientations[0][1]}
	orientations[8+1] = [3]int{orientations[1][2], orientations[1][0], orientations[1][1]}
	orientations[8+2] = [3]int{orientations[2][2], orientations[2][0], orientations[2][1]}
	orientations[8+3] = [3]int{orientations[3][2], orientations[3][0], orientations[3][1]}
	// orient around -y and rotate
	orientations[8+4] = [3]int{orientations[4][2], orientations[4][0], orientations[4][1]}
	orientations[8+5] = [3]int{orientations[5][2], orientations[5][0], orientations[5][1]}
	orientations[8+6] = [3]int{orientations[6][2], orientations[6][0], orientations[6][1]}
	orientations[8+7] = [3]int{orientations[7][2], orientations[7][0], orientations[7][1]}

	// orient around +z and rotate
	orientations[16+0] = [3]int{orientations[0][1], orientations[0][2], orientations[0][0]}
	orientations[16+1] = [3]int{orientations[1][1], orientations[1][2], orientations[1][0]}
	orientations[16+2] = [3]int{orientations[2][1], orientations[2][2], orientations[2][0]}
	orientations[16+3] = [3]int{orientations[3][1], orientations[3][2], orientations[3][0]}
	// orient around -z and rotate
	orientations[16+4] = [3]int{orientations[4][1], orientations[4][2], orientations[4][0]}
	orientations[16+5] = [3]int{orientations[5][1], orientations[5][2], orientations[5][0]}
	orientations[16+6] = [3]int{orientations[6][1], orientations[6][2], orientations[6][0]}
	orientations[16+7] = [3]int{orientations[7][1], orientations[7][2], orientations[7][0]}

	return orientations
}

func generateInverseRotations(beacon [3]int) [24][3]int {

	orientations := [24][3]int{}

	// screw it I'm not making this fancy with matrix algebra
	// orient around +x and rotate
	orientations[0] = [3]int{beacon[0], beacon[1], beacon[2]}
	orientations[1] = [3]int{beacon[0], -beacon[2], beacon[1]}
	orientations[2] = [3]int{beacon[0], -beacon[1], -beacon[2]}
	orientations[3] = [3]int{beacon[0], beacon[2], -beacon[1]}
	// orient around -x and rotate
	orientations[4] = [3]int{-beacon[0], -beacon[1], beacon[2]}
	orientations[5] = [3]int{-beacon[0], beacon[2], beacon[1]}
	orientations[6] = [3]int{-beacon[0], beacon[1], -beacon[2]}
	orientations[7] = [3]int{-beacon[0], -beacon[2], -beacon[1]}

	// todo: figure out inverses for 8-23
	// orient around +y and rotate
	orientations[8+0] = [3]int{orientations[0][2], orientations[0][0], orientations[0][1]}
	orientations[8+1] = [3]int{orientations[1][2], orientations[1][0], orientations[1][1]}
	orientations[8+2] = [3]int{orientations[2][2], orientations[2][0], orientations[2][1]}
	orientations[8+3] = [3]int{orientations[3][2], orientations[3][0], orientations[3][1]}
	// orient around -y and rotate
	orientations[8+4] = [3]int{orientations[4][2], orientations[4][0], orientations[4][1]}
	orientations[8+5] = [3]int{orientations[5][2], orientations[5][0], orientations[5][1]}
	orientations[8+6] = [3]int{orientations[6][2], orientations[6][0], orientations[6][1]}
	orientations[8+7] = [3]int{orientations[7][2], orientations[7][0], orientations[7][1]}

	// orient around +z and rotate
	orientations[16+0] = [3]int{orientations[0][1], orientations[0][2], orientations[0][0]}
	orientations[16+1] = [3]int{orientations[1][1], orientations[1][2], orientations[1][0]}
	orientations[16+2] = [3]int{orientations[2][1], orientations[2][2], orientations[2][0]}
	orientations[16+3] = [3]int{orientations[3][1], orientations[3][2], orientations[3][0]}
	// orient around -z and rotate
	orientations[16+4] = [3]int{orientations[4][1], orientations[4][2], orientations[4][0]}
	orientations[16+5] = [3]int{orientations[5][1], orientations[5][2], orientations[5][0]}
	orientations[16+6] = [3]int{orientations[6][1], orientations[6][2], orientations[6][0]}
	orientations[16+7] = [3]int{orientations[7][1], orientations[7][2], orientations[7][0]}

	return orientations
}

func part1(v []string) {

	res := 0
	scanners := make([][][3]int, 0, len(v)) // indices are scanner, beacon, coordinates
	//scanners := make([]map[[3]int]bool, 0, len(v)) // indices are scanner and coordinates
	for i := 0; i < len(v); i++ {
		// scan scanner lists
		beacons := strings.Split(v[i], "\n")[1:]
		detectedBeacons := make([][3]int, 0, len(beacons))
		//detectedBeacons := make(map[[3]int]bool)
		for _, beacon := range beacons {
			coords := strings.Split(beacon, ",")
			intcoords := [3]int{}
			intcoords[0], _ = strconv.Atoi(coords[0])
			intcoords[1], _ = strconv.Atoi(coords[1])
			intcoords[2], _ = strconv.Atoi(coords[2])
			detectedBeacons = append(detectedBeacons, intcoords)
			//detectedBeacons[intcoords] = true
		}
		scanners = append(scanners, detectedBeacons)
	}

	// key: scanner, values: neighbour they have 12 overlaps with, and their
	//   relative orientation compared to them (out of 24)
	overlapMap := make(map[int][2]int)
	// key: scanner, values: delta after rotation compared to them
	deltaMap := make(map[int][3]int)

	// connectionsMade := 0                 // should be len(scanners)-1 when done
	// connect each unconnected scanner to some other one, and only one, and record their relative orientation and position
OuterFindOverlap:
	for i := 1; i < len(scanners); i++ {
		scanner := scanners[i]

		allPerms := make([][24][3]int, 0, len(scanner))
		// generate orientations
		for _, beacon := range scanner {
			allPerms = append(allPerms, generateOrientations(beacon))
		}

		// for each pair, check the 24 different possible orientations to match
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}
			relativeScanner := scanners[j]
			for rottype := 0; rottype < 24; rottype++ {
				possibleDeltas := make(map[[3]int]int)
				for k := range allPerms {
					beacon := allPerms[k][rottype]
					for _, relativeBeacon := range relativeScanner {
						delta := [3]int{
							beacon[0] - relativeBeacon[0],
							beacon[1] - relativeBeacon[1],
							beacon[2] - relativeBeacon[2]}

						_, exists := possibleDeltas[delta]
						if !exists {
							possibleDeltas[delta] = 0
						}
						possibleDeltas[delta]++
					}
				}

				for delta, beaconsMatched := range possibleDeltas {
					if beaconsMatched >= 12 {
						overlapMap[i] = [2]int{j, rottype}
						deltaMap[i] = delta
						continue OuterFindOverlap
					}
				}
			}

		}
	}

	// make delta and orientation maps of all scanners relative to scanner 0
	zeroedRotMap := make(map[int]int)
	zeroedDeltMap := make(map[int][3]int)
	zeroedRotMap[0] = 0
	zeroedDeltMap[0] = [3]int{0, 0, 0}
	testOrientations := generateOrientations([3]int{1, 2, 3})
	// you know what I'm also making this easier on myself
	for i := 0; i < len(scanners); i++ {
		for scanner, neighbour := range overlapMap {
			_, alreadyMapped := zeroedRotMap[scanner]
			if alreadyMapped {
				continue
			}
			zeroedNeighbourRot, exists := zeroedRotMap[neighbour[0]]
			if exists {
				// apply zeroedNeighbourRotation, then own rotation
				neighbourTestOrientation := testOrientations[zeroedNeighbourRot]
				relativeSelfRotation := neighbour[1]
				selfTestOrientation := generateOrientations(neighbourTestOrientation)[relativeSelfRotation]
				zeroedSelfRotation := -1
				for j, testOrientation := range testOrientations {
					if testOrientation == selfTestOrientation {
						zeroedSelfRotation = j
						break
					}
				}
				zeroedRotMap[scanner] = zeroedSelfRotation

				// brute force way is to actually do both, and compare to which results from just doing one

				// then save that to zeroedRotMap, and zeroedDeltMap
			}
		}
	}
	// add all beacons to map

	// count map

	fmt.Println(res)
}

func part2(v []string) {

	res := 0
	for i := 0; i < len(v); i++ {

	}
	fmt.Println(res)
}

func main() {

	content, err := ioutil.ReadFile("Day 19/input.txt")

	if err != nil {
		log.Fatal(err)
	}
	strcontent := string(content)
	v := strings.Split(strcontent, "\n\n")
	part1(v)
	part2(v)
}
