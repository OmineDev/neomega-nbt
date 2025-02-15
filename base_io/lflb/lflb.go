package lflb

// FeedEof -> TermOK & BackN TermFail
// Feed -> FeedMore TermOK & BackN TermFail

type StatusFlag byte

const (
	TERM = StatusFlag(0x1) // Term or FeedMore
	OK   = StatusFlag(0x2) // when Term Set, Term OK or Term Fail
)

const (
	FEED_MORE = StatusFlag(0)
	TERM_FAIL = TERM
	TERM_OK   = TERM | OK
	FLAG      = 0x3
)

type Status int

const (
	RET_FEED_MORE = Status(FEED_MORE)
	RET_TERM_FAIL = Status(TERM_FAIL)
	RET_TERM_OK   = Status(TERM_OK)
)

type Finity interface {
	Feed(b byte) Status
	FeedEof() Status
	Reset()
}
