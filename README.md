# student-register-cli

A command-line student register written in Go. Add, find, update, and delete students, calculate the class average, and show the top scorer.

Built as a practice project for Go basics: slices, loops, conditionals, functions, multiple return values, and multi-file packages.

## Features

- Add a student with a name and score
- List all students
- Find a student by name (case-insensitive)
- Update a student's score
- Delete a student
- Show the average score
- Show the top scorer
- Show a student's initials

## Requirements

- Go installed (check with `go version`)

## Usage

```
go run .
```

The program displays a numbered menu. Input is read one word at a time with `fmt.Scan`, so names cannot contain spaces.

## Testing

```
go test -v
```

## Project Structure

```
main.go         menu loop and input prompts
students.go     storage and student operations
utils.go        average, top scorer, initials, menu display
student_test.go unit tests
```

## Limitations

- Data is kept in memory and lost on exit.
- Names cannot contain spaces.
- If two students share a name, `find` returns the first match.