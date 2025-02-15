package caster

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

type ErrCannotCastType struct {
	Data any
}

func (e ErrCannotCastType) Error() string {
	return fmt.Sprintf("cannot handle value: %v of type: %T", e.Data, e.Data)
}

func DefaultCaster(orig any) any {
	val := reflect.ValueOf(orig)
	kind := val.Kind()
	if val.Kind() == reflect.Ptr {
		return DefaultCaster(val.Elem().Interface())
	}
	switch kind {
	case reflect.Int8:
		return int8(val.Int())
	case reflect.Uint8:
		return int8(uint8(val.Uint()))
	case reflect.Bool:
		if val.Bool() {
			return int8(1)
		}
		return int8(0)
	case reflect.Int16:
		return int16(val.Int())
	case reflect.Int32:
		return int32(val.Int())
	case reflect.Int64:
		return int64(val.Int())
	case reflect.Float32:
		return float32(val.Float())
	case reflect.Float64:
		return float64(val.Float())
	case reflect.String:
		return val.String()
	case reflect.Array:
		switch val.Type().Elem().Kind() {
		case reflect.Int8:
			n := val.Len()
			out := make([]int8, n)
			for i := 0; i < n; i++ {
				out[i] = int8(val.Index(i).Int())
			}
			return out
		case reflect.Uint8:
			n := val.Len()
			out := make([]int8, n)
			for i := 0; i < n; i++ {
				out[i] = int8(uint8(val.Index(i).Uint()))
			}
			return out
		case reflect.Int32:
			n := val.Len()
			out := make([]int32, n)
			for i := 0; i < n; i++ {
				out[i] = int32(val.Index(i).Int())
			}
			return out
		case reflect.Uint32:
			n := val.Len()
			out := make([]int32, n)
			for i := 0; i < n; i++ {
				out[i] = int32(uint32(val.Index(i).Uint()))
			}
			return out
		case reflect.Int64:
			n := val.Len()
			out := make([]int64, n)
			for i := 0; i < n; i++ {
				out[i] = int64(val.Index(i).Int())
			}
			return out
		case reflect.Uint64:
			n := val.Len()
			out := make([]int64, n)
			for i := 0; i < n; i++ {
				out[i] = int64(uint64(val.Index(i).Uint()))
			}
			return out
		default:
			n := val.Len()
			out := make([]any, n)
			for i := 0; i < n; i++ {
				out[i] = val.Index(i).Interface()
			}
			return out
		}
	case reflect.Slice:
		switch val.Type().Elem().Kind() {
		case reflect.Int32, reflect.Uint32:
			return unsafe.Slice((*int32)(unsafe.Pointer(val.Pointer())), val.Len())
		case reflect.Int8, reflect.Uint8:
			return unsafe.Slice((*int8)(unsafe.Pointer(val.Pointer())), val.Len())
		case reflect.Int64, reflect.Uint64:
			return unsafe.Slice((*int64)(unsafe.Pointer(val.Pointer())), val.Len())
		default:
			n := val.Len()
			out := make([]any, n)
			for i := 0; i < n; i++ {
				out[i] = val.Index(i).Interface()
			}
			return out
		}
	case reflect.Struct:
		out := map[string]any{}
		takeStructMember(val, out)
		return out
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
		return unsortedMap
	}
	return nil
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
