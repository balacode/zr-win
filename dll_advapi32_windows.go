// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-03-23 11:40:22 5FF1F3               [zr-win/dll_advapi32_windows.go]
// -----------------------------------------------------------------------------

package win

import "syscall" // standard
import "unsafe"  // standard

var advapi32 = syscall.NewLazyDLL("advapi32.dll")
var advapiRegOpenKeyEx = advapi32.NewProc("RegOpenKeyExW")
var advapiRegQueryValueEx = advapi32.NewProc("RegQueryValueExW")

// RegOpenKeyEx library: advapi32.dll
func RegOpenKeyEx(
	hKey HKEY,
	lpSubKey string,
	ulOptions DWORD,
	samDesired REGSAM,
	phkResult PHKEY,
) LONG {
	//
	var ret, _, _ = advapiRegOpenKeyEx.Call(
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
	var ret, _, _ = advapiRegQueryValueEx.Call(
		uintptr(hKey),                       // [in] HKEY
		UintptrFromString(&lpValueName),     // [in] LPCTSTR
		uintptr(unsafe.Pointer(lpReserved)), // LPDWORD
		uintptr(unsafe.Pointer(lpType)),     // [out] LPDWORD
		uintptr(unsafe.Pointer(lpData)),     // [out] LPBYTE
		uintptr(unsafe.Pointer(lpcbData)),   // [in/out] LPDWORD
	)
	return LONG(ret)
} //                                                             RegQueryValueEx

//end
