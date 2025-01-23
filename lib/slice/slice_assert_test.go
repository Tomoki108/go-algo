package slice

import "testing"

func TestEqualsWithAtMostOneDiff(t *testing.T) {
	tests := []struct {
		name     string
		sl1, sl2 []int
		want     bool
	}{
		{"same", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"elements math", []int{1, 2, 3}, []int{1, 3, 2}, false},
		{"one diff", []int{1, 2, 3}, []int{1, 2, 4}, true},
		{"two diff", []int{1, 2, 3}, []int{1, 4, 2}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsWithAtMostOneDiff(tt.sl1, tt.sl2)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestEqualsWithOneInsertion(t *testing.T) {
	tests := []struct {
		name    string
		longer  []int
		shorter []int
		want    bool
	}{
		{"equal", []int{1, 2, 4, 3}, []int{1, 2, 3}, true},
		{"different", []int{0, 2, 4, 3}, []int{1, 2, 3}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EqualsWithOneInsertion(tt.longer, tt.shorter)
			if got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}
