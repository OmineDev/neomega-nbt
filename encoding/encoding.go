package encoding

import "neomega_nbt/base_io"

// Encoding is an encoding variant of NBT. In general, there are three different encodings of NBT, which are
// all the same except for the way basic types are written.
type WriteEncoding interface {
	WriteInt16(w base_io.Writer, x int16) error
	WriteInt32(w base_io.Writer, x int32) error
	WriteInt64(w base_io.Writer, x int64) error
	WriteFloat32(w base_io.Writer, x float32) error
	WriteFloat64(w base_io.Writer, x float64) error
	WriteString(w base_io.Writer, x string) error
}

type ReadEncoding interface {
	Int16(r base_io.Reader) (int16, error)
	Int32(r base_io.Reader) (int32, error)
	Int64(r base_io.Reader) (int64, error)
	Float32(r base_io.Reader) (float32, error)
	Float64(r base_io.Reader) (float64, error)
	String(r base_io.Reader) (string, error)
	// Int32Slice(r base_io.Reader) ([]int32, error)
	// Int64Slice(r base_io.Reader) ([]int64, error)
}
