package ptr

import (
	"reflect"
	"testing"
)

func equal(t *testing.T, expected, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %#v, actual %#v", expected, actual)
	}
}

func TestInt(t *testing.T) {
	equal(t, int(10), *Int(10))
}

func TestInt8(t *testing.T) {
	equal(t, int8(10), *Int8(10))
}

func TestInt16(t *testing.T) {
	equal(t, int16(10), *Int16(10))
}

func TestInt32(t *testing.T) {
	equal(t, int32(10), *Int32(10))
}

func TestInt64(t *testing.T) {
	equal(t, int64(10), *Int64(10))
}

func TestUInt(t *testing.T) {
	equal(t, uint(10), *UInt(10))
}

func TestUInt8(t *testing.T) {
	equal(t, uint8(10), *UInt8(10))
}

func TestUInt16(t *testing.T) {
	equal(t, uint16(10), *UInt16(10))
}

func TestUInt32(t *testing.T) {
	equal(t, uint32(10), *UInt32(10))
}

func TestUInt64(t *testing.T) {
	equal(t, uint64(10), *UInt64(10))
}

func TestByte(t *testing.T) {
	equal(t, byte(10), *Byte(10))
}

func TestRune(t *testing.T) {
	equal(t, rune('🤔'), *Rune('🤔'))
}

func TestBool(t *testing.T) {
	equal(t, bool(true), *Bool(true))
}

func TestString(t *testing.T) {
	equal(t, string("10"), *String("10"))
}

func TestFloat32(t *testing.T) {
	equal(t, float32(10.1), *Float32(10.1))
}

func TestFloat64(t *testing.T) {
	equal(t, float64(10.1), *Float64(10.1))
}

func TestComplex64(t *testing.T) {
	equal(t, complex(float32(10.0), float32(100.0)), *Complex64(complex(float32(10.0), float32(100.0))))
}

func TestComplex128(t *testing.T) {
	equal(t, complex(float64(10.0), float64(100.0)), *Complex128(complex(float64(10.0), float64(100.0))))
}
