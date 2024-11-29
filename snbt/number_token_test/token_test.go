package main

import (
	"fa/lflb"
	"fa/lflb/lflbops"
	"fa/lflb/sources"
	"fmt"
	"math"
	"testing"
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
		o.dataT = 4
		if b != '.' {
			o.tailOffset += 1
			o.tailVal = o.tailVal * 10
			o.tailVal += int64(b - '0')
		}
	case 6:
		o.negExp = true
	case 7:
		o.dataT = 4
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

func TestNumber(t *testing.T) {
	seq := "-123,123,-1,1,-1b,1b,103b,-104b,-1s,1s,103s,-104s,-1l,1l,103l,-104l,-56.78,3.2,-3.,-.2,-234.,-.456,123E-2,-.456E1,3.E1,45.58E-12,123.456E-8f,123.456E10d,true,false,3b,end"
	number := &NumberFinity{}
	src := sources.NewBytesSourceFromString(seq)
	assertVal := func(val any, T byte) {
		switch T {
		case 'I':
			getV, ok1 := number.Val().(int32)
			wantV, ok2 := val.(int32)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'B':
			getV, ok1 := number.Val().(int8)
			wantV, ok2 := val.(int8)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}

		case 'S':
			getV, ok1 := number.Val().(int16)
			wantV, ok2 := val.(int16)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}

		case 'L':
			getV, ok1 := number.Val().(int64)
			wantV, ok2 := val.(int64)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if getV != wantV {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'F':
			getV, ok1 := number.Val().(float32)
			wantV, ok2 := val.(float32)
			if !ok1 || !ok2 {
				t.FailNow()
			}
			if math.Abs(float64(getV-wantV)) > 0.0001 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		case 'D':
			getV, ok1 := number.Val().(float64)
			wantV, ok2 := val.(float64)
			if !ok1 || !ok2 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
			if math.Abs(float64(getV-wantV)) > 1e-14 {
				t.Logf("want %v, get %v\n", wantV, getV)
				t.FailNow()
			}
		}
	}

	number.Reset()
	lflb.ReadFinity(src, number)
	assertVal(int32(-123), 'I')

	readNext := func() {
		lflb.ReadFinity(src, lflbops.FinityAny{})
		number.Reset()
		lflb.ReadFinity(src, number)
	}
	readNext()
	assertVal(int32(123), 'I')
	readNext()
	assertVal(int32(-1), 'I')
	readNext()
	assertVal(int32(1), 'I')

	readNext()
	assertVal(int8(-1), 'B')
	readNext()
	assertVal(int8(1), 'B')
	readNext()
	assertVal(int8(103), 'B')
	readNext()
	assertVal(int8(-104), 'B')
	readNext()
	assertVal(int16(-1), 'S')
	readNext()
	assertVal(int16(1), 'S')
	readNext()
	assertVal(int16(103), 'S')
	readNext()
	assertVal(int16(-104), 'S')
	readNext()
	assertVal(int64(-1), 'L')
	readNext()
	assertVal(int64(1), 'L')
	readNext()
	assertVal(int64(103), 'L')
	readNext()
	assertVal(int64(-104), 'L')
	readNext()
	assertVal(float32(-56.78), 'F')
	readNext()
	assertVal(float32(3.2), 'F')
	readNext()
	assertVal(float32(-3.0), 'F')
	readNext()
	assertVal(float32(-0.2), 'F')
	readNext()
	assertVal(float32(-234.0), 'F')
	readNext()
	assertVal(float32(-0.456), 'F')
	readNext()
	assertVal(float32(1.23), 'F')
	readNext()
	assertVal(float32(-4.56), 'F')
	readNext()
	assertVal(float32(30), 'F')
	readNext()
	assertVal(float32(45.58e-12), 'F')
	readNext()
	assertVal(float32(123.456e-8), 'F')
	readNext()
	assertVal(float64(123.456e10), 'D')
	readNext()
	assertVal(int8(1), 'B')
	readNext()
	assertVal(int8(0), 'B')
	readNext()
	assertVal(int8(3), 'B')
	readNext()
	fmt.Println(src)
}
