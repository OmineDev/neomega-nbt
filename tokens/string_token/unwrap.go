package string_token

import "snbt/lflb"

type UnwrapString struct {
	data []byte
}

func (o *UnwrapString) Feed(b byte) lflb.Status {
	if AcceptUnwrapStringFeed(b) {
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	} else {
		if len(o.data) > 0 {
			return lflb.RET_TERM_OK | lflb.Status(1)<<2
		}
		return lflb.RET_TERM_FAIL
	}
}

func AcceptUnwrapStringFeed(b byte) bool {
	switch b >> 6 {
	case 0:
		if (1<<((b<<2)>>2))&288063250384289792 > 0 {
			return true
		}
	case 1:
		if (1<<((b<<2)>>2))&576460745995190270 > 0 {
			return true
		}
	case 2:
		if (uint64(1)<<((b<<2)>>2))&18446744073709551615 > 0 {
			return true
		}
	case 3:
		if (uint64(1)<<((b<<2)>>2))&18446744073709551615 > 0 {
			return true
		}
	}
	return false
}
func (o *UnwrapString) FeedEof() lflb.Status {
	if len(o.data) > 0 {
		return lflb.RET_TERM_OK
	}
	return lflb.RET_TERM_FAIL
}
func (o *UnwrapString) Val() string {
	return string(o.data)
}
func (o *UnwrapString) String() string { return "[unwarp string]" }
func (o *UnwrapString) Reset() {
	o.data = o.data[:0]
}
