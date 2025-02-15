package encoding

import (
	"fmt"
)

// BufferOverrunError is returned when the data buffer passed in when reading is overrun, meaning one of the
// reading operations extended beyond the end of the slice.
type BufferOverrunError struct {
	Op string
}

// Error ...
func (err BufferOverrunError) Error() string {
	return fmt.Sprintf("nbt: unexpected buffer end during op: '%v'", err.Op)
}

// FailedWriteError is returned if a Write operation failed on an offsetWriter, meaning some of the data could
// not be written to the io.Writer.
type FailedWriteError struct {
	Op  string
	Err error
}

// Error ...
func (err FailedWriteError) Error() string {
	return fmt.Sprintf("nbt: failed write during op '%v': %v", err.Op, err.Err)
}

// InvalidVarintError is returned if a varint(32/64) is encountered that does
// not end after 5 or 10 bytes respectively.
type InvalidVarintError struct {
	N int
}

// Error ...
func (err InvalidVarintError) Error() string {
	return fmt.Sprintf("nbt: varint did not terminate after %v bytes", err.N)
}
