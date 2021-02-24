// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                            zr-win/[api_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package win

// # Constants
// # Simple Types
// # Pointer Types
// # Handle Types
// # Structure Types
// # Other Types
// # Windows Macro Functions
// # Helper Functions
//   WindowMessageName(msg UINT) string

import (
	"unsafe"
)

// -----------------------------------------------------------------------------
// # Constants

// -----------------------------------------------------------------------------
// Button Control Styles (from WinUser.h)

const (
	// BS_DEFPUSHBUTTON Win32 API constant.
	BS_DEFPUSHBUTTON = 0x00000001

	// BS_PUSHBUTTON Win32 API constant.
	BS_PUSHBUTTON = 0x00000000

	// not used:
	// BS_CHECKBOX        = 0x00000002
	// BS_AUTOCHECKBOX    = 0x00000003
	// BS_RADIOBUTTON     = 0x00000004
	// BS_3STATE          = 0x00000005
	// BS_AUTO3STATE      = 0x00000006
	// BS_GROUPBOX        = 0x00000007
	// BS_USERBUTTON      = 0x00000008
	// BS_AUTORADIOBUTTON = 0x00000009
	// BS_PUSHBOX         = 0x0000000A
	// BS_OWNERDRAW       = 0x0000000B
	// BS_TYPEMASK        = 0x0000000F
	// BS_LEFTTEXT        = 0x00000020
	// BS_TEXT            = 0x00000000 // WINVER >= 0x0400
	// BS_ICON            = 0x00000040
	// BS_BITMAP          = 0x00000080
	// BS_LEFT            = 0x00000100
	// BS_RIGHT           = 0x00000200
	// BS_CENTER          = 0x00000300
	// BS_TOP             = 0x00000400
	// BS_BOTTOM          = 0x00000800
	// BS_VCENTER         = 0x00000C00
	// BS_PUSHLIKE        = 0x00001000
	// BS_MULTILINE       = 0x00002000
	// BS_NOTIFY          = 0x00004000
	// BS_FLAT            = 0x00008000
	// BS_RIGHTBUTTON     = BS_LEFTTEXT
)

// -----------------------------------------------------------------------------
// Predefined Clipboard Formats (from WinUser.h)

const (
	// CF_TEXT clipboard format
	CF_TEXT = 1

	// CF_UNICODETEXT clipboard format
	CF_UNICODETEXT = 13

	// not used:
	// CF_BITMAP       = 2
	// CF_METAFILEPICT = 3
	// CF_SYLK         = 4
	// CF_DIF          = 5
	// CF_TIFF         = 6
	// CF_OEMTEXT      = 7
	// CF_DIB          = 8
	// CF_PALETTE      = 9
	// CF_PENDATA      = 10
	// CF_RIFF         = 11
	// CF_WAVE         = 12
	// CF_ENHMETAFILE  = 14
	// CF_HDROP        = 15 // WINVER >= 0x0400
	// CF_LOCALE       = 16
	// CF_DIBV5        = 17 // WINVER >= 0x0500
	// CF_MAX          = 18 // WINVER >= 0x0500
	// for WINVER >= 0x0400 CF_MAX = 17
	// else                 CF_MAX = 15
	// CF_OWNERDISPLAY    = 0x0080
	// CF_DSPTEXT         = 0x0081
	// CF_DSPBITMAP       = 0x0082
	// CF_DSPMETAFILEPICT = 0x0083
	// CF_DSPENHMETAFILE  = 0x008E
	// CF_PRIVATEFIRST    = 0x0200 // 'Private' formats don't get GlobalFree()
	// CF_PRIVATELAST     = 0x02FF
	// CF_GDIOBJFIRST     = 0x0300 // 'GDIOBJ' formats do get DeleteObject()
	// CF_GDIOBJLAST      = 0x03FF
)

// -----------------------------------------------------------------------------
// from WinGDI.h

const (
	// CLIP_DEFAULT_PRECIS Win32 API constant.
	CLIP_DEFAULT_PRECIS = 0

	// not used:
	// CLIP_CHARACTER_PRECIS = 1
	// CLIP_STROKE_PRECIS    = 2
	// CLIP_MASK             = 0xf
	// CLIP_LH_ANGLES        = (1 << 4)
	// CLIP_TT_ALWAYS        = (2 << 4)
	// CLIP_DFA_DISABLE      = (4 << 4) // _WIN32_WINNT >= _WIN32_WINNT_LONGHORN
	// CLIP_EMBEDDED         = (8 << 4)

	// CLR_INVALID Win32 API constant.
	CLR_INVALID = 0xFFFFFFFF
)

// -----------------------------------------------------------------------------
//  Code Page Default Values (from WinNls.h)

const (
	// CP_UTF8 Win32 API constant.
	CP_UTF8 = 65001 // UTF-8 translation

	// not used:
	// CP_ACP        = 0     // default to ANSI code page
	// CP_OEMCP      = 1     // default to OEM  code page
	// CP_MACCP      = 2     // default to MAC  code page
	// CP_THREAD_ACP = 3     // current thread's ANSI code page
	// CP_SYMBOL     = 42    // SYMBOL translations
	// CP_UTF7       = 65000 // UTF-7 translation
)

// -----------------------------------------------------------------------------
// Window class styles

const (
	// CS_DBLCLKS Win32 API constant.
	CS_DBLCLKS = 0x0008

	// CS_HREDRAW Win32 API constant.
	CS_HREDRAW = 0x0002

	// CS_VREDRAW Win32 API constant.
	CS_VREDRAW = 0x0001

	// not used:
	// CS_BYTEALIGNCLIENT = 0x1000
	// CS_BYTEALIGNWINDOW = 0x2000
	// CS_CLASSDC         = 0x0040
	// CS_DROPSHADOW      = 0x00020000
	// CS_GLOBALCLASS     = 0x4000
	// CS_IME             = 0x00010000
	// CS_NOCLOSE         = 0x0200
	// CS_OWNDC           = 0x0020
	// CS_PARENTDC        = 0x0080
	// CS_SAVEBITS        = 0x0800
)

// -----------------------------------------------------------------------------

const (
	// CW_USEDEFAULT Win32 API constant.
	CW_USEDEFAULT = ^0x7fffffff // C: ((int)0x80000000)

	// DC_BRUSH = 18

	// DEFAULT_CHARSET Win32 API constant.
	DEFAULT_CHARSET = 1 // from WinGdi.h

	// DEFAULT_PITCH Win32 API constant.
	DEFAULT_PITCH = 0 // from WinGdi.h
)

// -----------------------------------------------------------------------------
// Dialog Styles (from WinUser.h)

const (
	// DS_MODALFRAME Win32 API constant.
	DS_MODALFRAME = 0x80 // WINVER < 0x0400 Can be combined with WS_CAPTION

	// DS_CENTER Win32 API constant.
	DS_CENTER = 0x0800 // WINVER >= 0x0400

	// not used:
	// DS_ABSALIGN      = 0x01
	// DS_SYSMODAL      = 0x02
	// DS_LOCALEDIT     = 0x20  // Edit items get Local storage.
	// DS_SETFONT       = 0x40  // User specified font for Dlg controls
	// DS_NOIDLEMSG     = 0x100 // WM_ENTERIDLE message will not be sent
	// DS_SETFOREGROUND = 0x200 // not in win3.1
	//
	// WINVER >= 0x0400:
	// DS_3DLOOK       = 0x0004
	// DS_FIXEDSYS     = 0x0008
	// DS_NOFAILCREATE = 0x0010
	// DS_CONTROL      = 0x0400
	// DS_CENTERMOUSE  = 0x1000
	// DS_CONTEXTHELP  = 0x2000
	// DS_SHELLFONT    = (DS_SETFONT | DS_FIXEDSYS)
	//
	// _WIN32_WCE >= 0x0500:
	// DS_USEPIXELS = 0x8000
)

const (
	// DT_LEFT Win32 API constant.
	DT_LEFT = 0x00000000
)

// -----------------------------------------------------------------------------
// Region Flags

const (
	// ERROR Win32 API constant.
	ERROR = 0

	// NULLREGION Win32 API constant.
	NULLREGION = 1

	// SIMPLEREGION Win32 API constant.
	SIMPLEREGION = 2

	// COMPLEXREGION Win32 API constant.
	COMPLEXREGION = 3

	// RGN_ERROR Win32 API constant.
	RGN_ERROR = ERROR
)

// -----------------------------------------------------------------------------

const (
	// ERROR_FILE_NOT_FOUND - The system cannot find the file specified.
	ERROR_FILE_NOT_FOUND = 2

	// ERROR_SUCCESS - The operation completed successfully.
	ERROR_SUCCESS = 0

	// ES_LEFT Win32 API constant.
	ES_LEFT = 0x0000

	// FALSE Win32 API constant.
	FALSE = 0

	// FF_DONTCARE Win32 API constant.
	FF_DONTCARE = (0 << 4) // Don't care or don't know.

	// FILE_ATTRIBUTE_DIRECTORY Win32 API constant.
	FILE_ATTRIBUTE_DIRECTORY = 0x00000010

	// FILE_ATTRIBUTE_NORMAL Win32 API constant.
	FILE_ATTRIBUTE_NORMAL = 0x00000080

	// FILE_SHARE_READ Win32 API constant.
	FILE_SHARE_READ = 0x00000001
)

// -----------------------------------------------------------------------------

const (
	// FILE_NOTIFY_CHANGE_FILE_NAME _ _
	FILE_NOTIFY_CHANGE_FILE_NAME = 0x00000001

	// FILE_NOTIFY_CHANGE_DIR_NAME _ _
	FILE_NOTIFY_CHANGE_DIR_NAME = 0x00000002

	// FILE_NOTIFY_CHANGE_ATTRIBUTES _ _
	FILE_NOTIFY_CHANGE_ATTRIBUTES = 0x00000004

	// FILE_NOTIFY_CHANGE_SIZE _ _
	FILE_NOTIFY_CHANGE_SIZE = 0x00000008

	// FILE_NOTIFY_CHANGE_LAST_WRITE _ _
	FILE_NOTIFY_CHANGE_LAST_WRITE = 0x00000010

	// FILE_NOTIFY_CHANGE_LAST_ACCESS _ _
	FILE_NOTIFY_CHANGE_LAST_ACCESS = 0x00000020

	// FILE_NOTIFY_CHANGE_CREATION _ _
	FILE_NOTIFY_CHANGE_CREATION = 0x00000040

	// FILE_NOTIFY_CHANGE_SECURITY _ _
	FILE_NOTIFY_CHANGE_SECURITY = 0x00000100
)

