package snbt

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

var ErrCannotAcceptType = errors.New("cannot accept input type")
var ErrCannotEncodeListElem = "when encode element in list, an error occours"
var ErrCannotEncodeCompoundValue = "when encode compound value, an error occours"

func EncodeTo(w io.Writer, input any, caster func(any) any) (err error) {
	writer := newWriterWithBuffer(w)
	err = encodeTo(writer, input, caster, false)
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
			v := data[k]
			writeString(w, k)
			w.WriteByte(':')
			w.WriteByte(' ')
			err = encodeTo(w, v, caster, false)
			if err != nil {
				return fmt.Errorf("%v: %v", ErrCannotEncodeCompoundValue, err)
			}
		}
		w.WriteByte('}')
		return nil
	default:
		return encodeToWithCast(w, input, caster, casted)
	}
}

func takeStructMember(val reflect.Value, out map[string]any) {
	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Type().Field(i)
		fieldValue := val.Field(i)
		tag := fieldType.Tag.Get("nbt")
		if fieldType.PkgPath != "" || tag == "-" {
			continue
		}
		if fieldType.Anonymous {
			// The field was anonymous, so we write that in the same compound tag as this one.
			takeStructMember(fieldValue, out)
			continue
		}
		tagName := fieldType.Name
		if strings.HasSuffix(tag, ",omitempty") {
			tag = strings.TrimSuffix(tag, ",omitempty")
			if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
				// The tag had the ',omitempty' tag, meaning it should be omitted if it has the zero
				// value. If this is reached, that was the case, and we skip it.
				continue
			}
		}
		if tag != "" {
			tagName = tag
		}
		out[tagName] = fieldValue.Interface()
	}
}

func encodeToWithCast(w *writerWithBuffer, orig any, caster func(any) any, casted bool) (err error) {
	// we don't want caster allocate anything like []any{} or map[string]any{}
	val := reflect.ValueOf(orig)
	kind := val.Kind()
	switch kind {
	case reflect.Array:
		switch val.Type().Elem().Kind() {
		case reflect.Int8:
			w.WriteString("[B; ")
			n := val.Cap()
			for i := 0; i < n; i++ {
				if i != 0 {
					w.WriteByte(',')
					w.WriteByte(' ')
				}
				writeInt(w, int8(val.Index(i).Int()))
			}
			w.WriteByte(']')
			return nil
		case reflect.Int32:
			w.WriteString("[I; ")
			n := val.Cap()
			for i := 0; i < n; i++ {
				if i != 0 {
					w.WriteByte(',')
					w.WriteByte(' ')
				}
				writeInt(w, int32(val.Index(i).Int()))
			}
			w.WriteByte(']')
			return nil
		case reflect.Int64:
			w.WriteString("[L; ")
			n := val.Cap()
			for i := 0; i < n; i++ {
				if i != 0 {
					w.WriteByte(',')
					w.WriteByte(' ')
				}
				writeInt(w, int64(val.Index(i).Int()))
			}
			w.WriteByte(']')
			return nil
		default:
			n := val.Cap()
			if n > 0 {
				decided := 0
				if caster != nil {
					new := caster(val.Index(0).Interface())
					if v, ok := new.(int8); ok {
						decided = 1
						w.WriteString("[B; ")
						writeInt(w, v)
					} else if v, ok := new.(int32); ok {
						decided = 2
						w.WriteString("[I; ")
						writeInt(w, v)
					} else if v, ok := new.(int64); ok {
						decided = 3
						w.WriteString("[L; ")
						writeInt(w, v)
					} else {
						w.WriteString("[ ")
						err = encodeTo(w, new, caster, false)
						if err != nil {
							return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
						}
					}
				}
				for i := 1; i < n; i++ {
					w.WriteByte(',')
					w.WriteByte(' ')
					v := caster(val.Index(i).Interface())
					switch decided {
					case 0:
						err = encodeTo(w, v, caster, false)
						if err != nil {
							return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
						}
					case 1:
						writeInt(w, v.(int8))
					case 2:
						writeInt(w, v.(int32))
					case 3:
						writeInt(w, v.(int64))
					}
				}
				w.WriteByte(']')
			} else {
				w.WriteByte('[')
				w.WriteByte(']')
			}

			return nil
		}
	case reflect.Slice:
		switch val.Type().Elem().Kind() {
		case reflect.Int32:
			ret := unsafe.Slice((*int32)(unsafe.Pointer(val.Pointer())), val.Len())
			return encodeTo(w, ret, caster, false)
		case reflect.Int8:
			ret := unsafe.Slice((*int8)(unsafe.Pointer(val.Pointer())), val.Len())
			return encodeTo(w, ret, caster, false)
		case reflect.Int64:
			ret := unsafe.Slice((*int64)(unsafe.Pointer(val.Pointer())), val.Len())
			return encodeTo(w, ret, caster, false)
		default:
			w.WriteString("[ ")
			for i := 0; i < val.Len(); i++ {
				if i != 0 {
					w.WriteByte(',')
					w.WriteByte(' ')
				}
				err = encodeTo(w, val.Index(i).Interface(), caster, false)
				if err != nil {
					return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
				}
			}
			w.WriteByte(']')
			return nil
		}
	case reflect.Struct:
		out := map[string]any{}
		takeStructMember(val, out)
		return encodeToWithCast(w, out, caster, false)
	case reflect.Map:
		stringK := true
		if val.Type().Key().Kind() != reflect.String {
			stringK = false
		}
		iter := val.MapRange()
		unsortedMap := map[string]any{}
		// take data
		for iter.Next() {
			key := ""
			if stringK {
				key = iter.Key().String()
			} else {
				key = fmt.Sprintf("%v", iter.Key().Interface())
			}
			unsortedMap[key] = iter.Value().Interface()
		}
		return encodeTo(w, unsortedMap, caster, false)
	default:
		if !casted && caster != nil {
			new := caster(orig)
			if new != nil {
				return encodeTo(w, new, caster, true)
			}
		}
		return ErrCannotAcceptType
	}
}

type writerWithBuffer struct {
	*bufio.Writer
	buf []byte
}

func (w *writerWithBuffer) printf(f string, data any) {
	fmt.Fprintf(w.Writer, f, data)
}

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
