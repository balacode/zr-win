// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-23 11:40:22 D44CD8                             [zr-win/module.go]
// -----------------------------------------------------------------------------

// Package win provides Windows API function and
// type wrappers used by the Zircon-Go library.
package win

import "fmt" // standard

import "github.com/balacode/zr" // Zircon-Go

// LB specifies a line break string.
// On Windows it is a pair of CR and LF.
// CR is decimal 13, hex 0D.
// LF is decimal 10, hex 0A.
const LB = "\r\n"

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// VL is zr.VerboseLog() but is used only for debugging.
var VL = zr.VerboseLog

//end
