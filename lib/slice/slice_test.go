package slice

import (
	"reflect"
	"testing"
)

func TestSplitByChunks(t *testing.T) {
	tests := []struct {
		name      string
		sl        []int
		chunkSize int
		want      [][]int
	}{
		{
			name:      "simple case",
			sl:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			chunkSize: 3,
			want: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10},
			},
		},
		{
			name:      "over size",
			sl:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			chunkSize: 10,
			want: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9},
			},
		},
		{
			name:      "empty slice",
			sl:        []int{},
			chunkSize: 3,
			want:      [][]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SplitByChunks(tt.sl, tt.chunkSize)
			if len(got) != len(tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}

			for i := 0; i < len(got); i++ {
				if !reflect.DeepEqual(got[i], tt.want[i]) {
					t.Errorf("got: %v, want: %v", got, tt.want)
				}
			}
		})
	}
}
