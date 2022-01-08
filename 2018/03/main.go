package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

const fabricLen = 1000

var re = regexp.MustCompile("#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)")

func main() {
	file, err := os.Open("claims.txt")
	if err != nil {
		log.Fatal(err)
	}

	fabric := make([][][]int, fabricLen)
	for i := 0; i < fabricLen; i++ {
		fabric[i] = make([][]int, fabricLen)
	}

	ids := make(map[int]struct{})

	var overlap int
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		matches := re.FindStringSubmatch(scanner.Text())

		id, _ := strconv.Atoi(matches[1])
		ids[id] = struct{}{}

		left, _ := strconv.Atoi(matches[2])
		top, _ := strconv.Atoi(matches[3])
		width, _ := strconv.Atoi(matches[4])
		height, _ := strconv.Atoi(matches[5])

		for i := left; i < left+width; i++ {
			for j := top; j < top+height; j++ {
				fabric[i][j] = append(fabric[i][j], id)
				if len(fabric[i][j]) == 2 {
					overlap++
				}
			}
		}
	}
	fmt.Println(overlap)

	for _, i := range fabric {
		for _, j := range i {
			if len(j) > 1 {
				for _, overlappingID := range j {
					delete(ids, overlappingID)
				}
			}
		}
	}
	for id := range ids {
		fmt.Println(id)
	}
}
