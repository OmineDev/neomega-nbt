package nbt

import (
	"fmt"
	"neomega_nbt/base_io"
	"neomega_nbt/encoding"
	"sort"
)

type ErrCannotCastType struct {
	Data any
}

func (e ErrCannotCastType) Error() string {
	return fmt.Sprintf("cannot handle value: %v of type: %T", e.Data, e.Data)
}

type ErrListTypeInconsist struct {
	Expect tagType
	Data   any
}

func (e ErrListTypeInconsist) Error() string {
	return fmt.Sprintf("value: %v of type: %T is not %v", e.Data, e.Data, e.Expect)
}

// type encodingInspector struct {
// 	reciver [4]byte // 4 is enough for encoding "":0b
// 	off     int
// }

// func (i *encodingInspector) Write(b []byte) (n int, err error) {
// 	for p, v := range b {
// 		i.reciver[i.off+p] = v
// 		i.off += 1
// 	}
// 	return len(b), nil
// }

// func (i *encodingInspector) WriteByte(b byte) (err error) {
// 	i.reciver[i.off] = b
// 	i.off += 1
// 	return nil
// }

// type trapList1Writor[W base_io.Writer] struct {
// 	onWriteLen func()
// 	underlay   W
// 	tagNameLen int
// 	off        int
// }

// func (w *trapList1Writor[W]) Write(b []byte) (n int, err error) {
// 	if w.off > w.tagNameLen {
// 		return w.underlay.Write(b)
// 	}
// 	if len(b) == 0 {
// 		return
// 	}
// 	if w.off == 0 {
// 		w.off++
// 		w.underlay.WriteByte(b[0])
// 		w.onWriteLen()
// 	}
// 	for _, v := range b[1:] {
// 		if w.off > w.tagNameLen {
// 			w.underlay.WriteByte(v)
// 		}
// 		w.off += 1
// 	}
// 	return len(b), nil
// }

// func (w *trapList1Writor[W]) WriteByte(b byte) (err error) {
// 	if w.off > w.tagNameLen {
// 		return w.underlay.WriteByte(b)
// 	}
// 	if w.off == 0 {
// 		w.off++
// 		w.underlay.WriteByte(b)
// 		w.onWriteLen()
// 		return
// 	} else {
// 		w.off += 1
// 	}
// 	return nil
// }

func EncodeTagAndValueTo[E encoding.WriteEncoding, W base_io.Writer](
	w W,
	tagName string,
	value any,
	caster func(any) any,
) (err error) {
	valueType, cleanedValue := cleanValue(value, caster, false)
	if valueType == 0 {
		return ErrCannotCastType{value}
	}
	return EncodeCleanedTagAndValueTo[E, W](w, valueType, tagName, cleanedValue, caster)
}

func EncodeCleanedTagAndValueTo[E encoding.WriteEncoding, W base_io.Writer](
	w W,
	valueType tagType,
	tagName string,
	cleanedValue any,
	caster func(any) any,
) (err error) {
	var e E
	w.WriteByte(byte(valueType))
	e.WriteString(w, tagName)
	return EncodeCleanedValueTo[E, W](w, valueType, cleanedValue, caster)
}

