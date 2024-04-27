package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the expression: ")
	expression, _ := reader.ReadString('\n')

	result, operand1Type, operand2Type := calculateResult(expression)
	if result != nil {
		fmt.Printf("Result: %s\n", formatResult(result, operand1Type, operand2Type))
	}
}

func calculateResult(expression string) (interface{}, string, string) {
	// Remove trailing newline character
	expression = strings.TrimSuffix(expression, "\n")

	tokens := strings.Split(expression, " ")
	if len(tokens) != 3 {
		fmt.Println("Invalid expression")
		os.Exit(1)
	}

	operand1Type := determineNumberType(tokens[0]) // Определяем тип первого операнда
	operator := tokens[1]
	operand2Type := determineNumberType(tokens[2]) // Определяем тип второго операнда

	if operand1Type != operand2Type {
		fmt.Println("Mixing of arabic and roman numerals is not allowed")
		return nil, "", ""
	}

	operand1 := parseOperand(tokens[0], operand1Type)
	operand2 := parseOperand(tokens[2], operand2Type)

	checkRange(operand1)
	checkRange(operand2)
	checkInteger(operand1, operand1Type)
	checkInteger(operand2, operand2Type)

	var result interface{}
	switch operator {
	case "+":
		result = add(operand1, operand2)
	case "-":
		result = subtract(operand1, operand2)
	case "*":
		result = multiply(operand1, operand2)
	case "/":
		result = divide(operand1, operand2)
	default:
		fmt.Println("Invalid operator:", operator)
		return nil, "", ""
	}

	return result, operand1Type, operand2Type
}

func formatResult(result interface{}, operand1Type string, operand2Type string) string {
	// Проверяем, является ли результат числом
	switch v := result.(type) {
	case float64:
		num := v
		if operand1Type == "roman" || operand2Type == "roman" {
			// Если результат меньше 1, выходим из программы с ошибкой
			if num < 1 {
				fmt.Println("Result cannot be less than 1 for Roman numerals")
				os.Exit(1)
			}
			return arabicToRoman(int(num))
		} else {
			if num < 1 || num > 10 {
				return fmt.Sprintf("%.2f", num)
			}
			return fmt.Sprintf("%.0f", num)
		}
	default:
		return fmt.Sprintf("%v", result)
	}
}

func arabicToRoman(num int) string {
	// Маппинг арабских чисел на римские
	romanMap := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}
	return romanMap[num]
}

func parseOperand(token string, numberType string) float64 {
	if numberType == "arabic" {
		operand, err := strconv.ParseFloat(token, 64)
		if err != nil {
			fmt.Println("Invalid operand:", token)
			os.Exit(1)
		}
		return operand
	}

	if numberType == "roman" {
		romanMap := map[string]int{
			"I":    1,
			"II":   2,
			"III":  3,
			"IV":   4,
			"V":    5,
			"VI":   6,
			"VII":  7,
			"VIII": 8,
			"IX":   9,
			"X":    10,
			// Добавьте другие римские числа по мере необходимости
		}

		if num, ok := romanMap[token]; ok {
			if num < 1 || num > 10 {
				fmt.Println("Roman numbers must be between I and X")
				os.Exit(1)
			}
			return float64(num)
		}

		fmt.Println("Invalid operand:", token)
		os.Exit(1)
	}

	return 0
}

func determineNumberType(token string) string {
	if isRomanNumeral(token) {
		return "roman"
	}

	if isArabicNumeral(token) {
		return "arabic"
	}

	fmt.Println("Invalid number format:", token)
	os.Exit(1)
	return ""
}

func isRomanNumeral(str string) bool {
	// Проверка на римское число
	match, _ := regexp.MatchString(`^[IVXLCDM]+$`, str)
	return match
}

func isArabicNumeral(str string) bool {
	// Проверка на арабское число
	_, err := strconv.Atoi(str)
	return err == nil
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Division by zero")
		os.Exit(1)
	}
	// Отбрасываем остаток и возвращаем только целую часть результата
	return math.Trunc(a / b)
}

func checkRange(num float64) {
	if num < 1 || num > 10 {
		fmt.Println("Numbers must be between 1 and 10")
		os.Exit(1)
	}
}

func checkInteger(num float64, numberType string) {
	if numberType == "arabic" && num != float64(int(num)) {
		fmt.Println("Numbers must be integers")
		os.Exit(1)
	}
}
