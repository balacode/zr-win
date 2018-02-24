// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-24 01:44:00 FA5BAB                [zr_win/dll_shell32_windows.go]
// -----------------------------------------------------------------------------

package win

import "syscall" // standard
import "unsafe"  // standard

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
