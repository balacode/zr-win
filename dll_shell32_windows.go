// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                    zr-win/[dll_shell32_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unsafe"
)

var (
	shell32 = syscall.NewLazyDLL("shell32.dll")

	shellDragAcceptFiles = shell32.NewProc("DragAcceptFiles")
	shellDragFinish      = shell32.NewProc("DragFinish")
	shellDragQueryFile   = shell32.NewProc("DragQueryFileW")
)

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
	ret, _, _ := shellDragQueryFile.Call(
		uintptr(hDrop),
		uintptr(iFile),
		uintptr(unsafe.Pointer(lpszFile)),
		uintptr(cch),
	)
	return UINT(ret)
} //                                                               DragQueryFile

// end
