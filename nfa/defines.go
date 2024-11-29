package nfa

import "fmt"

type State int

type CharType uint8

const (
	COMMONCHAR  = CharType(0)
	EPSILONCHAR = CharType(1)
	EXTENDCHAR  = CharType(2)
)

type Char uint32 // 0x00 FF(charType) FF FF

func (c Char) Type() CharType {
	return CharType(c >> 16)
}

func (c Char) IsCommon() bool {
	return c&0x00FF0000 == 0
}

func (c Char) Common() byte {
	return byte(c)
}

func (c Char) IsEpsilon() bool {
	return c&0x00FF0000 == 0x00010000
}

func (c Char) IsExtend() bool {
	return c&0x00FF0000 == 0x00020000
}

func (c Char) Extend() ExtendType {
	return ExtendType(c)
}

func (c Char) String() string {
	if c.IsCommon() {
		return fmt.Sprintf("%v", string([]byte{c.Common()}))
	}
	if c.IsEpsilon() {
		return "{ε}"
	}
	if c.IsExtend() {
		return c.Extend().String()
	}
	panic(c)
}
