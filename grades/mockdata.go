package grades

func init() {
	students = []Student{
		Student{
			ID:        1,
			FirstName: "Averill",
			LastName:  "Simen",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 01 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 02 Homework",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        2,
			FirstName: "Mark",
			LastName:  "Simen",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 01 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 02 Homework",
					Type:  GradeQuiz,
					Score: 99,
				},
			},
		},
		Student{
			ID:        3,
			FirstName: "Lucas",
			LastName:  "Abel",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 01 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 02 Homework",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        4,
			FirstName: "Osvaldo",
			LastName:  "Abel",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 01 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 02 Homework",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		Student{
			ID:        5,
			FirstName: "Kanguanda",
			LastName:  "Abel",
			Grades: []Grade{
				Grade{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				Grade{
					Title: "Week 01 Homework",
					Type:  GradeHomework,
					Score: 94,
				},
				Grade{
					Title: "Quiz 02 Homework",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
	}
}
