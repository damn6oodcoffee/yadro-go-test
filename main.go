package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/damn6oodcoffee/yadro-go-test/sortable"
)

func main() {
	var size int
	var matrix [][]int64

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	size, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic("can't convert input to int")
	}

	matrix = make([][]int64, size)
	for r := 0; r < size; r++ {
		matrix[r] = make([]int64, size)
		scanner.Scan()
		row := strings.Fields(scanner.Text())
		for c, s := range row {
			val, err := strconv.Atoi(s)
			if err != nil {
				panic("can't convert input to a slice of type int")
			}
			matrix[r][c] = int64(val)
		}
	}

	fmt.Println(matrix)
	if sortable.IsSortable(matrix) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
