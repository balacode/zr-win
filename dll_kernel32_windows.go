// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-24 01:44:00 0968BF               [zr_win/dll_kernel32_windows.go]
// -----------------------------------------------------------------------------

package win

import "syscall" // standard
import "unsafe"  // standard

var kernel32 = syscall.NewLazyDLL("kernel32.dll")

//
var kernelAllocConsole = kernel32.NewProc("AllocConsole")
var kernelCreateFileW = kernel32.NewProc("CreateFileW")
var kernelFindClose = kernel32.NewProc("FindClose")
var kernelFindFirstFileExW = kernel32.NewProc("FindFirstFileExW")
var kernelFindNextFileNameW = kernel32.NewProc("FindNextFileNameW")
var kernelFormatMessageW = kernel32.NewProc("FormatMessageW")
var kernelGetConsoleWindow = kernel32.NewProc("GetConsoleWindow")
var kernelGetFileSize = kernel32.NewProc("GetFileSize")
var kernelGetLastError = kernel32.NewProc("GetLastError")
var kernelGetModuleHandleW = kernel32.NewProc("GetModuleHandleW")
var kernelGetStdHandle = kernel32.NewProc("GetStdHandle")
var kernelGetTickCount = kernel32.NewProc("GetTickCount")
var kernelGetVersionEx = kernel32.NewProc("GetVersionExW")
var kernelGetVolumeInformation = kernel32.NewProc("GetVolumeInformationW")
var kernelGetWindowsDirectory = kernel32.NewProc("GetWindowsDirectoryW")
var kernelGlobalAlloc = kernel32.NewProc("GlobalAlloc")
var kernelGlobalFree = kernel32.NewProc("GlobalFree")
var kernelGlobalLock = kernel32.NewProc("GlobalLock")
var kernelGlobalSize = kernel32.NewProc("GlobalSize")
var kernelGlobalUnlock = kernel32.NewProc("GlobalUnlock")
var kernelMulDiv = kernel32.NewProc("MulDiv")
var kernelMultiByteToWideChar = kernel32.NewProc("MultiByteToWideChar")
var kernelReadFile = kernel32.NewProc("ReadFile")
var kernelWaitForMultipleObjects = kernel32.NewProc("WaitForMultipleObjects")
var kernelWideCharToMultiByte = kernel32.NewProc("WideCharToMultiByte")

var kernelFindCloseChangeNotification = kernel32.NewProc(
	"FindCloseChangeNotification",
)
var kernelFindFirstChangeNotification = kernel32.NewProc(
	"FindFirstChangeNotificationW",
)
var kernelFindNextChangeNotification = kernel32.NewProc(
	"FindNextChangeNotification",
)
var kernelGetConsoleScreenBufferInfoEx = kernel32.NewProc(
	"GetConsoleScreenBufferInfoEx",
)
var kernelSetConsoleScreenBufferSize = kernel32.NewProc(
	"SetConsoleScreenBufferSize",
)

// [UNUSED]
// var kernelWriteConsoleInputW   = kernel32.NewProc("WriteConsoleInputW")

// AllocConsole library: kernel32.dll
func AllocConsole() BOOL {
	var ret, _, _ = kernelAllocConsole.Call()
	return BOOLResult(ret)
} //                                                                AllocConsole

// CreateFile library: kernel32.dll
func CreateFile(
	FileName string,
	dwDesiredAccess DWORD,
	dwShareMode DWORD,
	lpSecurityAttributes *SECURITY_ATTRIBUTES,
	dwCreationDisposition DWORD,
	dwFlagsAndAttributes DWORD,
	hTemplateFile HANDLE,
) HANDLE {
	var ret, _, _ = kernelCreateFileW.Call(
		UintptrFromString(&FileName),                  // [in] LPCTSTR
		uintptr(dwDesiredAccess),                      // [in] DWORD
		uintptr(dwShareMode),                          // [in] DWORD
		uintptr(unsafe.Pointer(lpSecurityAttributes)), // [in] LPSECURITY_ATT...
		uintptr(dwCreationDisposition),                // [in] DWORD
		uintptr(dwFlagsAndAttributes),                 // [in] DWORD
		uintptr(hTemplateFile),                        // [in] HANDLE
	)
	return HANDLE(ret)
} //                                                                  CreateFile

// FindClose library: kernel32.dll
func FindClose(hFindFile HANDLE) BOOL {
	var ret, _, _ = kernelFindClose.Call(uintptr(hFindFile)) // [in/out] HANDLE
	return BOOLResult(ret)
} //                                                                   FindClose

// FindCloseChangeNotification __
func FindCloseChangeNotification(hChangeHandle HANDLE) BOOL {
	var ret, _, _ = kernelFindCloseChangeNotification.Call(
		uintptr(hChangeHandle), // HANDLE
	)
	return BOOLResult(ret)
} //                                                 FindCloseChangeNotification

