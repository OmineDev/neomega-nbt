package encoding

import (
	"neomega_nbt/base_io"
	"unsafe"
)

type NetworkLittleEndian struct{ LittleEndian }

// WriteInt32 ...
func (NetworkLittleEndian) WriteInt32(w base_io.Writer, x int32) error {
	ux := uint32(x) << 1
	if x < 0 {
		ux = ^ux
	}
	for ux >= 0x80 {
		if err := w.WriteByte(byte(ux) | 0x80); err != nil {
			return FailedWriteError{Op: "WriteInt32"}
		}
		ux >>= 7
	}
	if err := w.WriteByte(byte(ux)); err != nil {
		return FailedWriteError{Op: "WriteInt32"}
	}
	return nil
}

// WriteInt64 ...
func (NetworkLittleEndian) WriteInt64(w base_io.Writer, x int64) error {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	for ux >= 0x80 {
		if err := w.WriteByte(byte(ux) | 0x80); err != nil {
			return FailedWriteError{Op: "WriteInt64"}
		}
		ux >>= 7
	}
	if err := w.WriteByte(byte(ux)); err != nil {
		return FailedWriteError{Op: "WriteInt64"}
	}
	return nil
}

// WriteString ...
func (NetworkLittleEndian) WriteString(w base_io.Writer, x string) error {
	// Netease
	// if len(x) > maxStringSize {
	// 	return InvalidStringError{Off: w.off, N: uint(len(x)), Err: errStringTooLong}
	// }
	ux := uint32(len(x))
	for ux >= 0x80 {
		if err := w.WriteByte(byte(ux) | 0x80); err != nil {
			return FailedWriteError{Op: "WriteString"}
		}
		ux >>= 7
	}
	if err := w.WriteByte(byte(ux)); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	// Use unsafe conversion from a string to a byte slice to prevent copying.
	if _, err := w.Write(*(*[]byte)(unsafe.Pointer(&x))); err != nil {
		return FailedWriteError{Op: "WriteString"}
	}
	return nil
}

// Int32 ...
func (NetworkLittleEndian) Int32(r base_io.Reader) (int32, error) {
	var ux uint32
	for i := uint(0); i < 35; i += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, BufferOverrunError{Op: "Int32"}
		}
		ux |= uint32(b&0x7f) << i
		if b&0x80 == 0 {
			x := int32(ux >> 1)
			if ux&1 != 0 {
				x = ^x
			}
			return x, nil
		}
	}
	return 0, InvalidVarintError{N: 5}
}

// Int64 ...
func (NetworkLittleEndian) Int64(r base_io.Reader) (int64, error) {
	var ux uint64
	for i := uint(0); i < 70; i += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, BufferOverrunError{Op: "Int64"}
		}
		ux |= uint64(b&0x7f) << i
		if b&0x80 == 0 {
			x := int64(ux >> 1)
			if ux&1 != 0 {
				x = ^x
			}
			return x, nil
		}
	}
	return 0, InvalidVarintError{N: 10}
}

// String ...
func (e NetworkLittleEndian) String(r base_io.Reader) (string, error) {
	length, err := e.stringLength(r)
	if err != nil {
		return "", err
	}
	// Netease
	// if length > maxStringSize {
	// 	return "", InvalidStringError{N: uint(length), Err: errStringTooLong}
	// }
	data := make([]byte, length)
	if _, err := r.Read(data); err != nil {
		return "", BufferOverrunError{Op: "String"}
	}
	return *(*string)(unsafe.Pointer(&data)), nil
}

// stringLength reads the length of a string as a varuint32.
func (NetworkLittleEndian) stringLength(r base_io.Reader) (uint32, error) {
	var ux uint32
	for i := uint(0); i < 35; i += 7 {
		b, err := r.ReadByte()
		if err != nil {
			return 0, BufferOverrunError{Op: "StringLength"}
		}
		ux |= uint32(b&0x7f) << i
		if b&0x80 == 0 {
			return ux, nil
		}
	}
	return 0, InvalidVarintError{N: 5}
}

// Int32Slice ...
func (e NetworkLittleEndian) Int32Slice(r base_io.Reader) ([]int32, error) {
	n, err := e.Int32(r)
	if err != nil {
		return nil, BufferOverrunError{Op: "Int32Slice"}
	}
	m := make([]int32, n)
	for i := int32(0); i < n; i++ {
		m[i], err = e.Int32(r)
		if err != nil {
			return nil, BufferOverrunError{Op: "Int32Slice"}
		}
	}
	return m, nil
}

// Int64Slice ...
func (e NetworkLittleEndian) Int64Slice(r base_io.Reader) ([]int64, error) {
	n, err := e.Int32(r)
	if err != nil {
		return nil, BufferOverrunError{Op: "Int64Slice"}
	}
	m := make([]int64, n)
	for i := int32(0); i < n; i++ {
		m[i], err = e.Int64(r)
		if err != nil {
			return nil, BufferOverrunError{Op: "Int64Slice"}
		}
	}
	return m, nil
}