func EncodeCleanedValueTo[E encoding.WriteEncoding, W base_io.Writer](
	w W,
	valueType tagType,
	cleanedValue any,
	caster func(any) any,
) (err error) {
	var e E
	switch valueType {
	case tagString:
		e.WriteString(w, cleanedValue.(string))
		return nil
	case tagInt:
		e.WriteInt32(w, cleanedValue.(int32))
		return nil
	case tagByte:
		w.WriteByte(byte(cleanedValue.(int8)))
		return nil
	case tagShort:
		e.WriteInt16(w, cleanedValue.(int16))
		return nil
	case tagLong:
		e.WriteInt64(w, cleanedValue.(int64))
		return nil
	case tagFloat:
		e.WriteFloat32(w, cleanedValue.(float32))
		return nil
	case tagDouble:
		e.WriteFloat64(w, cleanedValue.(float64))
		return nil
	case tagIntArray:
		data := cleanedValue.([]int32)
		e.WriteInt32(w, int32(len(data)))
		for _, v := range data {
			e.WriteInt32(w, int32(v))
		}
		return nil
	case tagByteArray:
		data := cleanedValue.([]int8)
		e.WriteInt32(w, int32(len(data)))
		for _, v := range data {
			w.WriteByte(byte(v))
		}
		return nil
	case tagLongArray:
		data := cleanedValue.([]int64)
		e.WriteInt32(w, int32(len(data)))
		for _, v := range data {
			e.WriteInt64(w, (v))
		}
		return nil
	case tagCompound:
		data := cleanedValue.(map[string]any)
		keys := make([]string, 0, len(data))
		for k := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			v := data[k]
			EncodeTagAndValueTo[E, W](w, k, v, caster)
		}
		w.WriteByte(byte(tagEnd))
		return nil
	case tagList:
		data := cleanedValue.([]any)
		// special case: 0 length
		if len(data) == 0 {
			w.WriteByte(byte(tagEnd))
			return nil
		}

		firstTagType, firstCleanedValue := cleanValue(data[0], caster, false)
		if firstTagType == 0 {
			return ErrCannotCastType{data[0]}
		}
		w.WriteByte(byte(firstTagType))
		e.WriteInt32(w, int32(len(data)))
		err = EncodeCleanedValueTo[E, W](w, firstTagType, firstCleanedValue, caster)
		if err != nil {
			return err
		}
		for _, rawV := range data[1:] {
			tT, cV := cleanValue(rawV, caster, false)
			if tT != firstTagType {
				return ErrListTypeInconsist{firstTagType, rawV}
			}
			err = EncodeCleanedValueTo[E, W](w, tT, cV, caster)
			if err != nil {
				return err
			}
		}
		return nil
	default:
		panic(fmt.Errorf("should not happen: %v, %T", cleanedValue, cleanedValue))
	}
}

// func encodeTo[E encoding.Encoding, W base_io.Writer](
// 	w W,
// 	tagName string,
// 	value any,
// 	caster func(any) any,
// 	casted bool,
// ) (err error) {
// 	var e E
// 	switch data := value.(type) {
// 	case string:
// 		w.WriteByte(byte(tagString))
// 		e.WriteString(w, tagName)
// 		e.WriteString(w, data)
// 		return nil
// 	case int32:
// 		w.WriteByte(byte(tagInt))
// 		e.WriteString(w, tagName)
// 		e.WriteInt32(w, data)
// 		return nil
// 	case int8:
// 		w.WriteByte(byte(tagByte))
// 		e.WriteString(w, tagName)
// 		w.WriteByte(byte(data))
// 		return nil
// 	case int16:
// 		w.WriteByte(byte(tagShort))
// 		e.WriteString(w, tagName)
// 		e.WriteInt16(w, data)
// 		return nil
// 	case int64:
// 		w.WriteByte(byte(tagLong))
// 		e.WriteString(w, tagName)
// 		e.WriteInt64(w, data)
// 		return nil
// 	case float32:
// 		w.WriteByte(byte(tagFloat))
// 		e.WriteString(w, tagName)
// 		e.WriteFloat32(w, data)
// 		return nil
// 	case float64:
// 		w.WriteByte(byte(tagDouble))
// 		e.WriteString(w, tagName)
// 		e.WriteFloat64(w, data)
// 		return nil
// 	case []int32:
// 		w.WriteByte(byte(tagIntArray))
// 		e.WriteString(w, tagName)
// 		e.WriteInt32(w, int32(len(data)))
// 		for _, v := range data {
// 			e.WriteInt32(w, int32(v))
// 		}
// 		return nil
// 	case []int8:
// 		w.WriteByte(byte(tagByteArray))
// 		e.WriteString(w, tagName)
// 		e.WriteInt32(w, int32(len(data)))
// 		for _, v := range data {
// 			w.WriteByte(byte(v))
// 		}
// 		return nil
// 	case []int64:
// 		w.WriteByte(byte(tagLongArray))
// 		e.WriteString(w, tagName)
// 		e.WriteInt32(w, int32(len(data)))
// 		for _, v := range data {
// 			e.WriteInt64(w, (v))
// 		}
// 		return nil
// 	case []any:
// 		w.WriteByte(byte(tagList))
// 		e.WriteString(w, tagName)

