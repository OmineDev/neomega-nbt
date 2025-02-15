package int_arr

import (
	"neomega_nbt/base_io/lflb"
)

type IntArray[I interface {
	int32 | int8 | int64
}] struct {
	state byte
	// 0x1 activate -> is reading a number
	// 0x2 upcoming -> should follow a number
	// 0x4 neg -> current number is negative
	// 0x8 true
	// 0x10 (16) false
	// 1,2 << 5 rue, 1,2,3 <<5 alse
	currD I
	data  []I // 存储解析后的数字
}

func (f *IntArray[I]) Feed(b byte) lflb.Status {
	if b == ']' {
		if f.state&0x2 != 0 {
			// should have number append, but not
			return lflb.RET_TERM_FAIL
		}
		if f.state&(0x8+0x10) != 0 {
			// in reading true/false
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

	switch b {
	case 't':
		if f.state&0x1 != 0 {
			return lflb.RET_TERM_FAIL
		}
		f.currD = 1
		f.state = (0x8 + 0x1)
		return lflb.RET_FEED_MORE
	case 'r':
		if f.state != (0x8 + 0x1) {
			return lflb.RET_TERM_FAIL
		}
		f.state = (0x8 + 0x1) + (0x1 << 5)
		return lflb.RET_FEED_MORE
	case 'u':
		if f.state != (0x8+0x1)+(0x1<<5) {
			return lflb.RET_TERM_FAIL
		}
		f.state = (0x8 + 0x1) + (0x2 << 5)
		return lflb.RET_FEED_MORE
	case 'f':
		if f.state&0x1 != 0 {
			return lflb.RET_TERM_FAIL
		}
		f.currD = 0
		f.state = (0x10 + 0x1)
		return lflb.RET_FEED_MORE
	case 'a':
		if f.state != (0x10 + 0x1) {
			return lflb.RET_TERM_FAIL
		}
		f.state = (0x10 + 0x1) + (0x1 << 5)
		return lflb.RET_FEED_MORE
	case 'l':
		if f.state != (0x10+0x1)+(0x1<<5) {
			return lflb.RET_TERM_FAIL
		}
		f.state = (0x10 + 0x1) + (0x2 << 5)
		return lflb.RET_FEED_MORE
	case 's':
		if f.state != (0x10+0x1)+(0x2<<5) {
			return lflb.RET_TERM_FAIL
		}
		f.state = (0x10 + 0x1) + (0x3 << 5)
		return lflb.RET_FEED_MORE
	case 'e':
		if f.state != (0x8+0x1)+(0x2<<5) && f.state != (0x10+0x1)+(0x3<<5) {
			return lflb.RET_TERM_FAIL
		}
		f.state = 1
		return lflb.RET_FEED_MORE
	}
	if f.state&(0x8+0x10) != 0 {
		// in reading true/false
		return lflb.RET_TERM_FAIL
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
	if b == ' ' || b == '\n' || b == '\t' || b == '\r' || b == 'B' || b == 'b' || b == 'L' || b == 'l' {
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