// FindFirstChangeNotification library: kernel32.dll
func FindFirstChangeNotification(
	lpPathName string,
	bWatchSubtree BOOL,
	dwNotifyFilter DWORD,
) HANDLE {
	var ret, _, _ = kernelFindFirstChangeNotification.Call(
		UintptrFromString(&lpPathName), // LPCSTR lpPathName
		uintptr(bWatchSubtree),         // BOOL bWatchSubtree
		uintptr(dwNotifyFilter),        // DWORD dwNotifyFilter
	)
	return HANDLE(ret)
} //                                                 FindFirstChangeNotification

// FindFirstFileEx library: kernel32.dll
func FindFirstFileEx(
	FileName string,
	fInfoLevelId FINDEX_INFO_LEVELS,
	lpFindFileData LPVOID,
	fSearchOp FINDEX_SEARCH_OPS,
	lpSearchFilter LPVOID,
	dwAdditionalFlags DWORD,
) HANDLE {
	var ret, _, _ = kernelFindFirstFileExW.Call(
		UintptrFromString(&FileName),            // [in] LPCTSTR
		uintptr(fInfoLevelId),                   // [in] FINDEX_INFO_LEVELS
		uintptr(unsafe.Pointer(lpFindFileData)), // [out] LPVOID
		uintptr(fSearchOp),                      // [in] FINDEX_SEARCH_OPS
		uintptr(unsafe.Pointer(lpSearchFilter)), // [in] LPVOID
		uintptr(dwAdditionalFlags),              // [in] DWORD
	)
	return HANDLE(ret)
} //                                                             FindFirstFileEx

// FindNextChangeNotification library: kernel32.dll
func FindNextChangeNotification(hChangeHandle HANDLE) BOOL {
	var ret, _, _ = kernelFindNextChangeNotification.Call(
		uintptr(hChangeHandle), // HANDLE
	)
	return BOOLResult(ret)
} //                                                  FindNextChangeNotification

// FindNextFileName library: kernel32.dll
func FindNextFileName(
	hFindStream HANDLE,
	StringLength LPDWORD,
	LinkName string,
) BOOL {
	var ret, _, _ = kernelFindNextFileNameW.Call(
		uintptr(hFindStream),                  // [in] HANDLE
		uintptr(unsafe.Pointer(StringLength)), // [in/out] LPDWORD
		UintptrFromString(&LinkName),          // [in/out] PWCHAR
	)
	return BOOLResult(ret)
} //                                                            FindNextFileName

// FormatMessage library: kernel32.dll
func FormatMessage(
	dwFlags DWORD,
	lpSource LPCVOID,
	dwMessageId DWORD,
	dwLanguageId DWORD,
	lpBuffer LPWSTR,
	nSize DWORD,
	Arguments *byte,
	//TODO: *va_list   type va_list = *c_char
) DWORD {
	var ret, _, _ = kernelFormatMessageW.Call(
		uintptr(dwFlags),                   // [in] DWORD
		uintptr(unsafe.Pointer(lpSource)),  // [in] LPCVOID
		uintptr(dwMessageId),               // [in] DWORD
		uintptr(dwLanguageId),              // [in] DWORD
		uintptr(unsafe.Pointer(lpBuffer)),  // [out] LPTSTR
		uintptr(nSize),                     // [in] DWORD
		uintptr(unsafe.Pointer(Arguments)), // [in] va_list*
	)
	return DWORD(ret)
} //                                                               FormatMessage

// GetConsoleScreenBufferInfoEx library: kernel32.dll
func GetConsoleScreenBufferInfoEx(
	hConsoleOutput HANDLE,
	lpConsoleScreenBufferInfoEx *CONSOLE_SCREEN_BUFFER_INFOEX,
) BOOL {
	var ret, _, _ = kernelGetConsoleScreenBufferInfoEx.Call(
		uintptr(hConsoleOutput),                              // [in] HANDLE
		uintptr(unsafe.Pointer(lpConsoleScreenBufferInfoEx)), // [out] PCONSO...
	)
	return BOOLResult(ret)
} //                                                GetConsoleScreenBufferInfoEx

// GetConsoleWindow library: kernel32.dll
func GetConsoleWindow() HWND {
	var ret, _, _ = kernelGetConsoleWindow.Call()
	return HWND(ret)
} //                                                            GetConsoleWindow

// GetFileSize library: kernel32.dll
func GetFileSize(
	hFile HANDLE,
	lpFileSizeHigh LPDWORD,
) DWORD {
	var ret, _, _ = kernelGetFileSize.Call(
		uintptr(hFile),                          // [in] HANDLE
		uintptr(unsafe.Pointer(lpFileSizeHigh)), // [out] LPDWORD
	)
	return DWORD(ret)
} //                                                                 GetFileSize

