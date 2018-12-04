package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
)

const idLen = 26

func main() {
	file, err := os.Open("boxes.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	part2(part1(file))
}

func part1(input io.Reader) (ids []string) {
	var doubles, triples int
	for scanner := bufio.NewScanner(input); scanner.Scan(); {
		ids = append(ids, scanner.Text())
		id := scanner.Bytes()
		sort.Slice(id, func(i, j int) bool { return id[i] < id[j] })
		var double, triple int
		for dupCount, i := 0, 1; i < idLen; i++ {
			if id[i] == id[i-1] {
				dupCount++
			}
			if i+1 == idLen || id[i] != id[i+1] {
				if dupCount == 1 {
					double = 1
				} else if dupCount == 2 {
					triple = 1
				}
				dupCount, i = 0, i+1
			}
		}
		doubles, triples = doubles+double, triples+triple
	}
	fmt.Println(doubles * triples)
	return
}

func part2(ids []string) {
	for i := 0; i < len(ids); i++ {
		for j := 0; j < len(ids); j++ {
			if i == j {
				continue
			}
			for k := 0; k < idLen; k++ {
				iString := string(ids[i][:k]) + string(ids[i][k+1:])
				jString := string(ids[j][:k]) + string(ids[j][k+1:])
				if iString == jString {
					fmt.Println(iString)
					return
				}
			}
		}
	}
}
