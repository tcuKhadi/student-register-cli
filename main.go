package main

import "fmt"

func main() {
	for {
		printMenu()

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			addStudentPrompt()

		} else if choice == 2 {
			listStudents()

		} else if choice == 3 {
			findStudentPrompt()

		} else if choice == 4 {
			updateScorePrompt()

		} else if choice == 5 {
			deleteStudentPrompt()

		} else if choice == 6 {
			fmt.Println("Average:", averageScore())

		} else if choice == 7 {
			name, score := topScorer()
			fmt.Println("Top:", name, score)

		} else if choice == 8 {
			initialsPrompt()

		} else if choice == 0 {
			fmt.Println("Bye.")
			break

		} else {
			fmt.Println("Invalid option.")
		}
	}
}

func addStudentPrompt() {
	var name string
	var score int

	fmt.Print("Name: ")
	fmt.Scan(&name)
	fmt.Print("Score: ")
	fmt.Scan(&score)
	addStudent(name, score)
}

func findStudentPrompt() {
	var name string
	fmt.Print("Name: ")
	fmt.Scan(&name)

	index := findStudent(name)
	if index == -1 {
		fmt.Println("Not found.")
	} else {
		fmt.Println("Found at index", index)
	}
}

func updateScorePrompt() {
	var name string
	var score int

	fmt.Print("Name: ")
	fmt.Scan(&name)

	index := findStudent(name)
	if index == -1 {
		fmt.Println("Not found.")
		return
	}
	fmt.Print("New score: ")
	fmt.Scan(&score)
	updateScore(index, score)
}

func deleteStudentPrompt() {
	var name string
	fmt.Print("Name: ")
	fmt.Scan(&name)

	index := findStudent(name)
	if index == -1 {
		fmt.Println("Not found.")
		return
	}
	deleteStudent(index)
}

func initialsPrompt() {
	var name string
	fmt.Print("Name: ")
	fmt.Scan(&name)

	first, second := initials(name)
	fmt.Println("Initials:", first, second)
}