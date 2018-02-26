// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-26 23:07:34 01B210                       [zr_win/func_windows.go]
// -----------------------------------------------------------------------------

package win

// # Constants
//
// # System Information Functions
//   FirmwareDate() string
//   FirmwareVersion() string
//   IsWindows() bool
//   OSClass() OSType
//   OSFolder() string
//   SystemVolumeID() uint32
//   VolumeID(path string) uint32
//
// # Registry Function
//   GetRegistryString(key, valueName string) string
//
// # Internal Constants
//
// # Internal Functions
//   getRegistryKey(key string) HKEY
//   getRegistrySubkey(key string) string
//   registryHardwareInfo(valueName string) string
//   windowsVersionInfo() OSVERSIONINFO

import "bytes"       // standard
import "fmt"         // standard
import "unsafe"      // standard
import str "strings" // standard

import "github.com/balacode/zr" // Zircon-Go

// -----------------------------------------------------------------------------
// # Constants

// OSType __
type OSType int

// OSUnknown indicates the operating system is unknown.
const OSUnknown = OSType(0)

// OSLinux indicates the operating system is any version of Linux.
const OSLinux = OSType('L')

// OSWindows indicates the operating system is any version of Windows.
const OSWindows = OSType('W')

// -----------------------------------------------------------------------------
// # System Information Functions

// FirmwareDate __
func FirmwareDate() string {
	if IsWindows() {
		return registryHardwareInfo("SystemBiosDate")
	}
	zr.IMPLEMENT()
	return ""
} //                                                                FirmwareDate

// FirmwareVersion __
func FirmwareVersion() string {
	if IsWindows() {
		return registryHardwareInfo("SystemBiosVersion")
	}
	zr.IMPLEMENT()
	return ""
} //                                                             FirmwareVersion

// IsWindows returns true if the operating system is Windows.
func IsWindows() bool {
	return OSClass() == OSWindows
} //                                                                   IsWindows

// OSClass __
func OSClass() OSType {
	switch windowsVersionInfo().dwPlatformId {
	// ancient Windows platforms e.g. Windows ME
	// (Win32s: could likely be platform 0)
	case 1:
		return OSUnknown
	//
	// Windows-NT based OS, e.g. Windows XP and above
	case 2:
		return OSWindows
	}
	return OSUnknown
} //                                                                     OSClass

// OSFolder __
func OSFolder() string {
	var buffer [MAX_PATH + 1]WCHAR // buffer to receive Windows directory
	var count = GetWindowsDirectory(
		&buffer[0], // buffer for Windows directory
		MAX_PATH+1, // size of buffer
	)
	var ret string
	for i := UINT(0); i < count; i++ {
		ret += string(buffer[i])
	}
	return ret
} //                                                                    OSFolder

// SystemVolumeID __
func SystemVolumeID() uint32 {
	return VolumeID(OSFolder())
} //                                                              SystemVolumeID

// VolumeID __
func VolumeID(path string) uint32 {
	switch len(path) {
	case 0:
		return 0
	// assume a drive letter
	case 1:
		path += ":\\"
	case 2:
		path += "\\"
	}
	var drive = path[:3]
	var ret = DWORD(0)
	var result = GetVolumeInformation(
		drive, // in  LPCTSTR lpRootPathName
		nil,   // out LPTSTR  lpVolumeNameBuffer
		0,     // in  DWORD   nVolumeNameSize
		&ret,  // out LPDWORD lpVolumeSerialNumber
		nil,   // out LPDWORD lpMaximumComponentLength
		nil,   // out LPDWORD lpFileSystemFlags
		nil,   // out LPTSTR  lpFileSystemNameBuffer
		0,     // in  DWORD   nFileSystemNameSize
	)
	if result == FALSE {
		return 0
	}
	return uint32(ret)
} //                                                                    VolumeID

// -----------------------------------------------------------------------------
// # Registry Function

