// https://exercism.org/tracks/go/exercises/scrabble-score
package main

import "fmt"

func main() {
	fmt.Println("Score:", calculateScore("cabbage"))
}

func calculateScore(word string) int {
	var score int
	for _, letter := range word {
		score += getScore(rune(letter))
	}
	return score
}

func getScore(letter rune) int {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
		return 1
	case 'd', 'g':
		return 2
	case 'b', 'c', 'm', 'p':
		return 3
	case 'f', 'h', 'v', 'w', 'y':
		return 4
	case 'k':
		return 5
	case 'j', 'x':
		return 8
	case 'q', 'z':
		return 10
	default:
		return 0
	}
}
