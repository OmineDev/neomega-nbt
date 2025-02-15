package nbt

import (
	"bytes"
	"encoding/hex"
	"neomega_nbt/caster"
	"neomega_nbt/encoding"
	"strings"
	"testing"
)

type MyString string
type MyTypeInt32 int32
type MyStructSub struct {
	SA MyString `nbt:"sa"`
}

type MyStruct struct {
	SB MyTypeInt32 `nbt:"sb"`
	SD MyStructSub
	SC uint8
}

func TestDecode(t *testing.T) {
	assertInOut := func(val string) {
		val = strings.ReplaceAll(strings.TrimSpace(val), " ", "")
		expBytes, _ := hex.DecodeString(val)
		if len(expBytes) == 0 {
			panic("invalid test input")
		}
		r := bytes.NewReader(expBytes)
		tagName, value, err := DecodeTagAndValue[encoding.BigEndian](r)
		if err != nil {
			t.Error(err)
		}
		w := bytes.NewBuffer(nil)
		EncodeTagAndValueTo[encoding.BigEndian](w, tagName, value, nil)
		out := w.Bytes()
		if !bytes.Equal(out, expBytes) {
			t.Errorf("%v %x!=%x", val, out, expBytes)
		}
	}
	assertInOut("08 00 02 65 78 00 04 74 65 73 74")
	assertInOut("03 00 02 65 78 00 00 00 08")
	assertInOut("01 00 02 65 78 7f")
	assertInOut("04 00 02 65 78 00 00 00 00 00 00 00 08")
	assertInOut("05 00 02 65 78 41 00 00 00")
	assertInOut("06 00 02 65 78 40 20 00 00 00 00 00 00")
	assertInOut("0b 00 02 65 78 00 00 00 03 00 00 00 01 00 00 00 02 00 00 00 03")
	assertInOut("07 00 02 65 78 00 00 00 03 01 09 03")
	assertInOut("0c 00 02 65 78 00 00 00 03 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00 03")
	assertInOut("09 00 02 65 78 01 00 00 00 03 01 01 01")
	assertInOut("09 00 02 65 78 00")
	assertInOut("0a 00 02 65 78 03 00 02 61 31 00 00 00 01 01 00 02 61 32 02 00")

	{
		my := MyStruct{
			SB: 1,
			SD: MyStructSub{"abc"},
			SC: 2,
		}
		w := bytes.NewBuffer(nil)
		EncodeTagAndValueTo[encoding.BigEndian](w, "ex", my, caster.DefaultCaster)
		out := w.Bytes()

		r := bytes.NewReader(out)
		tagName, value, err := DecodeTagAndValue[encoding.BigEndian](r)
		if err != nil {
			t.Error(err)
		}
		var back MyStruct
		err = caster.MapNBT(value, &back)
		if err != nil {
			t.Error(err)
		}
		if tagName != "ex" {
			t.Error("tag name error")
		}
		if back.SD.SA != "abc" {
			t.Error("back.SD.SA")
		}
		if back != my {
			t.Error("nbt name error")
		}
	}

}
