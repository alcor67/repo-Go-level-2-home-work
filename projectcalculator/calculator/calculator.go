package calculator

import (
	"errors"
	"fmt"
	"math"
)

func Calc(op string, x, y float64) (float64, error) {

	switch op {
	case "+":
		return x + y, nil

	case "-":
		return x - y, nil
	case "*":
		return x * y, nil

	case "**":
		result := math.Pow(x, y)
		if math.IsNaN(result) {
			return result, errors.New("result is NaN")
		}

		if math.IsInf(result, 0) {
			return result, errors.New("result is Inf")
		}

		return result, nil

	case "/", "div":
		if y == 0 {
			return 0, fmt.Errorf("can not divide by zero")
		} else {
			return x / y, nil
		}

	case "%":
		if y < 1 {
			return 0, errors.New("can not divide by zero")
		} else {
			return float64(int(x) % int(y)), nil
		}
	default:
		return 0, errors.New("operation not supported, valid operation is + - * / % **")
	}
}
