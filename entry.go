package neomega_nbt

import (
	"bufio"
	"bytes"
	"io"
	"neomega_nbt/base_io"
	"neomega_nbt/base_io/lflb"
	"neomega_nbt/base_io/lflb/sources"
	"neomega_nbt/caster"
	"neomega_nbt/encoding"
	"neomega_nbt/nbt"
	"neomega_nbt/snbt"
)

var Caster = caster.DefaultCaster

func SNBTDecode(inStr string) (v any, err error) {
	return snbt.DecodeFrom(sources.NewBytesSourceFromString(inStr))
}

func SNBTDecodeFrom(source lflb.Source) (v any, err error) {
	return snbt.DecodeFrom(source)
}

func SNBTEncodeTo(w io.Writer, input any, caster func(any) any) (err error) {
	return snbt.EncodeTo(w, input, caster)
}

func SNBTEncode(input any, caster func(any) any) (out []byte, err error) {
	buf := bytes.NewBuffer(nil)
	err = snbt.EncodeTo(buf, input, caster)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func NBTDecode[E encoding.ReadEncoding](data []byte) (tagName string, value any, err error) {
	return nbt.DecodeTagAndValue[E](bytes.NewReader(data))
}

func NBTDecodeFrom[E encoding.ReadEncoding](r io.Reader) (tagName string, value any, err error) {
	return nbt.DecodeTagAndValue[E](base_io.NewReader(r))
}

func NBTEncodeTo[E encoding.WriteEncoding](w io.Writer, tag string, value any, caster func(any) any) (err error) {
	buf := bufio.NewWriter(w)
	err = nbt.EncodeTagAndValueTo[E](buf, tag, value, caster)
	buf.Flush()
	return err
}

func NBTEncode[E encoding.WriteEncoding](tag string, value any, caster func(any) any) (out []byte, err error) {
	buf := bytes.NewBuffer(nil)
	err = nbt.EncodeTagAndValueTo[E](buf, tag, value, caster)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
