package arrs

import "fmt"

type matrix210 [2][10]int
type matrix410 [4][10]int

// Main for TestMatrix
func Main() {
	fmt.Println("testing arrays\n...")

	var _matrix1 [2][10]int
	var _matrix2 [4][10]int

	m1 := matrix210(_matrix1)
	m2 := matrix410(_matrix2)

	printMatrix(m1.toSlices())
	fmt.Println("---")
	printMatrix(m2.toSlices())

}

func (matrix matrix210) toSlices() [][]int {
	slices := make([][]int, len(matrix))
	for i := range matrix {
		slices[i] = matrix[i][:]
	}
	return slices
}

func (matrix matrix410) toSlices() [][]int {
	slices := make([][]int, len(matrix))
	for i := range matrix {
		slices[i] = matrix[i][:]
	}
	return slices
}

func printMatrix(m [][]int) {
	for y := range m {
		fmt.Printf("[%d] ", y)
		row := m[y]
		for x := range row {
			fmt.Printf("[%d]", m[y][x])
		}
		fmt.Println("")
	}
}