// -----------------------------------------------------------------------------

const (
	// FindExInfoStandard Win32 API constant.
	FindExInfoStandard = 0

	// FindExInfoBasic Win32 API constant.
	FindExInfoBasic = 1

	// FindExInfoMaxInfoLevel Win32 API constant.
	FindExInfoMaxInfoLevel = 2
)

// -----------------------------------------------------------------------------

const (
	// FindExSearchNameMatch Win32 API constant.
	FindExSearchNameMatch = 0

	// FindExSearchLimitToDirectories Win32 API constant.
	FindExSearchLimitToDirectories = 1

	// FindExSearchLimitToDevices Win32 API constant.
	FindExSearchLimitToDevices = 2

	// FindExSearchMaxSearchOp Win32 API constant.
	FindExSearchMaxSearchOp = 3
)

// -----------------------------------------------------------------------------

const (
	// FORMAT_MESSAGE_FROM_SYSTEM Win32 API constant.
	FORMAT_MESSAGE_FROM_SYSTEM = 0x00001000

	// FORMAT_MESSAGE_IGNORE_INSERTS Win32 API constant.
	FORMAT_MESSAGE_IGNORE_INSERTS = 0x00000200

	// FW_NORMAL Win32 API constant.
	FW_NORMAL = 400

	// GENERIC_READ Win32 API constant.
	GENERIC_READ = 0x80000000

	// GMEM_MOVEABLE Win32 API constant.
	GMEM_MOVEABLE = 0x0002

	// GMEM_ZEROINIT Win32 API constant.
	GMEM_ZEROINIT = 0x0040

	// HGDI_ERROR Win32 API constant.
	HGDI_ERROR = 0xFFFFFFFF
)

// -----------------------------------------------------------------------------

const (
	// HKEY_CLASSES_ROOT Win32 API constant.
	HKEY_CLASSES_ROOT = HKEY(0x80000000)

	// HKEY_CURRENT_USER Win32 API constant.
	HKEY_CURRENT_USER = HKEY(0x80000001)

	// HKEY_LOCAL_MACHINE Win32 API constant.
	HKEY_LOCAL_MACHINE = HKEY(0x80000002)

	// HKEY_USERS Win32 API constant.
	HKEY_USERS = HKEY(0x80000003)

	// HKEY_PERFORMANCE_DATA Win32 API constant.
	HKEY_PERFORMANCE_DATA = HKEY(0x80000004)

	// HKEY_CURRENT_CONFIG Win32 API constant.
	HKEY_CURRENT_CONFIG = HKEY(0x80000005)

	// HKEY_DYN_DATA Win32 API constant.
	HKEY_DYN_DATA = HKEY(0x80000006)
)

// -----------------------------------------------------------------------------

const (
	// HWND_DESKTOP Win32 API constant.
	HWND_DESKTOP = 0

	// HWND_NOTOPMOST Win32 API constant.
	HWND_NOTOPMOST = -2

	// HWND_TOPMOST Win32 API constant.
	HWND_TOPMOST = -1
)

// -----------------------------------------------------------------------------
// Dialog Box Command IDs (from WinUser.h)

const (
	// IDOK Win32 API constant.
	IDOK = 1

	// IDCANCEL Win32 API constant.
	IDCANCEL = 2

	// IDABORT Win32 API constant.
	IDABORT = 3

	// IDRETRY Win32 API constant.
	IDRETRY = 4

	// IDIGNORE Win32 API constant.
	IDIGNORE = 5

	// IDYES Win32 API constant.
	IDYES = 6

	// IDNO Win32 API constant.
	IDNO = 7

	// IDCLOSE Win32 API constant.
	IDCLOSE = 8

	// IDHELP Win32 API constant.
	IDHELP = 9

	// IDTRYAGAIN Win32 API constant.
	IDTRYAGAIN = 10

	// IDCONTINUE Win32 API constant.
	IDCONTINUE = 11

	// IDTIMEOUT Win32 API constant.
	IDTIMEOUT = 32000
)

// -----------------------------------------------------------------------------
// Cursor Constants

const (
	// IDC_ARROW Win32 API constant.
	IDC_ARROW = 32512

	// IDC_IBEAM Win32 API constant.
	IDC_IBEAM = 32513

	// not used:
	// IDC_WAIT        = 32514
	// IDC_CROSS       = 32515
	// IDC_UPARROW     = 32516
	// IDC_SIZE        = 32640
	// IDC_ICON        = 32641
	// IDC_SIZENWSE    = 32642
	// IDC_SIZENESW    = 32643
	// IDC_SIZEWE      = 32644
	// IDC_SIZENS      = 32645
	// IDC_SIZEALL     = 32646
	// IDC_NO          = 32648
	// IDC_HAND        = 32649
	// IDC_APPSTARTING = 32650
	// IDC_HELP        = 32651
)

// -----------------------------------------------------------------------------

const (
	// IDC_STATIC Win32 API constant.
	IDC_STATIC = 0xffffffff

	// IDI_APPLICATION Win32 API constant.
	IDI_APPLICATION = 32512

	// INVALID_HANDLE_VALUE _ _
	INVALID_HANDLE_VALUE = 0xffffffff // TODO: HANDLE(-1)

	// INFINITE _ _
	INFINITE = 0xFFFFFFFF // infinite timeout

	// KEY_EVENT Win32 API constant.
	KEY_EVENT = 0x0001 // Event contains key event record
)

// -----------------------------------------------------------------------------

const (
	// KEY_QUERY_VALUE Windows API constant.
	KEY_QUERY_VALUE = 0x0001

	// KEY_ENUMERATE_SUB_KEYS Win32 API constant.
	KEY_ENUMERATE_SUB_KEYS = 0x0008

	// KEY_NOTIFY Win32 API constant.
	KEY_NOTIFY = 0x0010

	// not used:
	// KEY_SET_VALUE      = 0x0002
	// KEY_CREATE_SUB_KEY = 0x0004
	// KEY_CREATE_LINK    = 0x0020
	// KEY_WOW64_32KEY    = 0x0200
	// KEY_WOW64_64KEY    = 0x0100
	// KEY_WOW64_RES      = 0x0300
)

// -----------------------------------------------------------------------------

const (
	// DELETE Win32 API constant.
	DELETE = 0x00010000

	// READ_CONTROL Win32 API constant.
	READ_CONTROL = 0x00020000

	// WRITE_DAC Win32 API constant.
	WRITE_DAC = 0x00040000

	// WRITE_OWNER Win32 API constant.
	WRITE_OWNER = 0x00080000

	// STANDARD_RIGHTS_REQUIRED Win32 API constant.
	STANDARD_RIGHTS_REQUIRED = 0x000F0000

	// STANDARD_RIGHTS_READ Win32 API constant.
	STANDARD_RIGHTS_READ = READ_CONTROL

	// STANDARD_RIGHTS_WRITE Win32 API constant.
	STANDARD_RIGHTS_WRITE = READ_CONTROL

	// STANDARD_RIGHTS_EXECUTE Win32 API constant.
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL

	// STANDARD_RIGHTS_ALL Win32 API constant.
	STANDARD_RIGHTS_ALL = 0x001F0000

	// SPECIFIC_RIGHTS_ALL Win32 API constant.
	SPECIFIC_RIGHTS_ALL = 0x0000FFFF

	// KEY_READ Win32 API constant.
	KEY_READ = REGSAM(0x20019)

	// LANG_NEUTRAL Win32 API constant.
	LANG_NEUTRAL = 0x00

	// LF_FACESIZE Win32 API constant.
	LF_FACESIZE = 32

	// LOGPIXELSY Win32 API constant.
	LOGPIXELSY = 90 // Logical pixels/inch in Y

	// MAPVK_VK_TO_VSC Win32 API constant.
	MAPVK_VK_TO_VSC = 0

	// MAX_PATH Win32 API constant.
	MAX_PATH = 260
)

// -----------------------------------------------------------------------------
// MessageBox() Flags (from WinUser.h)

const (
	// MB_OK Win32 API constant.
	MB_OK = 0x00000000

	// MB_OKCANCEL Win32 API constant.
	MB_OKCANCEL = 0x00000001

	// MB_ABORTRETRYIGNORE Win32 API constant.
	MB_ABORTRETRYIGNORE = 0x00000002

	// MB_YESNOCANCEL Win32 API constant.
	MB_YESNOCANCEL = 0x00000003

	// MB_YESNO Win32 API constant.
	MB_YESNO = 0x00000004

	// MB_RETRYCANCEL Win32 API constant.
	MB_RETRYCANCEL = 0x00000005

	// MB_CANCELTRYCONTINUE Win32 API constant.
	MB_CANCELTRYCONTINUE = 0x00000006

	// MB_USERICON Win32 API constant.
	MB_USERICON = 0x00000080

	// MB_ICONHAND Win32 API constant.
	MB_ICONHAND = 0x00000010

	// MB_ICONQUESTION Win32 API constant.
	MB_ICONQUESTION = 0x00000020

	// MB_ICONEXCLAMATION Win32 API constant.
	MB_ICONEXCLAMATION = 0x00000030

	// MB_ICONASTERISK Win32 API constant.
	MB_ICONASTERISK = 0x00000040

	// MB_ICONINFORMATION Win32 API constant.
	MB_ICONINFORMATION = MB_ICONASTERISK

	// MB_ICONWARNING Win32 API constant.
	MB_ICONWARNING = MB_ICONEXCLAMATION

	// MB_ICONERROR Win32 API constant.
	MB_ICONERROR = MB_ICONHAND

	// MB_ICONSTOP Win32 API constant.
	MB_ICONSTOP = MB_ICONHAND

	// MB_DEFBUTTON1 Win32 API constant.
	MB_DEFBUTTON1 = 0x00000000

	// MB_DEFBUTTON2 Win32 API constant.
	MB_DEFBUTTON2 = 0x00000100

	// MB_DEFBUTTON3 Win32 API constant.
	MB_DEFBUTTON3 = 0x00000200

	// MB_DEFBUTTON4 Win32 API constant.
	MB_DEFBUTTON4 = 0x00000300
)

// -----------------------------------------------------------------------------

