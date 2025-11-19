package myknowledge

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j) // thêm giá trị của col vào row
		}
		matrix = append(matrix, row) // thêm giá trị của cả 1 row vào ma trận
	}
	return matrix
}
func TestMatrix() {

	matrix := createMatrix(10, 10)

	fmt.Println(matrix)
	fmt.Println("----------------------------------------------")

	for i := 0; i < len(matrix); i++ {
		fmt.Printf("%d \n", matrix[i])
	}
	fmt.Println("==============================================")
}
