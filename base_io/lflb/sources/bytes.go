package sources

import (
	"bytes"
)

type BytesSource struct {
	data []byte
	p    int
}

func NewBytesSourceFromString(inp string) *BytesSource {
	return &BytesSource{
		data: []byte(inp),
		p:    0,
	}
}

func (bs *BytesSource) Reset() {
	bs.p = 0
}

func (bs *BytesSource) This() (b byte, eof bool) {
	if bs.p == len(bs.data) {
		return 0, true
	}
	b = bs.data[bs.p]
	return b, false
}

func (bs *BytesSource) IsEof() (eof bool) {
	return bs.p == len(bs.data)
}

func (bs *BytesSource) Next() {
	bs.p += 1
}
func (bs *BytesSource) Back(n int) error {
	bs.p -= n
	return nil
}
func (bs *BytesSource) ThisThenNext() (b byte, eof bool) {
	if bs.p == len(bs.data) {
		return 0, true
	}
	b = bs.data[bs.p]
	bs.p += 1
	return b, false
}

func (bs *BytesSource) String() string {
	s, e := bs.p-10, bs.p+10
	if s < 0 {
		s = 0
	}
	if e >= len(bs.data) {
		e = len(bs.data)
	}
	p := []byte("")
	for range bs.p - s {
		p = append(p, ' ')
	}
	asciiStr := bytes.Clone(bs.data[s:e])
	asciiWithUnkownStr := bytes.Clone(bs.data[s:e])
	hasNonAscii := false
	for i, c := range asciiWithUnkownStr {
		if c >= 128 {
			asciiWithUnkownStr[i] = '?'
			hasNonAscii = true
		}
	}
	p = append(p, '^')
	if hasNonAscii {
		return string(asciiWithUnkownStr) + " (" + string(asciiStr) + ")\n" + string(p)
	} else {
		return string(asciiWithUnkownStr) + "\n" + string(p)
	}
}

// type ReadSource struct {
// 	data     []byte
// 	p        int
// 	underlay io.ByteReader
// }

// func NewReadSourceFrom(r io.ByteReader) lflb.Source {
// 	s, ok := r.(lflb.Source)
// 	if ok {
// 		return s
// 	}
// 	return &ReadSource{underlay: r, cache: []byte{}}
// }

// func (bs *ReadSource) ensure() {
// 	if bs.p != len(bs.data) {
// 		return
// 	}
// 	b, err := bs.underlay.ReadByte()
// 	if err != nil {
// 		return
// 	}
// 	bs.data = append(bs.data, b)
// 	if len(bs.data) > 1024 {
// 		bs.data = bs.data[512:]
// 		bs.p -= 512
// 	}
// }

// func (bs *ReadSource) This() (b byte, eof bool) {
// 	bs.ensure()
// 	if bs.p == len(bs.data) {
// 		return 0, true
// 	}
// 	b = bs.data[bs.p]
// 	return b, false
// }

// func (bs *ReadSource) IsEof() (eof bool) {
// 	bs.ensure()
// 	return bs.p == len(bs.data)
// }

// func (bs *ReadSource) Next() {
// 	bs.ensure()
// 	bs.p += 1
// }
// func (bs *BytesSource) Back(n int) error {
// 	bs.p -= n
// 	if bs.p < 0 {
// 		panic()
// 	}
// }
// func (bs *BytesSource) ThisThenNext() (b byte, eof bool) {
// 	if bs.p == len(bs.data) {
// 		return 0, true
// 	}
// 	b = bs.data[bs.p]
// 	bs.p += 1
// 	return b, false
// }

// func (bs *BytesSource) String() string {
// 	s, e := bs.p-10, bs.p+10
// 	if s < 0 {
// 		s = 0
// 	}
// 	if e >= len(bs.data) {
// 		e = len(bs.data)
// 	}
// 	p := []byte("")
// 	for range bs.p - s {
// 		p = append(p, ' ')
// 	}
// 	asciiStr := bytes.Clone(bs.data[s:e])
// 	asciiWithUnkownStr := bytes.Clone(bs.data[s:e])
// 	hasNonAscii := false
// 	for i, c := range asciiWithUnkownStr {
// 		if c >= 128 {
// 			asciiWithUnkownStr[i] = '?'
// 			hasNonAscii = true
// 		}
// 	}
// 	p = append(p, '^')
// 	if hasNonAscii {
// 		return string(asciiWithUnkownStr) + " (" + string(asciiStr) + ")\n" + string(p)
// 	} else {
// 		return string(asciiWithUnkownStr) + "\n" + string(p)
// 	}
// }