const (
	// MF_CHECKED Win32 API constant.
	MF_CHECKED = 0x00000008

	// MF_POPUP Win32 API constant.
	MF_POPUP = 0x00000010

	// MF_SEPARATOR Win32 API constant.
	MF_SEPARATOR = 0x00000800

	// MF_STRING Win32 API constant.
	MF_STRING = 0x00000000

	// MF_UNCHECKED Win32 API constant.
	MF_UNCHECKED = 0x00000000

	// MFT_STRING Win32 API constant.
	MFT_STRING = 0

	// MIIM_SUBMENU Win32 API constant.
	MIIM_SUBMENU = 0x00000004

	// MM_TEXT Win32 API constant.
	MM_TEXT = 1

	// MSGFLT_ADD Win32 API constant.
	MSGFLT_ADD = 1

	// NO_ERROR Win32 API constant.
	NO_ERROR = 0 // dderror

	// OFN_FILEMUSTEXIST Win32 API constant.
	OFN_FILEMUSTEXIST = 0x00001000

	// OFN_PATHMUSTEXIST Win32 API constant.
	OFN_PATHMUSTEXIST = 0x00000800

	// OPEN_ALWAYS Win32 API constant.
	OPEN_ALWAYS = 4

	// OUT_TT_ONLY_PRECIS Win32 API constant.
	OUT_TT_ONLY_PRECIS = 7
)

// -----------------------------------------------------------------------------

const (
	// NULL _ _
	NULL = 0
)

// -----------------------------------------------------------------------------
// PeekMessage() wRemoveMsg parameters

const (
	// PM_NOREMOVE Win32 API constant.
	PM_NOREMOVE = 0x000

	// PM_REMOVE Win32 API constant.
	PM_REMOVE = 0x001

	// PM_NOYIELD Win32 API constant.
	PM_NOYIELD = 0x002
)

// -----------------------------------------------------------------------------

const (
	// PROOF_QUALITY Win32 API constant.
	PROOF_QUALITY = 2

	// SRCCOPY Win32 API constant.
	SRCCOPY = 0x00CC0020 // dest = source

	// SS_LEFT Win32 API constant.
	SS_LEFT = 0x00000000

	// STD_INPUT_HANDLE Win32 API constant.
	STD_INPUT_HANDLE = -10

	// SUBLANG_DEFAULT Win32 API constant.
	SUBLANG_DEFAULT = 0x01

	// SW_HIDE Win32 API constant.
	SW_HIDE = 0

	// SW_SHOW Win32 API constant.
	SW_SHOW = 5

	// SWP_NOMOVE Win32 API constant.
	SWP_NOMOVE = 0x0002

	// SWP_NOSIZE Win32 API constant.
	SWP_NOSIZE = 0x0001

	// SYNCHRONIZE Win32 API constant.
	SYNCHRONIZE = 0x00100000

	// TRUE Win32 API constant.
	TRUE = 1
)

// -----------------------------------------------------------------------------

const (
	// VK_LBUTTON virtual key
	VK_LBUTTON = 0x01

	// VK_BACK virtual key
	VK_BACK = 0x08

	// VK_SHIFT virtual key
	VK_SHIFT = 0x10

	// VK_CONTROL virtual key
	VK_CONTROL = 0x11

	// VK_MENU virtual key
	VK_MENU = 0x12

	// VK_ESCAPE virtual key
	VK_ESCAPE = 0x1B

	// VK_PRIOR virtual key
	VK_PRIOR = 0x21

	// VK_NEXT virtual key
	VK_NEXT = 0x22

	// VK_END virtual key
	VK_END = 0x23

	// VK_HOME virtual key
	VK_HOME = 0x24

	// VK_LEFT virtual key
	VK_LEFT = 0x25

	// VK_UP virtual key
	VK_UP = 0x26

	// VK_RIGHT virtual key
	VK_RIGHT = 0x27

	// VK_DOWN virtual key
	VK_DOWN = 0x28

	// VK_DELETE virtual key
	VK_DELETE = 0x2E

	// VK_APPS virtual key
	VK_APPS = 0x5D

	// VK_ADD virtual key
	VK_ADD = 0x6B

	// VK_SUBTRACT virtual key
	VK_SUBTRACT = 0x6D

	// VK_F1 virtual key
	VK_F1 = 0x70

	// VK_F2 virtual key
	VK_F2 = 0x71

	// VK_F3 virtual key
	VK_F3 = 0x72

	// VK_F4 virtual key
	VK_F4 = 0x73

	// VK_F5 virtual key
	VK_F5 = 0x74

	// VK_F6 virtual key
	VK_F6 = 0x75

	// VK_F7 virtual key
	VK_F7 = 0x76

	// VK_F8 virtual key
	VK_F8 = 0x77

	// VK_F9 virtual key
	VK_F9 = 0x78

	// VK_F10 virtual key
	VK_F10 = 0x79

	// VK_F11 virtual key
	VK_F11 = 0x7A

	// VK_F12 virtual key
	VK_F12 = 0x7B

	// VK_LSHIFT virtual key
	VK_LSHIFT = 0xA0

	// VK_RSHIFT virtual key
	VK_RSHIFT = 0xA1

	// VK_OEM_PLUS virtual key
	VK_OEM_PLUS = 0xBB

	// VK_OEM_MINUS virtual key
	VK_OEM_MINUS = 0xBD

	// VK_OEM_4 virtual key
	VK_OEM_4 = 0xDB

	// VK_OEM_6 virtual key
	VK_OEM_6 = 0xDD
)

// -----------------------------------------------------------------------------

const (
	// WAIT_OBJECT _ _
	WAIT_OBJECT = 0

	// WAIT_OBJECT_0 _ _
	WAIT_OBJECT_0 = 0

	// WAIT_TIMEOUT _ _
	WAIT_TIMEOUT = 258
)

// -----------------------------------------------------------------------------

const (
	// WHEEL_DELTA Win32 API constant.
	WHEEL_DELTA = 120
)

// -----------------------------------------------------------------------------
// Window messages used by Zircon-Go lib and applications:

