package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	data, err := ioutil.ReadFile("logs.txt")
	if err != nil {
		log.Fatal(err)
	}
	solve(strings.Split(string(data), "\n"))
}

func solve(data []string) {
	sort.Strings(data)
	guards := make(map[int][60]int)
	var sleepsAt time.Time
	var guardID int
	for _, entry := range data {
		parts := strings.Split(entry, " ")
		switch parts[2] {
		case "Guard":
			guardID, _ = strconv.Atoi(parts[3][1:])
			var ok bool
			_, ok = guards[guardID]
			if !ok {
				guards[guardID] = [60]int{}
			}
		case "falls":
			var err error
			sleepsAt, err = time.Parse("[2006-01-02 03:04]", parts[0]+" "+parts[1])
			if err != nil {
				log.Fatal(err)
			}
		case "wakes":
			wakesUpAt, err := time.Parse("[2006-01-02 03:04]", parts[0]+" "+parts[1])
			if err != nil {
				log.Fatal(err)
			}

			minutes := guards[guardID]
			minutesSleeping := int(wakesUpAt.Sub(sleepsAt).Minutes())
			for i := 0; i < minutesSleeping; i++ {
				minutes[(i+sleepsAt.Minute())%60]++
			}
			guards[guardID] = minutes
		}
	}
	var chosenGuard, minutesSleeping int
	for id, guard := range guards {
		var sum int
		for _, sleep := range guard {
			sum += sleep
		}
		if sum > minutesSleeping {
			minutesSleeping = sum
			chosenGuard = id
		}
	}

	var bestMinute, mostTimes int

	for minute, times := range guards[chosenGuard] {
		if times > mostTimes {
			bestMinute = minute
			mostTimes = times
		}
	}
	fmt.Println(bestMinute * chosenGuard)

	for id, guard := range guards {
		for minute, times := range guard {
			if times > mostTimes {
				mostTimes = times
				bestMinute = minute
				chosenGuard = id
			}
		}
	}
	fmt.Println(bestMinute * chosenGuard)
}
