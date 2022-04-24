// https://exercism.org/tracks/go/exercises/hamming
package main

import "fmt"

func main() {
	distance := Distance("GAGCCTACTAACGGGAT", "CATCGTAATGACGGCCT")
	if distance == -1 {
		fmt.Println("The strings must be the same length.")
	} else {
		fmt.Println("Hamming Distance: ", distance)
	}
}

func Distance(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}
	return count
}
