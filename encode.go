package snbt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"unsafe"
)

func EncodeTo(w io.Writer, input any) (err error) {
	writer := newWriterWithBuffer(w)
	err = encodeTo(writer, input, nil, false)
	if err != nil {
		return err
	}
	return writer.close()
}

func Encode(input any, caster func(any) any) (out []byte, err error) {
	buf := bytes.NewBuffer([]byte{})
	writer := newWriterWithBuffer(buf)
	err = encodeTo(writer, input, caster, false)
	if err != nil {
		return nil, err
	}
	err = writer.close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

var ErrCannotAcceptType = errors.New("cannot accept input type")
var ErrCannotEncodeListElem = "when encode element in list, an error occours"
var ErrCannotEncodeCompoundValue = "when encode compound value, an error occours"

func encodeTo(w *writerWithBuffer, input any, caster func(any) any, casted bool) (err error) {
	switch data := input.(type) {
	case string:
		writeString(w, data)
		return nil
	case int32:
		writeInt(w, data)
		return nil
	case int8:
		writeInt(w, data)
		w.WriteByte('b')
		return nil
	case int16:
		writeInt(w, data)
		w.WriteByte('s')
		return nil
	case int64:
		writeInt(w, data)
		w.WriteByte('l')
		return nil
	case float32:
		w.WriteString(strconv.FormatFloat(float64(data), 'f', -1, 32))
		w.WriteByte('f')
		return nil
	case float64:
		w.WriteString(strconv.FormatFloat(float64(data), 'f', -1, 64))
		w.WriteByte('d')
		return nil
	case []int32:
		w.WriteString("[I; ")
		for i, v := range data {
			if i != 0 {
				w.WriteByte(',')
				w.WriteByte(' ')
			}
			writeInt(w, v)
		}
		w.WriteByte(']')
		return nil
	case []int8:
		w.WriteString("[B; ")
		for i, v := range data {
			if i != 0 {
				w.WriteByte(',')
				w.WriteByte(' ')
			}
			writeInt(w, v)
		}
		w.WriteByte(']')
		return nil
	case []int64:
		w.WriteString("[L; ")
		for i, v := range data {
			if i != 0 {
				w.WriteByte(',')
				w.WriteByte(' ')
			}
			writeInt(w, v)
		}
		w.WriteByte(']')
		return nil
	case []any:
		w.WriteString("[ ")
		for i, v := range data {
			if i != 0 {
				w.WriteByte(',')
				w.WriteByte(' ')
			}
			err = encodeTo(w, v, caster, false)
			if err != nil {
				return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
			}
		}
		w.WriteByte(']')
		return nil
	case map[string]any:
		w.WriteString("{")
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if i != 0 {
				w.WriteByte(',')
				w.WriteByte(' ')
			}
			writeString(w, k)
			w.WriteByte(':')
			w.WriteByte(' ')
			err = encodeTo(w, data[k], caster, false)
			if err != nil {
				return fmt.Errorf("%v: %v", ErrCannotEncodeCompoundValue, err)
			}
		}
		w.WriteByte('}')
		return nil
	default:
		if !casted && caster != nil {
			input = caster(input)
			if input != nil {
				return encodeTo(w, input, caster, true)
			}
		}
		return ErrCannotAcceptType
	}
}

type writerWithBuffer struct {
	*bufio.Writer
	buf []byte
}

// func (w *writerWithBuffer) printf(f string, data any) {
// 	fmt.Fprintf(w.Writer, f, data)
// }

func writeString(w *writerWithBuffer, data string) {
	w.WriteByte('"')
	bs := unsafe.Slice(unsafe.StringData(data), len(data))
	for _, b := range bs {
		if b == '\\' || b == '"' {
			w.WriteByte('\\')
		}
		w.WriteByte(b)
	}
	w.WriteByte('"')
}

func writeInt[T interface{ int8 | int16 | int32 | int64 }](w *writerWithBuffer, i T) {
	if i < 0 {
		w.WriteByte('-')
		i = -i
	}
	off := 0
	for i > 10 {
		w.buf[off] = uint8(i%10) + '0'
		off += 1
		i = i / 10
	}
	w.WriteByte(uint8(i) + '0')
	for off != 0 {
		w.WriteByte(w.buf[off-1])
		off -= 1
	}
}

func (w *writerWithBuffer) close() (err error) {
	return w.Writer.Flush()
}

func newWriterWithBuffer(w io.Writer) *writerWithBuffer {
	writer, ok := w.(*writerWithBuffer)
	if ok {
		return writer
	}
	wbuf, ok := w.(*bufio.Writer)
	if ok {
		return &writerWithBuffer{
			Writer: wbuf,
			buf:    make([]byte, 32),
		}
	}
	return &writerWithBuffer{
		Writer: bufio.NewWriter(w),
		buf:    make([]byte, 32),
	}
}
