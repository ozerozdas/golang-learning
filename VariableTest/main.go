package main

import "fmt"

func main() {
	// Run the function you want
	TwoFer() // https://exercism.org/tracks/go/exercises/two-fer
}

func TwoFer() string {
	var list = []string{"Alice", "Bob", "", "Ozer"}
	for index, value := range list {
		if value == "" {
			list[index] = "you"
		}
		fmt.Printf("One for %v, one for me.\n", list[index])
	}
	return fmt.Sprintf("One for %s, one for me.", list[0])
}
