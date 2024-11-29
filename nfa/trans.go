package nfa

import (
	"fmt"
)

// - Char
//   - common: a,b,c..,0,1,2,+,-,.,*
//   - epsilon (/ )
//   - extend:
//   - - /d: 0,1,2,...,9
//   - - /c: 0,1,2,...,9,a,b,..,z,A,B,..,Z
//   - - /*: Any

type TransitCond [4]uint64

func (tc TransitCond) Accept(b byte) bool {
	return tc[b/64]&(1<<(b%64)) == (1 << (b % 64))
}

func (tc TransitCond) Allow(b byte) TransitCond {
	tc[b/64] |= (1 << (b % 64))
	return tc
}

func (tc TransitCond) Flip() TransitCond {
	return [4]uint64{^tc[0] | ^tc[1] | ^tc[2] | ^tc[3]}
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
