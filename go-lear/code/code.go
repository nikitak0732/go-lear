package code

import (
	"context"
	"errors"
)

var ErrNegativeNumber = errors.New("negative number is not allowed")

//go:generate mockgen -source=code.go -destination=mocks/mock_validator.go -package=mocks Validator
type Validator interface {
	Validator(ctx context.Context, value int) error
}

func AddPositive(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, ErrNegativeNumber
	}
	return a + b, nil
}

func AddPositiveWithValidator(ctx context.Context, validator Validator, a, b int) (int, error) {
	if err := validator.Validator(ctx, a); err != nil {
		return 0, err
	}

	if err := validator.Validator(ctx, b); err != nil {
		return 0, err
	}

	if a < 0 || b < 0 {
		return 0, ErrNegativeNumber
	}
	return a + b, nil
}
