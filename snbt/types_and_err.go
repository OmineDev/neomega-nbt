package snbt

import "errors"

var ErrNoData = errors.New("no data")
var ErrNotSNBT = errors.New("not snbt formate")
var ErrStringDNotTerminated = errors.New("string should terminated with \"")
var ErrStringSNotTerminated = errors.New("string should terminated with '")
var ErrNotInt32Arr = errors.New("data not match int32 array ([I; 123, 456, ...]) restriction")
var ErrNotInt8Arr = errors.New("data not match int8 array ([I; 12, 45, ...]) restriction")
var ErrNotInt64Arr = errors.New("data not match int64 array ([I; 123, 456, ...]) restriction")
var ErrListNotTerminated = errors.New("list has no right ]")
var ErrListNoRightComma = errors.New(`list has no "," after element`)
var ErrListElementTypeMismatch = errors.New(`elements type in list are not same`)
var ErrCompoundHasNoValidKey = errors.New("compound has no valid key")
var ErrCompoundHasNoColon = errors.New(`compound has no ":"`)
var ErrListElementError = "at decoding list element, a error occours:"
var ErrCompoundValError = "at decoding compound element, a error occours:"

type SNBType uint8

const (
	SNBTUnknown = SNBType(0)
	SNBTString  = SNBType(iota + 1)
	SNBTInt32
	SNBTInt8
	SNBTInt16
	SNBTInt64
	SNBTFloat32
	SNBTFloat64
	SNBTInt32Arr
	SNBTInt8Arr
	SNBTInt64Arr
	SNBTList
	SNBTCompound
)

func GetSNBTValueTypeID(data any) SNBType {
	switch data.(type) {
	case string:
		return SNBTString
	case int32:
		return SNBTInt32
	case int8:
		return SNBTInt8
	case int16:
		return SNBTInt16
	case int64:
		return SNBTInt64
	case float32:
		return SNBTFloat32
	case float64:
		return SNBTFloat64
	case []int32:
		return SNBTInt32Arr
	case []int8:
		return SNBTInt8Arr
	case []int64:
		return SNBTInt64Arr
	case []any:
		return SNBTList
	case map[string]any:
		return SNBTCompound
	}
	return SNBTUnknown
}
