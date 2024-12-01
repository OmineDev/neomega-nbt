package left_container

import (
	"snbt/lflb"
	"snbt/lflb/sources"
	whitespace "snbt/tokens/white_space"
	"testing"
)

func TestLeftContainer(t *testing.T) {
	seq := `" ' { [ [I; [B; [L; [I `
	leftContainer := &LeftContainerFinity{}
	src := sources.NewBytesSourceFromString(seq)
	assertType := func(tp uint8) {
		leftContainer.Reset()
		lflb.ReadFinity(src, leftContainer)
		if leftContainer.ContainerType != tp {
			t.Logf("want: %v getL %v \n", tp, leftContainer.ContainerType)
			t.FailNow()
		}
		lflb.ReadFinity(src, whitespace.FinityVariyLenWhiteSpace{})
	}
	assertType(1)
	assertType(2)
	assertType(3)
	assertType(4)
	assertType(6)
	assertType(5)
	assertType(7)
	assertType(4)
}
