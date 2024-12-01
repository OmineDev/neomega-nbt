package int_arr

import (
	"snbt/lflb"
	"snbt/lflb/lflbops"
	"snbt/lflb/sources"
	whitespace "snbt/tokens/white_space"
	"testing"
)

func consumeWhiteSpaceAndComma[S lflb.Source](src S) {
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	lflb.ReadFinity(src, lflbops.Specific(','))
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
}

func assertArr[S lflb.Source, I interface{ int8 | int32 | int64 }](
	src S, fi interface {
		lflb.Finity
		Val() []I
	}, val []I, t *testing.T) {
	fi.Reset()
	if !lflb.ReadFinity(src, fi) {
		t.Logf("fail to read: %v\n", val)
		t.FailNow()
	}
	getV := fi.Val()
	if len(getV) != len(val) {
		t.Logf("length mismatch: %v!=%v", len(getV), len(val))
		t.FailNow()
	}
	for i, v := range getV {
		if v != val[i] {
			t.Logf("value mismatch: %v!=%v", v, val[i])
			t.FailNow()
		}
	}
}

func TestString(t *testing.T) {
	seq := "12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   ] "
	src := sources.NewBytesSourceFromString(seq)
	assertFail := func(b bool) {
		if b {
			t.FailNow()
		}
	}
	assertFail(false)
	assertArr(src, &IntArray[int32]{}, []int32{12, -34, -567, 89, 10, 11, -12, -13, 14, 15}, t)

	seq = "12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15   "
	src = sources.NewBytesSourceFromString(seq)
	assertFail(lflb.ReadFinity(src, &IntArray[int32]{}))
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "a,12, -34 ,-567, 89,10,11 , -12 ,-13 ,14 ,15]   "
	src = sources.NewBytesSourceFromString(seq)
	assertFail(lflb.ReadFinity(src, &IntArray[int32]{}))
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = "12, --34] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail(lflb.ReadFinity(src, &IntArray[int32]{}))
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "12,,3] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail(lflb.ReadFinity(src, &IntArray[int32]{}))
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}

	seq = "12,3,] "
	src = sources.NewBytesSourceFromString(seq)
	assertFail(lflb.ReadFinity(src, &IntArray[int32]{}))
	if v, _ := src.This(); v != '1' {
		t.FailNow()
	}
}
