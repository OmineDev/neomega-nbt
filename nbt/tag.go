package nbt

const (
	tagEnd tagType = iota
	tagByte
	tagShort
	tagInt
	tagLong
	tagFloat
	tagDouble
	tagByteArray
	tagString
	tagList
	tagCompound
	tagIntArray
	tagLongArray
)

// tagType represents the type of NBT tag.
type tagType byte

// String converts a tagType to its string representation. This looks like `TAG_` + `<tag type>`, such as `TAG_Byte`.
func (t tagType) String() string {
	switch t {
	case tagEnd:
		return "TAG_End"
	case tagByte:
		return "TAG_Byte"
	case tagShort:
		return "TAG_Short"
	case tagInt:
		return "TAG_Int"
	case tagLong:
		return "TAG_Long"
	case tagFloat:
		return "TAG_Float"
	case tagDouble:
		return "TAG_Double"
	case tagByteArray:
		return "TAG_ByteArray"
	case tagString:
		return "TAG_String"
	case tagList:
		return "TAG_List"
	case tagCompound:
		return "TAG_Compound"
	case tagIntArray:
		return "TAG_IntArray"
	case tagLongArray:
		return "TAG_LongArray"
	default:
		panic("unknown tag")
	}
}

func cleanValue(value any, caster func(any) any, casted bool) (tag tagType, stdValue any) {
	switch data := value.(type) {
	case string:
		return tagString, data
	case int32:
		return tagInt, data
	case int8:
		return tagByte, data
	case int16:
		return tagShort, data
	case int64:
		return tagLong, data
	case float32:
		return tagFloat, data
	case float64:
		return tagDouble, data
	case []int32:
		return tagIntArray, data
	case []int8:
		return tagByteArray, data
	case []int64:
		return tagLongArray, data
	case []any:
		return tagList, data
	case map[string]any:
		return tagCompound, data
	default:
		if caster == nil || casted {
			return 0, nil
		}
		c := caster(value)
		if c == nil {
			return 0, nil
		}
		return cleanValue(c, caster, true)
	}
}

// func tagFromType(p any, caster func(any) any) (t tagType, ok bool) {
// 	switch p.(type) {
// 	case int8:
// 		return tagByte, true
// 	case int16:
// 		return tagInt16, true
// 	case int32:
// 		return tagInt32, true
// 	case int64:
// 		return tagInt64, true
// 	case float32:
// 		return tagFloat32, true
// 	case float64:
// 		return tagFloat64, true
// 	case []int32:
// 		return tagInt32Array, true
// 	case []int8:
// 		return tagByteArray, true
// 	case []int64:
// 		return tagInt64Array, true
// 	case []any:
// 		return tagList, true
// 	case map[string]any:
// 		return tagCompound, true
// 	}
// 	val := reflect.ValueOf(p)
// 	kind := val.Kind()
// 	switch kind {
// 	case reflect.Array:
// 		switch val.Type().Elem().Kind() {
// 		case reflect.Int8:
// 			return tagByteArray, true
// 		case reflect.Int32:
// 			return tagInt32Array, true
// 		case reflect.Int64:
// 			return tagInt64Array, true
// 		default:
// 			n := val.Cap()
// 			if n > 0 {
// 				if caster != nil {
// 					new := caster(val.Index(0).Interface())
// 					if _, ok := new.(int8); ok {
// 						return tagByteArray, true
// 					} else if _, ok := new.(int32); ok {
// 						return tagInt32Array, true
// 					} else if _, ok := new.(int64); ok {
// 						return tagInt64Array, true
// 					} else {
// 						return tagList, true
// 					}
// 				}
// 			}
// 		}
// 	case reflect.Slice:
// 		switch val.Type().Elem().Kind() {
// 		case reflect.Int32:
// 			return tagInt32Array, true
// 		case reflect.Int8:
// 			return tagByteArray, true
// 		case reflect.Int64:
// 			return tagInt64Array, true
// 		default:
// 			n := val.Cap()
// 			if n > 0 {
// 				if caster != nil {
// 					new := caster(val.Index(0).Interface())
// 					if _, ok := new.(int8); ok {
// 						return tagByteArray, true
// 					} else if _, ok := new.(int32); ok {
// 						return tagInt32Array, true
// 					} else if _, ok := new.(int64); ok {
// 						return tagInt64Array, true
// 					} else {
// 						return tagList, true
// 					}
// 				}
// 			}
// 		}
// 	}
// }
