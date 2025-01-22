package permutation

import (
	"reflect"
	"testing"
)

func BenchmarkPermute(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permute([]int{}, []int{1, 2, 3})
	}
}

func BenchmarkNextPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loop := true
		permutation := []int{1, 2, 3}
		for loop {
			loop = NextPermutation(permutation)
		}
	}
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name     string
		options  []int
		expected [][]int
	}{
		{
			name:    "example1",
			options: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name:    "duplicated elements",
			options: []int{1, 2, 2},
			expected: [][]int{
				{1, 2, 2},
				{2, 1, 2},
				{2, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Permute([]int{}, tt.options)
			if len(actual) != len(tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}
		})
	}
}

func BenchmarkPermute2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Permute2([]int{}, [][]int{{1, 2}, {3, 4}})
	}
}

func TestPermute2(t *testing.T) {
	tests := []struct {
		name     string
		options  [][]int
		expected [][]int
	}{
		{
			name:    "example1",
			options: [][]int{{1, 2}, {3, 4}},
			expected: [][]int{
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
			},
		},
		{
			name:    "duplicated elements",
			options: [][]int{{1, 2}, {2, 3}},
			expected: [][]int{
				{1, 2},
				{1, 3},
				{2, 2},
				{2, 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Permute2([]int{}, tt.options)
			if len(actual) != len(tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}
		})
	}
}
