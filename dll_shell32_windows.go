// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 3FBAB0                [zr-win/dll_shell32_windows.go]
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unsafe"
)

var shell32 = syscall.NewLazyDLL("shell32.dll")
var shellDragAcceptFiles = shell32.NewProc("DragAcceptFiles")
var shellDragFinish = shell32.NewProc("DragFinish")
var shellDragQueryFile = shell32.NewProc("DragQueryFileW")

// DragAcceptFiles shell32.dll
func DragAcceptFiles(hWnd HWND, fAccept BOOL) {
	shellDragAcceptFiles.Call(uintptr(hWnd), uintptr(fAccept))
} //                                                             DragAcceptFiles

// DragFinish library: shell32.dll
func DragFinish(hDrop HDROP) {
	shellDragFinish.Call(uintptr(hDrop))
} //                                                                  DragFinish

// DragQueryFile library: shell32.dll
func DragQueryFile(hDrop HDROP, iFile UINT, lpszFile LPTSTR, cch UINT) UINT {
	var ret, _, _ = shellDragQueryFile.Call(
		uintptr(hDrop),
		uintptr(iFile),
		uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch),
	)
	return UINT(ret)
} //                                                               DragQueryFile

//end
