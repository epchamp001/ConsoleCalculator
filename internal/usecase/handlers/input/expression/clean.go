package expression

import "strings"

func CleanExpression(expression string) string {
	return strings.ReplaceAll(expression, " ", "")
}
