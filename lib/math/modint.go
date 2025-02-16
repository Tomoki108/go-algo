package math

type ModInt struct {
	val, modulo int
}

func NewModInt(v, m int) ModInt {
	return ModInt{val: Mod(v, m), modulo: m}
}

func (mi ModInt) Val() int {
	return mi.val
}

func (mi ModInt) Add(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val+a.val, mi.modulo)
}

func (mi ModInt) AddI(a int) ModInt {
	return mi.Add(NewModInt(a, mi.modulo))
}

func (mi ModInt) Sub(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val-a.val, mi.modulo)
}

func (mi ModInt) SubI(a int) ModInt {
	return mi.Sub(NewModInt(a, mi.modulo))
}

func (mi ModInt) Mul(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	return NewModInt(mi.val*a.val, mi.modulo)
}

func (mi ModInt) MulI(a int) ModInt {
	return mi.Mul(NewModInt(a, mi.modulo))
}

// mod mi.moduloでのaの逆元を求め、mi.valに掛ける。逆元が存在するaを渡すこと
func (mi ModInt) Div(a ModInt) ModInt {
	if mi.modulo != a.modulo {
		panic("different modulo")
	}
	inverseElm, err := InverseElmByGCD(a.val, mi.modulo)
	if err != nil {
		panic(err)
	}

	return NewModInt(mi.val*inverseElm, mi.modulo)
}

// mod mi.moduloでのaの逆元を求め、mi.valに掛ける。逆元が存在するaを渡すこと
func (mi ModInt) DivI(a int) ModInt {
	return mi.Div(NewModInt(a, mi.modulo))
}

// 指数expはexp mod Mに置き換えられないので、int型のまま受け取る
func (mi ModInt) Pow(exp int) ModInt {
	return NewModInt(ModExponentiation(mi.val, exp, mi.modulo), mi.modulo)
}