// GetLastError library: kernel32.dll
func GetLastError() DWORD {
	var ret, _, _ = kernelGetLastError.Call()
	return DWORD(ret)
} //                                                                GetLastError

// GetModuleHandle library: kernel32.dll
func GetModuleHandle(lpModuleName string) HMODULE {
	var ptr = uintptr(0)
	if lpModuleName != "" {
		ptr = UintptrFromString(&lpModuleName)
	}
	var ret, _, _ = kernelGetModuleHandleW.Call(ptr) // [in] LPCTSTR
	return HMODULE(ret)
} //                                                             GetModuleHandle

// GetStdHandle library: kernel32.dll
func GetStdHandle(nStdHandle DWORD) HANDLE {
	var ret, _, _ = kernelGetStdHandle.Call(uintptr(nStdHandle))
	return HANDLE(ret)
} //                                                                GetStdHandle

// GetTickCount library: kernel32.dll
func GetTickCount() DWORD {
	var ret, _, _ = kernelGetTickCount.Call()
	return DWORD(ret)
} //                                                                GetTickCount

// GetVersionEx library: kernel32.dll
func GetVersionEx(lpVersionInfo *OSVERSIONINFO) BOOL {
	var ret, _, _ = kernelGetVersionEx.Call(
		uintptr(unsafe.Pointer(lpVersionInfo)), // [in/out] LPOSVERSIONINFO
	)
	return BOOLResult(ret)
} //                                                                GetVersionEx

// GetVolumeInformation library: kernel32.dll
func GetVolumeInformation(
	lpRootPathName string,
	lpVolumeNameBuffer LPTSTR,
	nVolumeNameSize DWORD,
	lpVolumeSerialNumber LPDWORD,
	lpMaximumComponentLength LPDWORD,
	lpFileSystemFlags LPDWORD,
	lpFileSystemNameBuffer LPTSTR,
	nFileSystemNameSize DWORD,
) BOOL {
	var ret, _, _ = kernelGetVolumeInformation.Call(
		UintptrFromString(&lpRootPathName),                // in  LPCTSTR
		uintptr(unsafe.Pointer(lpVolumeNameBuffer)),       // out LPTSTR
		uintptr(nVolumeNameSize),                          // in  DWORD
		uintptr(unsafe.Pointer(lpVolumeSerialNumber)),     // out LPDWORD
		uintptr(unsafe.Pointer(lpMaximumComponentLength)), // out LPDWORD
		uintptr(unsafe.Pointer(lpFileSystemFlags)),        // out LPDWORD
		uintptr(unsafe.Pointer(lpFileSystemNameBuffer)),   // out LPTSTR
		uintptr(nFileSystemNameSize))                      // in  DWORD
	return BOOLResult(ret)
} //                                                        GetVolumeInformation

// GetWindowsDirectory library: kernel32.dll
func GetWindowsDirectory(lpBuffer LPTSTR, uSize UINT) UINT {
	var ret, _, _ = kernelGetWindowsDirectory.Call(
		uintptr(unsafe.Pointer(lpBuffer)), // [out] LPWSTR
		uintptr(uSize),                    // [in]  UINT
	)
	return UINT(ret)
} //                                                         GetWindowsDirectory

// GlobalAlloc library: kernel32.dll
func GlobalAlloc(uFlags UINT, dwBytes SIZE_T) HGLOBAL {
	var ret, _, _ = kernelGlobalAlloc.Call(
		uintptr(uFlags),  // [in] UINT
		uintptr(dwBytes), // [in] SIZE_T
	)
	return HGLOBAL(ret)
} //                                                                 GlobalAlloc

// GlobalFree library: kernel32.dll
func GlobalFree(hMem HGLOBAL) HGLOBAL {
	var ret, _, _ = kernelGlobalFree.Call(uintptr(hMem))
	return HGLOBAL(ret)
} //                                                                  GlobalFree

// GlobalLock library: kernel32.dll
func GlobalLock(hMem HGLOBAL) unsafe.Pointer { // (returns LPVOID)
	var ret, _, _ = kernelGlobalLock.Call(uintptr(hMem))
	if ret == NULL {
		//TODO: mod.Error("GlobalLock failed"
	}
	// turn 'ret' to unsafe.Pointer without 'go vet' triggering a warning:
	var ptr unsafe.Pointer
	ptr = unsafe.Pointer(uintptr(ptr) + uintptr(uint(ret)))
	//
	return unsafe.Pointer(ptr)
} //                                                                  GlobalLock

