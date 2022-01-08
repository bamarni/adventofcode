package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const target = 2020

func main() {
	file, err := os.Open("expenses.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var expenses []int
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		expense := scanner.Text()

		amount, err := strconv.Atoi(expense)
		if err != nil {
			log.Fatal(err)
		}
		expenses = append(expenses, amount)
	}

	fmt.Println(part1(expenses))
	fmt.Println(part2(expenses))
}

func part1(expenses []int) int {
	expenseMap := make(map[int]bool)

	for _, amount := range expenses {
		if expenseMap[target-amount] {
			return amount * (target - amount)
		}
		expenseMap[amount] = true
	}
	return 0
}

func part2(expenses []int) int {
	for i, amount := range expenses {
		subTarget := target - amount
		expenseMap := make(map[int]bool)

		for j := i + 1; j < len(expenses); j++ {
			if expenseMap[subTarget-expenses[j]] {
				return amount * (subTarget - expenses[j]) * expenses[j]
			}
			expenseMap[expenses[j]] = true
		}
	}
	return 0
}
