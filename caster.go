package snbt

import (
	"reflect"
	"strings"
	"unsafe"
)

func StructCaster(val reflect.Value, out map[string]any) {
	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Type().Field(i)
		fieldValue := val.Field(i)
		tag := fieldType.Tag.Get("nbt")
		if fieldType.PkgPath != "" || tag == "-" {
			continue
		}
		if fieldType.Anonymous {
			// The field was anonymous, so we write that in the same compound tag as this one.
			StructCaster(fieldValue, out)
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

func DefaultCaster(orig any) any {
	val := reflect.ValueOf(orig)
	kind := val.Kind()
	switch kind {
	case reflect.Int8:
		return int8(val.Int())
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
			n := val.Cap()
			data := make([]int8, n)
			for i := 0; i < n; i++ {
				data[i] = int8(val.Index(i).Int())
			}
			return data
		case reflect.Int32:
			n := val.Cap()
			data := make([]int32, n)
			for i := 0; i < n; i++ {
				data[i] = int32(val.Index(i).Int())
			}
			return data
		case reflect.Int64:
			n := val.Cap()
			data := make([]int64, n)
			for i := 0; i < n; i++ {
				data[i] = int64(val.Index(i).Int())
			}
			return data
		}
	case reflect.Slice:
		switch val.Type().Elem().Kind() {
		case reflect.Int32:
			ret := unsafe.Slice((*int32)(unsafe.Pointer(val.Pointer())), val.Len())
			return ret
		case reflect.Int8:
			ret := unsafe.Slice((*int8)(unsafe.Pointer(val.Pointer())), val.Len())
			return ret
		case reflect.Int64:
			ret := unsafe.Slice((*int64)(unsafe.Pointer(val.Pointer())), val.Len())
			return ret
		default:
			shallowCpy := make([]any, val.Len())
			for i := 0; i < val.Len(); i++ {
				shallowCpy[i] = val.Index(i).Interface()
			}
			return shallowCpy
		}
	case reflect.Struct:
		shallowCpy := map[string]any{}
		StructCaster(val, shallowCpy)
		return shallowCpy
	case reflect.Map:
		shallowCpy := map[string]any{}
		if val.Type().Key().Kind() != reflect.String {
			return nil
		}
		iter := val.MapRange()
		for iter.Next() {
			shallowCpy[iter.Key().String()] = iter.Value().Interface()
		}
		return shallowCpy
	}
	return nil
}
