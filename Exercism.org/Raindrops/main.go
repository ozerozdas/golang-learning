// https://exercism.org/tracks/go/exercises/raindrops
package main

import "fmt"

func main() {
	for i := 1; i <= 40; i++ {
		fmt.Println("Loop index: ", i, " - ", Convert(i))
	}
}

func Convert(n int) string {
	var output string
	if n%3 == 0 {
		output += "Pling"
	}
	if n%5 == 0 {
		output += "Plang"
	}
	if n%7 == 0 {
		output += "Plong"
	}
	if output == "" {
		output = fmt.Sprintf("%d", n)
	}
	return output
}
