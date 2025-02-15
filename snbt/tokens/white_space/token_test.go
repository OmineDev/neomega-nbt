package whitespace

import (
	"neomega_nbt/base_io/lflb"
	"neomega_nbt/base_io/lflb/sources"
	"neomega_nbt/snbt/tokens"
	"testing"
)

func TestWhiteSpace(t *testing.T) {
	seq := "a \t\n\v\f\rb \t\n\v\f\rcd"
	src := sources.NewBytesSourceFromString(seq)

	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	lflb.ReadFinity(src, tokens.FinityAny{})
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	lflb.ReadFinity(src, tokens.FinityAny{})
	if !lflb.ReadFinity(src, FinityVariyLenWhiteSpace{}) {
		t.FailNow()
	}
	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, tokens.FinityAny{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityVariyLenWhiteSpace{}) {
		t.FailNow()
	}
	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, tokens.FinityAny{}) {
		t.FailNow()
	}
}
