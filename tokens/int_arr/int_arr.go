package int_arr

import "snbt/lflb"

type IntArray[I interface {
	int32 | int8 | int64
}] struct {
	state byte // 0x1 activate 0x2 upcoming 0x4 neg
	currD I
	data  []I // 存储解析后的数字
}

func (f *IntArray[I]) Feed(b byte) lflb.Status {
	if b == ']' {
		if f.state&0x2 != 0 {
			// should have number append, but not
			return lflb.RET_TERM_FAIL
		}
		if f.state&0x1 != 0 {
			// has a val
			if f.state&0x4 == 0 {
				f.data = append(f.data, f.currD)
			} else {
				f.data = append(f.data, -f.currD)
			}
		}
		return lflb.RET_TERM_OK
	}
	if b == ',' {
		if f.state&0x2 != 0 || f.state&0x1 == 0 {
			// should have number before, but not
			return lflb.RET_TERM_FAIL
		}
		// handle value
		if f.state&0x4 == 0 {
			f.data = append(f.data, f.currD)
		} else {
			f.data = append(f.data, -f.currD)
		}
		f.currD = 0
		f.state = 0x2
		return lflb.RET_FEED_MORE
	}
	if b == ' ' {
		return lflb.RET_FEED_MORE
	}
	if b == '-' {
		if f.state&0x5 != 0 {
			return lflb.RET_TERM_FAIL
		}
		// neg
		f.state |= 0x4
		f.state &= ^uint8(0x2)
		return lflb.RET_FEED_MORE
	}

	if b >= '0' && b <= '9' {
		f.state |= 0x1
		f.state &= ^uint8(0x2)
		f.currD = f.currD*10 + I(b-'0')
		return lflb.RET_FEED_MORE
	}
	return lflb.RET_TERM_FAIL
}

func (f *IntArray[I]) FeedEof() lflb.Status {
	return lflb.RET_TERM_FAIL
}

func (f *IntArray[I]) Val() []I {
	return f.data
}

func (o *IntArray[I]) String() string { return "[int array]" }
func (o *IntArray[I]) Reset() {
	o.data = nil
	o.state = 0
	o.currD = 0
}
