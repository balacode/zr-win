// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-24 01:44:01 8BDCFD                             [zr_win/module.go]
// -----------------------------------------------------------------------------

// Package win provides Windows API function and
// type wrappers used by the Zirconium library.
package win

import "fmt" // standard

import "github.com/balacode/zr" // Zirconium

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
