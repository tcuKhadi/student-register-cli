package main

import (
	"fmt"
	"strings"
)

func averageScore() float64 {
	if len(scores) == 0 {
		return 0
	}
	total := 0
	for _, s := range scores {
		total += s
	}
	return float64(total) / float64(len(scores))
}

func topScorer() (string, int) {
	if len(scores) == 0 {
		return "", 0
	}

	bestName := names[0]
	bestScore := scores[0]
	for i, s := range scores {
		if s > bestScore {
			bestScore = s
			bestName = names[i]
		}
	}
	return bestName, bestScore
}

func initials(name string) (string, string) {
	words := strings.Fields(strings.ToUpper(name))
	if len(words) == 0 {
		return "_", "_"
	}

	first := words[0][:1]
	if len(words) == 1 {
		return first, "_"
	}
	return first, words[1][:1]
}

func printMenu() {
	fmt.Println("\n1. Add Student")
	fmt.Println("2. List Students")
	fmt.Println("3. Find Student")
	fmt.Println("4. Update Score")
	fmt.Println("5. Delete Student")
	fmt.Println("6. Average Score")
	fmt.Println("7. Top Scorer")
	fmt.Println("8. Initials")
	fmt.Println("0. Exit")
	fmt.Print("Choose: ")
}
