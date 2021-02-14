// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                   zr-win/[dll_advapi32_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unsafe"
)

var (
	advapi32 = syscall.NewLazyDLL("advapi32.dll")

	advapiRegOpenKeyEx    = advapi32.NewProc("RegOpenKeyExW")
	advapiRegQueryValueEx = advapi32.NewProc("RegQueryValueExW")
)

// RegOpenKeyEx library: advapi32.dll
func RegOpenKeyEx(
	hKey HKEY,
	lpSubKey string,
	ulOptions DWORD,
	samDesired REGSAM,
	phkResult PHKEY,
) LONG {
	//
	ret, _, _ := advapiRegOpenKeyEx.Call(
		uintptr(hKey),                      // [in]  HKEY
		UintptrFromString(&lpSubKey),       // [in]  LPCTSTR
		uintptr(ulOptions),                 // DWORD
		uintptr(samDesired),                // [in]  REGSAM
		uintptr(unsafe.Pointer(phkResult)), // [out] PHKEY
	)
	return LONG(ret)
} //                                                                RegOpenKeyEx

// RegQueryValueEx library: advapi32.dll
func RegQueryValueEx(
	hKey HKEY,
	lpValueName string,
	lpReserved LPDWORD,
	lpType LPDWORD,
	lpData LPBYTE,
	lpcbData LPDWORD,
) LONG {
	ret, _, _ := advapiRegQueryValueEx.Call(
		uintptr(hKey),                       // [in] HKEY
		UintptrFromString(&lpValueName),     // [in] LPCTSTR
		uintptr(unsafe.Pointer(lpReserved)), // LPDWORD
		uintptr(unsafe.Pointer(lpType)),     // [out] LPDWORD
		uintptr(unsafe.Pointer(lpData)),     // [out] LPBYTE
		uintptr(unsafe.Pointer(lpcbData)),   // [in/out] LPDWORD
	)
	return LONG(ret)
} //                                                             RegQueryValueEx

// end
