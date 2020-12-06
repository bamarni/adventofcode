package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	north int = iota
	east
	south
	west

	vertical, horizontal, turnA, turnB, intersection = '|', '-', '/', '\\', '+'
)

var (
	area       [][]byte
	carts      []*Cart
	directions = map[byte]int{
		'^': north,
		'v': south,
		'<': west,
		'>': east,
	}
)

type Cart struct {
	x, y, direction, turn int
	crashed               bool
}

func (c *Cart) Move() {
	if c.crashed {
		return
	}
	switch c.direction {
	case north:
		c.y--
	case south:
		c.y++
	case west:
		c.x--
	case east:
		c.x++
	}
	switch area[c.y][c.x] {
	case intersection:
		c.direction = (c.direction + c.turn) % 4
		c.turn = (c.turn<<1 + 1) % 7
	case turnA:
		switch c.direction {
		case east, west:
			c.direction = (c.direction + 3) % 4
		case north, south:
			c.direction = (c.direction + 1) % 4
		}
	case turnB:
		switch c.direction {
		case west, east:
			c.direction = (c.direction + 1) % 4
		case south, north:
			c.direction = (c.direction + 3) % 4
		}
	}
	for _, cart := range carts {
		if cart != c && cart.x == c.x && cart.y == c.y {
			c.crashed = true
			cart.crashed = true
		}
	}
}

func main() {
	file, _ := os.Open("tracks.txt")
	for y, scanner := 0, bufio.NewScanner(file); scanner.Scan(); y++ {
		area = append(area, make([]byte, len(scanner.Bytes())))
		for x := 0; x < len(scanner.Bytes()); x++ {
			char := scanner.Bytes()[x]
			if direction, ok := directions[char]; ok {
				carts = append(carts, &Cart{x, y, direction, 3, false})
				if direction == north || direction == south {
					char = vertical
				} else {
					char = horizontal
				}
			}
			area[y][x] = char
		}
	}
	for {
		sort.Slice(carts, func(i, j int) bool {
			return carts[i].y < carts[j].y || carts[i].y == carts[j].y && carts[i].x < carts[j].x
		})
		for i := 0; i < len(carts); i++ {
			carts[i].Move()
			if carts[i].crashed {
				// Part 1
				//fmt.Printf("%d,%d\n", carts[i].x, carts[i].y)
				//os.Exit(0)
				carts = append(carts[:i], carts[i+1:]...)
				i--
			}
		}
		if len(carts) == 1 {
			fmt.Printf("%d,%d\n", carts[0].x, carts[0].y)
			os.Exit(0)
		}
	}
}
