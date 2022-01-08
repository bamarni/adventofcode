package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Node struct {
	parent   *Node
	children []*Node
	metadata []int
}

func (n *Node) value() (sum int) {
	for _, val := range n.metadata {
		if len(n.children) == 0 {
			sum += val
		} else if val-1 < len(n.children) {
			sum += n.children[val-1].value()
		}
	}
	return
}

var (
	current     = &Node{}
	metadataSum int
)

func main() {
	input, err := ioutil.ReadFile("license.txt")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range strings.Split(string(input), " ") {
		number, err := strconv.Atoi(item)
		if err != nil {
			log.Fatal(err)
		}
		switch {
		case current.children == nil:
			current.children = make([]*Node, 0, number)

		case current.metadata == nil:
			current.metadata = make([]int, 0, number)

		case len(current.children) < cap(current.children):
			current.children = append(current.children, &Node{
				parent:   current,
				children: make([]*Node, 0, number),
			})
			current = current.children[len(current.children)-1]

		default:
			if len(current.metadata) < cap(current.metadata) {
				current.metadata = append(current.metadata, number)
				metadataSum += number
			}
			if len(current.metadata) == cap(current.metadata) && current.parent != nil {
				current = current.parent
			}
		}
	}
	fmt.Println(metadataSum)
	fmt.Println(current.value())
}