const (
	// WM_ACTIVATE Win32 API message constant.
	WM_ACTIVATE = 0x0006

	// WM_CHAR Win32 API message constant.
	WM_CHAR = 0x0102

	// WM_COMMAND Win32 API message constant.
	WM_COMMAND = 0x0111

	// WM_COPYDATA Win32 API message constant.
	WM_COPYDATA = 0x004A

	// WM_CREATE Win32 API message constant.
	WM_CREATE = 0x0001

	// WM_DESTROY Win32 API message constant.
	WM_DESTROY = 0x0002

	// WM_DROPFILES Win32 API message constant.
	WM_DROPFILES = 0x0233

	// WM_INITDIALOG Win32 API message constant.
	WM_INITDIALOG = 0x0110

	// WM_KEYDOWN Win32 API message constant.
	WM_KEYDOWN = 0x0100

	// WM_KILLFOCUS Win32 API message constant.
	WM_KILLFOCUS = 0x0008

	// WM_LBUTTONDOWN Win32 API message constant.
	WM_LBUTTONDOWN = 0x0201

	// WM_LBUTTONUP Win32 API message constant.
	WM_LBUTTONUP = 0x0202

	// WM_MOUSEMOVE Win32 API message constant.
	WM_MOUSEMOVE = 0x0200

	// WM_MOUSEWHEEL Win32 API message constant.
	WM_MOUSEWHEEL = 0x020A

	// WM_MOVE Win32 API message constant.
	WM_MOVE = 0x0003

	// WM_NULL Win32 API message constant.
	WM_NULL = 0x0000

	// WM_PAINT Win32 API message constant.
	WM_PAINT = 0x000F

	// WM_SETFOCUS Win32 API message constant.
	WM_SETFOCUS = 0x0007

	// WM_SETFONT Win32 API message constant.
	WM_SETFONT = 0x0030

	// WM_SIZE Win32 API message constant.
	WM_SIZE = 0x0005

	// All other window messages:

	// WM_ENABLE Win32 API message constant.
	WM_ENABLE = 0x000A

	// WM_SETREDRAW Win32 API message constant.
	WM_SETREDRAW = 0x000B

	// WM_SETTEXT Win32 API message constant.
	WM_SETTEXT = 0x000C

	// WM_GETTEXT Win32 API message constant.
	WM_GETTEXT = 0x000D

	// WM_GETTEXTLENGTH Win32 API message constant.
	WM_GETTEXTLENGTH = 0x000E

	// WM_CLOSE Win32 API message constant.
	WM_CLOSE = 0x0010

	// WM_QUERYENDSESSION Win32 API message constant.
	WM_QUERYENDSESSION = 0x0011

	// WM_QUIT Win32 API message constant.
	WM_QUIT = 0x0012

	// WM_QUERYOPEN Win32 API message constant.
	WM_QUERYOPEN = 0x0013

	// WM_ERASEBKGND Win32 API message constant.
	WM_ERASEBKGND = 0x0014

	// WM_SYSCOLORCHANGE Win32 API message constant.
	WM_SYSCOLORCHANGE = 0x0015

	// WM_ENDSESSION Win32 API message constant.
	WM_ENDSESSION = 0x0016

	// WM_SHOWWINDOW Win32 API message constant.
	WM_SHOWWINDOW = 0x0018

	// WM_WININICHANGE Win32 API message constant.
	WM_WININICHANGE = 0x001A

	// WM_DEVMODECHANGE Win32 API message constant.
	WM_DEVMODECHANGE = 0x001B

	// WM_ACTIVATEAPP Win32 API message constant.
	WM_ACTIVATEAPP = 0x001C

	// WM_FONTCHANGE Win32 API message constant.
	WM_FONTCHANGE = 0x001D

	// WM_TIMECHANGE Win32 API message constant.
	WM_TIMECHANGE = 0x001E

	// WM_CANCELMODE Win32 API message constant.
	WM_CANCELMODE = 0x001F

	// WM_SETCURSOR Win32 API message constant.
	WM_SETCURSOR = 0x0020

	// WM_MOUSEACTIVATE Win32 API message constant.
	WM_MOUSEACTIVATE = 0x0021

	// WM_CHILDACTIVATE Win32 API message constant.
	WM_CHILDACTIVATE = 0x0022

	// WM_QUEUESYNC Win32 API message constant.
	WM_QUEUESYNC = 0x0023

	// WM_GETMINMAXINFO Win32 API message constant.
	WM_GETMINMAXINFO = 0x0024

	// WM_PAINTICON Win32 API message constant.
	WM_PAINTICON = 0x0026

	// WM_ICONERASEBKGND Win32 API message constant.
	WM_ICONERASEBKGND = 0x0027

	// WM_NEXTDLGCTL Win32 API message constant.
	WM_NEXTDLGCTL = 0x0028

	// WM_SPOOLERSTATUS Win32 API message constant.
	WM_SPOOLERSTATUS = 0x002A

	// WM_DRAWITEM Win32 API message constant.
	WM_DRAWITEM = 0x002B

	// WM_MEASUREITEM Win32 API message constant.
	WM_MEASUREITEM = 0x002C

	// WM_DELETEITEM Win32 API message constant.
	WM_DELETEITEM = 0x002D

	// WM_VKEYTOITEM Win32 API message constant.
	WM_VKEYTOITEM = 0x002E

	// WM_CHARTOITEM Win32 API message constant.
	WM_CHARTOITEM = 0x002F

	// WM_GETFONT Win32 API message constant.
	WM_GETFONT = 0x0031

	// WM_SETHOTKEY Win32 API message constant.
	WM_SETHOTKEY = 0x0032

	// WM_GETHOTKEY Win32 API message constant.
	WM_GETHOTKEY = 0x0033

	// WM_QUERYDRAGICON Win32 API message constant.
	WM_QUERYDRAGICON = 0x0037

	// WM_COMPAREITEM Win32 API message constant.
	WM_COMPAREITEM = 0x0039

	// WM_GETOBJECT Win32 API message constant.
	WM_GETOBJECT = 0x003D

	// WM_COMPACTING Win32 API message constant.
	WM_COMPACTING = 0x0041

	// WM_COMMNOTIFY Win32 API message constant. No longer supported.
	WM_COMMNOTIFY = 0x0044

	// WM_WINDOWPOSCHANGING Win32 API message constant.
	WM_WINDOWPOSCHANGING = 0x0046

	// WM_WINDOWPOSCHANGED Win32 API message constant.
	WM_WINDOWPOSCHANGED = 0x0047

	// WM_POWER Win32 API message constant.
	WM_POWER = 0x0048

	// WM_CANCELJOURNAL Win32 API message constant.
	WM_CANCELJOURNAL = 0x004B

	// WM_NOTIFY Win32 API message constant.
	WM_NOTIFY = 0x004E

	// WM_INPUTLANGCHANGEREQUEST Win32 API message constant.
	WM_INPUTLANGCHANGEREQUEST = 0x0050

	// WM_INPUTLANGCHANGE Win32 API message constant.
	WM_INPUTLANGCHANGE = 0x0051

	// WM_TCARD Win32 API message constant.
	WM_TCARD = 0x0052

	// WM_HELP Win32 API message constant.
	WM_HELP = 0x0053

	// WM_USERCHANGED Win32 API message constant.
	WM_USERCHANGED = 0x0054

	// WM_NOTIFYFORMAT Win32 API message constant.
	WM_NOTIFYFORMAT = 0x0055

	// WM_CONTEXTMENU Win32 API message constant.
	WM_CONTEXTMENU = 0x007B

	// WM_STYLECHANGING Win32 API message constant.
	WM_STYLECHANGING = 0x007C

	// WM_STYLECHANGED Win32 API message constant.
	WM_STYLECHANGED = 0x007D

	// WM_DISPLAYCHANGE Win32 API message constant.
	WM_DISPLAYCHANGE = 0x007E

	// WM_GETICON Win32 API message constant.
	WM_GETICON = 0x007F

	// WM_SETICON Win32 API message constant.
	WM_SETICON = 0x0080

	// WM_NCCREATE Win32 API message constant.
	WM_NCCREATE = 0x0081

	// WM_NCDESTROY Win32 API message constant.
	WM_NCDESTROY = 0x0082

	// WM_NCCALCSIZE Win32 API message constant.
	WM_NCCALCSIZE = 0x0083

	// WM_NCHITTEST Win32 API message constant.
	WM_NCHITTEST = 0x0084

	// WM_NCPAINT Win32 API message constant.
	WM_NCPAINT = 0x0085

	// WM_NCACTIVATE Win32 API message constant.
	WM_NCACTIVATE = 0x0086

	// WM_GETDLGCODE Win32 API message constant.
	WM_GETDLGCODE = 0x0087

	// WM_SYNCPAINT Win32 API message constant.
	WM_SYNCPAINT = 0x0088

	// WM_NCMOUSEMOVE Win32 API message constant.
	WM_NCMOUSEMOVE = 0x00A0

	// WM_NCLBUTTONDOWN Win32 API message constant.
	WM_NCLBUTTONDOWN = 0x00A1

	// WM_NCLBUTTONUP Win32 API message constant.
	WM_NCLBUTTONUP = 0x00A2

	// WM_NCLBUTTONDBLCLK Win32 API message constant.
	WM_NCLBUTTONDBLCLK = 0x00A3

	// WM_NCRBUTTONDOWN Win32 API message constant.
	WM_NCRBUTTONDOWN = 0x00A4

	// WM_NCRBUTTONUP Win32 API message constant.
	WM_NCRBUTTONUP = 0x00A5

	// WM_NCRBUTTONDBLCLK Win32 API message constant.
	WM_NCRBUTTONDBLCLK = 0x00A6

	// WM_NCMBUTTONDOWN Win32 API message constant.
	WM_NCMBUTTONDOWN = 0x00A7

	// WM_NCMBUTTONUP Win32 API message constant.
	WM_NCMBUTTONUP = 0x00A8

	// WM_NCMBUTTONDBLCLK Win32 API message constant.
	WM_NCMBUTTONDBLCLK = 0x00A9

	// WM_NCXBUTTONDOWN Win32 API message constant.
	WM_NCXBUTTONDOWN = 0x00AB

	// WM_NCXBUTTONUP Win32 API message constant.
	WM_NCXBUTTONUP = 0x00AC

	// WM_NCXBUTTONDBLCLK Win32 API message constant.
	WM_NCXBUTTONDBLCLK = 0x00AD

	// WM_INPUT_DEVICE_CHANGE Win32 API message constant.
	WM_INPUT_DEVICE_CHANGE = 0x00FE

	// WM_INPUT Win32 API message constant.
	WM_INPUT = 0x00FF

	// WM_KEYFIRST Win32 API message constant.
	WM_KEYFIRST = 0x0100

	// WM_KEYUP Win32 API message constant.
	WM_KEYUP = 0x0101

	// WM_DEADCHAR Win32 API message constant.
	WM_DEADCHAR = 0x0103

	// WM_SYSKEYDOWN Win32 API message constant.
	WM_SYSKEYDOWN = 0x0104

	// WM_SYSKEYUP Win32 API message constant.
	WM_SYSKEYUP = 0x0105

	// WM_SYSCHAR Win32 API message constant.
	WM_SYSCHAR = 0x0106

	// WM_SYSDEADCHAR Win32 API message constant.
	WM_SYSDEADCHAR = 0x0107

	// WM_UNICHAR Win32 API message constant.
	WM_UNICHAR = 0x0109

	// WM_KEYLAST Win32 API message constant.
	WM_KEYLAST = 0x0109 // depends on Windows version

	// WM_IME_STARTCOMPOSITION Win32 API message constant.
	WM_IME_STARTCOMPOSITION = 0x010D

	// WM_IME_ENDCOMPOSITION Win32 API message constant.
	WM_IME_ENDCOMPOSITION = 0x010E

	// WM_IME_COMPOSITION Win32 API message constant.
	WM_IME_COMPOSITION = 0x010F

	// WM_IME_KEYLAST Win32 API message constant.
	WM_IME_KEYLAST = 0x010F

	// WM_SYSCOMMAND Win32 API message constant.
	WM_SYSCOMMAND = 0x0112

	// WM_TIMER Win32 API message constant.
	WM_TIMER = 0x0113

	// WM_HSCROLL Win32 API message constant.
	WM_HSCROLL = 0x0114

	// WM_VSCROLL Win32 API message constant.
	WM_VSCROLL = 0x0115

	// WM_INITMENU Win32 API message constant.
	WM_INITMENU = 0x0116

	// WM_INITMENUPOPUP Win32 API message constant.
	WM_INITMENUPOPUP = 0x0117

	// WM_GESTURE Win32 API message constant.
	WM_GESTURE = 0x0119

	// WM_GESTURENOTIFY Win32 API message constant.
	WM_GESTURENOTIFY = 0x011A

	// WM_MENUSELECT Win32 API message constant.
	WM_MENUSELECT = 0x011F

	// WM_MENUCHAR Win32 API message constant.
	WM_MENUCHAR = 0x0120

	// WM_ENTERIDLE Win32 API message constant.
	WM_ENTERIDLE = 0x0121

	// WM_MENURBUTTONUP Win32 API message constant.
	WM_MENURBUTTONUP = 0x0122

	// WM_MENUDRAG Win32 API message constant.
	WM_MENUDRAG = 0x0123

	// WM_MENUGETOBJECT Win32 API message constant.
	WM_MENUGETOBJECT = 0x0124

	// WM_UNINITMENUPOPUP Win32 API message constant.
	WM_UNINITMENUPOPUP = 0x0125

	// WM_MENUCOMMAND Win32 API message constant.
	WM_MENUCOMMAND = 0x0126

	// WM_CHANGEUISTATE Win32 API message constant.
	WM_CHANGEUISTATE = 0x0127

	// WM_UPDATEUISTATE Win32 API message constant.
	WM_UPDATEUISTATE = 0x0128

	// WM_QUERYUISTATE Win32 API message constant.
	WM_QUERYUISTATE = 0x0129

	// WM_CTLCOLORMSGBOX Win32 API message constant.
	WM_CTLCOLORMSGBOX = 0x0132

	// WM_CTLCOLOREDIT Win32 API message constant.
	WM_CTLCOLOREDIT = 0x0133

	// WM_CTLCOLORLISTBOX Win32 API message constant.
	WM_CTLCOLORLISTBOX = 0x0134

	// WM_CTLCOLORBTN Win32 API message constant.
	WM_CTLCOLORBTN = 0x0135

	// WM_CTLCOLORDLG Win32 API message constant.
	WM_CTLCOLORDLG = 0x0136

	// WM_CTLCOLORSCROLLBAR Win32 API message constant.
	WM_CTLCOLORSCROLLBAR = 0x0137

	// WM_CTLCOLORSTATIC Win32 API message constant.
	WM_CTLCOLORSTATIC = 0x0138

	// WM_MOUSEFIRST Win32 API message constant.
	WM_MOUSEFIRST = 0x0200

	// WM_LBUTTONDBLCLK Win32 API message constant.
	WM_LBUTTONDBLCLK = 0x0203

	// WM_RBUTTONDOWN Win32 API message constant.
	WM_RBUTTONDOWN = 0x0204

	// WM_RBUTTONUP Win32 API message constant.
	WM_RBUTTONUP = 0x0205

	// WM_RBUTTONDBLCLK Win32 API message constant.
	WM_RBUTTONDBLCLK = 0x0206

	// WM_MBUTTONDOWN Win32 API message constant.
	WM_MBUTTONDOWN = 0x0207

	// WM_MBUTTONUP Win32 API message constant.
	WM_MBUTTONUP = 0x0208

	// WM_MBUTTONDBLCLK Win32 API message constant.
	WM_MBUTTONDBLCLK = 0x0209

	// WM_XBUTTONDOWN Win32 API message constant.
	WM_XBUTTONDOWN = 0x020B

	// WM_XBUTTONUP Win32 API message constant.
	WM_XBUTTONUP = 0x020C

	// WM_XBUTTONDBLCLK Win32 API message constant.
	WM_XBUTTONDBLCLK = 0x020D

	// WM_MOUSEHWHEEL Win32 API message constant.
	WM_MOUSEHWHEEL = 0x020E

	// WM_MOUSELAST Win32 API message constant.
	WM_MOUSELAST = 0x020E // depends on Windows version

	// WM_PARENTNOTIFY Win32 API message constant.
	WM_PARENTNOTIFY = 0x0210

	// WM_ENTERMENULOOP Win32 API message constant.
	WM_ENTERMENULOOP = 0x0211

	// WM_EXITMENULOOP Win32 API message constant.
	WM_EXITMENULOOP = 0x0212

	// WM_NEXTMENU Win32 API message constant.
	WM_NEXTMENU = 0x0213

	// WM_SIZING Win32 API message constant.
	WM_SIZING = 0x0214

	// WM_CAPTURECHANGED Win32 API message constant.
	WM_CAPTURECHANGED = 0x0215

	// WM_MOVING Win32 API message constant.
	WM_MOVING = 0x0216

	// WM_POWERBROADCAST Win32 API message constant.
	WM_POWERBROADCAST = 0x0218

	// WM_DEVICECHANGE Win32 API message constant.
	WM_DEVICECHANGE = 0x0219

	// WM_MDICREATE Win32 API message constant.
	WM_MDICREATE = 0x0220

	// WM_MDIDESTROY Win32 API message constant.
	WM_MDIDESTROY = 0x0221

	// WM_MDIACTIVATE Win32 API message constant.
	WM_MDIACTIVATE = 0x0222

	// WM_MDIRESTORE Win32 API message constant.
	WM_MDIRESTORE = 0x0223

	// WM_MDINEXT Win32 API message constant.
	WM_MDINEXT = 0x0224

	// WM_MDIMAXIMIZE Win32 API message constant.
	WM_MDIMAXIMIZE = 0x0225

	// WM_MDITILE Win32 API message constant.
	WM_MDITILE = 0x0226

	// WM_MDICASCADE Win32 API message constant.
	WM_MDICASCADE = 0x0227

	// WM_MDIICONARRANGE Win32 API message constant.
	WM_MDIICONARRANGE = 0x0228

	// WM_MDIGETACTIVE Win32 API message constant.
	WM_MDIGETACTIVE = 0x0229

	// WM_MDISETMENU Win32 API message constant.
	WM_MDISETMENU = 0x0230

	// WM_ENTERSIZEMOVE Win32 API message constant.
	WM_ENTERSIZEMOVE = 0x0231

	// WM_EXITSIZEMOVE Win32 API message constant.
	WM_EXITSIZEMOVE = 0x0232

	// WM_MDIREFRESHMENU Win32 API message constant.
	WM_MDIREFRESHMENU = 0x0234

	// WM_TOUCH Win32 API message constant.
	WM_TOUCH = 0x0240

	// WM_IME_SETCONTEXT Win32 API message constant.
	WM_IME_SETCONTEXT = 0x0281

	// WM_IME_NOTIFY Win32 API message constant.
	WM_IME_NOTIFY = 0x0282

	// WM_IME_CONTROL Win32 API message constant.
	WM_IME_CONTROL = 0x0283

	// WM_IME_COMPOSITIONFULL Win32 API message constant.
	WM_IME_COMPOSITIONFULL = 0x0284

	// WM_IME_SELECT Win32 API message constant.
	WM_IME_SELECT = 0x0285

	// WM_IME_CHAR Win32 API message constant.
	WM_IME_CHAR = 0x0286

	// WM_IME_REQUEST Win32 API message constant.
	WM_IME_REQUEST = 0x0288

	// WM_IME_KEYDOWN Win32 API message constant.
	WM_IME_KEYDOWN = 0x0290

	// WM_IME_KEYUP Win32 API message constant.
	WM_IME_KEYUP = 0x0291

	// WM_NCMOUSEHOVER Win32 API message constant.
	WM_NCMOUSEHOVER = 0x02A0

	// WM_MOUSEHOVER Win32 API message constant.
	WM_MOUSEHOVER = 0x02A1

	// WM_NCMOUSELEAVE Win32 API message constant.
	WM_NCMOUSELEAVE = 0x02A2

	// WM_MOUSELEAVE Win32 API message constant.
	WM_MOUSELEAVE = 0x02A3

	// WM_WTSSESSION_CHANGE Win32 API message constant.
	WM_WTSSESSION_CHANGE = 0x02B1

	// WM_TABLET_FIRST Win32 API message constant.
	WM_TABLET_FIRST = 0x02c0

	// WM_TABLET_LAST Win32 API message constant.
	WM_TABLET_LAST = 0x02df

	// WM_CUT Win32 API message constant.
	WM_CUT = 0x0300

	// WM_COPY Win32 API message constant.
	WM_COPY = 0x0301

	// WM_PASTE Win32 API message constant.
	WM_PASTE = 0x0302

	// WM_CLEAR Win32 API message constant.
	WM_CLEAR = 0x0303

	// WM_UNDO Win32 API message constant.
	WM_UNDO = 0x0304

	// WM_RENDERFORMAT Win32 API message constant.
	WM_RENDERFORMAT = 0x0305

	// WM_RENDERALLFORMATS Win32 API message constant.
	WM_RENDERALLFORMATS = 0x0306

	// WM_DESTROYCLIPBOARD Win32 API message constant.
	WM_DESTROYCLIPBOARD = 0x0307

	// WM_DRAWCLIPBOARD Win32 API message constant.
	WM_DRAWCLIPBOARD = 0x0308

	// WM_PAINTCLIPBOARD Win32 API message constant.
	WM_PAINTCLIPBOARD = 0x0309

	// WM_VSCROLLCLIPBOARD Win32 API message constant.
	WM_VSCROLLCLIPBOARD = 0x030A

	// WM_SIZECLIPBOARD Win32 API message constant.
	WM_SIZECLIPBOARD = 0x030B

	// WM_ASKCBFORMATNAME Win32 API message constant.
	WM_ASKCBFORMATNAME = 0x030C

	// WM_CHANGECBCHAIN Win32 API message constant.
	WM_CHANGECBCHAIN = 0x030D

	// WM_HSCROLLCLIPBOARD Win32 API message constant.
	WM_HSCROLLCLIPBOARD = 0x030E

	// WM_QUERYNEWPALETTE Win32 API message constant.
	WM_QUERYNEWPALETTE = 0x030F

	// WM_PALETTEISCHANGING Win32 API message constant.
	WM_PALETTEISCHANGING = 0x0310

	// WM_PALETTECHANGED Win32 API message constant.
	WM_PALETTECHANGED = 0x0311

	// WM_HOTKEY Win32 API message constant.
	WM_HOTKEY = 0x0312

	// WM_PRINT Win32 API message constant.
	WM_PRINT = 0x0317

	// WM_PRINTCLIENT Win32 API message constant.
	WM_PRINTCLIENT = 0x0318

	// WM_APPCOMMAND Win32 API message constant.
	WM_APPCOMMAND = 0x0319

	// WM_THEMECHANGED Win32 API message constant.
	WM_THEMECHANGED = 0x031A

	// WM_CLIPBOARDUPDATE Win32 API message constant.
	WM_CLIPBOARDUPDATE = 0x031D

	// WM_DWMCOMPOSITIONCHANGED Win32 API message constant.
	WM_DWMCOMPOSITIONCHANGED = 0x031E

	// WM_DWMNCRENDERINGCHANGED Win32 API message constant.
	WM_DWMNCRENDERINGCHANGED = 0x031F

	// WM_DWMCOLORIZATIONCOLORCHANGED Win32 API message constant.
	WM_DWMCOLORIZATIONCOLORCHANGED = 0x0320

	// WM_DWMWINDOWMAXIMIZEDCHANGE Win32 API message constant.
	WM_DWMWINDOWMAXIMIZEDCHANGE = 0x0321

	// WM_DWMSENDICONICTHUMBNAIL Win32 API message constant.
	WM_DWMSENDICONICTHUMBNAIL = 0x0323

	// WM_DWMSENDICONICLIVEPREVIEWBITMAP Win32 API message constant.
	WM_DWMSENDICONICLIVEPREVIEWBITMAP = 0x0326

	// WM_GETTITLEBARINFOEX Win32 API message constant.
	WM_GETTITLEBARINFOEX = 0x033F

	// WM_HANDHELDFIRST Win32 API message constant.
	WM_HANDHELDFIRST = 0x0358

	// WM_HANDHELDLAST Win32 API message constant.
	WM_HANDHELDLAST = 0x035F

	// WM_AFXFIRST Win32 API message constant.
	WM_AFXFIRST = 0x0360

	// WM_AFXLAST Win32 API message constant.
	WM_AFXLAST = 0x037F

	// WM_PENWINFIRST Win32 API message constant.
	WM_PENWINFIRST = 0x0380

	// WM_PENWINLAST Win32 API message constant.
	WM_PENWINLAST = 0x038F

	// WM_USER Win32 API message constant.
	WM_USER = 0x0400

	// WM_APP Win32 API message constant.
	WM_APP = 0x8000

	// WM_SETTINGCHANGE Win32 API message constant.
	WM_SETTINGCHANGE = WM_WININICHANGE
)

