// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                           zr-win/[func_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
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
//   rootKeys = []struct
//
// # Internal Functions
//   getRegistryKey(key string) HKEY
//   getRegistrySubkey(key string) string
//   registryHardwareInfo(valueName string) string
//   windowsVersionInfo() OSVERSIONINFO

import (
	"bytes"
	"fmt"
	"strings"
	"unsafe"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Constants

// OSType _ _
type OSType int

const (
	// OSUnknown indicates the operating system is unknown.
	OSUnknown = OSType(0)

	// OSLinux indicates the operating system is any version of Linux.
	OSLinux = OSType('L')

	// OSWindows indicates the operating system is any version of Windows.
	OSWindows = OSType('W')
)

// -----------------------------------------------------------------------------
// # System Information Functions

// FirmwareDate _ _
func FirmwareDate() string {
	if IsWindows() {
		return registryHardwareInfo("SystemBiosDate")
	}
	zr.IMPLEMENT()
	return ""
} //                                                                FirmwareDate

// FirmwareVersion _ _
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

// OSClass _ _
func OSClass() OSType {
	switch windowsVersionInfo().dwPlatformId {
	// ancient Windows platforms e.g. Windows ME
	// (Win32s: could likely be platform 0)
	case 1:
		{
			return OSUnknown
		}
	// Windows-NT based OS, e.g. Windows XP and above
	case 2:
		return OSWindows
	}
	return OSUnknown
} //                                                                     OSClass

// OSFolder _ _
func OSFolder() string {
	var buffer [MAX_PATH + 1]WCHAR // buffer to receive Windows directory
	count := GetWindowsDirectory(
		&buffer[0], // buffer for Windows directory
		MAX_PATH+1, // size of buffer
	)
	var ret string
	for i := UINT(0); i < count; i++ {
		ret += string(buffer[i])
	}
	return ret
} //                                                                    OSFolder

// SystemVolumeID _ _
func SystemVolumeID() uint32 {
	return VolumeID(OSFolder())
} //                                                              SystemVolumeID

// VolumeID _ _
func VolumeID(path string) uint32 {
	switch len(path) {
	case 0:
		{
			return 0
		}
	// assume a drive letter
	case 1:
		{
			path += ":\\"
		}
	case 2:
		path += "\\"
	}
	var (
		drive  = path[:3]
		ret    = DWORD(0)
		result = GetVolumeInformation(
			drive, // in  LPCTSTR lpRootPathName
			nil,   // out LPTSTR  lpVolumeNameBuffer
			0,     // in  DWORD   nVolumeNameSize
			&ret,  // out LPDWORD lpVolumeSerialNumber
			nil,   // out LPDWORD lpMaximumComponentLength
			nil,   // out LPDWORD lpFileSystemFlags
			nil,   // out LPTSTR  lpFileSystemNameBuffer
			0)     // in  DWORD   nFileSystemNameSize
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
	rootKey := getRegistryKey(key)
	subkey := getRegistrySubkey(key)
	//
	// open the specified registry location
	openKey := HKEY(0)
	result := RegOpenKeyEx(
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
	lpData := LPBYTE(unsafe.Pointer(&ar[0]))
	bufSize := DWORD(BufferSize)
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
		// TODO: add SystemErrorName(result) (in GetRegistryString)
		return erv
	}
	// copy the array to a string
	var (
		retBuf = bytes.NewBuffer(make([]byte, 0, bufSize))
		ws     = retBuf.WriteString
	)
	for i := 0; ar[i] != 0; i++ {
		ws(string(ar[i]))
	}
	return retBuf.String()
} //                                                           GetRegistryString

// -----------------------------------------------------------------------------
// # Internal Constants

// rootKeys _ _
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

// -----------------------------------------------------------------------------
// # Internal Functions

// getRegistryKey _ _
func getRegistryKey(key string) HKEY {
	const erv = HKEY(0)
	if len(key) == 0 {
		zr.Error("Key is blank")
		return erv
	}
	for _, iter := range rootKeys {
		if strings.HasPrefix(strings.ToUpper(key), iter.name) {
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
		if strings.HasPrefix(strings.ToUpper(key), iter.name) {
			pos := len(iter.name)
			ret := key[pos+1:]
			return ret
		}
	}
	zr.Error("Invalid key^", key)
	return erv
} //                                                           getRegistrySubkey

// registryHardwareInfo _ _
func registryHardwareInfo(valueName string) string {
	ret := GetRegistryString(
		`HKEY_LOCAL_MACHINE\HARDWARE\DESCRIPTION\System`,
		valueName,
	)
	ret = strings.TrimSpace(ret)
	return ret
} //                                                        registryHardwareInfo

// windowsVersionInfo _ _
func windowsVersionInfo() OSVERSIONINFO {
	var ret OSVERSIONINFO
	ret.dwOSVersionInfoSize = DWORD(unsafe.Sizeof(ret))
	result := GetVersionEx(&ret) != 0
	if !result {
		// TODO: handle failure to get OS version (in windowsVersionInfo)
	}
	return ret
} //                                                          windowsVersionInfo

// end
