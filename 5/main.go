package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var shortest int

	for j := 96; j <= 122; j++ {
		input, err := ioutil.ReadFile("polymers.txt")
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(input)-1; i++ {
			if input[i] == byte(j) || input[i]+32 == byte(j) {
				input = append(input[:i], input[i+1:]...)
				if i > 0 {
					i--
				}
			}
			if input[i] != input[i+1] {
				low := bytes.ToLower(input[i : i+2])
				if low[0] == low[1] {
					input = append(input[:i], input[i+2:]...)
					if i > 0 {
						i--
					}
					i--
				}
			}
		}
		if j == 96 {
			fmt.Println(len(input))
			shortest = len(input)
		} else if len(input) < shortest {
			shortest = len(input)
		}
	}
	fmt.Println(shortest)
}