// -----------------------------------------------------------------------------

const (
	// WS_OVERLAPPED Window Style (Win32)
	WS_OVERLAPPED = 0x00000000

	// WS_EX_ACCEPTFILES Window Style (Win32)
	WS_EX_ACCEPTFILES = 0x00000010

	// WS_TABSTOP Window Style (Win32)
	WS_TABSTOP = 0x00010000

	// WS_MAXIMIZEBOX Window Style (Win32)
	WS_MAXIMIZEBOX = 0x00010000

	// WS_MINIMIZEBOX Window Style (Win32)
	WS_MINIMIZEBOX = 0x00020000

	// WS_THICKFRAME Window Style (Win32)
	WS_THICKFRAME = 0x00040000

	// WS_SYSMENU Window Style (Win32)
	WS_SYSMENU = 0x00080000

	// WS_BORDER Window Style (Win32)
	WS_BORDER = 0x00800000

	// WS_CAPTION Window Style (Win32)
	WS_CAPTION = 0x00C00000

	// WS_VISIBLE Window Style (Win32)
	WS_VISIBLE = 0x10000000

	// WS_CHILD Window Style (Win32)
	WS_CHILD = 0x40000000

	// WS_POPUP Window Style (Win32)
	WS_POPUP = 0x80000000

	// WS_OVERLAPPEDWINDOW Window Style (Win32)
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU |
		WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
)

