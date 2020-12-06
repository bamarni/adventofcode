package main

import (
	"fmt"
	"os"
	"strconv"
)

const nbRecipes = 165061

var recipes = []int{3, 7}
var elves = []int{0, 1}
var part2 = []int{1, 6, 5, 0, 6, 1}

func main() {
	for {
		var sum int
		for _, elve := range elves {
			sum += recipes[elve]
		}
		for _, digitStr := range strconv.Itoa(sum) {
			digit, _ := strconv.Atoi(string(digitStr))
			recipes = append(recipes, digit)
		}
		for i, elve := range elves {
			elves[i] = (elve + recipes[elve] + 1) % len(recipes)
		}
		// Part 1
		//if len(recipes) >= nbRecipes+10 {
		//	part1 := recipes[nbRecipes : nbRecipes+10]
		//	fmt.Println(part1)
		//	os.Exit(0)
		//}
		if len(recipes) >= len(part2)+1 {
			check(len(recipes))
			check(len(recipes) - 1)
		}
	}
}

func check(index int) {
	for i, val := range recipes[index-len(part2) : index] {
		if val != part2[i] {
			return
		}
	}
	fmt.Println(index - len(part2))
	os.Exit(0)
}
