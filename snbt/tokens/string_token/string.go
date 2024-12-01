package string_token

import "fa/lflb"

// abc"
type StringD struct {
	slash bool
	data  []byte
}

func (o *StringD) Feed(b byte) lflb.Status {
	if o.slash {
		o.slash = false
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	}
	if b == '\\' {
		o.slash = true
		return lflb.RET_FEED_MORE
	}
	if b == '"' {
		return lflb.RET_TERM_OK
	}
	o.data = append(o.data, b)
	return lflb.RET_FEED_MORE
}

func (o *StringD) FeedEof() lflb.Status {
	return lflb.RET_TERM_FAIL
}

func (o *StringD) Val() string {
	return string(o.data)
}
func (o *StringD) String() string { return "[string'_']" }
func (o *StringD) Reset() {
	o.slash = false
	o.data = nil
}

// abc'
type StringS struct {
	slash bool
	data  []byte
}

func (o *StringS) Feed(b byte) lflb.Status {
	if o.slash {
		o.slash = false
		o.data = append(o.data, b)
		return lflb.RET_FEED_MORE
	}
	if b == '\\' {
		o.slash = true
		return lflb.RET_FEED_MORE
	}
	if b == '\'' {
		return lflb.RET_TERM_OK
	}
	o.data = append(o.data, b)
	return lflb.RET_FEED_MORE
}

func (o *StringS) FeedEof() lflb.Status {
	return lflb.RET_TERM_FAIL
}

func (o *StringS) Val() string {
	return string(o.data)
}
func (o *StringS) String() string { return "[string\"_\"]" }
func (o *StringS) Reset() {
	o.slash = false
	o.data = nil
}
