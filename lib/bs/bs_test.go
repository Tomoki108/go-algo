package bs

import "testing"

func TestSearchDescending(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		x        int
		expected int
	}{
		{
			name:     "Find last element >= x",
			slice:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			x:        7,
			expected: 3,
		},
		{
			name:     "All elements < x",
			slice:    []int{10, 9, 8, 7, 6},
			x:        11,
			expected: 5,
		},
		{
			name:     "All elements >= x",
			slice:    []int{10, 9, 8, 7, 6},
			x:        5,
			expected: 4,
		},
		{
			name:     "Mixed elements",
			slice:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			x:        4,
			expected: 6,
		},
		{
			name:     "Single element matching",
			slice:    []int{5},
			x:        5,
			expected: 0,
		},
		{
			name:     "Single element not matching",
			slice:    []int{10},
			x:        5,
			expected: 1,
		},
		{
			name:     "No matching element",
			slice:    []int{10, 9, 8, 7, 6},
			x:        4,
			expected: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SearchDescending(len(test.slice), func(i int) bool {
				return test.slice[i] >= test.x
			})
			if result != test.expected {
				t.Errorf("got %d, want %d", result, test.expected)
			}
		})
	}
}
