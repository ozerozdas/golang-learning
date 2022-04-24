// https://exercism.org/tracks/go/exercises/two-fer
package main

import "fmt"

func main() {
	var list = []string{"Alice", "Bob", "", "Ozer"}
	for index, value := range list {
		if value == "" {
			list[index] = "you"
		}
		fmt.Printf("One for %v, one for me.\n", list[index])
	}
}