// -----------------------------------------------------------------------------
// # Simple Types

type (
	// ATOM type
	ATOM uint16

	// BOOL type
	BOOL int32

	// BYTE type
	BYTE byte

	// CHAR type
	CHAR byte

	// COLORREF type
	COLORREF uint32

	// DWORD type
	DWORD uint32

	// INT type
	INT int32

	// LANGID type
	LANGID WORD

	// LONG type
	LONG int32

	// SHORT type
	SHORT uint16

	// UINT type
	UINT uint32

	// ULONG type
	ULONG uint32

	// WCHAR type
	WCHAR uint16

	// WORD type
	WORD uint16
)

// -----------------------------------------------------------------------------
// # Pointer Types

type (
	// LPARAM type
	LPARAM uintptr

	// WPARAM type
	WPARAM uintptr

	// DLGPROC type
	DLGPROC uintptr

	// INT_PTR type
	INT_PTR uintptr

	// LONG_PTR type
	LONG_PTR uintptr

	// LPBOOL type
	LPBOOL *BOOL

	// LPBYTE type
	LPBYTE *DWORD

	// LPCCH type
	LPCCH *CHAR

	// LPCSTR type
	LPCSTR *CHAR

	// LPCTSTR type
	LPCTSTR *WCHAR

	// LPCVOID type
	LPCVOID unsafe.Pointer

	// LPCWCH type
	LPCWCH *WCHAR

	// LPCWSTR type
	LPCWSTR *WCHAR

	// LPDWORD type
	LPDWORD *DWORD

	// LPSTR type
	LPSTR *CHAR // always use Unicode

	// LPTSTR type
	LPTSTR *WCHAR

	// LPVOID type
	LPVOID unsafe.Pointer

	// LPWSTR type
	LPWSTR *WCHAR

	// LRESULT type
	LRESULT LONG_PTR

	// PHKEY type
	PHKEY *HKEY

	// PWSTR type
	PWSTR *WCHAR

	// SIZE_T type
	SIZE_T ULONG_PTR

	// UINT_PTR type
	UINT_PTR uintptr

	// ULONG_PTR type
	ULONG_PTR uintptr

	// WNDPROC type
	WNDPROC uintptr
)

// -----------------------------------------------------------------------------
// # Handle Types

type (
	// HANDLE handle type
	HANDLE uintptr

	// HACCEL handle type (Accelerator)
	HACCEL HANDLE

	// HBITMAP handle type (GDI Bitmap)
	HBITMAP HANDLE

	// HBRUSH handle type (GDI Brush)
	HBRUSH HANDLE

	// HCURSOR handle type (Cursor)
	HCURSOR HANDLE

	// HDC handle type (Device Context)
	HDC HANDLE

	// HDROP handle type
	HDROP HANDLE

	// HFONT handle type
	HFONT HANDLE

	// HGDIOBJ handle type
	HGDIOBJ HANDLE

	// HGLOBAL handle type
	HGLOBAL HANDLE

	// HICON handle type
	HICON HANDLE

	// HINSTANCE handle type
	HINSTANCE HANDLE

	// HKEY handle type
	HKEY HANDLE

	// HMENU handle type
	HMENU HANDLE

	// HMODULE handle type
	HMODULE HANDLE

	// HRGN handle type
	HRGN HANDLE

	// HWND handle type
	HWND HANDLE
)

// -----------------------------------------------------------------------------
// # Other Types

type (
	// ACCESS_MASK Win32 API type.
	ACCESS_MASK DWORD

	// FINDEX_INFO_LEVELS Win32 API type.
	FINDEX_INFO_LEVELS int32

	// FINDEX_SEARCH_OPS Win32 API type.
	FINDEX_SEARCH_OPS int32

	// HOOKPROC Win32 API type.
	HOOKPROC func(int, WPARAM, LPARAM) LRESULT

	// REGSAM Win32 API type. Requested Key access mask type.
	REGSAM ACCESS_MASK
)

// -----------------------------------------------------------------------------
// # Structure Types

type (
	// CONSOLE_SCREEN_BUFFER_INFOEX structure
	CONSOLE_SCREEN_BUFFER_INFOEX struct {
		CbSize               ULONG
		DwSize               COORD
		DwCursorPosition     COORD
		WAttributes          WORD
		SrWindow             SMALL_RECT
		DwMaximumWindowSize  COORD
		WPopupAttributes     WORD
		BFullscreenSupported BOOL
		ColorTable           [16]COLORREF
	} //                                            CONSOLE_SCREEN_BUFFER_INFOEX

	// COORD structure
	COORD struct {
		X SHORT
		Y SHORT
	} //                                                                   COORD

	// DLGITEMTEMPLATE structure
	DLGITEMTEMPLATE struct {
		Style           DWORD
		DwExtendedStyle DWORD
		X               SHORT
		Y               SHORT
		Cx              SHORT
		Cy              SHORT
		Id              WORD
	} //                                                         DLGITEMTEMPLATE

	// DLGTEMPLATE structure
	DLGTEMPLATE struct {
		Style           DWORD
		DwExtendedStyle DWORD
		Cdit            WORD
		X               SHORT
		Y               SHORT
		Cx              SHORT
		Cy              SHORT
	} //                                                             DLGTEMPLATE

	// FILETIME structure
	FILETIME struct {
		DwLowDateTime  DWORD
		DwHighDateTime DWORD
	} //                                                                FILETIME

	// INPUT_RECORD structure
	INPUT_RECORD struct {
		EventType WORD
		Event     [4]uint32
	} //                                                            INPUT_RECORD

	// LOGFONT structure
	LOGFONT struct {
		LfHeight         LONG
		LfWidth          LONG
		LfEscapement     LONG
		LfOrientation    LONG
		LfWeight         LONG
		LfItalic         BYTE
		LfUnderline      BYTE
		LfStrikeOut      BYTE
		LfCharSet        BYTE
		LfOutPrecision   BYTE
		LfClipPrecision  BYTE
		LfQuality        BYTE
		LfPitchAndFamily BYTE
		LfFaceName       [LF_FACESIZE]WCHAR
	} //                                                                 LOGFONT

	// MENUITEMINFO structure
	MENUITEMINFO struct {
		CbSize        UINT
		FMask         UINT
		FType         UINT
		FState        UINT
		WID           UINT
		HSubMenu      HMENU
		HbmpChecked   HBITMAP
		HbmpUnchecked HBITMAP
		DwItemData    ULONG_PTR
		DwTypeData    LPWSTR
		Cch           UINT
		HbmpItem      HBITMAP
	} //                                                            MENUITEMINFO

	// MSG structure
	MSG struct {
		HWnd    HWND
		Message UINT
		WParam  WPARAM
		LParam  LPARAM
		Time    DWORD
		Pt      POINT
	} //                                                                     MSG

	// OPENFILENAME structure
	OPENFILENAME struct {
		LStructSize       DWORD
		HwndOwner         HWND
		HInstance         HINSTANCE
		LpstrFilter       LPCWSTR
		LpstrCustomFilter LPWSTR
		NMaxCustFilter    DWORD
		NFilterIndex      DWORD
		LpstrFile         LPWSTR
		NMaxFile          DWORD
		LpstrFileTitle    LPWSTR
		NMaxFileTitle     DWORD
		LpstrInitialDir   LPCWSTR
		LpstrTitle        LPCWSTR
		Flags             DWORD
		NFileOffset       WORD
		NFileExtension    WORD
		LpstrDefExt       LPCWSTR
		LCustData         LPARAM
		LpfnHook          *HOOKPROC
		LpTemplateName    LPCWSTR
		// if _WIN32_WINNT >= 0x0500
		//    pvReserved  LPVOID
		//    dwReserved  DWORD
		//    FlagsEx     DWORD
		// endif
	} //                                                            OPENFILENAME

	// OSVERSIONINFO structure
	OSVERSIONINFO struct {
		dwOSVersionInfoSize DWORD
		dwMajorVersion      DWORD
		dwMinorVersion      DWORD
		dwBuildNumber       DWORD
		dwPlatformId        DWORD
		szCSDVersion        [128]WCHAR
	} //                                                           OSVERSIONINFO

	// OVERLAPPED structure
	OVERLAPPED struct {
		Internal     ULONG_PTR
		InternalHigh ULONG_PTR
		Offset       DWORD
		OffsetHigh   DWORD
		HEvent       HANDLE
	} //                                                              OVERLAPPED

	// PAINTSTRUCT structure
	PAINTSTRUCT struct {
		HDC         HDC
		FErase      BOOL
		RcPaint     RECT
		FRestore    BOOL
		FIncUpdate  BOOL
		RgbReserved [32]BYTE
	} //                                                             PAINTSTRUCT

	// POINT structure
	POINT struct {
		X LONG
		Y LONG
	} //                                                                   POINT

	// RECT structure
	RECT struct {
		Left   LONG
		Top    LONG
		Right  LONG
		Bottom LONG
	} //                                                                    RECT

	// SECURITY_ATTRIBUTES structure
	SECURITY_ATTRIBUTES struct {
		NLength              DWORD
		LpSecurityDescriptor LPVOID
		BInheritHandle       BOOL
	} //                                                     SECURITY_ATTRIBUTES

	// SIZE structure
	SIZE struct {
		Cx int32
		Cy int32
	} //                                                                    SIZE

	// SMALL_RECT structure
	SMALL_RECT struct {
		Left   int16
		Top    int16
		Right  int16
		Bottom int16
	} //                                                              SMALL_RECT

	// TEXTMETRIC structure contains information about a font
	TEXTMETRIC struct {
		TmHeight           LONG
		TmAscent           LONG
		TmDescent          LONG
		TmInternalLeading  LONG
		TmExternalLeading  LONG
		TmAveCharWidth     LONG
		TmMaxCharWidth     LONG
		TmWeight           LONG
		TmOverhang         LONG
		TmDigitizedAspectX LONG
		TmDigitizedAspectY LONG
		TmFirstChar        WCHAR
		TmLastChar         WCHAR
		TmDefaultChar      WCHAR
		TmBreakChar        WCHAR
		TmItalic           BYTE
		TmUnderlined       BYTE
		TmStruckOut        BYTE
		TmPitchAndFamily   BYTE
		TmCharSet          BYTE
	} //                                                              TEXTMETRIC

	// WIN32_FIND_DATA structure
	WIN32_FIND_DATA struct {
		DwFileAttributes   DWORD
		FtCreationTime     FILETIME
		FtLastAccessTime   FILETIME
		FtLastWriteTime    FILETIME
		NFileSizeHigh      DWORD
		NFileSizeLow       DWORD
		DwReserved0        DWORD
		DwReserved1        DWORD
		CFileName          [MAX_PATH]WCHAR
		CAlternateFileName [14]WCHAR
	} //                                                         WIN32_FIND_DATA

	// WNDCLASSEX structure
	WNDCLASSEX struct {
		CbSize        UINT
		Style         UINT
		LpfnWndProc   WNDPROC
		CbClsExtra    int32
		CbWndExtra    int32
		HInstance     HINSTANCE
		HIcon         HICON
		HCursor       HCURSOR
		HbrBackground HBRUSH
		LpszMenuName  LPCWSTR
		LpszClassName LPCWSTR
		HIconSm       HICON
	} //                                                              WNDCLASSEX
)