// 		// special case: 0 length
// 		if len(data) == 0 {
// 			w.WriteByte(byte(tagEnd))
// 			return nil
// 		}
// 		inspectpr := &encodingInspector{}
// 		e.WriteString(inspectpr, "")
// 		tagNameLen := inspectpr.off
// 		fmt.Println(tagNameLen)

// 		trapHead := &trapList1Writor[W]{
// 			onWriteLen: func() {},
// 			underlay:   w,
// 			tagNameLen: tagNameLen,
// 			off:        0,
// 		}
// 		encodeTo[E](trapHead, "", trapHead, caster, false)
// 		return nil
// 		// writeString(w, "[ ")
// 		// for i, v := range data {
// 		// 	if i != 0 {
// 		// 		w.WriteByte(',')
// 		// 		w.WriteByte(' ')
// 		// 	}
// 		// 	err = encodeTo(w, v, caster, false, intBuf)
// 		// 	if err != nil {
// 		// 		return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
// 		// 	}
// 		// }
// 		// w.WriteByte(']')
// 		// return nil
// 	case map[string]any:
// 		w.WriteByte(byte(tagCompound))
// 		e.WriteString(w, tagName)
// 		keys := make([]string, 0, len(data))
// 		for k := range data {
// 			keys = append(keys, k)
// 		}
// 		sort.Strings(keys)
// 		for _, k := range keys {
// 			v := data[k]
// 			encodeTo[E, W](w, k, v, caster, false)
// 		}
// 		w.WriteByte(byte(tagEnd))
// 		return nil
// 	default:
// 		panic(value)
// 		// 	return encodeToWithCast[W, E](w, input, caster, casted, intBuf)
// 	}
// }

// func takeStructMember(val reflect.Value, out map[string]any) {
// 	for i := 0; i < val.NumField(); i++ {
// 		fieldType := val.Type().Field(i)
// 		fieldValue := val.Field(i)
// 		tag := fieldType.Tag.Get("nbt")
// 		if fieldType.PkgPath != "" || tag == "-" {
// 			continue
// 		}
// 		if fieldType.Anonymous {
// 			// The field was anonymous, so we write that in the same compound tag as this one.
// 			takeStructMember(fieldValue, out)
// 			continue
// 		}
// 		tagName := fieldType.Name
// 		if strings.HasSuffix(tag, ",omitempty") {
// 			tag = strings.TrimSuffix(tag, ",omitempty")
// 			if reflect.DeepEqual(fieldValue.Interface(), reflect.Zero(fieldValue.Type()).Interface()) {
// 				// The tag had the ',omitempty' tag, meaning it should be omitted if it has the zero
// 				// value. If this is reached, that was the case, and we skip it.
// 				continue
// 			}
// 		}
// 		if tag != "" {
// 			tagName = tag
// 		}
// 		out[tagName] = fieldValue.Interface()
// 	}
// }

