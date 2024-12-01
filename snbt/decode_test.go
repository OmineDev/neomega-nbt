package snbt

import (
	"fa/lflb"
	"fa/lflb/lflbops"
	"fa/lflb/sources"
	whitespace "fa/snbt/tokens/white_space"
	"math"
	"reflect"
	"testing"
)

func consumeWhiteSpaceAndComma[S lflb.Source](src S) {
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	lflb.ReadFinity(src, lflbops.Specific(','))
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
}

func assertArr[S lflb.Source, I any](
	src S, val []I, t *testing.T) {
	getV, err := DecodeFrom(src)
	if err != nil {
		t.Logf("fail to decode: %v\n", src)
		t.FailNow()
	}

	if reflect.ValueOf(getV).Len() != len(val) {
		t.Logf("length mismatch: %v!=%v", reflect.ValueOf(getV).Len(), len(val))
		t.FailNow()
	}
	v := reflect.ValueOf(getV)
	for i := 0; i < v.Len(); i++ {
		getV := v.Index(i).Interface()
		wantV := val[i]
		if reflect.TypeOf(wantV).Kind() != reflect.TypeOf(getV).Kind() {
			t.Logf("type mismatch: %v!=%v", wantV, getV)
			t.FailNow()
		}
		if !reflect.DeepEqual(wantV, getV) {
			t.Logf("value mismatch: %v!=%v", wantV, getV)
			t.FailNow()
		}
	}
}