// GlobalSize library: kernel32.dll
func GlobalSize(hMem HGLOBAL) SIZE_T {
	var ret, _, _ = kernelGlobalSize.Call(uintptr(hMem))
	return SIZE_T(ret)
} //                                                                  GlobalSize

// GlobalUnlock library: kernel32.dll
func GlobalUnlock(hMem HGLOBAL) BOOL {
	var ret, _, _ = kernelGlobalUnlock.Call(uintptr(hMem))
	return BOOLResult(ret)
} //                                                                GlobalUnlock

// MulDiv library: kernel32.dll
func MulDiv(nNumber, nNumerator, nDenominator INT) INT {
	var ret, _, _ = kernelMulDiv.Call(
		uintptr(nNumber),
		uintptr(nNumerator),
		uintptr(nDenominator),
	)
	return INT(ret)
} //                                                                      MulDiv

// MultiByteToWideChar library: kernel32.dll
func MultiByteToWideChar(
	CodePage UINT,
	dwFlags DWORD,
	lpMultiByteStr LPCCH,
	cbMultiByte INT,
	lpWideCharStr LPWSTR,
	cchWideChar INT,
) INT {
	var ret, _, _ = kernelMultiByteToWideChar.Call(
		uintptr(CodePage),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpMultiByteStr)),
		uintptr(cbMultiByte),
		uintptr(unsafe.Pointer(lpWideCharStr)),
		uintptr(cchWideChar),
	)
	return INT(ret)
} //                                                         MultiByteToWideChar

// ReadFile library: kernel32.dll
func ReadFile(
	hFile HANDLE,
	lpBuffer LPVOID,
	nNumberOfBytesToRead DWORD,
	lpNumberOfBytesRead LPDWORD,
	lpOverlapped *OVERLAPPED,
) BOOL {
	var ret, _, _ = kernelReadFile.Call(
		uintptr(hFile),
		uintptr(lpBuffer),
		uintptr(nNumberOfBytesToRead),
		uintptr(unsafe.Pointer(lpNumberOfBytesRead)),
		uintptr(unsafe.Pointer(lpOverlapped)),
	)
	return BOOLResult(ret)
} //                                                                    ReadFile

/*
//TODO: How to pass COORD struct when Call() needs uintptr?
// SetConsoleScreenBufferSize library: kernel32.dll
func SetConsoleScreenBufferSize(hConsoleOutput HANDLE, dwSize COORD) BOOL {
    var ret, _, _ = kernelSetConsoleScreenBufferSize.Call(
        uintptr(hConsoleOutput), // [in] HANDLE
        uintptr(dwSize),         // [in] COORD
    )
    return BOOLResult(ret)
} //                                               SetConsoleScreenBufferSize
*/

// WaitForMultipleObjects library: kernel32.dll
func WaitForMultipleObjects(
	nCount DWORD,
	lpHandles *HANDLE,
	bWaitAll BOOL,
	dwMilliseconds DWORD,
) DWORD {
	var ret, _, _ = kernelWaitForMultipleObjects.Call(
		uintptr(nCount),                    // DWORD
		uintptr(unsafe.Pointer(lpHandles)), // *HANDLE
		uintptr(bWaitAll),                  // BOOL
		uintptr(dwMilliseconds),            // DWORD
	)
	return DWORD(ret)
} //                                                      WaitForMultipleObjects

// WideCharToMultiByte library: kernel32.dll
func WideCharToMultiByte(
	CodePage UINT,
	dwFlags DWORD,
	lpWideCharStr LPCWCH,
	cchWideChar INT,
	lpMultiByteStr LPSTR,
	cbMultiByte INT,
	lpDefaultChar LPCCH,
	lpUsedDefaultChar LPBOOL,
) INT {
	var ret, _, _ = kernelWideCharToMultiByte.Call(
		uintptr(CodePage),
		uintptr(dwFlags),
		uintptr(unsafe.Pointer(lpWideCharStr)),
		uintptr(cchWideChar),
		uintptr(unsafe.Pointer(lpMultiByteStr)),
		uintptr(cbMultiByte),
		uintptr(unsafe.Pointer(lpDefaultChar)),
		uintptr(unsafe.Pointer(lpUsedDefaultChar)),
	)
	return INT(ret)
} //                                                         WideCharToMultiByte

/*
// WriteConsoleInput library: kernel32.dll [UNUSED]
func WriteConsoleInput(
	hConsoleInput HANDLE,
	lpBuffer *INPUT_RECORD,
	nLength DWORD,
	lpNumberOfEventsWritten LPDWORD,
) BOOL {
	var ret, _, _ = kernelWriteConsoleInputW.Call(
		uintptr(hConsoleInput),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nLength),
		uintptr(unsafe.Pointer(lpNumberOfEventsWritten)),
	)
	return BOOLResult(ret)
} //                                                           WriteConsoleInput
*/

//end
