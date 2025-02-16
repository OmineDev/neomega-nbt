//go:build !armbe && !arm64be && !ppc64 && !mips && !mips64 && !mips64p32 && !ppc && !sparc && !sparc64 && !s390 && !s390x

package encoding

import (
	"encoding/binary"
	"math"
	"neomega_nbt/base_io"
	"unsafe"
)

type LittleEndian struct{}

// WriteInt16 ...
func (LittleEndian) WriteInt16(w base_io.Writer, x int16) error {
	b := *(*[2]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteInt16"}
	}
	return nil
}

// WriteInt32 ...
func (LittleEndian) WriteInt32(w base_io.Writer, x int32) error {
	b := *(*[4]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteInt32"}
	}
	return nil
}

// WriteInt64 ...
func (LittleEndian) WriteInt64(w base_io.Writer, x int64) error {
	b := *(*[8]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteInt64"}
	}
	return nil
}

// WriteFloat32 ...
func (LittleEndian) WriteFloat32(w base_io.Writer, x float32) error {
	b := *(*[4]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteFloat32"}
	}
	return nil
}

// WriteFloat64 ...
func (LittleEndian) WriteFloat64(w base_io.Writer, x float64) error {
	b := *(*[8]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteFloat64"}
	}
	return nil
}

// WriteString ...
func (e LittleEndian) WriteString(w base_io.Writer, x string) error {
	if x == "" {
		e.WriteInt16(w, 0)
		return nil
	}
	if err := e.WriteInt16(w, int16(uint16(len(x)))); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	// Use unsafe conversion from a string to a byte slice to prevent copying.
	b := *(*[]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	return nil
}

// Int16 ...
func (LittleEndian) Int16(r base_io.Reader) (int16, error) {
	b := make([]byte, 2)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Int16"}
	}
	return *(*int16)(unsafe.Pointer(&b[0])), nil
}

// Int32 ...
func (LittleEndian) Int32(r base_io.Reader) (int32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Int32"}
	}
	return *(*int32)(unsafe.Pointer(&b[0])), nil
}

// Int64 ...
func (LittleEndian) Int64(r base_io.Reader) (int64, error) {
	b := make([]byte, 8)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float64"}
	}
	return *(*int64)(unsafe.Pointer(&b[0])), nil
}

// Float32 ...
func (LittleEndian) Float32(r base_io.Reader) (float32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float32"}
	}
	return *(*float32)(unsafe.Pointer(&b[0])), nil
}

// Float64 ...
func (LittleEndian) Float64(r base_io.Reader) (float64, error) {
	b := make([]byte, 8)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float64"}
	}
	return *(*float64)(unsafe.Pointer(&b[0])), nil
}

// String ...
func (e LittleEndian) String(r base_io.Reader) (string, error) {
	strLen, err := e.Int16(r)
	if err != nil {
		return "", BufferOverrunError{Op: "String"}
	}
	b := make([]byte, uint16(strLen))
	if _, err := r.Read(b); err != nil {
		return "", BufferOverrunError{Op: "String"}
	}
	return *(*string)(unsafe.Pointer(&b)), nil
}

// // Int32Slice ...
// func (e LittleEndian) Int32Slice(r base_io.Reader) ([]int32, error) {
// 	n, err := e.Int32(r)
// 	if err != nil {
// 		return nil, BufferOverrunError{Op: "Int32Slice"}
// 	}
// 	b := make([]byte, n*4)
// 	if _, err := r.Read(b); err != nil {
// 		return nil, BufferOverrunError{Op: "Int32Slice"}
// 	}
// 	if n == 0 {
// 		return []int32{}, nil
// 	}
// 	return unsafe.Slice((*int32)(unsafe.Pointer(&b[0])), n), nil
// }

// // Int64Slice ...
// func (e LittleEndian) Int64Slice(r base_io.Reader) ([]int64, error) {
// 	n, err := e.Int32(r)
// 	if err != nil {
// 		return nil, BufferOverrunError{Op: "Int64Slice"}
// 	}
// 	b := make([]byte, n*8)
// 	if _, err := r.Read(b); err != nil {
// 		return nil, BufferOverrunError{Op: "Int64Slice"}
// 	}
// 	if n == 0 {
// 		return []int64{}, nil
// 	}
// 	return unsafe.Slice((*int64)(unsafe.Pointer(&b[0])), n), nil
// }

type BigEndian struct{}

