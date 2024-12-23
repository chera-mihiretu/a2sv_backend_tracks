package main

import (
	"testing"
)

func AveraegeCalculate(t *testing.T) {

	ss := map[string]int{
		"Math":      90,
		"English":   85,
		"Science":   80,
		"History":   75,
		"Geography": 92,
	}

	expected := 84.4
	var student_test Student

	student_test.Subjects = ss
	student_test.CalculateAverage()

	actual := student_test.Average

	if actual == int(expected) {
		t.Errorf("Average = %v; Expected %v", actual, expected)
	}
}
