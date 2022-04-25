// https://courses.calhoun.io/courses/cor_gophercises
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	_ = csvFilename

	var doYouWantStart string
	fmt.Println("You have 30 seconds to answer each question. Do you want to start the quiz? (y/n)")
	fmt.Scanf("%s\n", &doYouWantStart)

	if doYouWantStart == "y" {
		fmt.Println("Starting the quiz...")
	} else {
		exit(fmt.Sprintf("Quiz will not start"))
	}

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemLoop:
	for _, p := range problems {
		fmt.Print(p.q, " = ")
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nYou're time is up!")
			break problemLoop
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("You got it right!")
				correct++
			} else {
				fmt.Printf("You got it wrong. The answer is %s.\n", p.a)
			}
		}

	}
	fmt.Println("You scored", correct, "out of", len(problems), "correct and you got", correct*100/len(problems), "point.")
	if correct >= len(problems)/2 {
		fmt.Println("You passed the quiz! Congratulations!")
	} else {
		fmt.Println("You failed the quiz!")
	}
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines)) // make a slice of problems
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
