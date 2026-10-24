package main

import "testing"

func reset() {
	names = []string{}
	scores = []int{}
}

func TestAddStudent(t *testing.T) {
	reset()
	addStudent("Mojalefa", 80)
	addStudent("Thabo", 65)

	if len(names) != 2 || len(scores) != 2 {
		t.Fatalf("expected 2 students, got %d names and %d scores", len(names), len(scores))
	}
	if names[1] != "Thabo" || scores[1] != 65 {
		t.Errorf("index 1: got %v %v, want Thabo 65", names[1], scores[1])
	}
}

func TestFindStudent(t *testing.T) {
	reset()
	addStudent("Mojalefa", 80)
	addStudent("Thabo", 65)

	if got := findStudent("thabo"); got != 1 {
		t.Errorf("case-insensitive find: got %d, want 1", got)
	}
	if got := findStudent("nobody"); got != -1 {
		t.Errorf("missing student: got %d, want -1", got)
	}
}

func TestUpdateScore(t *testing.T) {
	reset()
	addStudent("Mojalefa", 80)

	updateScore(0, 95)
	if scores[0] != 95 {
		t.Errorf("valid update: got %d, want 95", scores[0])
	}

	updateScore(5, 10)  // out of range, must not panic
	updateScore(-1, 10) // negative, must not panic
	if len(scores) != 1 || scores[0] != 95 {
		t.Errorf("invalid update changed data: %v", scores)
	}
}

func TestDeleteStudent(t *testing.T) {
	reset()
	addStudent("A", 1)
	addStudent("B", 2)
	addStudent("C", 3)

	deleteStudent(1) // middle
	if len(names) != 2 || names[0] != "A" || names[1] != "C" || scores[1] != 3 {
		t.Errorf("after middle delete: names=%v scores=%v", names, scores)
	}

	deleteStudent(1) // last
	if len(names) != 1 || names[0] != "A" || scores[0] != 1 {
		t.Errorf("after last delete: names=%v scores=%v", names, scores)
	}

	deleteStudent(0) // only remaining
	if len(names) != 0 || len(scores) != 0 {
		t.Errorf("after only delete: names=%v scores=%v", names, scores)
	}

	deleteStudent(0) // empty list, must not panic
}

func TestAverageScore(t *testing.T) {
	reset()
	if got := averageScore(); got != 0 {
		t.Errorf("empty list: got %v, want 0", got)
	}

	addStudent("A", 80)
	addStudent("B", 65)
	if got := averageScore(); got != 72.5 {
		t.Errorf("two students: got %v, want 72.5", got)
	}
}

func TestTopScorer(t *testing.T) {
	reset()
	if n, s := topScorer(); n != "" || s != 0 {
		t.Errorf("empty list: got %q %d, want \"\" 0", n, s)
	}

	addStudent("A", 50)
	addStudent("B", 90)
	addStudent("C", 90)
	if n, s := topScorer(); n != "B" || s != 90 {
		t.Errorf("tie: got %q %d, want B 90 (first wins)", n, s)
	}
}

func TestInitials(t *testing.T) {
	cases := []struct {
		in, first, second string
	}{
		{"tifa lockhart", "T", "L"},
		{"Giga", "G", "_"},
		{"", "_", "_"},
		{"   ", "_", "_"},
		{"  spaced   out ", "S", "O"},
	}

	for _, c := range cases {
		a, b := initials(c.in)
		if a != c.first || b != c.second {
			t.Errorf("initials(%q): got %v %v, want %v %v", c.in, a, b, c.first, c.second)
		}
	}
}