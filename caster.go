package snbt

import (
	"reflect"
)

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
	}
	return nil
}
