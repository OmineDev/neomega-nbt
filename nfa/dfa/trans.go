package dfa

import (
	"fmt"
)

type TransitCond [4]uint64

func (tc TransitCond) Allow(b byte) TransitCond {
	tc[b/64] |= (1 << (b % 64))
	return tc
}

func (tc TransitCond) Union(tc2 TransitCond) TransitCond {
	tc[0] |= tc2[0]
	tc[1] |= tc2[1]
	tc[2] |= tc2[2]
	tc[3] |= tc2[3]
	return tc
}

func (tc TransitCond) String() string {
	return fmt.Sprintf("%X %X %X %X", (tc[3]), (tc[2]), (tc[1]), (tc[0]))
}
