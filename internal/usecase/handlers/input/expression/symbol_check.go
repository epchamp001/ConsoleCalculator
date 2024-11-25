package expression

import (
	"fmt"
	"unicode"
)

func CheckSymbolOrder(expression string) error {
	var prev rune
	for i, char := range expression {
		if i == 0 {
			if !(unicode.IsDigit(char) || char == '-' || char == '(') {
				return fmt.Errorf("an expression cannot start with a character: '%c'", char)
			}
		} else {
			if err := checkInvalidCombinations(prev, char, i); err != nil {
				return err
			}
		}
		prev = char
	}

	if !(unicode.IsDigit(prev) || prev == ')') {
		return fmt.Errorf("the expression cannot end with the character '%c'", prev)
	}

	return nil
}

func checkInvalidCombinations(prev, char rune, pos int) error {
	if isOperator(prev) && isOperator(char) {
		return fmt.Errorf("two operators in a row '%c%c' in position %d", prev, char, pos)
	}

	if unicode.IsDigit(char) && prev == ')' {
		return fmt.Errorf("the number immediately after ')' in position %d", pos)
	}

	if char == '(' && (unicode.IsDigit(prev) || prev == ')') {
		return fmt.Errorf("'(' after the number or ')' in position %d", pos)
	}

	if char == ')' && (isOperator(prev) || prev == '(') {
		return fmt.Errorf("')' after '%c' in position %d", prev, pos)
	}

	return nil
}

func isOperator(char rune) bool {
	return char == '+' || char == '-' || char == '*' || char == '/' || char == '%'
}
