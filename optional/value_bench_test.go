package optional

import (
	"testing"
)

var (
	aryByteResult []byte
	_             = aryByteResult

	boolResult bool
	_          = boolResult

	notOkBool = Value[bool]{}

	okTrue = Value[bool]{
		v: true, ok: true,
	}

	optBoolResult Value[bool]
	_             = optBoolResult

	optPBoolResult Value[*bool]
	_              = optPBoolResult

	optFunc Value[func()]
	_       = optFunc

	errResult error
	_         = errResult
)

func BenchmarkContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = Contains(true, okTrue)
	}
}

func BenchmarkFlatMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = FlatMap(func(bool) (zero Value[bool]) {
			return
		}, okTrue)
	}
}

func BenchmarkFlatten(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = Flatten(Value[Value[bool]]{})
	}
}

func BenchmarkMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = Map(func(bool) (zero bool) {
			return
		}, okTrue)
	}
}

func BenchmarkOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = Of(true, true)
	}
}

func BenchmarkOfAssert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = OfAssert[bool, bool](true)
	}
}

func BenchmarkOfFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optFunc = OfFunc(func() {})
	}
}

func BenchmarkOfIndex(b *testing.B) {
	m := map[bool]bool{true: true}

	for i := 0; i < b.N; i++ {
		optBoolResult = OfIndex(m, true)
	}
}

func BenchmarkOfNonZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = OfNonZero(true)
	}
}

func BenchmarkOfNotOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = OfNotOk[bool]()
	}
}

func BenchmarkOfOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = OfOk(true)
	}
}

func BenchmarkOfPointee(b *testing.B) {
	t := true

	for i := 0; i < b.N; i++ {
		optBoolResult = OfPointee(&t)
	}
}

func BenchmarkOfPointer(b *testing.B) {
	t := true

	for i := 0; i < b.N; i++ {
		optPBoolResult = OfPointer(&t)
	}
}

//func BenchmarkOfReceive(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//
//	}
//}

func BenchmarkUnzipWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult, optBoolResult = UnzipWith(okTrue, func(bool) (zero bool, zero2 bool) {
			return
		})
	}
}

func BenchmarkValue_Do(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = okTrue.Do(func(bool) {})
	}
}

func BenchmarkValue_Filter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = okTrue.Filter(func(bool) (zero bool) {
			return
		})
	}
}

func BenchmarkValue_IsOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = okTrue.IsOk()
	}
}

func BenchmarkValue_MarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aryByteResult, _ = okTrue.MarshalJSON()
	}
}

func BenchmarkValue_MustOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = okTrue.MustOk()
	}
}

func BenchmarkValue_OfOk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = OfOk(true)
	}
}

func BenchmarkValue_Ok(b *testing.B) {
	t := true

	for i := 0; i < b.N; i++ {
		boolResult = okTrue.Ok(&t)
	}
}

func BenchmarkValue_Or(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = okTrue.Or(true)
	}
}

func BenchmarkValue_OrDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = okTrue.OrTake(func() (zero bool) {
			return
		})
	}
}

func BenchmarkValue_OrZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		boolResult = okTrue.OrZero()
	}
}

func BenchmarkValue_SelfOr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = okTrue.SelfOr(notOkBool)
	}
}

func BenchmarkValue_SelfOrDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = okTrue.SelfOrTake(func() (zero Value[bool]) {
			return
		})
	}
}

func BenchmarkValue_UnmarshalJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		errResult = okTrue.UnmarshalJSON([]byte(""))
	}
}

func BenchmarkValue_Unpack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, boolResult = okTrue.Unpack()
	}
}

func BenchmarkZipWith(b *testing.B) {
	for i := 0; i < b.N; i++ {
		optBoolResult = ZipWith(okTrue, notOkBool, func(_, _ bool) (zero bool) {
			return
		})
	}
}
