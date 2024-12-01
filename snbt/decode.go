package snbt

import (
	"fa/lflb"
	"fa/snbt/tokens/int_arr"
	"fa/snbt/tokens/left_container.go"
	"fa/snbt/tokens/number"
	"fa/snbt/tokens/string_token"
	"fmt"
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

func DecodeFrom[S lflb.Source](src S) (v any, err error) {
	// consume white space
	// equal to: lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	// ReadWhiteSpace(src)

	// return nil if is blank
	if _, eof := src.This(); eof {
		return nil, ErrNoData
	}

	// check is left container: 1" 2' 3{ 4[ 5[I; 6[B; 7[L
	leftContainer := &left_container.LeftContainerFinity{}
	lflb.ReadFinity(src, leftContainer)
	switch leftContainer.ContainerType {
	case 1:
		// "
		stringD := &string_token.StringD{}
		if !lflb.ReadFinity(src, stringD) {
			return nil, ErrStringDNotTerminated
		}
		return stringD.Val(), nil
	case 2:
		// '
		stringS := &string_token.StringS{}
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
		return DecodeListFrom(src)
	case 3:
		return DecodeCompoundFrom(src)
	}

	// check is number
	// 123 is number
	// 123_abc is unwrap string
	numberToken := &number.NumberFinity{}
	if ok, counter := lflb.ReadFinityWithCounter(src, numberToken); ok {
		if b, eof := src.This(); eof || !string_token.AcceptUnwrapStringFeed(b) {
			return numberToken.Val(), nil
		} else {
			// sth like 123_abc, now @ _abc
			src.Back(counter - 1)
		}
	}

	// check is unwrap string
	unwrapString := &string_token.UnwrapString{}
	if lflb.ReadFinity(src, unwrapString) {
		return unwrapString.Val(), nil
	}

	// not snbt dat
	return nil, ErrNotSNBT
}

func DecodeListFrom[S lflb.Source](src S) (v []any, err error) {
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
		elem, valErr := DecodeFrom(src)
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

func DecodeCompoundFrom[S lflb.Source](src S) (v map[string]any, err error) {
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
		elem, valErr := DecodeFrom(src)
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