// WriteInt16 ...
func (BigEndian) WriteInt16(w base_io.Writer, x int16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(x))
	if _, err := w.Write(b); err != nil {
		return FailedWriteError{Op: "WriteInt16"}
	}
	return nil
}

// WriteInt32 ...
func (BigEndian) WriteInt32(w base_io.Writer, x int32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteInt32"}
	}
	return nil
}

// WriteInt64 ...
func (BigEndian) WriteInt64(w base_io.Writer, x int64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteInt64"}
	}
	return nil
}

// WriteFloat32 ...
func (BigEndian) WriteFloat32(w base_io.Writer, x float32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteFloat32"}
	}
	return nil
}

// WriteFloat64 ...
func (BigEndian) WriteFloat64(w base_io.Writer, x float64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(x))
	if _, err := w.Write(b[:]); err != nil {
		return FailedWriteError{Op: "WriteFloat64"}
	}
	return nil
}

// WriteString ...
func (e BigEndian) WriteString(w base_io.Writer, x string) error {
	if err := e.WriteInt16(w, int16(uint16(len(x)))); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	// Use unsafe conversion from a string to a byte slice to prevent copying.
	b := *(*[]byte)(unsafe.Pointer(&x))
	if _, err := w.Write(b); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	return nil
}

// Int16 ...
func (BigEndian) Int16(r base_io.Reader) (int16, error) {
	b := make([]byte, 2)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Int16"}
	}
	return int16(binary.BigEndian.Uint16(b)), nil
}

// Int32 ...
func (BigEndian) Int32(r base_io.Reader) (int32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Int32"}
	}
	return int32(binary.BigEndian.Uint32(b)), nil
}

// Int64 ...
func (BigEndian) Int64(r base_io.Reader) (int64, error) {
	b := make([]byte, 8)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float64"}
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

// Float32 ...
func (BigEndian) Float32(r base_io.Reader) (float32, error) {
	b := make([]byte, 4)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float32"}
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
}

// Float64 ...
func (BigEndian) Float64(r base_io.Reader) (float64, error) {
	b := make([]byte, 8)
	if _, err := r.Read(b); err != nil {
		return 0, BufferOverrunError{Op: "Float64"}
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
}

// String ...
func (e BigEndian) String(r base_io.Reader) (string, error) {
	strLen, err := e.Int16(r)
	if err != nil {
		return "", BufferOverrunError{Op: "String"}
	}
	b := make([]byte, uint16(strLen))
	if _, err := r.Read(b); err != nil {
		return "", BufferOverrunError{Op: "String"}
	}
	return *(*string)(unsafe.Pointer(&b)), nil
}

// // Int32Slice ...
// func (e BigEndian) Int32Slice(r base_io.Reader) ([]int32, error) {
// 	n, err := e.Int32(r)
// 	if err != nil {
// 		return nil, BufferOverrunError{Op: "Int32Slice"}
// 	}
// 	b := make([]byte, n*4)
// 	if _, err := r.Read(b); err != nil {
// 		return nil, BufferOverrunError{Op: "Int32Slice"}
// 	}
// 	if n == 0 {
// 		return []int32{}, nil
// 	}
// 	// Manually rotate the bytes, so we can just re-interpret this as a slice.
// 	for i := int32(0); i < n; i++ {
// 		off := i * 4
// 		b[off], b[off+3] = b[off+3], b[off]
// 		b[off+1], b[off+2] = b[off+2], b[off+1]
// 	}
// 	return unsafe.Slice((*int32)(unsafe.Pointer(&b[0])), n), nil
// }

// // Int64Slice ...
// func (e BigEndian) Int64Slice(r base_io.Reader) ([]int64, error) {
// 	n, err := e.Int32(r)
// 	if err != nil {
// 		return nil, BufferOverrunError{Op: "Int64Slice"}
// 	}
// 	b := make([]byte, n*8)
// 	if _, err := r.Read(b); err != nil {
// 		return nil, BufferOverrunError{Op: "Int64Slice"}
// 	}
// 	if n == 0 {
// 		return []int64{}, nil
// 	}
// 	// Manually rotate the bytes, so we can just re-interpret this as a slice.
// 	for i := int32(0); i < n; i++ {
// 		off := i * 4
// 		b[off], b[off+7] = b[off+7], b[off]
// 		b[off+1], b[off+6] = b[off+6], b[off+1]
// 		b[off+2], b[off+5] = b[off+5], b[off+2]
// 		b[off+3], b[off+4] = b[off+4], b[off+3]
// 	}
// 	return unsafe.Slice((*int64)(unsafe.Pointer(&b[0])), n), nil
// }
