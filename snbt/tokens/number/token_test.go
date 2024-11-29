package number

import (
	"fa/lflb"
	"fa/lflb/lflbops"
	"fa/lflb/sources"
	"fmt"
	"math"
	"testing"
)

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
