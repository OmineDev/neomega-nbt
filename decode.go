package snbt

import (
	"fmt"
	"snbt/lflb"
	"snbt/tokens/int_arr"
	"snbt/tokens/left_container.go"
	"snbt/tokens/number"
	"snbt/tokens/string_token"
)

func ReadWhiteSpace[S lflb.Source](src S) {
	// consume white space
	var b byte
	var eof bool
	for {
		if b, eof = src.This(); eof {
			return
		} else if b>>6 != 0 || ((1<<((b<<2)>>2))&4294983168 == 0) {
			return
		}
		src.Next()
	}

}

type preAlloc struct {
	*left_container.LeftContainerFinity
	*string_token.StringD
	*string_token.StringS
	*string_token.UnwrapString
	*number.NumberFinity
}

func DecodeFrom[S lflb.Source](src S) (v any, err error) {
	pa := &preAlloc{
		LeftContainerFinity: &left_container.LeftContainerFinity{},
		StringD:             &string_token.StringD{},
		StringS:             &string_token.StringS{},
		UnwrapString:        &string_token.UnwrapString{},
		NumberFinity:        &number.NumberFinity{},
	}
	return decodeFrom(src, pa)
}

func decodeFrom[S lflb.Source](src S, preAlloc *preAlloc) (v any, err error) {
	// consume white space
	// equal to: lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	// ReadWhiteSpace(src)

	// return nil if is blank
	if _, eof := src.This(); eof {
		return nil, ErrNoData
	}

	// check is left container: 1" 2' 3{ 4[ 5[I; 6[B; 7[L
	leftContainer := preAlloc.LeftContainerFinity
	leftContainer.Reset()
	lflb.ReadFinity(src, leftContainer)
	switch leftContainer.ContainerType {
	case 1:
		// "
		stringD := preAlloc.StringD
		stringD.Reset()
		if !lflb.ReadFinity(src, stringD) {
			return nil, ErrStringDNotTerminated
		}
		return stringD.Val(), nil
	case 2:
		// '
		stringS := preAlloc.StringS
		stringS.Reset()
		if !lflb.ReadFinity(src, stringS) {
			return nil, ErrStringSNotTerminated
		}
		return stringS.Val(), nil
	case 5:
		//[B;
		intArr := &int_arr.IntArray[int8]{}
		if !lflb.ReadFinity(src, intArr) {
			return nil, ErrNotInt8Arr
		}
		return intArr.Val(), nil
	case 6:
		// [I;
		intArr := &int_arr.IntArray[int32]{}
		if !lflb.ReadFinity(src, intArr) {
			return nil, ErrNotInt32Arr
		}
		return intArr.Val(), nil
	case 7:
		// [L;
		intArr := &int_arr.IntArray[int64]{}
		if !lflb.ReadFinity(src, intArr) {
			return nil, ErrNotInt64Arr
		}
		return intArr.Val(), nil
	case 4:
		// [
		return DecodeListFrom(src, preAlloc)
	case 3:
		return DecodeCompoundFrom(src, preAlloc)
	}

	// check is number
	// 123 is number
	// 123_abc is unwrap string
	numberToken := preAlloc.NumberFinity
	numberToken.Reset()
	if ok, counter := lflb.ReadFinityWithCounter(src, numberToken); ok {
		if b, eof := src.This(); eof || !string_token.AcceptUnwrapStringFeed(b) {
			if numberToken.IsInt32Overflow() {
				if src.IsEof() {
					src.Back(counter)
				} else {
					src.Back(counter - 1)
				}
			} else {
				return numberToken.Val(), nil
			}
		} else {
			// sth like 123_abc, now @ _abc
			if src.IsEof() {
				src.Back(counter)
			} else {
				src.Back(counter - 1)
			}
		}
	}

	// check is unwrap string
	unwrapString := preAlloc.UnwrapString
	unwrapString.Reset()
	if lflb.ReadFinity(src, unwrapString) {
		return unwrapString.Val(), nil
	}

	// not snbt dat
	return nil, ErrNotSNBT
}

func DecodeListFrom[S lflb.Source](src S, preAlloc *preAlloc) (v []any, err error) {
	var b byte
	var eof bool
	out := []any{}
	// dt := SNBTUnknown
	for {
		ReadWhiteSpace(src)
		if b, eof = src.This(); b == ']' {
			src.Next()
			return out, nil
		} else if eof {
			return nil, ErrListNotTerminated
		}
		elem, valErr := decodeFrom(src, preAlloc)
		if valErr == nil {
			out = append(out, elem)
		} else {
			return nil, fmt.Errorf("%v %v", ErrListElementError, valErr)
		}
		// et := GetSNBTValueTypeID(elem)
		// if dt == SNBTUnknown {
		// 	dt = et
		// } else {
		// 	if dt != et {
		// 		return nil, ErrListElementTypeMismatch
		// 	}
		// }
		ReadWhiteSpace(src)
		if b, eof = src.This(); b == ',' {
			src.Next()
		} else if b == ']' {
			src.Next()
			return out, nil
		} else if eof {
			return nil, ErrListNotTerminated
		} else {
			return nil, ErrListNoRightComma
		}
	}
}

func DecodeCompoundFrom[S lflb.Source](src S, preAlloc *preAlloc) (v map[string]any, err error) {
	var b byte
	var eof bool
	out := map[string]any{}
	keyString := &string_token.AnyString{}
	var key string
	for {
		ReadWhiteSpace(src)
		if b, eof = src.This(); b == '}' {
			src.Next()
			return out, nil
		} else if eof {
			return nil, ErrListNotTerminated
		}
		keyString.Reset()
		if !lflb.ReadFinity(src, keyString) {
			return nil, ErrCompoundHasNoValidKey
		}
		key = keyString.Val()
		ReadWhiteSpace(src)
		if b, _ = src.This(); b != ':' {
			return nil, ErrCompoundHasNoColon
		}
		src.Next()
		ReadWhiteSpace(src)
		elem, valErr := decodeFrom(src, preAlloc)
		if valErr == nil {
			out[key] = elem
		} else {
			return nil, fmt.Errorf("%v %v", ErrCompoundValError, valErr)
		}
		ReadWhiteSpace(src)
		if b, eof = src.This(); b == ',' {
			src.Next()
		} else if b == '}' {
			src.Next()
			return out, nil
		} else if eof {
			return nil, ErrListNotTerminated
		} else {
			return nil, ErrListNoRightComma
		}
	}
}
