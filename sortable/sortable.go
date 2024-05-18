package sortable

func IsSortable(matrix [][]int64) bool {
	n := len(matrix)
	rowSums := make([]int64, n)
	colSums := make([]int64, n)
	excludedRows := make(map[int]struct{})

	for i := 0; i < n; i++ {
		rowSums[i] = rowSum(matrix, i)
		colSums[i] = columnSum(matrix, i)
	}

	for colorIndex := 0; colorIndex < n; colorIndex++ {
		foundMatchedSum := false
		for i := 0; i < n; i++ {
			if _, ok := excludedRows[i]; ok {
				continue
			}
			if colSums[colorIndex] == rowSums[i] {
				excludedRows[i] = struct{}{}
				foundMatchedSum = true
				break
			}
		}
		if !foundMatchedSum {
			return false
		}
	}

	return true
}

func columnSum(matrix [][]int64, colNum int) int64 {
	n := len(matrix)
	var sum int64
	for i := 0; i < n; i++ {
		sum += matrix[i][colNum]
	}
	return sum
}

func rowSum(matrix [][]int64, rowNum int) int64 {
	n := len(matrix)
	var sum int64
	for j := 0; j < n; j++ {
		sum += matrix[rowNum][j]
	}
	return sum
}
