package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dice = []int{1, 2, 3, 4, 5, 6} // list of possible dice values

func rollDice() int {
	return dice[rand.Intn(len(dice))] // random index between 0 and 5
}

func main() {
	rand.Seed(time.Now().UnixNano()) // seed the random number generator
	fmt.Println("Dice 1: ", rollDice())
	fmt.Println("Dice 2: ", rollDice())
}