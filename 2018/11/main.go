package main

import (
	"fmt"
)

const serialNumber = 9435

func main() {
	area := make([][]int, 300)
	for y := range area {
		area[y] = make([]int, 300)
	}

	for y := range area {
		for x := range area[y] {
			rackID := x + 1 + 10
			powerLevel := (rackID*(y+1) + serialNumber) * rackID
			area[y][x] = powerLevel/100%10 - 5
		}
	}

	var highest, highestLen, highestX, highestY int

	for squareLen := 3; squareLen < len(area); squareLen++ {
		for y := 0; y+squareLen <= len(area); y++ {
			for x := 0; x+squareLen <= len(area); x++ {
				var powerLevel int
				for i := 0; i < squareLen; i++ {
					for j := 0; j < squareLen; j++ {
						powerLevel += area[y+i][x+j]
					}
				}
				if powerLevel > highest {
					highest = powerLevel
					highestX = x + 1
					highestY = y + 1
					highestLen = squareLen
				}
			}
		}
		if squareLen == 3 {
			fmt.Printf("%d,%d\n", highestX, highestY)
		}
	}
	fmt.Printf("%d,%d,%d\n", highestX, highestY, highestLen)
}
