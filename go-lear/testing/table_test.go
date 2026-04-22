package testing

import (
	"context"
	"errors"
	m "test-module/code"
	"test-module/code/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAddPositive(t *testing.T) {
	tests := []struct {
		name      string
		a, b      int
		wantSum   int
		wantError bool
	}{
		{
			name:      "Оба +",
			a:         2,
			b:         3,
			wantSum:   5,
			wantError: false,
		},
		{
			name:      "Первый -",
			a:         -2,
			b:         3,
			wantSum:   0,
			wantError: true,
		},
		{
			name:      "Второй -",
			a:         2,
			b:         -3,
			wantSum:   0,
			wantError: true,
		},
		{
			name:      "оба -",
			a:         -2,
			b:         -3,
			wantSum:   0,
			wantError: true,
		},
		{
			name:      "ноль разрешен",
			a:         0,
			b:         5,
			wantSum:   5,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.AddPositive(tt.a, tt.b)
			if tt.wantError && err == nil {
				t.Errorf("нет ошибки")
			}
			if tt.wantError && err != nil {
				t.Errorf("Есть ошибка %v", err)
			}
			if got != tt.wantSum {
				t.Errorf("AddPositive(%d, %d)= %d want %d", tt.a, tt.b, got, tt.wantSum)
			}
		})
	}
}

func Test(t *testing.T) {
	tests := []struct {
		name         string
		a, b, wanSum int
		wantError    bool
	}{
		{"Оба +", 2, 3, 5, false},
		{"Оба -", -2, -3, 0, true},
		{"Первый -", -2, 3, 0, true},
		{"Второй -", 2, -3, 0, true},
		{"Ноль и число", 0, 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := m.AddPositive(tt.a, tt.b)

			if tt.wantError {
				assert.Error(t, err)
				// ИСПРАВЛЕНО: порядок аргументов: expected, actual
				assert.Equal(t, m.ErrNegativeNumber, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.wanSum, got)
			}
		})
	}
}

func GoMockTestValidator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ctx := context.Background()
	tests := []struct {
		name          string
		a, b          int
		validateError error
		wantSum       int
		wantError     error
	}{
		{
			name:          "success",
			a:             5,
			b:             3,
			validateError: nil,
			wantSum:       8,
			wantError:     nil,
		},
		{
			name:          "validator fails on first arg",
			a:             5,
			b:             3,
			validateError: errors.New("validation failed"),
			wantSum:       0,
			wantError:     errors.New("validation failed"),
		},
		{
			name:          "negative number",
			a:             -1,
			b:             5,
			validateError: nil,
			wantSum:       0,
			wantError:     m.ErrNegativeNumber,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockValidator := mocks.NewMockValidator(ctrl)

			mockValidator.EXPECT().Validator(ctx, tt.a).Return(tt.validateError)
			if tt.validateError == nil {
				mockValidator.EXPECT().Validator(ctx, tt.b).Return(nil)
			}

			gotsum, gotErr := m.AddPositiveWithValidator(ctx, mockValidator, tt.a, tt.b)
			if gotErr != tt.wantError {
				t.Errorf("Неподходящая ошибка. Вышла %s. Ожидали %s", gotErr, tt.wantError)
			}
			if gotsum != tt.wantSum {
				t.Errorf("Неверный результат. Вышло %d. Ожидали %d", gotsum, tt.wantSum)
			}
		})
	}
}
