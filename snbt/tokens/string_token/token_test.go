package string_token

import (
	"nbt/base_io/lflb"
	"nbt/base_io/lflb/sources"
	"nbt/snbt/tokens"
	whitespace "nbt/snbt/tokens/white_space"
	"testing"
)

func consumeWhiteSpaceAndComma[S lflb.Source](src S) {
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	lflb.ReadFinity(src, tokens.Specific(','))
	lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
}

func TestString(t *testing.T) {
	seq := "  abc_123.+-ABC, a_c+123.测试"
	unwarpString := &UnwrapString{}
	src := sources.NewBytesSourceFromString(seq)
	assertFail := func(b bool) {
		if b {
			t.FailNow()
		}
	}
	assertVal := func(fi interface {
		lflb.Finity
		Val() string
	}, v string) {
		fi.Reset()
		if !lflb.ReadFinity(src, fi) {
			t.FailNow()
		}
		gv := fi.Val()
		if gv != v {
			t.Logf("want: %v get: %v", v, gv)
			t.FailNow()
		}
	}

	assertFail(lflb.ReadFinity(src, unwarpString))
	consumeWhiteSpaceAndComma(src)
	assertVal(unwarpString, "abc_123.+-ABC")
	consumeWhiteSpaceAndComma(src)
	assertVal(unwarpString, "a_c+123.测试")

	seq = ` "abc()  _+:'\"\\测试" "abc()  _+:'\"\\测试`
	src = sources.NewBytesSourceFromString(seq)
	stringD := &StringD{}
	consumeWhiteSpaceAndComma(src)
	lflb.ReadFinity(src, tokens.Specific('"'))
	assertVal(stringD, `abc()  _+:'"\测试`)
	consumeWhiteSpaceAndComma(src)
	lflb.ReadFinity(src, tokens.Specific('"'))
	stringD.Reset()
	assertFail(lflb.ReadFinity(src, stringD))

	seq = ` 'abc()  _+:"\'\\测试' 'abc()  _+:"\'\\测试`
	src = sources.NewBytesSourceFromString(seq)
	stringS := &StringS{}
	consumeWhiteSpaceAndComma(src)
	lflb.ReadFinity(src, tokens.Specific('\''))
	assertVal(stringS, `abc()  _+:"'\测试`)
	consumeWhiteSpaceAndComma(src)
	lflb.ReadFinity(src, tokens.Specific('\''))
	stringS.Reset()
	assertFail(lflb.ReadFinity(src, stringS))
	if v, _ := src.This(); v != 'a' {
		t.FailNow()
	}

	seq = `'abc()  _+:"\'\\测试' "abc()  _+:'\"\\测试" abc_123.+-ABC, `
	src = sources.NewBytesSourceFromString(seq)
	stringA := &AnyString{}
	consumeWhiteSpaceAndComma(src)
	assertVal(stringA, `abc()  _+:"'\测试`)
	consumeWhiteSpaceAndComma(src)
	assertVal(stringA, `abc()  _+:'"\测试`)
	consumeWhiteSpaceAndComma(src)
	assertVal(stringA, `abc_123.+-ABC`)

	seq = `'abc()  _+:"\'\\测试 `
	src = sources.NewBytesSourceFromString(seq)
	stringA.Reset()
	assertFail(lflb.ReadFinity(src, stringA))
	if v, _ := src.This(); v != '\'' {
		t.FailNow()
	}

	seq = `"abc()  _+:'\"\\测试 `
	src = sources.NewBytesSourceFromString(seq)
	stringA.Reset()
	assertFail(lflb.ReadFinity(src, stringA))
	if v, _ := src.This(); v != '"' {
		t.FailNow()
	}

	seq = `abc_123.+-ABC`
	src = sources.NewBytesSourceFromString(seq)
	stringA.Reset()
	assertVal(stringA, `abc_123.+-ABC`)
}
