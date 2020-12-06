package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/crypto/ssh/terminal"
)

var areaSize int
var points []*Point
var re = regexp.MustCompile(`position=<\s*(\-?\d+), \s*(\-?\d+)> velocity=<\s*(\-?\d+), \s*(\-?\d+)>`)

type Point struct {
	position, velocity [2]int
}

func (p *Point) next() {
	p.position[0] += p.velocity[0]
	p.position[1] += p.velocity[1]
}

func (p *Point) prev() {
	p.position[0] -= p.velocity[0]
	p.position[1] -= p.velocity[1]
}

func main() {
	fmt.Println("Press left or right arrow to move backwards of forwards in time")

	file, _ := os.Open("points.txt")

	for i, scanner := 0, bufio.NewScanner(file); scanner.Scan(); i++ {
		matches := re.FindStringSubmatch(scanner.Text())
		posX, _ := strconv.Atoi(matches[1])
		posY, _ := strconv.Atoi(matches[2])
		velocityX, _ := strconv.Atoi(matches[3])
		velocityY, _ := strconv.Atoi(matches[4])
		points = append(points, &Point{[2]int{posX, posY}, [2]int{velocityX, velocityY}})
	}

	oldState, _ := terminal.MakeRaw(0)
	defer terminal.Restore(0, oldState)

	stdinBuf := make([]byte, 3)
	for sec := 0; ; sec++ {
		var nw, se [2]int
		for i, point := range points {
			if i == 0 || nw[0] > point.position[0] {
				nw[0] = point.position[0]
			}
			if i == 0 || nw[1] > point.position[1] {
				nw[1] = point.position[1]
			}
			if i == 0 || se[0] < point.position[0] {
				se[0] = point.position[0]
			}
			if i == 0 || se[1] < point.position[1] {
				se[1] = point.position[1]
			}
		}

		if sec == 0 || areaSize > distance(nw, se) {
			areaSize = distance(nw, se)
			for _, point := range points {
				point.next()
			}
		} else {
			os.Stdin.Read(stdinBuf)

			if stdinBuf[0] == 27 && stdinBuf[1] == 91 {
				goBack := stdinBuf[2] == 68
				for _, point := range points {
					if goBack {
						point.prev()
					} else {
						point.next()
					}
				}
				if goBack {
					sec = sec - 2
				}
			} else {
				break
			}
			draw(nw, se)
			fmt.Printf("Second %d", sec+1)
			fmt.Println("\r")
		}
	}
}

func distance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func draw(nw, se [2]int) {
	for y := nw[1]; y <= se[1]; y++ {
		for x := nw[0]; x <= se[0]; x++ {
			toPrint := "."
			for _, point := range points {
				if point.position[0] == x && point.position[1] == y {
					toPrint = "#"
				}
			}
			fmt.Print(toPrint)
		}
		fmt.Println("\r")
	}
}
