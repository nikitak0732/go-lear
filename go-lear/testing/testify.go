package testing

// func Test(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		a, b, wanSum int
// 		wantError    bool
// 	}{
// 		{"Оба +", 2, 3, 5, false},
// 		{"Оба -", -2, -3, 0, true},
// 		{"Первый -", -2, 3, 0, true},
// 		{"Второй -", 2, -3, 0, true},
// 		{"Ноль и число", 0, 5, 5, false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := m.AddPositive(tt.a, tt.b)

// 			if tt.wantError {
// 				assert.Error(t, err)
// 				// ИСПРАВЛЕНО: порядок аргументов: expected, actual
// 				assert.Equal(t, m.ErrNegativeNumber, err)
// 			} else {
// 				assert.NoError(t, err)
// 				assert.Equal(t, tt.wanSum, got)
// 			}
// 		})
// 	}
// }
