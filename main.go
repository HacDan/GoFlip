package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Missing arguements: Please enter the number of flips you'd like")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("Too many arguements: Please only supply one number as an arguement")
		os.Exit(1)
	}

	flips, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Wrong arguement type: Please supply whole integers")
		os.Exit(1)
	}

	i := 1

	heads := 0
	tails := 0
	rand.Seed(time.Now().UTC().UnixNano())

	for i <= flips {
		result := flipCoin()
		if result == 0 {
			tails = tails + 1
		} else {
			heads = heads + 1
		}

		i = i + 1
	}

	fmt.Println("Results")
	fmt.Println("Heads: ", heads)
	fmt.Println("Tails: ", tails)
}

func flipCoin() int {
	return rand.Intn(2)
}
