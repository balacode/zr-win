// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                                 zr-win/[module.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

// Package win provides Windows API function and
// type wrappers used by the Zircon-Go library.
package win

import (
	"fmt"

	"github.com/balacode/zr"
)

var (
	// PL is fmt.Println() but is used only for debugging.
	PL = fmt.Println

	// VL is zr.VerboseLog() but is used only for debugging.
	VL = zr.VerboseLog
)

// end
