// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-11-29 10:45:07 DC57C3                zr-win/[api_helpers_windows.go]
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
	str "strings"
	"syscall"
	"unsafe"
)

// -----------------------------------------------------------------------------
// # Call Stack Functions

// ShowResultStatus specifies if BOOLResult should print (to the console)
// the result of the last Win32 API function call. Used for debugging.
var ShowResultStatus = false

// FuncName __
func FuncName(callDepth int) string {
	var programCounter, _, _, _ = runtime.Caller(callDepth)
	var funcName = runtime.FuncForPC(programCounter).Name()
	return funcName
} //                                                                    FuncName

// LineNo __
func LineNo(callDepth int) int {
	var _, _, lineNo, _ = runtime.Caller(callDepth)
	return lineNo
} //                                                                      LineNo

// -----------------------------------------------------------------------------
// # Helper Functions

// BOOLResult converts a function call result to a BOOL value.
func BOOLResult(result uintptr) BOOL {
	if result == FALSE && ShowResultStatus {
		var err = GetLastError()
		fmt.Printf("win.%s() == FALSE. GetLastError() == %d - %s"+LB,
			FuncName(2), err, ErrorName(err))
	}
	return BOOL(result)
} //                                                                  BOOLResult

// ErrorName returns the Windows error description given an error number.
// It calls FormatMessage() Win32 API function to get the message text.
func ErrorName(errNo DWORD) string {
	var err = errNo
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

// PWSTRFromString __
func PWSTRFromString(s string) *WCHAR {
	if s == "" {
		return (*WCHAR)(nil)
	}
	if str.Contains(s, "\x00") {
		var count = str.Count(s, "")
		var ret = make([]WCHAR, count)
		for i, r := range s {
			ret[i] = WCHAR(r)
		}
		return (*WCHAR)(&ret[0])
	}
	var ret, _ = syscall.UTF16PtrFromString(s)
	return (*WCHAR)(ret)
} //                                                             PWSTRFromString

// StringFromPWSTR returns a Go string given a Windows API
// PWSTR (pointer to a wide string).
func StringFromPWSTR(s LPWSTR) string {
	var ptr = unsafe.Pointer(s)
	var retBuf bytes.Buffer
	for {
		var ch = rune(*((*WCHAR)(ptr)))
		if ch == 0 {
			break
		}
		ptr = unsafe.Pointer(uintptr(ptr) + 2)
		retBuf.WriteRune(ch)
	}
	return retBuf.String()
} //                                                             StringFromPWSTR

// UintptrFromString __
func UintptrFromString(s *string) uintptr {
	if *s == "" {
		return 0
	}
	var ret *uint16
	// Some Windows API functions like GetTextExtentPoint32() panic when given
	// a string containing NUL. This block checks & returns the part before NUL.
	var zeroAt = str.Index(*s, "\x00")
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

// WriteWSTR __
func WriteWSTR(ptr *uintptr, s string) {
	for _, r := range s + "\x00" {
		var wch = WCHAR(r)
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

//end
