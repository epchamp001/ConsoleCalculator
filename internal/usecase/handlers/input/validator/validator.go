package validator

import (
	expr "Calculator/internal/usecase/handlers/input/expression"
	"fmt"
	"strings"
)

func ValidateExpression(expression string) (string, error) {
	newExpression := expr.CleanExpression(expression)
	if strings.Count(newExpression, "(") != strings.Count(newExpression, ")") {
		return "", fmt.Errorf("incorrect number of brackets ()")
	}

	if ok, err := expr.IsValidExpression(newExpression); !ok {
		return "", err
	}

	if err := expr.CheckSymbolOrder(newExpression); err != nil {
		return "", err
	}

	return newExpression, nil
}
