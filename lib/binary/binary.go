package binary

// k桁目のビットが1かどうかを判定（一番右を0桁目とする）
func IsBitPop(num uint64, k int) bool {
	// 1 << k はビットマスク。1をk桁左にシフトすることで、k桁目のみが1で他の桁が0の二進数を作る。
	// numとビットマスクの論理積（各桁について、numとビットマスクが両方trueならtrue）を作り、その結果が0でないかどうかで判定できる
	return (num & (1 << k)) != 0
}

// k桁目のビットが立っていれば0に、立っていなければ1にする（一番右を0桁目とする）
func BitFlip(num uint64, k int) uint64 {
	if IsBitPop(num, k) {
		return num & ^(1 << k) // &^ はビットクリア演算子。A &^ Bは、AからBのビットが立っている桁を0にしたものを返す。
	} else {
		return num | (1 << k) // | は論理和演算子。A | Bは、少なくともどちらか一方のビットが立っている桁を1にしたものを返す。
	}
}
