package nbt

import (
	"bufio"
	"bytes"
	"io"
	"nbt/base_io"
	"nbt/base_io/lflb"
	"nbt/base_io/lflb/sources"
	"nbt/caster"
	"nbt/encoding"
	"nbt/nbt"
	"nbt/snbt"
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

func NBTDecode[E encoding.Encoding](data []byte, caster func(any) any) (tagName string, value any, err error) {
	return nbt.DecodeTagAndValue[E](bytes.NewReader(data))
}

func NBTDecodeFrom[E encoding.Encoding](r io.Reader, caster func(any) any) (tagName string, value any, err error) {
	return nbt.DecodeTagAndValue[E](base_io.NewReader(r))
}

func NBTEncodeTo[E encoding.Encoding](w io.Writer, tag string, value any, caster func(any) any) (err error) {
	buf := bufio.NewWriter(w)
	err = nbt.EncodeTagAndValueTo[E](buf, tag, value, caster)
	buf.Flush()
	return err
}

func NBTEncode[E encoding.Encoding](tag string, value any, caster func(any) any) (out []byte, err error) {
	buf := bytes.NewBuffer(nil)
	err = nbt.EncodeTagAndValueTo[E](buf, tag, value, caster)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
