package whitespace

import (
	"fa/lflb"
	"fa/lflb/lflbops"
	"fa/lflb/sources"
	"testing"
)

func TestWhiteSpace(t *testing.T) {
	seq := "a \t\n\v\f\rb \t\n\v\f\rcd"
	src := sources.NewBytesSourceFromString(seq)

	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	lflb.ReadFinity(src, lflbops.FinityAny{})
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
	lflb.ReadFinity(src, lflbops.FinityAny{})
	if !lflb.ReadFinity(src, FinityVariyLenWhiteSpace{}) {
		t.FailNow()
	}
	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, lflbops.FinityAny{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, FinityVariyLenWhiteSpace{}) {
		t.FailNow()
	}
	if lflb.ReadFinity(src, FinityWhiteSpace{}) {
		t.FailNow()
	}
	if !lflb.ReadFinity(src, lflbops.FinityAny{}) {
		t.FailNow()
	}
}
