package lflbops

import "fa/lflb"

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
