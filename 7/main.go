package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	nodes       []byte
	edges       = make(map[byte][]byte)
	workPool    = make(map[byte]int)
	elapsedTime int
)

func main() {
	input, err := ioutil.ReadFile("requirements.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range bytes.Split(input, []byte("\n")) {
		from, to := line[36], line[5]
		addNode(from)
		addNode(to)
		edges[from] = append(edges[from], to)
	}

	/* Part 1
	for i := 0; i < len(nodes); i++ {
		node := nodes[i]

		if len(edges[node]) == 0 {
			fmt.Printf("%c", node)

			removeNode(i)
			i = -1
		}
	} */

	for len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			if len(edges[node]) == 0 {
				if _, alreadyInProgress := workPool[node]; alreadyInProgress {
					continue
				}
				if len(workPool) == 5 {
					break
				}
				workPool[node] += int(node - 4)
				i = -1
			}
		}

		for node := range workPool {
			workPool[node]--
			if workPool[node] == 0 {
				delete(workPool, node)
				for i := 0; i < len(nodes); i++ {
					if nodes[i] == node {
						removeNode(i)
						break
					}
				}
			}
		}
		elapsedTime++
	}
	fmt.Println(elapsedTime)
}

func addNode(node byte) {
	for i := range nodes {
		if nodes[i] == node {
			return
		}
		if nodes[i] > node {
			nodes = append(nodes, 0)
			copy(nodes[i+1:], nodes[i:])
			nodes[i] = node
			return
		}
	}
	nodes = append(nodes, node)
}

func removeNode(id int) {
	for i := range edges {
		for j := range edges[i] {
			if edges[i][j] == nodes[id] {
				edges[i] = append(edges[i][:j], edges[i][j+1:]...)
				break
			}
		}
	}
	nodes = append(nodes[:id], nodes[id+1:]...)
}
