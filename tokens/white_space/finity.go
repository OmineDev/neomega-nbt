package whitespace

import "snbt/lflb"

type FinityWhiteSpace struct{}

func (o FinityWhiteSpace) Feed(b byte) lflb.Status {
	if b>>6 == 0 && ((1<<((b<<2)>>2))&4294983168 > 0) {
		return lflb.RET_TERM_OK
	}
	return lflb.RET_TERM_FAIL
}
func (o FinityWhiteSpace) FeedEof() lflb.Status { return lflb.RET_TERM_FAIL }
func (o FinityWhiteSpace) String() string       { return "[white space]" }
func (o FinityWhiteSpace) Reset()               {}

type FinityVariyLenWhiteSpace struct{}

func (o FinityVariyLenWhiteSpace) Feed(b byte) lflb.Status {
	if b>>6 == 0 && ((1<<((b<<2)>>2))&4294983168 > 0) {
		return lflb.RET_FEED_MORE
	}
	return lflb.RET_TERM_OK | (1 << 2)
}
func (o FinityVariyLenWhiteSpace) FeedEof() lflb.Status { return lflb.RET_TERM_OK }
func (o FinityVariyLenWhiteSpace) String() string       { return "[white space*]" }
func (o FinityVariyLenWhiteSpace) Reset()               {}
