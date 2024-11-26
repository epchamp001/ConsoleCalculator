package operation

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateExpression(expression string) (float64, error) {
	for strings.Contains(expression, "(") {
		openIndex := strings.LastIndex(expression, "(")
		closeIndex := strings.Index(expression[openIndex:], ")") + openIndex

		subExpression := expression[openIndex+1 : closeIndex]
		result, err := CalculateExpression(subExpression)
		if err != nil {
			return 0, err
		}

		expression = expression[:openIndex] + fmt.Sprintf("%.8f", result) + expression[closeIndex+1:]
	}

	return evaluteFlatExpression(expression)
}

func evaluteFlatExpression(expression string) (float64, error) {
	var err error
	for strings.ContainsAny(expression, "%*/") {
		if _, parseErr := strconv.ParseFloat(expression, 64); parseErr == nil {
			break
		}

		expression, err = calculateSimpleExpression(expression, "*/%")
		if err != nil {
			return 0, err
		}
	}

	for strings.ContainsAny(expression, "+-") {
		if _, parseErr := strconv.ParseFloat(expression, 64); parseErr == nil {
			break
		}

		expression, err = calculateSimpleExpression(expression, "+-")
		if err != nil {
			return 0, err
		}
	}

	return strconv.ParseFloat(expression, 64)
}

func calculateSimpleExpression(expression, operators string) (string, error) {
	for i := 0; i < len(expression); i++ {
		char := expression[i]

		if strings.ContainsRune(operators, rune(char)) {
			left, leftIndex := findLeftOperand(expression, i)
			right, rightIndex := findRightOperand(expression, i)

			leftNum, err := strconv.ParseFloat(left, 64)
			if err != nil {
				return "", fmt.Errorf("invalid operand: %s", left)
			}

			rightNum, err := strconv.ParseFloat(right, 64)
			if err != nil {
				return "", fmt.Errorf("invalid operand: %s", right)
			}

			var result float64
			switch char {
			case '+':
				result = Add(leftNum, rightNum)
			case '-':
				result = Subtract(leftNum, rightNum)
			case '*':
				result = Multiply(leftNum, rightNum)
			case '/':
				result, err = Divide(leftNum, rightNum)
				if err != nil {
					return "", err
				}
			case '%':
				remainder, err := Remainder(int64(leftNum), int64(rightNum))
				if err != nil {
					return "", err
				}
				result = float64(remainder)
			}

			newExpression := expression[:leftIndex] + fmt.Sprintf("%.8f", result) + expression[rightIndex:]
			return newExpression, nil
		}
	}
	return expression, nil
}

func findLeftOperand(expression string, index int) (string, int) {
	start := index - 1
	for start >= 0 && (expression[start] >= '0' && expression[start] <= '9' || expression[start] == '.') {
		start--
	}
	return expression[start+1 : index], start + 1
}

func findRightOperand(expression string, index int) (string, int) {
	end := index + 1
	for end < len(expression) && (expression[end] >= '0' && expression[end] <= '9' || expression[end] == '.') {
		end++
	}
	return expression[index+1 : end], end
}
