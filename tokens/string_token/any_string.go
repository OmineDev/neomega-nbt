package string_token

import "snbt/lflb"

type AnyString struct {
	state byte //0 -> unknown 1-> stringD 2-> stringS 3-> unwrap 5-> stringD+slash 6-> stringS+slash
	data  []byte
}

func (o *AnyString) Feed(b byte) lflb.Status {
	switch o.state {
	case 0:
		// unknow
		switch b {
		case '"':
			o.state = 1
			return lflb.RET_FEED_MORE
		case '\'':
			o.state = 2
			return lflb.RET_FEED_MORE
		default:
			if AcceptUnwrapStringFeed(b) {
				o.state = 3
				if o.data == nil {
					o.data = make([]byte, 0, 8)
				}
				o.data = append(o.data, b)
				return lflb.RET_FEED_MORE
			}
			return lflb.RET_TERM_FAIL
		}
	case 1:
		// stringD
		if b == '\\' {
			o.state = 5 // stringD + slash
			return lflb.RET_FEED_MORE
		}
		if b == '"' {
			return lflb.RET_TERM_OK
		}
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	case 5:
		o.state = 1
		if o.data == nil {
			o.data = make([]byte, 0, 8)
		}
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	case 2:
		// stringS
		if b == '\\' {
			o.state = 6 // stringS + slash
			return lflb.RET_FEED_MORE
		}
		if b == '\'' {
			return lflb.RET_TERM_OK
		}
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	case 6:
		o.state = 2
		if o.data == nil {
			o.data = make([]byte, 0, 8)
		}
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	case 3:
		if AcceptUnwrapStringFeed(b) {
			o.data = append(o.data, b)
			return lflb.RET_FEED_MORE
		} else {
			if len(o.data) > 0 {
				return lflb.RET_TERM_OK | lflb.Status(1)<<2
			}
			return lflb.RET_TERM_FAIL
		}
	default:
		return lflb.RET_TERM_FAIL
	}
}

func (o *AnyString) FeedEof() lflb.Status {
	if o.state != 3 {
		return lflb.RET_TERM_FAIL
	}
	if len(o.data) > 0 {
		return lflb.RET_TERM_OK
	}
	return lflb.RET_TERM_FAIL
}

func (o *AnyString) Val() string {
	return string(o.data)
}
func (o *AnyString) String() string { return "[string]" }
func (o *AnyString) Reset() {
	o.state = 0
	o.data = o.data[:0]
}