// GetRegistryString reads a string value from
// the named  registry key and value name
// -----------------------------------------------------------------------------
// details: use hasRegistryValue()
//          to test for existence of the key and value name,
//          because this function logs an error if the key is not found
//
// params.: key - string specifying the registry key to be opened,
//                it should begin with one of the following:
//                 HKEY_CLASSES_ROOT,
//                 HKEY_CURRENT_CONFIG,
//                 HKEY_CURRENT_USER,
//                 HKEY_DYN_DATA,
//                 HKEY_LOCAL_MACHINE,
//                 HKEY_PERFORMANCE_DATA,
//                 HKEY_USERS
//
//          valueName - name of registry value to enquire
//
// returns: the stored string value, or a zero-length
//          string if value not found
//
// example: GetRegistryString(
//   `HKEY_LOCAL_MACHINE\SOFTWARE\Microsoft\Windows\CurrentVersion`,
//   "Version")
// -----------------------------------------------------------------------------
func GetRegistryString(key, valueName string) string {
	const erv = ""
	if len(key) == 0 {
		zr.Error("Key is blank")
		return erv
	}
	// prepare a buffer for reading the result
	const BufferSize = 1024
	var ar [BufferSize]WCHAR
	//
	// split registry key and subkey
	var rootKey = getRegistryKey(key)
	var subkey = getRegistrySubkey(key)
	//
	// open the specified registry location
	var openKey = HKEY(0)
	var result = RegOpenKeyEx(
		rootKey,  // hKey       [in]  HKEY
		subkey,   // lpSubKey   [in]  LPCTSTR
		0,        // ulOptions        DWORD
		KEY_READ, // samDesired [in]  REGSAM
		&openKey, // phkResult  [out] PHKEY
	)
	if result != ERROR_SUCCESS {
		zr.Error(result, "failed opening key:^", key,
			"root:", fmt.Sprintf("0x%X", rootKey),
			"subkey:^", subkey)
		return erv
	}
	// read the specified registry value
	var lpData = LPBYTE(unsafe.Pointer(&ar[0]))
	var bufSize = DWORD(BufferSize)
	result = RegQueryValueEx(
		openKey,   // hKey        HKEY
		valueName, // lpValueName string
		nil,       // lpReserved  LPDWORD
		nil,       // lpType      LPDWORD
		lpData,    // lpData      LPBYTE
		&bufSize,  // lpcbData    LPDWORD
	)
	if result != ERROR_SUCCESS && result != ERROR_FILE_NOT_FOUND {
		zr.Error(zr.EFailedReading, "registry key^", key, "value^", valueName)
		//TODO: add SystemErrorName(result) (in GetRegistryString)
		return erv
	}
	// copy the array to a string
	var retBuf = bytes.NewBuffer(make([]byte, 0, bufSize))
	var ws = retBuf.WriteString
	for i := 0; ar[i] != 0; i++ {
		ws(string(ar[i]))
	}
	return retBuf.String()
} //                                                           GetRegistryString

// -----------------------------------------------------------------------------
// # Internal Constants

// rootKeys __
var rootKeys = []struct {
	name string
	key  HKEY
}{
	{"HKEY_CLASSES_ROOT", HKEY_CLASSES_ROOT},
	{"HKEY_CURRENT_CONFIG", HKEY_CURRENT_CONFIG},
	{"HKEY_CURRENT_USER", HKEY_CURRENT_USER},
	{"HKEY_DYN_DATA", HKEY_DYN_DATA},
	{"HKEY_LOCAL_MACHINE", HKEY_LOCAL_MACHINE},
	{"HKEY_PERFORMANCE_DATA", HKEY_PERFORMANCE_DATA},
	{"HKEY_USERS", HKEY_USERS},
}

// SPACES is a string of all white-space characters,
// which includes spaces, tabs, and newline characters.
const SPACES = " \a\b\f\n\r\t\v"

// -----------------------------------------------------------------------------
// # Internal Functions

// getRegistryKey __
func getRegistryKey(key string) HKEY {
	const erv = HKEY(0)
	if len(key) == 0 {
		zr.Error("Key is blank")
		return erv
	}
	for _, iter := range rootKeys {
		if str.HasPrefix(str.ToUpper(key), iter.name) {
			return iter.key
		}
	}
	zr.Error("Invalid key^", key)
	return erv
} //                                                              getRegistryKey

// getRegistrySubkey returns the subkey value
// from the specified registry key string
func getRegistrySubkey(key string) string {
	const erv = ""
	for _, iter := range rootKeys {
		if str.HasPrefix(str.ToUpper(key), iter.name) {
			var pos = len(iter.name)
			var ret = key[pos+1:]
			return ret
		}
	}
	zr.Error("Invalid key^", key)
	return erv
} //                                                           getRegistrySubkey

// registryHardwareInfo __
func registryHardwareInfo(valueName string) string {
	var ret = GetRegistryString(
		`HKEY_LOCAL_MACHINE\HARDWARE\DESCRIPTION\System`,
		valueName,
	)
	ret = str.Trim(ret, SPACES)
	return ret
} //                                                        registryHardwareInfo

// windowsVersionInfo __
func windowsVersionInfo() OSVERSIONINFO {
	var ret OSVERSIONINFO
	ret.dwOSVersionInfoSize = DWORD(unsafe.Sizeof(ret))
	var result = GetVersionEx(&ret) != 0
	if !result {
		//TODO: handle failure to get OS version (in windowsVersionInfo)
	}
	return ret
} //                                                          windowsVersionInfo

//end
