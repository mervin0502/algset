package id3

import (
	"testing"

	"mervin.me/algset/container"
)

func TestNewDecisionTree(t *testing.T) {
	age := []float64{0, 0, 1, 2, 2, 2, 1, 0, 0, 2, 0, 1, 1, 2}
	income := []float64{2, 2, 2, 1, 0, 0, 0, 1, 0, 1, 1, 1, 2, 1}
	student := []float64{0, 0, 0, 0, 1, 1, 1, 0, 1, 1, 1, 0, 1, 0}
	rating := []float64{0, 1, 0, 0, 0, 1, 1, 0, 0, 0, 1, 1, 0, 1}

	label := []string{"0", "0", "1", "1", "1", "0", "1", "0", "1", "1", "1", "1", "1", "0"}

	t.Logf("%v", label)
	var ageT, incomeT, studentT, ratingT *container.Tensor
	ageT = container.NewTensor(age)
	ageT.Label = "age"
	incomeT = container.NewTensor(income)
	incomeT.Label = "income"
	studentT = container.NewTensor(student)
	studentT.Label = "student"
	ratingT = container.NewTensor(rating)
	ratingT.Label = "rating"

	var data *container.Tensor
	data = container.NewTensor([]*container.Tensor{ageT, incomeT, studentT, ratingT})
	data.Label = "data"
	// t.Logf("%s", data.String())

	dt := NewDecisionTree(data, label)
	dt.BFS()
}
