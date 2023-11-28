package main

import (
	"fmt"
)

func main() {
	var operand1, operand2 float64
	var operator string

	fmt.Print("Enter the first operand: ")
	fmt.Scan(&operand1)

	fmt.Print("Enter the second operand: ")
	fmt.Scan(&operand2)

	fmt.Print("Enter the operator (+, -, *, /, %): ")
	fmt.Scan(&operator)

	result := performOperation(operand1, operand2, operator)
	fmt.Printf("Result: %f %s %f = %f\n", operand1, operator, operand2, result)
}

func performOperation(operand1, operand2 float64, operator string) float64 {
	var result float64

	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 != 0 {
			result = operand1 / operand2
		} else {
			fmt.Println("Error: Division by zero")
			// You can handle this error case as per your application's requirements
		}
	case "%":
		if operand2 != 0 {
			result = float64(int(operand1) % int(operand2))
		} else {
			fmt.Println("Error: Modulo by zero")
			// You can handle this error case as per your application's requirements
		}
	default:
		fmt.Println("Error: Invalid operator")
		// You can handle this error case as per your application's requirements
	}

	return result
}
