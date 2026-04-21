package code

import "errors"

var ErrNegativeNumber = errors.New("negative number is not allowed")

func AddPositive(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrNegativeNumber
	}
	return a + b, nil
}
