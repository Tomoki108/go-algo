package combination

import (
	"reflect"
	"testing"
)

func TestPickN(t *testing.T) {
	tests := []struct {
		name     string
		options  []int
		n        int
		expected [][]int
	}{
		{
			name:    "example1",
			options: []int{1, 2, 3, 4},
			n:       2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		{
			name:    "example2",
			options: []int{1, 2, 3, 4},
			n:       3,
			expected: [][]int{
				{1, 2, 3},
				{1, 2, 4},
				{1, 3, 4},
				{2, 3, 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := PickN([]int{}, tt.options, tt.n)
			if len(actual) != len(tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}
		})
	}
}

func TestGrouping(t *testing.T) {
	tests := []struct {
		name     string
		sl       []int
		expected [][][]int
	}{
		{
			name: "example1",
			sl:   []int{1, 2, 3},
			expected: [][][]int{
				{{1, 2, 3}},
				{{1, 2}, {3}},
				{{1, 3}, {2}},
				{{1}, {2, 3}},
				{{1}, {2}, {3}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Grouping(tt.sl)
			if len(actual) != len(tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}
		})
	}
}

func TestGroupingDistintBySize(t *testing.T) {
	tests := []struct {
		name     string
		sl       []int
		size     int
		expected [][][]int
	}{
		{
			name: "example1",
			sl:   []int{1, 2, 3},
			size: 2,
			expected: [][][]int{
				{{1, 2}, {3}},
				{{1, 3}, {2}},
				{{1}, {2, 3}},
				{{2, 3}, {1}},
				{{2}, {1, 3}},
				{{3}, {1, 2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := GroupingDistintBySize(tt.sl, tt.size)
			if len(actual) != len(tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}

			if !reflect.DeepEqual(actual, tt.expected) {
				t.Fatalf("got %v, expect %v", actual, tt.expected)
			}
		})
	}
}
