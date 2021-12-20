package grades

import (
	"fmt"
	"sync"
)

const (
	GradeTest     = GradeType("Test")
	GradeHomework = GradeType("Homework")
	GradeQuiz     = GradeType("Quiz")
)

type GradeType string
type Students []Student

func (s Students) GetByID(id int) (*Student, error) {
	for i := range s {
		if s[i].ID == id {
			return &s[i], nil
		}
	}
	return nil, fmt.Errorf("Student with ID %v not found", id)
}

var (
	students      Students
	studentsMutex sync.Mutex
)

type Student struct {
	ID        int
	FirstName string
	LastName  string
	Grades    []Grade
}

type Grade struct {
	Title string
	Type  GradeType
	Score float32
}

func (s Student) Average() float32 {
	var result float32
	for _, grade := range s.Grades {
		result += grade.Score

	}

	return result / float32(len(s.Grades))
}
