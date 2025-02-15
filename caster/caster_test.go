package caster

import (
	"testing"
)

type MyString string
type MyTypeInt8 int8
type MyTypeUInt8 uint8
type MyTypeInt16 int16
type MyTypeInt32 int32
type MyTypeInt64 int64
type MyTypeFloat32 float32
type MyTypeFloat64 float64
type MyBool bool

type MyArrInt32 [3]MyTypeInt32
type MyArrInt8 [3]MyTypeInt8
type MyArrUInt8 [3]MyTypeUInt8
type MyArrInt64 [3]MyTypeInt64

type MySliceInt32 []MyTypeInt32
type MySliceInt8 []MyTypeInt8
type MySliceInt64 []MyTypeInt64

type MyAny any
type MyArrAny [3]MyAny
type MySliceAny []MyAny
type MyMapAny map[string]MyAny
type MyStructSub struct {
	SA MyString `nbt:"sa"`
}

type MyStruct struct {
	SB MyTypeInt32 `nbt:"sb"`
	MyStructSub
}

func TestCaster(t *testing.T) {
	{
		v := DefaultCaster(MyString("abc")).(string)
		if v != "abc" {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(MyTypeInt8(3)).(int8)
		if v != int8(3) {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(MyTypeInt8(3)).(int8)
		if v != int8(3) {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(MyArrInt8{1, 2, 3}).([]int8)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyArrInt8{1, 2, 3}).([]int8)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyArrUInt8{1, 2, 255}).([]int8)
		if v[2] != -1 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyArrInt32{1, 2, 3}).([]int32)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyArrInt32{1, 2, 3}).([]int32)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyArrAny{1, MyAny("abc"), 3}).([]any)
		if v[1] != MyAny("abc") || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MySliceInt8{1, 2, 3}).([]int8)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MySliceInt32{1, 2, 3}).([]int32)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MySliceInt32{1, 2, 3}).([]int32)
		if v[1] != 2 || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MySliceAny{1, MyAny("abc"), 3}).([]any)
		if v[1] != MyAny("abc") || len(v) != 3 {
			t.Error(v)
		}
	}
	{
		v := DefaultCaster(&MyStruct{1, MyStructSub{"abc"}}).(map[string]any)
		if v["sa"] != MyString("abc") || v["sb"] != MyTypeInt32(1) || len(v) != 2 {
			t.Error(v)
		}
	}
}