// func encodeToWithCast[W base_io.Writer, E encoding.Encoding](w W, orig any, caster func(any) any, casted bool, intBuf []byte) (err error) {
// 	// we don't want caster allocate anything like []any{} or map[string]any{}
// 	val := reflect.ValueOf(orig)
// 	kind := val.Kind()
// 	switch kind {
// 	case reflect.Array:
// 		switch val.Type().Elem().Kind() {
// 		case reflect.Int8:
// 			writeString(w, "[B; ")
// 			n := val.Cap()
// 			for i := 0; i < n; i++ {
// 				if i != 0 {
// 					w.WriteByte(',')
// 					w.WriteByte(' ')
// 				}
// 				writeInt(w, int8(val.Index(i).Int()), intBuf)
// 				w.WriteByte('b')
// 			}
// 			w.WriteByte(']')
// 			return nil
// 		case reflect.Int32:
// 			writeString(w, "[I; ")
// 			n := val.Cap()
// 			for i := 0; i < n; i++ {
// 				if i != 0 {
// 					w.WriteByte(',')
// 					w.WriteByte(' ')
// 				}
// 				writeInt(w, int32(val.Index(i).Int()), intBuf)
// 			}
// 			w.WriteByte(']')
// 			return nil
// 		case reflect.Int64:
// 			writeString(w, "[L; ")
// 			n := val.Cap()
// 			for i := 0; i < n; i++ {
// 				if i != 0 {
// 					w.WriteByte(',')
// 					w.WriteByte(' ')
// 				}
// 				writeInt(w, int64(val.Index(i).Int()), intBuf)
// 				w.WriteByte('l')
// 			}
// 			w.WriteByte(']')
// 			return nil
// 		default:
// 			n := val.Cap()
// 			if n > 0 {
// 				decided := 0
// 				if caster != nil {
// 					new := caster(val.Index(0).Interface())
// 					if v, ok := new.(int8); ok {
// 						decided = 1
// 						writeString(w, "[B; ")
// 						writeInt(w, v, intBuf)
// 					} else if v, ok := new.(int32); ok {
// 						decided = 2
// 						writeString(w, "[I; ")
// 						writeInt(w, v, intBuf)
// 					} else if v, ok := new.(int64); ok {
// 						decided = 3
// 						writeString(w, "[L; ")
// 						writeInt(w, v, intBuf)
// 					} else {
// 						writeString(w, "[ ")
// 						err = encodeTo(w, new, caster, false, intBuf)
// 						if err != nil {
// 							return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
// 						}
// 					}
// 				}
// 				for i := 1; i < n; i++ {
// 					w.WriteByte(',')
// 					w.WriteByte(' ')
// 					v := caster(val.Index(i).Interface())
// 					switch decided {
// 					case 0:
// 						err = encodeTo(w, v, caster, false, intBuf)
// 						if err != nil {
// 							return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
// 						}
// 					case 1:
// 						writeInt(w, v.(int8), intBuf)
// 					case 2:
// 						writeInt(w, v.(int32), intBuf)
// 					case 3:
// 						writeInt(w, v.(int64), intBuf)
// 					}
// 				}
// 				w.WriteByte(']')
// 			} else {
// 				w.WriteByte('[')
// 				w.WriteByte(']')
// 			}

// 			return nil
// 		}
// 	case reflect.Slice:
// 		switch val.Type().Elem().Kind() {
// 		case reflect.Int32:
// 			ret := unsafe.Slice((*int32)(unsafe.Pointer(val.Pointer())), val.Len())
// 			return encodeTo[W, E](w, ret, caster, false, intBuf)
// 		case reflect.Int8:
// 			ret := unsafe.Slice((*int8)(unsafe.Pointer(val.Pointer())), val.Len())
// 			return encodeTo[W, E](w, ret, caster, false, intBuf)
// 		case reflect.Int64:
// 			ret := unsafe.Slice((*int64)(unsafe.Pointer(val.Pointer())), val.Len())
// 			return encodeTo[W, E](w, ret, caster, false, intBuf)
// 		default:
// 			writeString(w, "[ ")
// 			for i := 0; i < val.Len(); i++ {
// 				if i != 0 {
// 					w.WriteByte(',')
// 					w.WriteByte(' ')
// 				}
// 				err = encodeTo[W, E](w, val.Index(i).Interface(), caster, false, intBuf)
// 				if err != nil {
// 					return fmt.Errorf("%v: %v", ErrCannotEncodeListElem, err)
// 				}
// 			}
// 			w.WriteByte(']')
// 			return nil
// 		}
// 	case reflect.Struct:
// 		out := map[string]any{}
// 		takeStructMember(val, out)
// 		return encodeToWithCast[W, E](w, out, caster, false, intBuf)
// 	case reflect.Map:
// 		stringK := true
// 		if val.Type().Key().Kind() != reflect.String {
// 			stringK = false
// 		}
// 		iter := val.MapRange()
// 		unsortedMap := map[string]any{}
// 		// take data
// 		for iter.Next() {
// 			key := ""
// 			if stringK {
// 				key = iter.Key().String()
// 			} else {
// 				key = fmt.Sprintf("%v", iter.Key().Interface())
// 			}
// 			unsortedMap[key] = iter.Value().Interface()
// 		}
// 		return encodeTo[W, E](w, unsortedMap, caster, false, intBuf)
// 	default:
// 		if !casted && caster != nil {
// 			new := caster(orig)
// 			if new != nil {
// 				return encodeTo[W, E](w, new, caster, true, intBuf)
// 			}
// 		}
// 		return ErrCannotAcceptType
// 	}
// }
