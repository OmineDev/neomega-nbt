package number

import (
	"math"
	"neomega_nbt/base_io/lflb"
)

type NumberFinity struct {
	goBack         int8
	state          uint8
	previousAccept bool

	// number token specific
	dataT      int8 //0 -> int32(I) 1 -> int8(B) 2 -> int16(S) 3-> int64(L) 4->float(F) 5->double(D)
	neg        bool
	baseVal    int64
	tailVal    int64
	tailOffset int32
	negExp     bool
	exp        int32
}

func (o *NumberFinity) handleMark(b, m uint8) {
	switch m {
	case 0:
		o.neg = true
	case 1:
		o.baseVal = o.baseVal * 10
		o.baseVal += int64(b - '0')
	case 2:
		o.dataT = 1
	case 3:
		o.dataT = 2
	case 4:
		o.dataT = 3
	case 5:
		o.dataT = 5
		if b != '.' {
			o.tailOffset += 1
			o.tailVal = o.tailVal * 10
			o.tailVal += int64(b - '0')
		}
	case 6:
		o.negExp = true
	case 7:
		o.dataT = 5
		o.exp = o.exp * 10
		o.exp += int32(b - '0')
	case 8:
		o.dataT = 4
	case 9:
		o.dataT = 5
	case 10:
		o.dataT = 1
		o.baseVal = 1
	case 11:
		o.dataT = 1
		o.baseVal = 0
	}
}

func (o *NumberFinity) Feed(b byte) lflb.Status {
	o.goBack += 1
	nextMarkTermAccept := NumberFeed(o.state, b)
	if termOrAcc := nextMarkTermAccept & 0x3; termOrAcc != 0 {
		switch termOrAcc {
		case 0x3:
			if nextMarkTermAccept&0x80 == 0 {
				m := uint8(nextMarkTermAccept) >> 2
				o.handleMark(b, m)
			}
			return lflb.RET_TERM_OK
		case 0x2:
			if o.previousAccept {
				return lflb.RET_TERM_OK | lflb.Status(o.goBack)<<2
			} else {
				return lflb.RET_TERM_FAIL
			}
		case 0x1:
			o.goBack = 0
			o.previousAccept = true
		}
	}

	o.state = uint8(nextMarkTermAccept >> 8)
	if nextMarkTermAccept&0x80 == 0 {
		m := uint8(nextMarkTermAccept) >> 2
		o.handleMark(b, m)
	}
	return lflb.RET_FEED_MORE
}
func (o *NumberFinity) FeedEof() lflb.Status {
	if o.previousAccept {
		return lflb.RET_TERM_OK | lflb.Status(o.goBack)<<2
	} else {
		return lflb.RET_TERM_FAIL
	}
}

func (o *NumberFinity) IsInt32Overflow() bool {
	if o.dataT != 0 {
		return false
	}
	return int64(int32(o.baseVal)) != o.baseVal
}

func (o *NumberFinity) Val() any {
	if o.neg {
		o.baseVal = -o.baseVal
		o.tailVal = -o.tailVal
	}
	switch o.dataT {
	case 0:
		return int32(o.baseVal)
	case 1:
		return int8(o.baseVal)
	case 2:
		return int16(o.baseVal)
	case 3:
		return int64(o.baseVal)
	case 4:
		if o.negExp {
			o.exp = -o.exp
		}
		return ((float32(o.baseVal) + (float32(o.tailVal) * float32(math.Pow10(-int(o.tailOffset))))) * float32(math.Pow10(int(o.exp))))
	case 5:
		if o.negExp {
			o.exp = -o.exp
		}
		return ((float64(o.baseVal) + (float64(o.tailVal) * float64(math.Pow10(-int(o.tailOffset))))) * float64(math.Pow10(int(o.exp))))
	}
	return nil
}

func (o *NumberFinity) String() string {
	return "[number]"
}
func (o *NumberFinity) Reset() {
	*o = NumberFinity{}
}
