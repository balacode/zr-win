// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 14:50:33 4B6091               zr-win/[dll_comdlg32_windows.go]
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unsafe"
)

var (
	comdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	comdlgGetOpenFileNameW = comdlg32.NewProc("GetOpenFileNameW")
	comdlgGetSaveFileNameW = comdlg32.NewProc("GetSaveFileNameW")
)

// GetOpenFileName library: comdlg32.dll
func GetOpenFileName(lpofn *OPENFILENAME) BOOL {
	ret, _, _ := comdlgGetOpenFileNameW.Call(uintptr(unsafe.Pointer(lpofn)))
	return BOOLResult(ret)
} //                                                             GetOpenFileName

// GetSaveFileName library: comdlg32.dll
func GetSaveFileName(lpofn *OPENFILENAME) BOOL {
	ret, _, _ := comdlgGetSaveFileNameW.Call(uintptr(unsafe.Pointer(lpofn)))
	return BOOLResult(ret)
} //                                                             GetSaveFileName

//end
