// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-24 01:44:00 CDE70B               [zr_win/dll_comdlg32_windows.go]
// -----------------------------------------------------------------------------

package win

import "syscall" // standard
import "unsafe"  // standard

var comdlg32 = syscall.NewLazyDLL("comdlg32.dll")
var comdlgGetOpenFileNameW = comdlg32.NewProc("GetOpenFileNameW")
var comdlgGetSaveFileNameW = comdlg32.NewProc("GetSaveFileNameW")

// GetOpenFileName library: comdlg32.dll
func GetOpenFileName(lpofn *OPENFILENAME) BOOL {
	var ret, _, _ = comdlgGetOpenFileNameW.Call(uintptr(unsafe.Pointer(lpofn)))
	return BOOLResult(ret)
} //                                                             GetOpenFileName

// GetSaveFileName library: comdlg32.dll
func GetSaveFileName(lpofn *OPENFILENAME) BOOL {
	var ret, _, _ = comdlgGetSaveFileNameW.Call(uintptr(unsafe.Pointer(lpofn)))
	return BOOLResult(ret)
} //                                                             GetSaveFileName

//end
