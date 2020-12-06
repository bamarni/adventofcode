package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("coordinates.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")

	var coordinates [][2]int
	var maxX, maxY int

	for _, line := range lines {
		point := strings.Split(line, ", ")
		x, err := strconv.Atoi(point[0])
		if err != nil {
			log.Fatal(err)
		}
		if x > maxX {
			maxX = x
		}
		y, err := strconv.Atoi(point[1])
		if err != nil {
			log.Fatal(err)
		}
		if y > maxY {
			maxY = y
		}
		coordinates = append(coordinates, [2]int{x, y})
	}

	area := make([][][]int, maxY+1)
	for y := range area {
		area[y] = make([][]int, maxX+1)
	}

	var size10000 int
	sizes := make(map[int]int)
	infiniteLocations := make(map[int]struct{})

	for y, row := range area {
		for x := range row {
			var sum, closestDist float64
			for i, coord := range coordinates {
				candidateDist := distance(x, y, coord[0], coord[1])
				sum += candidateDist
				if closestDist == 0 || closestDist > candidateDist {
					area[y][x] = []int{i}
					closestDist = candidateDist
				} else if closestDist == candidateDist {
					area[y][x] = append(area[y][x], i)
				}
			}
			if sum < 10000 {
				size10000++
			}
			if len(area[y][x]) == 1 {
				sizes[area[y][x][0]]++
				if x*y == 0 || x == maxX || y == maxY {
					infiniteLocations[area[y][x][0]] = struct{}{}
				}
			}
		}
	}

	var maxSize int
	for loc, size := range sizes {
		if _, ok := infiniteLocations[loc]; !ok && size > maxSize {
			maxSize = size
		}
	}
	fmt.Println(maxSize)
	fmt.Println(size10000)
}

func distance(x1, y1, x2, y2 int) float64 {
	return math.Abs(float64(x1-x2)) + math.Abs(float64(y1-y2))
}