// -----------------------------------------------------------------------------
// # Windows Macro Functions

// DialogBox Win32 API macro.
func DialogBox(
	hInstance HINSTANCE,
	TemplateName string,
	hWndParent HWND,
	lpDialogFunc DLGPROC,
) INT_PTR {
	ret, _, _ := userDialogBoxParamW.Call(
		uintptr(hInstance),
		UintptrFromString(&TemplateName),
		uintptr(hWndParent),
		uintptr(lpDialogFunc),
		0,
	)
	return INT_PTR(ret)
} //                                                                   DialogBox

// DialogBoxIndirect Win32 API macro.
func DialogBoxIndirect(
	hInstance HINSTANCE,
	lpTemplate *DLGTEMPLATE,
	hWndParent HWND,
	lpDialogFunc DLGPROC,
) INT_PTR {
	return DialogBoxIndirectParam(
		hInstance,
		lpTemplate,
		hWndParent,
		lpDialogFunc,
		0,
	)
} //                                                           DialogBoxIndirect

// GET_WHEEL_DELTA_WPARAM Win32 API macro.
func GET_WHEEL_DELTA_WPARAM(wParam WPARAM) int16 {
	return int16(HIWORD(uint32(wParam)))
} //                                                      GET_WHEEL_DELTA_WPARAM

// HIWORD Win32 API macro.
func HIWORD(val uint32) uint16 {
	return uint16(val >> 16 & 0xffff)
} //                                                                      HIWORD

// LOWORD Win32 API macro.
func LOWORD(val uint32) uint16 {
	return uint16(val)
} //                                                                      LOWORD

// MAKEINTRESOURCE Win32 API macro.
func MAKEINTRESOURCE(id WORD) *WCHAR {
	// turn 'id' to unsafe.Pointer without 'go vet' triggering a warning:
	var ptr unsafe.Pointer
	ptr = unsafe.Pointer(uintptr(ptr) + uintptr(uint(id)))
	return (*WCHAR)(ptr)
} //                                                             MAKEINTRESOURCE

// MAKELANGID Win32 API macro.
func MAKELANGID(p, s WORD) LANGID {
	return LANGID(s<<10 | p)
} //                                                                  MAKELANGID

// RGB Win32 API macro.
func RGB(red, green, blue int) COLORREF {
	return COLORREF(uint32(red) | uint32(green)<<8 | uint32(blue)<<16)
} //                                                                         RGB

// -----------------------------------------------------------------------------
// # Helper Functions

