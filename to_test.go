package ptr

import (
	"testing"
	"time"
)

func TestToInt(t *testing.T) {
	equal(t, int(10), ToInt(Int(10)))
	equal(t, int(0), ToInt(nil))
}

func TestToInt8(t *testing.T) {
	equal(t, int8(10), ToInt8(Int8(10)))
	equal(t, int8(0), ToInt8(nil))
}

func TestToInt16(t *testing.T) {
	equal(t, int16(10), ToInt16(Int16(10)))
	equal(t, int16(0), ToInt16(nil))
}

func TestToInt32(t *testing.T) {
	equal(t, int32(10), ToInt32(Int32(10)))
	equal(t, int32(0), ToInt32(nil))
}

func TestToInt64(t *testing.T) {
	equal(t, int64(10), ToInt64(Int64(10)))
	equal(t, int64(0), ToInt64(nil))
}

func TestToUInt(t *testing.T) {
	equal(t, uint(10), ToUInt(UInt(10)))
	equal(t, uint(0), ToUInt(nil))
}

func TestToUInt8(t *testing.T) {
	equal(t, uint8(10), ToUInt8(UInt8(10)))
	equal(t, uint8(0), ToUInt8(nil))
}

func TestToUInt16(t *testing.T) {
	equal(t, uint16(10), ToUInt16(UInt16(10)))
	equal(t, uint16(0), ToUInt16(nil))
}

func TestToUInt32(t *testing.T) {
	equal(t, uint32(10), ToUInt32(UInt32(10)))
	equal(t, uint32(0), ToUInt32(nil))
}

func TestToUInt64(t *testing.T) {
	equal(t, uint64(10), ToUInt64(UInt64(10)))
	equal(t, uint64(0), ToUInt64(nil))
}

func TestToByte(t *testing.T) {
	equal(t, byte(10), ToByte(Byte(10)))
	equal(t, byte(0), ToByte(nil))
}

func TestToRune(t *testing.T) {
	equal(t, rune('🤔'), ToRune(Rune('🤔')))
	equal(t, rune(0), ToRune(nil))
}

func TestToBool(t *testing.T) {
	equal(t, bool(true), ToBool(Bool(true)))
	equal(t, bool(false), ToBool(nil))
}

func TestToString(t *testing.T) {
	equal(t, string("10"), ToString(String("10")))
	equal(t, string(""), ToString(nil))
}

func TestToFloat32(t *testing.T) {
	equal(t, float32(10.1), ToFloat32(Float32(10.1)))
	equal(t, float32(0.0), ToFloat32(nil))
}

func TestToFloat64(t *testing.T) {
	equal(t, float64(10.1), ToFloat64(Float64(10.1)))
	equal(t, float64(0.0), ToFloat64(nil))
}

func TestToComplex64(t *testing.T) {
	equal(t, complex(float32(10.0), float32(100.0)), ToComplex64(Complex64(complex(float32(10.0), float32(100.0)))))
	equal(t, complex(float32(0.0), float32(0.0)), ToComplex64(nil))
}

func TestToComplex128(t *testing.T) {
	equal(t, complex(float64(10.0), float64(100.0)), ToComplex128(Complex128(complex(float64(10.0), float64(100.0)))))
	equal(t, complex(float64(0.0), float64(0.0)), ToComplex128(nil))
}

func TestToDuration(t *testing.T) {
	equal(t, time.Second, ToDuration(Duration(time.Second)))
	equal(t, time.Duration(0), ToDuration(nil))
}
