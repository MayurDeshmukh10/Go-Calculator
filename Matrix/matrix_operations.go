package Matrix

import "fmt"

func GetMatrixInput() ([][]float64, [][]float64, int, int) {
	var matrix1 [][]float64
	var matrix2 [][]float64
	var rows int
	var columns int
	var value float64
	fmt.Println("Enter No. of Rows :")
	fmt.Scanln(&rows)
	fmt.Println("Enter No. of Columns :")
	fmt.Scanln(&columns)
	fmt.Println("Enter values of Matrix 1 :")
	for i := 0; i < rows; i++ {
		var temp []float64
		for j := 0; j < columns; j++ {
			fmt.Scanln(&value)
			temp = append(temp, value)
		}
		matrix1 = append(matrix1, temp)
	}
	fmt.Println("Enter values of Matrix 2 :")
	for i := 0; i < rows; i++ {
		var temp []float64
		for j := 0; j < columns; j++ {
			fmt.Scanln(&value)
			temp = append(temp, value)
		}
		matrix2 = append(matrix2, temp)
	}
	return matrix1, matrix2, rows, columns
}

func DisplayMatrix(matrix [][]float64, rows, columns int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		fmt.Println()
	}
}

func MatrixAddition(matrix1, matrix2 [][]float64, rows, columns int) [][]float64 {
	var result [][]float64
	for i := 0; i < rows; i++ {
		var temp []float64
		for j := 0; j < columns; j++ {
			temp = append(temp, matrix1[i][j]+matrix2[i][j])
		}
		result = append(result, temp)
	}
	return result
}

func MatrixSubtraction(matrix1, matrix2 [][]float64, rows, columns int) [][]float64 {
	var result [][]float64
	for i := 0; i < rows; i++ {
		var temp []float64
		for j := 0; j < columns; j++ {
			temp = append(temp, matrix1[i][j]-matrix2[i][j])
		}
		result = append(result, temp)
	}
	return result
}

func MatrixMultiplication(matrix1, matrix2 [][]float64, rows, columns int) [][]float64 {
	var result [][]float64
	for i := 0; i < rows; i++ {
		var temp []float64
		for j := 0; j < columns; j++ {
			var value float64 = 0
			for k := 0; k < columns; k++ {
				value += matrix1[i][k] * matrix2[k][j]
			}
			temp = append(temp, value)
		}
		result = append(result, temp)
	}
	return result
}
