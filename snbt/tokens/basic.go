package tokens

import (
	"fmt"
	"neomega_nbt/base_io/lflb"
)

type FinityNone struct{}

func (o FinityNone) Feed(b byte) lflb.Status { return lflb.RET_TERM_OK | (1 << 2) }
func (o FinityNone) FeedEof() lflb.Status    { return lflb.RET_TERM_OK }
func (o FinityNone) String() string          { return "[none]" }
func (o FinityNone) Reset()                  {}

type FinityAnyOrNone struct{}

func (o FinityAnyOrNone) Feed(b byte) lflb.Status { return lflb.RET_TERM_OK }
func (o FinityAnyOrNone) FeedEof() lflb.Status    { return lflb.RET_TERM_OK }
func (o FinityAnyOrNone) String() string          { return "[any|none]" }
func (o FinityAnyOrNone) Reset()                  {}

type FinityAny struct{}

func (o FinityAny) Feed(b byte) lflb.Status { return lflb.RET_TERM_OK }
func (o FinityAny) FeedEof() lflb.Status    { return lflb.RET_TERM_FAIL }
func (o FinityAny) String() string          { return "[any]" }
func (o FinityAny) Reset()                  {}

type Specific byte

func (o Specific) Feed(b byte) lflb.Status {
	if b == byte(o) {
		return lflb.RET_TERM_OK
	} else {
		return lflb.RET_TERM_FAIL
	}
}
func (o Specific) FeedEof() lflb.Status { return lflb.RET_TERM_FAIL }
func (o Specific) String() string       { return fmt.Sprintf("[%v]", string([]byte{byte(o)})) }
func (o Specific) Reset()               {}
