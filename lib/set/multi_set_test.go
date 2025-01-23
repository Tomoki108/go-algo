package set

import (
	"reflect"
	"testing"

	"github.com/liyue201/gostl/utils/comparator"
)

// TestMultiSetInt は、int 型のテストです。
func TestMultiSetInt(t *testing.T) {
	ms := NewMultiSet[int](comparator.IntComparator)

	// 1) 初期状態のテスト
	if ms.Size() != 0 {
		t.Errorf("initial Size() = %d, want 0", ms.Size())
	}
	if len(ms.Values()) != 0 {
		t.Errorf("initial Values() should be empty, got: %v", ms.Values())
	}

	// 2) 挿入テスト
	ms.Insert(5)
	ms.Insert(5)
	ms.Insert(10)
	ms.Insert(5)

	// SizeとCountの検証
	if got, want := ms.Size(), 4; got != want {
		t.Errorf("Size() = %d, want %d", got, want)
	}
	if got, want := ms.Count(5), 3; got != want {
		t.Errorf("Count(5) = %d, want %d", got, want)
	}
	if got, want := ms.Count(10), 1; got != want {
		t.Errorf("Count(10) = %d, want %d", got, want)
	}
	if got, want := ms.Count(1), 0; got != want {
		t.Errorf("Count(1) = %d, want %d (not inserted yet)", got, want)
	}

	// 3) Erase (削除) のテスト
	ms.Erase(5) // 5を1つだけ削除
	if got, want := ms.Count(5), 2; got != want {
		t.Errorf("after Erase(5), Count(5) = %d, want %d", got, want)
	}
	if got, want := ms.Size(), 3; got != want {
		t.Errorf("after Erase(5), Size() = %d, want %d", got, want)
	}

	ms.Erase(10) // 10を1つ削除（これで10は0になる）
	if got, want := ms.Count(10), 0; got != want {
		t.Errorf("after Erase(10), Count(10) = %d, want %d", got, want)
	}

	// 4) Values のテスト
	//  いまマルチセットには 5(×2) が残っている
	gotVals := ms.Values()
	wantVals := []int{5, 5}
	if !reflect.DeepEqual(gotVals, wantVals) {
		t.Errorf("Values() = %v, want %v", gotVals, wantVals)
	}

	// 5) Clear のテスト
	ms.Clear()
	if got, want := ms.Size(), 0; got != want {
		t.Errorf("after Clear(), Size() = %d, want %d", got, want)
	}
	if got := ms.Values(); len(got) != 0 {
		t.Errorf("after Clear(), Values() should be empty, got: %v", got)
	}
}

// TestMultiSetString は、string 型のテストです。
func TestMultiSetString(t *testing.T) {
	ms := NewMultiSet[string](comparator.StringComparator)

	// 挿入
	ms.Insert("apple")
	ms.Insert("banana")
	ms.Insert("apple")

	// サイズ/カウント確認
	if got, want := ms.Size(), 3; got != want {
		t.Errorf("Size() = %d, want %d", got, want)
	}
	if got, want := ms.Count("apple"), 2; got != want {
		t.Errorf("Count(apple) = %d, want %d", got, want)
	}
	if got, want := ms.Count("banana"), 1; got != want {
		t.Errorf("Count(banana) = %d, want %d", got, want)
	}
	if got, want := ms.Count("cherry"), 0; got != want {
		t.Errorf("Count(cherry) = %d, want %d (not inserted yet)", got, want)
	}

	// Erase
	ms.Erase("apple")
	if got, want := ms.Count("apple"), 1; got != want {
		t.Errorf("after Erase(apple), Count(apple) = %d, want %d", got, want)
	}

	// Clear
	ms.Clear()
	if got, want := ms.Size(), 0; got != want {
		t.Errorf("after Clear(), Size() = %d, want %d", got, want)
	}
	if got := ms.Values(); len(got) != 0 {
		t.Errorf("after Clear(), Values() should be empty, got: %v", got)
	}
}
