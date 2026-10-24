package main

import (
	"fmt"
	"strings"
)

var names []string 
var scores []int

func addStudent(name string, score int) {
	names = append(names, name)
	scores = append(scores, score)
}

func listStudents() {
	if len(names) == 0 {
		fmt.Println("No students yet.")
		return
	}
	for i, name := range names {
		fmt.Printf("%d. %v - %v\n", i, name, scores[i])
	}
}

func findStudent(name string) int {
	target := strings.ToUpper(name) 
	for i, name := range names {
		if strings.ToUpper(name) == target {
			return i
		}
	}
	return -1
}

func updateScore(index int, score int) {
	if index < 0 || index >= len(scores) {
		fmt.Println("Invalid Index.")
		return
	}
	scores[index] = score
}

func deleteStudent(index int) {
	if index < 0 || index >= len(names) {
		fmt.Println("Invalid index.")
		return
	}
	names = append(names[:index], names[index+1:]...)
	scores = append(scores[:index], scores[index+1:]...)
}