package ptr

import (
	"testing"
)

func TestToIntDef(t *testing.T) {
	equal(t, int(10), ToIntDef(Int(10), 0))
	equal(t, int(5), ToIntDef(nil, 5))
}

func TestToInt8Def(t *testing.T) {
	equal(t, int8(10), ToInt8Def(Int8(10), 5))
	equal(t, int8(5), ToInt8Def(nil, 5))
}

func TestToInt16Def(t *testing.T) {
	equal(t, int16(10), ToInt16Def(Int16(10), 5))
	equal(t, int16(5), ToInt16Def(nil, 5))
}

func TestToInt32Def(t *testing.T) {
	equal(t, int32(10), ToInt32Def(Int32(10), 5))
	equal(t, int32(5), ToInt32Def(nil, 5))
}

func TestToInt64Def(t *testing.T) {
	equal(t, int64(10), ToInt64Def(Int64(10), 5))
	equal(t, int64(5), ToInt64Def(nil, 5))
}

func TestToUIntDef(t *testing.T) {
	equal(t, uint(10), ToUIntDef(UInt(10), 5))
	equal(t, uint(5), ToUIntDef(nil, 5))
}

func TestToUInt8Def(t *testing.T) {
	equal(t, uint8(10), ToUInt8Def(UInt8(10), 5))
	equal(t, uint8(5), ToUInt8Def(nil, 5))
}

func TestToUInt16Def(t *testing.T) {
	equal(t, uint16(10), ToUInt16Def(UInt16(10), 5))
	equal(t, uint16(5), ToUInt16Def(nil, 5))
}

func TestToUInt32Def(t *testing.T) {
	equal(t, uint32(10), ToUInt32Def(UInt32(10), 5))
	equal(t, uint32(5), ToUInt32Def(nil, 5))
}

func TestToUInt64Def(t *testing.T) {
	equal(t, uint64(10), ToUInt64Def(UInt64(10), 5))
	equal(t, uint64(5), ToUInt64Def(nil, 5))
}

func TestToByteDef(t *testing.T) {
	equal(t, byte(10), ToByteDef(Byte(10), 5))
	equal(t, byte(5), ToByteDef(nil, 5))
}

func TestToRuneDef(t *testing.T) {
	equal(t, rune('🤔'), ToRuneDef(Rune('🤔'), '😏'))
	equal(t, rune('😏'), ToRuneDef(nil, '😏'))
}

func TestToBoolDef(t *testing.T) {
	equal(t, bool(true), ToBoolDef(Bool(true), false))
	equal(t, bool(false), ToBoolDef(nil, false))
}

func TestToStringDef(t *testing.T) {
	equal(t, string("🤔"), ToStringDef(String("🤔"), "😏"))
	equal(t, string("😏"), ToStringDef(nil, "😏"))
}

func TestToFloat32Def(t *testing.T) {
	equal(t, float32(10.1), ToFloat32Def(Float32(10.1), 5.1))
	equal(t, float32(5.1), ToFloat32Def(nil, 5.1))
}

func TestToFloat64Def(t *testing.T) {
	equal(t, float64(10.1), ToFloat64Def(Float64(10.1), 5.1))
	equal(t, float64(5.1), ToFloat64Def(nil, 5.1))
}

func TestToComplex64Def(t *testing.T) {
	equal(t, complex(float32(10.0), float32(100.0)), ToComplex64Def(Complex64(complex(float32(10.0), float32(100.0))), complex(float32(5.0), float32(50.0))))
	equal(t, complex(float32(5.0), float32(50.0)), ToComplex64Def(nil, complex(float32(5.0), float32(50.0))))
}

func TestToComplex128Def(t *testing.T) {
	equal(t, complex(float64(10.0), float64(100.0)), ToComplex128Def(Complex128(complex(float64(10.0), float64(100.0))), complex(float64(5.0), float64(50.0))))
	equal(t, complex(float64(5.0), float64(50.0)), ToComplex128Def(nil, complex(float64(5.0), float64(50.0))))
}