func TestSnbtDecode(t *testing.T) {
	seq := " \t\n\v\f\r  "
	src := sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != ErrNotSNBT {
		t.Logf("get v: %v\n", v)
		t.FailNow()
	}

	seq = ""
	src = sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != ErrNoData || v != nil {
		t.Logf("get v: %v\n", v)
		t.FailNow()
	}

	seq = "abc"
	src = sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != nil || v != "abc" {
		t.Logf("get v: %v\n", v)
		t.FailNow()
	}

	// should be number, not string
	seq = "-123"
	src = sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != nil || v != int32(-123) {
		t.Logf("get v: %v\n", v)
		t.FailNow()
	}

	seq = "123_abc"
	src = sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != nil || v != "123_abc" {
		t.Logf("get v: %v\n", v)
		t.FailNow()
	}

	// number tests
	seq = "-123,123,-1,1,-1b,1b,103b,-104b,-1s,1s,103s,-104s,-1l,1l,103l,-104l,-56.78,3.2,-3.,-.2,-234.,-.456,123E-2,-.456E1,3.E1,45.58E-12,123.456E-8f,123.456E10d,true,false,3b,end"
	src = sources.NewBytesSourceFromString(seq)

	assertNumber := func(val any, T byte) {
		number, err := DecodeFrom(src)
		if err != nil {
			t.Logf("want: %v, get err: %v\n", val, err)
			t.FailNow()
		}
		switch T {
		case 'I':
			getV, ok1 := number.(int32)
			wantV, ok2 := val.(int32)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'B':
			getV, ok1 := number.(int8)
			wantV, ok2 := val.(int8)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}

		case 'S':
			getV, ok1 := number.(int16)
			wantV, ok2 := val.(int16)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}

		case 'L':
			getV, ok1 := number.(int64)
			wantV, ok2 := val.(int64)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'F':
			getV, ok1 := number.(float32)
			wantV, ok2 := val.(float32)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if math.Abs(float64(getV-wantV)) > 0.0001 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'D':
			getV, ok1 := number.(float64)
			wantV, ok2 := val.(float64)
			if !ok1 || !ok2 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
			if math.Abs(float64(getV-wantV)) > 1e-14 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		}
	}
	assertNumber(int32(-123), 'I')
	readNext := func() {
		lflb.ReadFinity(src, lflbops.FinityAny{})
	}
	readNext()
	assertNumber(int32(123), 'I')
	readNext()
	assertNumber(int32(-1), 'I')
	readNext()
	assertNumber(int32(1), 'I')

	readNext()
	assertNumber(int8(-1), 'B')
	readNext()
	assertNumber(int8(1), 'B')
	readNext()
	assertNumber(int8(103), 'B')
	readNext()
	assertNumber(int8(-104), 'B')
	readNext()
	assertNumber(int16(-1), 'S')
	readNext()
	assertNumber(int16(1), 'S')
	readNext()
	assertNumber(int16(103), 'S')
	readNext()
	assertNumber(int16(-104), 'S')
	readNext()
	assertNumber(int64(-1), 'L')
	readNext()
	assertNumber(int64(1), 'L')
	readNext()
	assertNumber(int64(103), 'L')
	readNext()
	assertNumber(int64(-104), 'L')
	readNext()
	assertNumber(float32(-56.78), 'F')
	readNext()
	assertNumber(float32(3.2), 'F')
	readNext()
	assertNumber(float32(-3.0), 'F')
	readNext()
	assertNumber(float32(-0.2), 'F')
	readNext()
	assertNumber(float32(-234.0), 'F')
	readNext()
	assertNumber(float32(-0.456), 'F')
	readNext()
	assertNumber(float32(1.23), 'F')
	readNext()
	assertNumber(float32(-4.56), 'F')
	readNext()
	assertNumber(float32(30), 'F')
	readNext()
	assertNumber(float32(45.58e-12), 'F')
	readNext()
	assertNumber(float32(123.456e-8), 'F')
	readNext()
	assertNumber(float64(123.456e10), 'D')
	readNext()
	assertNumber(int8(1), 'B')
	readNext()
	assertNumber(int8(0), 'B')
	readNext()
	assertNumber(int8(3), 'B')
	readNext()
	if v, _ := src.This(); v != 'e' {
		t.FailNow()
	}

	seq = `"abc()  _+:'\"\\测试" "abc()  _+:'\"\\测试`

	assertVal := func(val any) {
		if v, err := DecodeFrom(src); err != nil || v != val {
			t.Logf("want: %v, get v: %v\n", val, v)
			t.FailNow()
		}
	}
	assertFail := func() {
		_, err := DecodeFrom(src)
		if err == nil {
			t.FailNow()
		}
	}

	src = sources.NewBytesSourceFromString(seq)
	assertVal(`abc()  _+:'"\测试`)
	consumeWhiteSpaceAndComma(src)
	lflb.ReadFinity(src, lflbops.Specific('\''))
	assertFail()
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = `'abc()  _+:"\'\\测试' "abc()  _+:'\"\\测试" abc_123.+-ABC, `
	src = sources.NewBytesSourceFromString(seq)
	assertVal(`abc()  _+:"'\测试`)
	consumeWhiteSpaceAndComma(src)
	assertVal(`abc()  _+:'"\测试`)
	consumeWhiteSpaceAndComma(src)
	assertVal(`abc_123.+-ABC`)

	seq = `'abc()  _+:"\'\\测试 `
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = `"abc()  _+:'\"\\测试 `
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = `abc_123.+-ABC`
	src = sources.NewBytesSourceFromString(seq)
	assertVal(`abc_123.+-ABC`)

	seq = "[I; 12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   ] "
	src = sources.NewBytesSourceFromString(seq)
	assertArr(src, []int32{12, -34, -567, 89, 10, 11, -12, -13, 14, 15}, t)

	seq = "[B; 12, -34 ,-57, 89,10,11 , -12 ,-13 ,14 ,15   ] "
	src = sources.NewBytesSourceFromString(seq)
	assertArr(src, []int8{12, -34, -57, 89, 10, 11, -12, -13, 14, 15}, t)

	seq = "[L; 12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   ] "
	src = sources.NewBytesSourceFromString(seq)
	assertArr(src, []int64{12, -34, -567, 89, 10, 11, -12, -13, 14, 15}, t)
	seq = "[I;12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   "
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}
	seq = "[I;a,12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15]   "
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = "[I;12, --34] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "[I;12,,3] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "[I;12,3,] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail()
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "[ 12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   ] "
	src = sources.NewBytesSourceFromString(seq)
	assertArr(src, []int32{12, -34, -567, 89, 10, 11, -12, -13, 14, 15}, t)

	// not allowed in snbt, but we allow this
	seq = "[ 12, 123_abc, -12.34E3f ,'bcd',[a,b,bc],{a:1,b:1,c:2}] "
	src = sources.NewBytesSourceFromString(seq)
	assertArr(src, []any{int32(12), "123_abc", float32(-12.34e3), "bcd", []any{"a", "b", "bc"}, map[string]any{
		"a": int32(1), "b": int32(1), "c": int32(2),
	}}, t)

	seq = "{123_abc:{a:1,b:1,c:2},\"嗨 我的世界wiki\": 123, '@{}<>!-=()[]*&^%$#/+`~.\";': 123}"
	src = sources.NewBytesSourceFromString(seq)
	if v, err := DecodeFrom(src); err != nil {
		t.Logf("get err: %v\n", err)
		t.FailNow()
	} else {
		if !reflect.DeepEqual(map[string]any{
			"嗨 我的世界wiki":                 int32(123),
			"@{}<>!-=()[]*&^%$#/+`~.\";": int32(123),
			"123_abc": map[string]any{
				"a": int32(1), "b": int32(1), "c": int32(2),
			},
		}, v) {
			t.FailNow()
		}
	}
}
