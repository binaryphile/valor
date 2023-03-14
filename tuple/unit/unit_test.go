package unit_test

import (
	"fmt"
	"github.com/binaryphile/valor/tuple/unit"
	"unsafe"
)

func Example() {
	fmt.Println(unsafe.Sizeof(unit.Unit))
	// Output: 0
}
