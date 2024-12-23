package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type Student struct {
	ID       int
	Name     string
	Subjects map[string]int
	Average  int
}

func (s *Student) CalculateAverage() {
	for _, value := range s.Subjects {
		s.Average += value
	}
}

func (s *Student) Header() (int, string) {
	maxSubjectLength := 0
	for subjects := range s.Subjects {
		if len(subjects) > maxSubjectLength {
			maxSubjectLength = len(subjects)
		}
	}

	return maxSubjectLength, fmt.Sprintf("%-*s %-10s", maxSubjectLength+4, "Subjects", "Scores")
}

func (s *Student) Display() {
	fmt.Println(s.Name)
	n, result := s.Header()
	fmt.Println(result)
	fmt.Println("---------------------------")

	for subject, score := range s.Subjects {
		row := fmt.Sprintf("%-*s : %-10d", n, subject, score)
		fmt.Println(row)
	}
}

func main() {
	students := make([]Student, 0)

	for {
		var nameHolder string
		fmt.Print("Enter Your Name : ")
		for fmt.Scan(&nameHolder); containInteger(nameHolder); {
			fmt.Println("Please Enter Valid Name")
		}
		students = append(students, Student{})
		last_index := len(students) - 1
		students[last_index].Name = nameHolder
		students[last_index].Subjects = make(map[string]int)

		fmt.Println("Enter The SubjectName and Score. eg: Math 98, Enter -1 to finish")
		for {
			var SubjectHolder string
			var ScoreHolder string

			fmt.Scan(&SubjectHolder)
			if SubjectHolder == "-1" {
				break
			}
			fmt.Scan(&ScoreHolder)
			score, err := strconv.Atoi(ScoreHolder)
			if err != nil || SubjectHolder == "" || ScoreHolder == "" || containInteger(SubjectHolder) || score > 100 || score < 0 {
				fmt.Println("Please Enter Valid Value")
				continue
			}

			students[last_index].Subjects[SubjectHolder] = score

		}

		students[last_index].CalculateAverage()

		students[last_index].Display()

	}
}

func containInteger(val string) bool {
	for _, char := range val {
		if unicode.IsDigit(char) {
			return true
		}
	}
	return false
}
