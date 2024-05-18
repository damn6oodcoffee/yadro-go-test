package sortable

import (
	"math/rand"
	"testing"
)

func TestExampleOne(t *testing.T) {
	matrix := [][]int64{
		{1, 2},
		{2, 1},
	}
	res := IsSortable(matrix)
	if res != true {
		t.Errorf("got '%v' from IsSortable(), expected 'true'", res)
	}
}

func TestExampleTwo(t *testing.T) {
	matrix := [][]int64{
		{10, 20, 30},
		{1, 1, 1},
		{0, 0, 1},
	}
	res := IsSortable(matrix)
	if res != false {
		t.Errorf("got '%v' from IsSortable(), expected 'false'", res)
	}
}

func TestTrivial(t *testing.T) {
	matrix := [][]int64{
		{9},
	}
	res := IsSortable(matrix)
	if res != true {
		t.Errorf("got '%v' from IsSortable(), expected 'true'", res)
	}
}

func TestSorted(t *testing.T) {
	matrix := [][]int64{
		{9, 0, 0, 0, 0},
		{0, 0, 0, 99, 0},
		{0, 2, 0, 0, 0},
		{0, 0, 11, 0, 0},
		{0, 0, 0, 0, 56},
	}
	res := IsSortable(matrix)
	if res != true {
		t.Errorf("got '%v' from IsSortable(), expected 'true'", res)
	}
}

func TestUnsortable(t *testing.T) {
	matrix := [][]int64{
		{9, 0, 23, 55, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 22, 0, 0, 12},
	}
	res := IsSortable(matrix)
	if res != false {
		t.Errorf("got '%v' from IsSortable(), expected 'false'", res)
	}
}

func TestMaxNumOfContainers(t *testing.T) {
	maxItems := int64(1e9)
	maxSize := 100

	matrix := make([][]int64, maxSize)
	for r := range matrix {
		matrix[r] = make([]int64, maxSize)
	}
	for i := 0; i < maxSize; i += 2 {
		matrix[i][i] = maxItems / 2
		matrix[i][i+1] = maxItems / 2
		matrix[i+1][i] = maxItems / 2
		matrix[i+1][i+1] = maxItems / 2
	}

	rand.Shuffle(maxSize, func(i, j int) {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	})

	res := IsSortable(matrix)
	if res != true {
		t.Errorf("got '%v' from IsSortable(), expected 'true'", res)
	}
}
