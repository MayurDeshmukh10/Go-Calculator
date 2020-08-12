package main

import (
	"fmt"

	"./Matrix"
)

func getInput() (float64, float64) {
	var input1 float64
	var input2 float64
	fmt.Println("Enter First Number")
	fmt.Scanln(&input1)
	fmt.Println("Enter Second Number")
	fmt.Scanln(&input2)
	return input1, input2
}

func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return x / y, nil
}

func main() {
	var choice int
	fmt.Println("Simple Calculator")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Matrix-Matrix Addition")
	fmt.Println("6. Matrix-Matrix Subtraction")
	fmt.Println("7. Matrix-Matrix Multiplication")
	fmt.Println("Select a Option :")
	fmt.Scanln(&choice)
	switch {
	case choice == 1:
		input1, input2 := getInput()
		output := input1 + input2
		fmt.Printf("Output of Addition - \n%v\n", output)

	case choice == 2:
		input1, input2 := getInput()
		output := input1 - input2
		fmt.Printf("Output of Subtraction - \n%v\n", output)

	case choice == 3:
		input1, input2 := getInput()
		output := input1 * input2
		fmt.Printf("Output of Multiplication - \n%v\n", output)

	case choice == 4:
		input1, input2 := getInput()
		output, err := divide(input1, input2)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("Output of Multiplication - \n%v\n", output)
		}

	case choice == 5:
		matrix1, matrix2, rows, columns := Matrix.GetMatrixInput()
		var result [][]float64 = Matrix.MatrixAddition(matrix1, matrix2, rows, columns)
		fmt.Println("Matrix-Matrix Addition Output - ")
		Matrix.DisplayMatrix(result, rows, columns)

	case choice == 6:
		matrix1, matrix2, rows, columns := Matrix.GetMatrixInput()
		var result [][]float64 = Matrix.MatrixSubtraction(matrix1, matrix2, rows, columns)
		fmt.Println("Matrix-Matrix Subtraction Output - ")
		Matrix.DisplayMatrix(result, rows, columns)

	case choice == 7:
		matrix1, matrix2, rows, columns := Matrix.GetMatrixInput()
		var result [][]float64 = Matrix.MatrixMultiplication(matrix1, matrix2, rows, columns)
		fmt.Println("Matrix-Matrix Multiplication Output - ")
		Matrix.DisplayMatrix(result, rows, columns)
	}

	// matrix1, matrix2, rows, columns := Matrix.GetMatrixInput()
	// var result [][]float64 = Matrix.MatrixMultiplication(matrix1, matrix2, rows, columns)
	// Matrix.DisplayMatrix(result, rows, columns)

}
