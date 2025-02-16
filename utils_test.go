package neomega_nbt

import (
	"fmt"
	"math"
	"testing"
)

func TestUtils(t *testing.T) {
	n := map[string]any{"x": int32(100)}
	x := ReadFromNBT(n, "x", math.MaxInt)
	fmt.Println(x)
}
