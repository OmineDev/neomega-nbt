package neomega_nbt

import "reflect"

func ReadFromNBT[T any](nbt any, key string, defaultV T) T {
	compound, ok := nbt.(map[string]any)
	if !ok {
		return defaultV
	}
	vI, ok := compound[key]
	if !ok {
		return defaultV
	}
	if v, ok := vI.(T); ok {
		return v
	}
	var r T
	elem := reflect.ValueOf(&r).Elem()
	elem.Set(reflect.ValueOf(vI).Convert(elem.Type()))
	return r
}
