package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter an expression: ")
	expression, _ := reader.ReadString('\n')

	result := calculate(expression)
	fmt.Printf("Result: %f\n", result)
}

func calculate(expression string) float64 {
	// Remove trailing newline character
	expression = strings.TrimSuffix(expression, "\n")

	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Invalid expression")
		os.Exit(1)
	}

	operand1, err := strconv.ParseFloat(tokens[0], 64)
	if err != nil {
		fmt.Println("Invalid operand 1:", err)
		os.Exit(1)
	}

	operator := tokens[1]

	operand2, err := strconv.ParseFloat(tokens[2], 64)
	if err != nil {
		fmt.Println("Invalid operand 2:", err)
		os.Exit(1)
	}

	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Division by zero")
			os.Exit(1)
		}
		result = operand1 / operand2
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}

	return result
}
