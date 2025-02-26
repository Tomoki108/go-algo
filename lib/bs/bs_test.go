package bs

import (
	"reflect"
	"testing"
)

func TestAscIntSearch(t *testing.T) {
	sl1 := []int{1, 4, 5, 5, 6, 20}
	sl2 := []int{5}
	sl3 := []int{4}

	tests := []struct {
		name string
		low  int
		high int
		f    func(num int) bool
		want int
	}{
		{
			name: "example1",
			low:  0,
			high: 4,
			f:    func(num int) bool { return sl1[num] >= 5 },
			want: 2,
		},
		{
			name: "example2",
			low:  0,
			high: 0,
			f:    func(num int) bool { return sl2[num] >= 5 },
			want: 0,
		},
		{
			name: "example3",
			low:  0,
			high: 0,
			f:    func(num int) bool { return sl3[num] >= 5 },
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AscIntSearch(tt.low, tt.high, tt.f); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestDescIntSearch(t *testing.T) {
	sl1 := []int{1, 1, 5, 5, 5, 6, 20, 21, 22, 22}
	sl2 := []int{5}
	sl3 := []int{9}

	tests := []struct {
		name string
		high int
		low  int
		f    func(num int) bool
		want int
	}{
		{
			name: "example1",
			high: 8,
			low:  0,
			f:    func(num int) bool { return sl1[num] <= 5 },
			want: 4,
		},
		{
			name: "example2",
			high: 0,
			low:  0,
			f:    func(num int) bool { return sl2[num] <= 5 },
			want: 0,
		},
		{
			name: "example3",
			high: 0,
			low:  0,
			f:    func(num int) bool { return sl3[num] <= 5 },
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DescIntSearch(tt.high, tt.low, tt.f); got != tt.want {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestRangeIntSearch(t *testing.T) {
	sl1 := []int{1, 1, 5, 5, 5, 6, 20, 21, 22, 22}
	sl2 := []int{5}
	sl3 := []int{9}

	tests := []struct {
		name string
		sl   []int
		l    int
		r    int
		want []int
	}{
		{
			name: "example1",
			sl:   sl1,
			l:    5,
			r:    21,
			want: []int{5, 5, 5, 6, 20, 21},
		},
		{
			name: "example2",
			sl:   sl2,
			l:    5,
			r:    5,
			want: []int{5},
		},
		{
			name: "example3",
			sl:   sl3,
			l:    5,
			r:    5,
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RangeIntSearch(tt.sl, tt.l, tt.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got: %v, want: %v", got, tt.want)
			}
		})
	}
}

func TestRangeSearch(t *testing.T) {
	type person struct {
		Name string
		Age  int
	}

	persons := []person{
		{"Alice", 20},
		{"Bob", 25},
		{"Charlie", 30},
		{"Dave", 35},
		{"Eve", 40},
	}

	// valuer extracts the Age field.
	ageValuer := func(p person) int { return p.Age }

	// Expect to extract persons with Age between 26 and 39.
	got := RangeSearch(persons, ageValuer, 26, 39)
	want := []person{
		{"Charlie", 30},
		{"Dave", 35},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("RangeSearch(persons, 26, 39) = %v, want %v", got, want)
	}
}
