package tree

import (
	"reflect"
	"testing"
)

func TestTreeSet(t *testing.T) {
	set := NewTreeSet[int]()
	set.Add(3)
	set.Add(1)
	set.Add(4)
	set.Add(2)

	expected1 := []int{1, 2, 3, 4}
	got1 := set.GetAll()
	if !reflect.DeepEqual(expected1, got1) {
		t.Errorf("got: %v, expected: %v", got1, expected1)
	}

	set.Remove(3)
	expected2 := []int{1, 2, 4}
	got2 := set.GetAll()
	if !reflect.DeepEqual(expected2, got2) {
		t.Errorf("got: %v, expected: %v", got2, expected2)
	}

	if !set.Contains(2) { // should be true
		t.Errorf("got: %v, expected: %v", false, true)
	}

	if set.Contains(3) { // should be false
		t.Errorf("got: %v, expected: %v", true, false)
	}
}

func TestMultiSet(t *testing.T) {
	ms := NewMultiSet[int]()
	ms.Add(3)
	ms.Add(1)
	ms.Add(4)
	ms.Add(3)
	ms.Add(2)

	expected1 := []int{1, 2, 3, 3, 4}
	got1 := ms.Range(1, 4)
	if !reflect.DeepEqual(expected1, got1) {
		t.Errorf("got: %v, expected: %v", got1, expected1)
	}

	ms.Remove(3)
	expected2 := []int{1, 2, 3, 4}
	got2 := ms.Range(1, 4)
	if !reflect.DeepEqual(expected2, got2) {
		t.Errorf("got: %v, expected: %v", got2, expected2)
	}
}
