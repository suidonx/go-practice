package math

import "testing"

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) = %d; want 3", result)
	}
}

// テーブル駆動テスト
// 複数の入力値を1つのテストケースでテストする。
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		name string
		a int
		b int 
		want int
	}{
		{"1+2=3", 1, 2, 3},
		{"-1+1=0", -1, 1, 0},
		{"0+0=0", 0, 0, 0},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}