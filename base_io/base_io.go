package base_io

import "io"

type Reader interface {
	io.Reader
	ReadByte() (byte, error)
}

type Writer interface {
	io.Writer
	WriteByte(b byte) (err error)
}

type uReader struct {
	io.Reader
	readByte func() (byte, error)
}

func (r *uReader) ReadByte() (byte, error) {
	return r.readByte()
}

// newOffsetReader returns a new offset reader for the io.Reader passed, setting the ReadByte and Next
// functions as appropriate for that particular reader.
func NewReader(r io.Reader) Reader {
	br, ok := r.(Reader)
	if ok {
		return br
	}
	reader := &uReader{Reader: r}
	if byteReader, ok := r.(io.ByteReader); ok {
		reader.readByte = func() (byte, error) {
			return byteReader.ReadByte()
		}
	} else {
		reader.readByte = func() (byte, error) {
			data := make([]byte, 1)
			_, err := io.ReadAtLeast(reader, data, 1)
			return data[0], err
		}
	}
	return reader
}
