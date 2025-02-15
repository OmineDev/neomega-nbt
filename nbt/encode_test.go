package nbt

import (
	"bytes"
	"encoding/hex"
	"nbt/encoding"
	"strings"
	"testing"
)

func TestEncode(t *testing.T) {
	assertOut := func(tag string, v any, expect string) {
		expect = strings.ReplaceAll(strings.TrimSpace(expect), " ", "")
		w := bytes.NewBuffer(nil)
		EncodeTagAndValueTo[encoding.BigEndian](w, tag, v, nil)
		out := w.Bytes()
		expBytes, _ := hex.DecodeString(expect)
		if !bytes.Equal(out, expBytes) {
			t.Errorf("%v %x!=%x", v, out, expBytes)
		}
	}
	assertOut("ex", "test", "08 00 02 65 78 00 04 74 65 73 74")
	assertOut("ex", int32(8), "03 00 02 65 78 00 00 00 08")
	assertOut("ex", int8(127), "01 00 02 65 78 7f")
	assertOut("ex", int64(8), "04 00 02 65 78 00 00 00 00 00 00 00 08")
	assertOut("ex", float32(8), "05 00 02 65 78 41 00 00 00")
	assertOut("ex", float64(8), "06 00 02 65 78 40 20 00 00 00 00 00 00")
	assertOut("ex", []int32{1, 2, 3}, "0b 00 02 65 78 00 00 00 03 00 00 00 01 00 00 00 02 00 00 00 03")
	assertOut("ex", []int8{1, 9, 3}, "07 00 02 65 78 00 00 00 03 01 09 03")
	assertOut("ex", []int64{1, 2, 3}, "0c 00 02 65 78 00 00 00 03 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00 03")
	assertOut("ex", []any{
		int8(1), int8(1), int8(1),
	}, "09 00 02 65 78 01 00 00 00 03 01 01 01")
	assertOut("ex", map[string]any{
		"a1": int32(1),
		"a2": int8(2),
	}, "0a 00 02 65 78 03 00 02 61 31 00 00 00 01 01 00 02 61 32 02 00")
}
