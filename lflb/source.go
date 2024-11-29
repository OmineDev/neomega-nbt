package lflb

type Source interface {
	This() (b byte, eof bool)
	Next()
	ThisThenNext() (b byte, eof bool)
	IsEof() bool
	Back(n int)
	String() string
}
