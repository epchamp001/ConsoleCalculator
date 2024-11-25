package expression

import (
	"fmt"
	"regexp"
)

func IsValidExpression(expression string) (bool, error) {
	re := regexp.MustCompile(`^[0-9+\-*/(). %]+$`)

	if re.MatchString(expression) {
		return true, nil
	}
	return false, fmt.Errorf("the expression contains invalid characters")
}
