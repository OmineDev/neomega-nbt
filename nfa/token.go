package nfa

import (
	"fmt"
)

type TokenType uint8

// Type:
// - Operator:
//   - ()*|&
//
// - Char
const (
	CHAR   = TokenType(0)
	LBR    = TokenType(1)
	OR     = TokenType(2)
	AND    = TokenType(3)
	REPEAT = TokenType(4)
	MARK   = TokenType(5)
	RBR    = TokenType(6)
)

type Token uint32

func (t Token) Type() TokenType {
	return TokenType(t >> 24)
}

func (t Token) IsChar() bool {
	return t.Type() == CHAR
}

func (t Token) Char() Char {
	return Char(t & 0xFFFFFF)
}

func (t Token) Mark() int {
	return int(t & 0xFFFFFF)
}

func (t Token) String() string {
	switch t.Type() {
	case CHAR:
		return t.Char().String()
	case OR:
		return "[|]"
	case AND:
		return "[&]"
	case REPEAT:
		return "[*]"
	case MARK:
		return fmt.Sprintf("[#%v]", t.Mark())
	case LBR:
		return "[(]"
	case RBR:
		return "[)]"
	default:
		panic(t)
	}
}

func NewToken(tokenType TokenType, char Char) Token {
	return Token(uint32(tokenType)<<24 | uint32(char)&0x00FFFFFF)
}

func NewMark(mark int) Token {
	return Token(uint32(MARK)<<24 | uint32(mark)&0x00FFFFFF)
}

func CommonCharToken(char byte) Token {
	return Token(char)
}

var Or = NewToken(OR, 0)
var And = NewToken(AND, 0)
var Repeat = NewToken(REPEAT, 0)
var LBr = NewToken(LBR, 0)
var RBr = NewToken(RBR, 0)

type TokenSeq []Token

func (ts TokenSeq) String() string {
	out := []byte{}
	for _, t := range ts {
		if t.IsChar() && t.Char().IsCommon() {
			out = append(out, t.Char().Common())
		} else {
			out = append(out, []byte(t.String())...)
		}
	}
	return string(out)
}
