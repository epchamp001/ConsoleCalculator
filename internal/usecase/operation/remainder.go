package operation

import (
	"fmt"
)

func Remainder(a, b int64) (int64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	}
	return a % b, nil
}
