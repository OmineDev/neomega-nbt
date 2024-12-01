package left_container

import (
	"snbt/lflb"
)

type LeftContainerFinity struct {
	goBack         int8
	state          uint8
	previousAccept bool

	// left container token specific
	ContainerType uint8 // 1-> "" 2-> '' 3-> {} 4-> [] 5-> [B;] 6-> [I;] 7-> [L;]
}

func (o *LeftContainerFinity) Feed(b byte) lflb.Status {
	o.goBack += 1
	nextMarkTermAccept := LeftContainerFeed(o.state, b)
	if termOrAcc := nextMarkTermAccept & 0x3; termOrAcc != 0 {
		switch termOrAcc {
		case 0x3:
			if nextMarkTermAccept&0x80 == 0 {
				m := uint8(nextMarkTermAccept) >> 2
				o.ContainerType = m + 1
			}
			return lflb.RET_TERM_OK
		case 0x2:
			if o.previousAccept {
				return lflb.RET_TERM_OK | lflb.Status(o.goBack)<<2
			} else {
				return lflb.RET_TERM_FAIL
			}
		case 0x1:
			o.goBack = 0
			o.previousAccept = true
		}
	}

	o.state = uint8(nextMarkTermAccept >> 8)
	if nextMarkTermAccept&0x80 == 0 {
		m := uint8(nextMarkTermAccept) >> 2
		o.ContainerType = m + 1
	}
	return lflb.RET_FEED_MORE
}
func (o *LeftContainerFinity) FeedEof() lflb.Status {
	if o.previousAccept {
		return lflb.RET_TERM_OK | lflb.Status(o.goBack)<<2
	} else {
		return lflb.RET_TERM_FAIL
	}
}

func (o *LeftContainerFinity) String() string {
	return "[left container]"
}
func (o *LeftContainerFinity) Reset() {
	*o = LeftContainerFinity{}
}
