// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                    zr-win/[api_helpers_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package win

// # Call Stack Functions
//   FuncName(callDepth int) string
//   LineNo(callDepth int) int
//
// # Helper Functions
//   BOOLResult(result uintptr) BOOL
//   ErrorName(errNo DWORD) string
//   PWSTRFromString(s string) *WCHAR
//   StringFromPWSTR(s LPWSTR) string
//   UintptrFromString(s *string) uintptr
//   WriteWSTR(ptr *uintptr, s string)
//
// # Internal
//   abort(funcName string, err error)

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

// -----------------------------------------------------------------------------
// # Call Stack Functions

// ShowResultStatus specifies if BOOLResult should print (to the console)
// the result of the last Win32 API function call. Used for debugging.
var ShowResultStatus = false

// FuncName _ _
func FuncName(callDepth int) string {
	programCounter, _, _, _ := runtime.Caller(callDepth)
	funcName := runtime.FuncForPC(programCounter).Name()
	return funcName
} //                                                                    FuncName

// LineNo _ _
func LineNo(callDepth int) int {
	_, _, lineNo, _ := runtime.Caller(callDepth)
	return lineNo
} //                                                                      LineNo

// -----------------------------------------------------------------------------
// # Helper Functions

// BOOLResult converts a function call result to a BOOL value.
func BOOLResult(result uintptr) BOOL {
	if result == FALSE && ShowResultStatus {
		err := GetLastError()
		fmt.Printf("win.%s() == FALSE. GetLastError() == %d - %s\r\n",
			FuncName(2), err, ErrorName(err))
	}
	return BOOL(result)
} //                                                                  BOOLResult

// ErrorName returns the Windows error description given an error number.
// It calls FormatMessage() Win32 API function to get the message text.
func ErrorName(errNo DWORD) string {
	err := errNo
	if err == 0 {
		err = GetLastError()
	}
	if err == 0 {
		return "no error"
	}
	// static buffer for the error description
	const BufferSize = 256
	var buf256 [BufferSize]WCHAR
	// get the error description from the OS
	FormatMessage(
		FORMAT_MESSAGE_FROM_SYSTEM|
			FORMAT_MESSAGE_IGNORE_INSERTS, // dwFlags
		nil, // lpSource
		err, // dwMessageId
		DWORD(MAKELANGID(LANG_NEUTRAL, SUBLANG_DEFAULT)), // dwLanguageID
		LPWSTR(&buf256[0]), // lpBuffer
		BufferSize,         // nSize
		nil,                // arguments
	)
	// build result string
	var retBuf bytes.Buffer
	for _, ch := range buf256 {
		if ch == 0 {
			break
		}
		retBuf.WriteRune(rune(ch))
	}
	return retBuf.String()
} //                                                                   ErrorName

// PWSTRFromString _ _
func PWSTRFromString(s string) *WCHAR {
	if s == "" {
		return (*WCHAR)(nil)
	}
	if strings.Contains(s, "\x00") {
		count := strings.Count(s, "")
		ret := make([]WCHAR, count)
		for i, r := range s {
			ret[i] = WCHAR(r)
		}
		return (*WCHAR)(&ret[0])
	}
	ret, _ := syscall.UTF16PtrFromString(s)
	return (*WCHAR)(ret)
} //                                                             PWSTRFromString

// StringFromPWSTR returns a Go string given a Windows API
// PWSTR (pointer to a wide string).
func StringFromPWSTR(s LPWSTR) string {
	ptr := unsafe.Pointer(s)
	var retBuf bytes.Buffer
	for {
		ch := rune(*((*WCHAR)(ptr)))
		if ch == 0 {
			break
		}
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
		retBuf.WriteRune(ch)
	}
	return retBuf.String()
} //                                                             StringFromPWSTR

// UintptrFromString _ _
func UintptrFromString(s *string) uintptr {
	if *s == "" {
		return 0
	}
	var ret *uint16
	// Some Windows API functions like GetTextExtentPoint32() panic when given
	// a string containing NUL. This block checks & returns the part before NUL.
	zeroAt := strings.Index(*s, "\x00")
	if zeroAt == -1 {
		ret, _ = syscall.UTF16PtrFromString(*s)
		return uintptr(unsafe.Pointer(ret))
	}
	if zeroAt == 0 {
		return 0
	}
	ret, _ = syscall.UTF16PtrFromString((*s)[:zeroAt])
	return uintptr(unsafe.Pointer(ret))
} //                                                           UintptrFromString

// WriteWSTR _ _
func WriteWSTR(ptr *uintptr, s string) {
	for _, r := range s + "\x00" {
		wch := WCHAR(r)
		*(*WCHAR)(unsafe.Pointer(ptr)) = wch
		*ptr += unsafe.Sizeof(wch)
	}
} //                                                                   WriteWSTR

// -----------------------------------------------------------------------------
// # Internal

// abort terminates the running application after printing the
// supplied name of the calling function and an error description.
func abort(funcName string, err error) {
	panic(fmt.Sprintf("%s failed: %v", funcName, err))
} //                                                                       abort

// end
