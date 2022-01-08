package main

import (
	"container/ring"
	"fmt"
)

const nbPlayers = 419
const lastMarble = 71052

var players [nbPlayers]int

func main() {
	circle := ring.New(1)
	circle.Value = 0

	var highestScore int
	for i := 1; i <= lastMarble; i++ {
		marble := ring.New(1)
		marble.Value = i

		if i%23 == 0 {
			circle = circle.Move(-8)
			marbleScore := circle.Unlink(1).Value.(int) + i
			players[i%nbPlayers] += marbleScore

			if players[i%nbPlayers] > highestScore {
				highestScore = players[i%nbPlayers]
			}
		} else {
			circle = circle.Next()
			circle.Link(marble)
		}
		circle = circle.Next()
	}
	fmt.Println(highestScore)
}
