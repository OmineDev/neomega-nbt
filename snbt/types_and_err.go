package snbt

import "errors"

var ErrNoData = errors.New("no data")
var ErrNotSNBT = errors.New("not snbt formate")
var ErrStringDNotTerminated = errors.New("string should terminated with \"")
var ErrStringSNotTerminated = errors.New("string should terminated with '")
var ErrNotInt32Arr = errors.New("data not match int32 array ([I; 123, 456, ...]) restriction")
var ErrNotInt8Arr = errors.New("data not match int8 array ([B; 12, 45, ...]) restriction")
var ErrNotInt64Arr = errors.New("data not match int64 array ([L; 123, 456, ...]) restriction")
var ErrListNotTerminated = errors.New("list has no right ]")
var ErrListNoRightComma = errors.New(`list has no "," after element`)
var ErrListElementTypeMismatch = errors.New(`elements type in list are not same`)
var ErrCompoundHasNoValidKey = errors.New("compound has no valid key")
var ErrCompoundHasNoColon = errors.New(`compound has no ":"`)
var ErrListElementError = "at decoding list element, a error occours:"
var ErrCompoundValError = "at decoding compound element, a error occours:"

type SNBType uint8

const (
	tagInvalid = SNBType(0)
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

func GetSNBTValueTypeID(data any) SNBType {
	switch data.(type) {
	case string:
		return tagString
	case int32:
		return tagInt
	case int8:
		return tagByte
	case int16:
		return tagShort
	case int64:
		return tagLong
	case float32:
		return tagFloat
	case float64:
		return tagDouble
	case []int32:
		return tagIntArray
	case []int8:
		return tagByteArray
	case []int64:
		return tagLongArray
	case []any:
		return tagList
	case map[string]any:
		return tagCompound
	}
	return tagInvalid
}
