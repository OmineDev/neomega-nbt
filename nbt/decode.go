package nbt

import (
	"fmt"
	"neomega_nbt/base_io"
	"neomega_nbt/encoding"
)

type ErrInvalidTagType struct {
	tag tagType
}

func (e ErrInvalidTagType) Error() string {
	return fmt.Sprintf("invalid tagType: %v", e.tag)
}

func DecodeValue[E encoding.ReadEncoding, R base_io.Reader](
	r R,
	tagT tagType,
) (any, error) {
	var e E
	switch tagT {
	case tagString:
		return e.String(r)
	case tagInt:
		return e.Int32(r)
	case tagByte:
		b, err := r.ReadByte()
		return int8(b), err
	case tagShort:
		return e.Int16(r)
	case tagLong:
		return e.Int64(r)
	case tagFloat:
		return e.Float32(r)
	case tagDouble:
		return e.Float64(r)
	case tagIntArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		out := make([]int32, dataLen)
		for i := uint32(0); i < uint32(dataLen); i++ {
			out[i], err = e.Int32(r)
			if err != nil {
				return nil, err
			}
		}
		return out, nil
	case tagByteArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		out := make([]int8, dataLen)
		for i := uint32(0); i < uint32(dataLen); i++ {
			b, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			out[i] = int8(b)
		}
		return out, nil
	case tagLongArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		out := make([]int64, dataLen)
		for i := uint32(0); i < uint32(dataLen); i++ {
			out[i], err = e.Int64(r)
			if err != nil {
				return nil, err
			}
		}
		return out, nil
	case tagCompound:
		out := map[string]any{}
		for {
			tagTypeB, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			if tagTypeB == byte(tagEnd) {
				return out, nil
			}
			tagName, err := e.String(r)
			if err != nil {
				return nil, err
			}
			v, err := DecodeValue[E, R](r, tagType(tagTypeB))
			if err != nil {
				return nil, err
			}
			out[tagName] = v
		}
	case tagList:
		tagTypeB, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		if tagTypeB == byte(tagEnd) {
			dataLen, err := e.Int32(r)
			if err != nil {
				return nil, err
			}
			if dataLen != 0 {
				fmt.Printf("length should be 0 (type is tag end), but get %v\n", dataLen)
			}
			return []any{}, nil
		}
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		out := make([]any, dataLen)
		for i := int32(0); i < dataLen; i++ {
			out[i], err = DecodeValue[E, R](r, tagType(tagTypeB))
			if err != nil {
				return nil, err
			}
		}
		return out, nil
	default:
		return nil, ErrInvalidTagType{tagT}
	}
}

func DecodeTagAndValue[E encoding.ReadEncoding, R base_io.Reader](
	r R,
) (string, any, error) {
	var e E
	tagTypeB, err := r.ReadByte()
	if err != nil {
		return "", nil, err
	}
	tagName, err := e.String(r)
	if err != nil {
		return "", nil, err
	}
	v, err := DecodeValue[E, R](r, tagType(tagTypeB))
	if err != nil {
		return "", nil, err
	}
	return tagName, v, nil
}

func DryDecodeValue[E encoding.ReadEncoding, R base_io.Reader](
	r R,
	tagT tagType,
) (any, error) {
	var e E
	switch tagT {
	case tagString:
		return e.String(r)
	case tagInt:
		return e.Int32(r)
	case tagByte:
		b, err := r.ReadByte()
		return int8(b), err
	case tagShort:
		return e.Int16(r)
	case tagLong:
		return e.Int64(r)
	case tagFloat:
		return e.Float32(r)
	case tagDouble:
		return e.Float64(r)
	case tagIntArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		for i := uint32(0); i < uint32(dataLen); i++ {
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case tagByteArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		for i := uint32(0); i < uint32(dataLen); i++ {
			_, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case tagLongArray:
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		for i := uint32(0); i < uint32(dataLen); i++ {
			_, err = e.Int64(r)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	case tagCompound:
		for {
			tagTypeB, err := r.ReadByte()
			if err != nil {
				return nil, err
			}
			if tagTypeB == byte(tagEnd) {
				return nil, nil
			}
			_, err = e.String(r)
			if err != nil {
				return nil, err
			}
			_, err = DecodeValue[E, R](r, tagType(tagTypeB))
			if err != nil {
				return nil, err
			}
		}
	case tagList:
		tagTypeB, err := r.ReadByte()
		if err != nil {
			return nil, err
		}
		if tagTypeB == byte(tagEnd) {
			return nil, nil
		}
		dataLen, err := e.Int32(r)
		if err != nil {
			return nil, err
		}
		for i := int32(0); i < dataLen; i++ {
			_, err = DecodeValue[E, R](r, tagType(tagTypeB))
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	default:
		return nil, ErrInvalidTagType{tagT}
	}
}

func DryDecodeTagAndValue[E encoding.ReadEncoding, R base_io.Reader](
	r R,
) error {
	var e E
	tagTypeB, err := r.ReadByte()
	if err != nil {
		return err
	}
	_, err = e.String(r)
	if err != nil {
		return err
	}
	_, err = DecodeValue[E, R](r, tagType(tagTypeB))
	if err != nil {
		return err
	}
	return nil
}
