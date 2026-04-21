package testing

import (
	m "test-module/code"
	"testing"
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
