package combination

import (
	"testing"
)

func TestCalcSurjectionNum(t *testing.T) {
	// 任意のf(n, k)を求める場合は、書き換えて行数横のボタンからテスト実行
	t.Logf("n人を区別のあるk個のグループに分ける場合の数: %d", CalcSurjectionNum(10, 5))

	// test
	{
		tests := []struct {
			name     string
			n        int
			k        int
			expected int
		}{
			{
				name:     "example1",
				n:        10,
				k:        5,
				expected: 5103000,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := CalcSurjectionNum(tt.n, tt.k)
				if actual != tt.expected {
					t.Fatalf("got %v, expect %v", actual, tt.expected)
				}
			})
		}
	}
}

func TestCalcStirlingNum(t *testing.T) {
	// 任意のS(n, k)を求める場合、書き換えて行数横のボタンからテスト実行
	t.Logf("n人を区別のないk個のグループに分ける場合の数: %d", CalcStirlingNum(10, 5))

	// test
	{
		tests := []struct {
			name     string
			n        int
			k        int
			expected int
		}{
			{
				name:     "example1",
				n:        10,
				k:        5,
				expected: 42525,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := CalcStirlingNum(tt.n, tt.k)
				if actual != tt.expected {
					t.Fatalf("got %v, expect %v", actual, tt.expected)
				}
			})
		}
	}
}

func TestBellNum(t *testing.T) {
	// 任意のB(n, k)を求める場合、書き換えて行数横のボタンからテスト実行
	t.Logf("n人を区別のないk個以下のグループに分ける場合の数: %d", CalcBellNum(12, 12))

	// test
	{
		tests := []struct {
			name     string
			n        int
			k        int
			expected int
		}{
			{
				name:     "example1",
				n:        12,
				k:        12,
				expected: 4213597,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				actual := CalcBellNum(tt.n, tt.k)
				if actual != tt.expected {
					t.Fatalf("got %v, expect %v", actual, tt.expected)
				}
			})
		}
	}
}