// WindowMessageName returns a string describing a window message.
// For example WindowMessageName(0x0006) returns "WM_ACTIVATE".
func WindowMessageName(msg UINT) string {
	messages := map[UINT]string{
		WM_ACTIVATE:    "WM_ACTIVATE",
		WM_CHAR:        "WM_CHAR",
		WM_COMMAND:     "WM_COMMAND",
		WM_COPYDATA:    "WM_COPYDATA",
		WM_CREATE:      "WM_CREATE",
		WM_DESTROY:     "WM_DESTROY",
		WM_DROPFILES:   "WM_DROPFILES",
		WM_INITDIALOG:  "WM_INITDIALOG",
		WM_KEYDOWN:     "WM_KEYDOWN",
		WM_KILLFOCUS:   "WM_KILLFOCUS",
		WM_LBUTTONDOWN: "WM_LBUTTONDOWN",
		WM_LBUTTONUP:   "WM_LBUTTONUP",
		WM_MOUSEMOVE:   "WM_MOUSEMOVE",
		WM_MOUSEWHEEL:  "WM_MOUSEWHEEL",
		WM_MOVE:        "WM_MOVE",
		WM_NULL:        "WM_NULL",
		WM_PAINT:       "WM_PAINT",
		WM_SETFOCUS:    "WM_SETFOCUS",
		WM_SETFONT:     "WM_SETFONT",
		WM_SIZE:        "WM_SIZE",
		// other messages:
		WM_ENABLE:                         "WM_ENABLE",
		WM_SETREDRAW:                      "WM_SETREDRAW",
		WM_SETTEXT:                        "WM_SETTEXT",
		WM_GETTEXT:                        "WM_GETTEXT",
		WM_GETTEXTLENGTH:                  "WM_GETTEXTLENGTH",
		WM_CLOSE:                          "WM_CLOSE",
		WM_QUERYENDSESSION:                "WM_QUERYENDSESSION",
		WM_QUIT:                           "WM_QUIT",
		WM_QUERYOPEN:                      "WM_QUERYOPEN",
		WM_ERASEBKGND:                     "WM_ERASEBKGND",
		WM_SYSCOLORCHANGE:                 "WM_SYSCOLORCHANGE",
		WM_ENDSESSION:                     "WM_ENDSESSION",
		WM_SHOWWINDOW:                     "WM_SHOWWINDOW",
		WM_WININICHANGE:                   "WM_WININICHANGE",
		WM_DEVMODECHANGE:                  "WM_DEVMODECHANGE",
		WM_ACTIVATEAPP:                    "WM_ACTIVATEAPP",
		WM_FONTCHANGE:                     "WM_FONTCHANGE",
		WM_TIMECHANGE:                     "WM_TIMECHANGE",
		WM_CANCELMODE:                     "WM_CANCELMODE",
		WM_SETCURSOR:                      "WM_SETCURSOR",
		WM_MOUSEACTIVATE:                  "WM_MOUSEACTIVATE",
		WM_CHILDACTIVATE:                  "WM_CHILDACTIVATE",
		WM_QUEUESYNC:                      "WM_QUEUESYNC",
		WM_GETMINMAXINFO:                  "WM_GETMINMAXINFO",
		WM_PAINTICON:                      "WM_PAINTICON",
		WM_ICONERASEBKGND:                 "WM_ICONERASEBKGND",
		WM_NEXTDLGCTL:                     "WM_NEXTDLGCTL",
		WM_SPOOLERSTATUS:                  "WM_SPOOLERSTATUS",
		WM_DRAWITEM:                       "WM_DRAWITEM",
		WM_MEASUREITEM:                    "WM_MEASUREITEM",
		WM_DELETEITEM:                     "WM_DELETEITEM",
		WM_VKEYTOITEM:                     "WM_VKEYTOITEM",
		WM_CHARTOITEM:                     "WM_CHARTOITEM",
		WM_GETFONT:                        "WM_GETFONT",
		WM_SETHOTKEY:                      "WM_SETHOTKEY",
		WM_GETHOTKEY:                      "WM_GETHOTKEY",
		WM_QUERYDRAGICON:                  "WM_QUERYDRAGICON",
		WM_COMPAREITEM:                    "WM_COMPAREITEM",
		WM_GETOBJECT:                      "WM_GETOBJECT",
		WM_COMPACTING:                     "WM_COMPACTING",
		WM_COMMNOTIFY:                     "WM_COMMNOTIFY",
		WM_WINDOWPOSCHANGING:              "WM_WINDOWPOSCHANGING",
		WM_WINDOWPOSCHANGED:               "WM_WINDOWPOSCHANGED",
		WM_POWER:                          "WM_POWER",
		WM_CANCELJOURNAL:                  "WM_CANCELJOURNAL",
		WM_NOTIFY:                         "WM_NOTIFY",
		WM_INPUTLANGCHANGEREQUEST:         "WM_INPUTLANGCHANGEREQUEST",
		WM_INPUTLANGCHANGE:                "WM_INPUTLANGCHANGE",
		WM_TCARD:                          "WM_TCARD",
		WM_HELP:                           "WM_HELP",
		WM_USERCHANGED:                    "WM_USERCHANGED",
		WM_NOTIFYFORMAT:                   "WM_NOTIFYFORMAT",
		WM_CONTEXTMENU:                    "WM_CONTEXTMENU",
		WM_STYLECHANGING:                  "WM_STYLECHANGING",
		WM_STYLECHANGED:                   "WM_STYLECHANGED",
		WM_DISPLAYCHANGE:                  "WM_DISPLAYCHANGE",
		WM_GETICON:                        "WM_GETICON",
		WM_SETICON:                        "WM_SETICON",
		WM_NCCREATE:                       "WM_NCCREATE",
		WM_NCDESTROY:                      "WM_NCDESTROY",
		WM_NCCALCSIZE:                     "WM_NCCALCSIZE",
		WM_NCHITTEST:                      "WM_NCHITTEST",
		WM_NCPAINT:                        "WM_NCPAINT",
		WM_NCACTIVATE:                     "WM_NCACTIVATE",
		WM_GETDLGCODE:                     "WM_GETDLGCODE",
		WM_SYNCPAINT:                      "WM_SYNCPAINT",
		WM_NCMOUSEMOVE:                    "WM_NCMOUSEMOVE",
		WM_NCLBUTTONDOWN:                  "WM_NCLBUTTONDOWN",
		WM_NCLBUTTONUP:                    "WM_NCLBUTTONUP",
		WM_NCLBUTTONDBLCLK:                "WM_NCLBUTTONDBLCLK",
		WM_NCRBUTTONDOWN:                  "WM_NCRBUTTONDOWN",
		WM_NCRBUTTONUP:                    "WM_NCRBUTTONUP",
		WM_NCRBUTTONDBLCLK:                "WM_NCRBUTTONDBLCLK",
		WM_NCMBUTTONDOWN:                  "WM_NCMBUTTONDOWN",
		WM_NCMBUTTONUP:                    "WM_NCMBUTTONUP",
		WM_NCMBUTTONDBLCLK:                "WM_NCMBUTTONDBLCLK",
		WM_NCXBUTTONDOWN:                  "WM_NCXBUTTONDOWN",
		WM_NCXBUTTONUP:                    "WM_NCXBUTTONUP",
		WM_NCXBUTTONDBLCLK:                "WM_NCXBUTTONDBLCLK",
		WM_INPUT_DEVICE_CHANGE:            "WM_INPUT_DEVICE_CHANGE",
		WM_INPUT:                          "WM_INPUT",
		WM_KEYUP:                          "WM_KEYUP",
		WM_DEADCHAR:                       "WM_DEADCHAR",
		WM_SYSKEYDOWN:                     "WM_SYSKEYDOWN",
		WM_SYSKEYUP:                       "WM_SYSKEYUP",
		WM_SYSCHAR:                        "WM_SYSCHAR",
		WM_SYSDEADCHAR:                    "WM_SYSDEADCHAR",
		WM_UNICHAR:                        "WM_UNICHAR",
		WM_IME_STARTCOMPOSITION:           "WM_IME_STARTCOMPOSITION",
		WM_IME_ENDCOMPOSITION:             "WM_IME_ENDCOMPOSITION",
		WM_IME_COMPOSITION:                "WM_IME_COMPOSITION",
		WM_SYSCOMMAND:                     "WM_SYSCOMMAND",
		WM_TIMER:                          "WM_TIMER",
		WM_HSCROLL:                        "WM_HSCROLL",
		WM_VSCROLL:                        "WM_VSCROLL",
		WM_INITMENU:                       "WM_INITMENU",
		WM_INITMENUPOPUP:                  "WM_INITMENUPOPUP",
		WM_GESTURE:                        "WM_GESTURE",
		WM_GESTURENOTIFY:                  "WM_GESTURENOTIFY",
		WM_MENUSELECT:                     "WM_MENUSELECT",
		WM_MENUCHAR:                       "WM_MENUCHAR",
		WM_ENTERIDLE:                      "WM_ENTERIDLE",
		WM_MENURBUTTONUP:                  "WM_MENURBUTTONUP",
		WM_MENUDRAG:                       "WM_MENUDRAG",
		WM_MENUGETOBJECT:                  "WM_MENUGETOBJECT",
		WM_UNINITMENUPOPUP:                "WM_UNINITMENUPOPUP",
		WM_MENUCOMMAND:                    "WM_MENUCOMMAND",
		WM_CHANGEUISTATE:                  "WM_CHANGEUISTATE",
		WM_UPDATEUISTATE:                  "WM_UPDATEUISTATE",
		WM_QUERYUISTATE:                   "WM_QUERYUISTATE",
		WM_CTLCOLORMSGBOX:                 "WM_CTLCOLORMSGBOX",
		WM_CTLCOLOREDIT:                   "WM_CTLCOLOREDIT",
		WM_CTLCOLORLISTBOX:                "WM_CTLCOLORLISTBOX",
		WM_CTLCOLORBTN:                    "WM_CTLCOLORBTN",
		WM_CTLCOLORDLG:                    "WM_CTLCOLORDLG",
		WM_CTLCOLORSCROLLBAR:              "WM_CTLCOLORSCROLLBAR",
		WM_CTLCOLORSTATIC:                 "WM_CTLCOLORSTATIC",
		WM_LBUTTONDBLCLK:                  "WM_LBUTTONDBLCLK",
		WM_RBUTTONDOWN:                    "WM_RBUTTONDOWN",
		WM_RBUTTONUP:                      "WM_RBUTTONUP",
		WM_RBUTTONDBLCLK:                  "WM_RBUTTONDBLCLK",
		WM_MBUTTONDOWN:                    "WM_MBUTTONDOWN",
		WM_MBUTTONUP:                      "WM_MBUTTONUP",
		WM_MBUTTONDBLCLK:                  "WM_MBUTTONDBLCLK",
		WM_XBUTTONDOWN:                    "WM_XBUTTONDOWN",
		WM_XBUTTONUP:                      "WM_XBUTTONUP",
		WM_XBUTTONDBLCLK:                  "WM_XBUTTONDBLCLK",
		WM_MOUSEHWHEEL:                    "WM_MOUSEHWHEEL",
		WM_PARENTNOTIFY:                   "WM_PARENTNOTIFY",
		WM_ENTERMENULOOP:                  "WM_ENTERMENULOOP",
		WM_EXITMENULOOP:                   "WM_EXITMENULOOP",
		WM_NEXTMENU:                       "WM_NEXTMENU",
		WM_SIZING:                         "WM_SIZING",
		WM_CAPTURECHANGED:                 "WM_CAPTURECHANGED",
		WM_MOVING:                         "WM_MOVING",
		WM_POWERBROADCAST:                 "WM_POWERBROADCAST",
		WM_DEVICECHANGE:                   "WM_DEVICECHANGE",
		WM_MDICREATE:                      "WM_MDICREATE",
		WM_MDIDESTROY:                     "WM_MDIDESTROY",
		WM_MDIACTIVATE:                    "WM_MDIACTIVATE",
		WM_MDIRESTORE:                     "WM_MDIRESTORE",
		WM_MDINEXT:                        "WM_MDINEXT",
		WM_MDIMAXIMIZE:                    "WM_MDIMAXIMIZE",
		WM_MDITILE:                        "WM_MDITILE",
		WM_MDICASCADE:                     "WM_MDICASCADE",
		WM_MDIICONARRANGE:                 "WM_MDIICONARRANGE",
		WM_MDIGETACTIVE:                   "WM_MDIGETACTIVE",
		WM_MDISETMENU:                     "WM_MDISETMENU",
		WM_ENTERSIZEMOVE:                  "WM_ENTERSIZEMOVE",
		WM_EXITSIZEMOVE:                   "WM_EXITSIZEMOVE",
		WM_MDIREFRESHMENU:                 "WM_MDIREFRESHMENU",
		WM_TOUCH:                          "WM_TOUCH",
		WM_IME_SETCONTEXT:                 "WM_IME_SETCONTEXT",
		WM_IME_NOTIFY:                     "WM_IME_NOTIFY",
		WM_IME_CONTROL:                    "WM_IME_CONTROL",
		WM_IME_COMPOSITIONFULL:            "WM_IME_COMPOSITIONFULL",
		WM_IME_SELECT:                     "WM_IME_SELECT",
		WM_IME_CHAR:                       "WM_IME_CHAR",
		WM_IME_REQUEST:                    "WM_IME_REQUEST",
		WM_IME_KEYDOWN:                    "WM_IME_KEYDOWN",
		WM_IME_KEYUP:                      "WM_IME_KEYUP",
		WM_NCMOUSEHOVER:                   "WM_NCMOUSEHOVER",
		WM_MOUSEHOVER:                     "WM_MOUSEHOVER",
		WM_NCMOUSELEAVE:                   "WM_NCMOUSELEAVE",
		WM_MOUSELEAVE:                     "WM_MOUSELEAVE",
		WM_WTSSESSION_CHANGE:              "WM_WTSSESSION_CHANGE",
		WM_TABLET_FIRST:                   "WM_TABLET_FIRST",
		WM_TABLET_LAST:                    "WM_TABLET_LAST",
		WM_CUT:                            "WM_CUT",
		WM_COPY:                           "WM_COPY",
		WM_PASTE:                          "WM_PASTE",
		WM_CLEAR:                          "WM_CLEAR",
		WM_UNDO:                           "WM_UNDO",
		WM_RENDERFORMAT:                   "WM_RENDERFORMAT",
		WM_RENDERALLFORMATS:               "WM_RENDERALLFORMATS",
		WM_DESTROYCLIPBOARD:               "WM_DESTROYCLIPBOARD",
		WM_DRAWCLIPBOARD:                  "WM_DRAWCLIPBOARD",
		WM_PAINTCLIPBOARD:                 "WM_PAINTCLIPBOARD",
		WM_VSCROLLCLIPBOARD:               "WM_VSCROLLCLIPBOARD",
		WM_SIZECLIPBOARD:                  "WM_SIZECLIPBOARD",
		WM_ASKCBFORMATNAME:                "WM_ASKCBFORMATNAME",
		WM_CHANGECBCHAIN:                  "WM_CHANGECBCHAIN",
		WM_HSCROLLCLIPBOARD:               "WM_HSCROLLCLIPBOARD",
		WM_QUERYNEWPALETTE:                "WM_QUERYNEWPALETTE",
		WM_PALETTEISCHANGING:              "WM_PALETTEISCHANGING",
		WM_PALETTECHANGED:                 "WM_PALETTECHANGED",
		WM_HOTKEY:                         "WM_HOTKEY",
		WM_PRINT:                          "WM_PRINT",
		WM_PRINTCLIENT:                    "WM_PRINTCLIENT",
		WM_APPCOMMAND:                     "WM_APPCOMMAND",
		WM_THEMECHANGED:                   "WM_THEMECHANGED",
		WM_CLIPBOARDUPDATE:                "WM_CLIPBOARDUPDATE",
		WM_DWMCOMPOSITIONCHANGED:          "WM_DWMCOMPOSITIONCHANGED",
		WM_DWMNCRENDERINGCHANGED:          "WM_DWMNCRENDERINGCHANGED",
		WM_DWMCOLORIZATIONCOLORCHANGED:    "WM_DWMCOLORIZATIONCOLORCHANGED",
		WM_DWMWINDOWMAXIMIZEDCHANGE:       "WM_DWMWINDOWMAXIMIZEDCHANGE",
		WM_DWMSENDICONICTHUMBNAIL:         "WM_DWMSENDICONICTHUMBNAIL",
		WM_DWMSENDICONICLIVEPREVIEWBITMAP: "WM_DWMSENDICONICLIVEPREVIEWBITMAP",
		WM_GETTITLEBARINFOEX:              "WM_GETTITLEBARINFOEX",
		WM_HANDHELDFIRST:                  "WM_HANDHELDFIRST",
		WM_HANDHELDLAST:                   "WM_HANDHELDLAST",
		WM_AFXFIRST:                       "WM_AFXFIRST",
		WM_AFXLAST:                        "WM_AFXLAST",
		WM_PENWINFIRST:                    "WM_PENWINFIRST",
		WM_PENWINLAST:                     "WM_PENWINLAST",
		WM_USER:                           "WM_USER",
		WM_APP:                            "WM_APP",
	}
	// messages that equal to existing messages:
	// WM_IME_KEYLAST
	// WM_SETTINGCHANGE
	// WM_MOUSEFIRST
	// WM_MOUSELAST
	if s, ok := messages[msg]; ok {
		return s
	}
	return "WM_...UNKNOWN"
} //                                                           WindowMessageName

// end
