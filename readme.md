snbt (like) decoder   
aims to prase snbt    
data that follows snbt standard can be parsed by this package,   
however, data can be parsed by this package may not be a a valid snbt string,   
e.g. [123, abc] is illegal snbt string (snbt requires all element in list have same type, you can enable it a decode.go, line 125)   

data mapping (especially list):   
```
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
        // event if all elem in list is same type, we still keep it []any to avoid confliction with [I;], [B;], [L;]
		return SNBTList
	case map[string]any:
		return SNBTCompound
	}
	return SNBTUnknown
}
```

there are some magic code (e.g. tokens/number/core) generated by gen_code using package fa (https://github.com/OmineDev/fa)
