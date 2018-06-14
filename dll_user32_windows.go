// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-09 01:03:18 1E0D93                 zr-win/[dll_user32_windows.go]
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unsafe"
)

var user32 = syscall.NewLazyDLL("user32.dll")
var userAppendMenuW = user32.NewProc("AppendMenuW")
var userBeginPaint = user32.NewProc("BeginPaint")
var userChangeWindowMessageFilter = user32.NewProc("ChangeWindowMessageFilter")
var userCheckMenuItem = user32.NewProc("CheckMenuItem")
var userCloseClipboard = user32.NewProc("CloseClipboard")
var userCreateCaret = user32.NewProc("CreateCaret")
var userCreateMenu = user32.NewProc("CreateMenu")
var userCreateWindowExW = user32.NewProc("CreateWindowExW")
var userDefWindowProcW = user32.NewProc("DefWindowProcW")
var userDestroyCaret = user32.NewProc("DestroyCaret")
var userDestroyMenu = user32.NewProc("DestroyMenu")
var userDestroyWindow = user32.NewProc("DestroyWindow")
var userDialogBoxIndirectParamW = user32.NewProc("DialogBoxIndirectParamW")
var userDialogBoxParamW = user32.NewProc("DialogBoxParamW")
var userDispatchMessageW = user32.NewProc("DispatchMessageW")
var userDrawTextW = user32.NewProc("DrawTextW")
var userEmptyClipboard = user32.NewProc("EmptyClipboard")
var userEndDialog = user32.NewProc("EndDialog")
var userEndPaint = user32.NewProc("EndPaint")
var userFillRect = user32.NewProc("FillRect")
var userGetActiveWindow = user32.NewProc("GetActiveWindow")
var userGetClientRect = user32.NewProc("GetClientRect")
var userGetClipboardData = user32.NewProc("GetClipboardData")
var userGetDC = user32.NewProc("GetDC")
var userGetDlgItem = user32.NewProc("GetDlgItem")
var userGetDlgItemTextW = user32.NewProc("GetDlgItemTextW")
var userGetKeyState = user32.NewProc("GetKeyState")
var userGetMenu = user32.NewProc("GetMenu")
var userGetMessageW = user32.NewProc("GetMessageW")
var userHideCaret = user32.NewProc("HideCaret")
var userInvalidateRect = user32.NewProc("InvalidateRect")
var userIsMenu = user32.NewProc("IsMenu")
var userIsWindow = user32.NewProc("IsWindow")
var userLoadAcceleratorsW = user32.NewProc("LoadAcceleratorsW")
var userLoadCursorW = user32.NewProc("LoadCursorW")
var userLoadIconW = user32.NewProc("LoadIconW")
var userMapVirtualKeyW = user32.NewProc("MapVirtualKeyW")
var userMessageBoxW = user32.NewProc("MessageBoxW")
var userOpenClipboard = user32.NewProc("OpenClipboard")
var userPeekMessageW = user32.NewProc("PeekMessageW")
var userPostQuitMessage = user32.NewProc("PostQuitMessage")
var userRegisterClassExW = user32.NewProc("RegisterClassExW")
var userReleaseDC = user32.NewProc("ReleaseDC")
var userSendMessageW = user32.NewProc("SendMessageW")
var userSetCaretPos = user32.NewProc("SetCaretPos")
var userSetClipboardData = user32.NewProc("SetClipboardData")
var userSetCursor = user32.NewProc("SetCursor")
var userSetDlgItemTextW = user32.NewProc("SetDlgItemTextW")
var userSetMenu = user32.NewProc("SetMenu")
var userSetMenuItemInfoW = user32.NewProc("SetMenuItemInfoW")
var userSetRect = user32.NewProc("SetRect")
var userSetWindowPos = user32.NewProc("SetWindowPos")
var userSetWindowTextW = user32.NewProc("SetWindowTextW")
var userShowCaret = user32.NewProc("ShowCaret")
var userShowWindow = user32.NewProc("ShowWindow")
var userTranslateAccelerator = user32.NewProc("TranslateAccelerator")
var userTranslateMessage = user32.NewProc("TranslateMessage")
var userUnregisterHotKey = user32.NewProc("UnregisterHotKey")
var userUpdateWindow = user32.NewProc("UpdateWindow")
var userWindowFromDC = user32.NewProc("WindowFromDC")

var userIsClipboardFormatAvailable = user32.NewProc(
	"IsClipboardFormatAvailable",
)

// unused for now:
// userRedrawWindow   = user32.NewProc("RedrawWindow")

// AppendMenu library: user32.dll
func AppendMenu(
	hMenu HMENU,
	uFlags UINT,
	uIDNewItem UINT_PTR,
	NewItem string,
) BOOL {
	var ret, _, _ = userAppendMenuW.Call(
		uintptr(hMenu),
		uintptr(uFlags),
		uintptr(uIDNewItem),
		UintptrFromString(&NewItem))
	return BOOLResult(ret)
} //                                                                  AppendMenu

// BeginPaint library: user32.dll
func BeginPaint(hWnd HWND, lpPaint *PAINTSTRUCT) HDC {
	var ret, _, _ = userBeginPaint.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)))
	return HDC(ret)
} //                                                                  BeginPaint

// CheckMenuItem library: user32.dll
func CheckMenuItem(hMenu HMENU, uIDCheckItem, uCheck UINT) DWORD {
	var ret, _, _ = userCheckMenuItem.Call(
		uintptr(hMenu),
		uintptr(uIDCheckItem),
		uintptr(uCheck))
	return DWORD(ret)
} //                                                               CheckMenuItem

// CloseClipboard library: user32.dll
func CloseClipboard() BOOL {
	var ret, _, _ = userCloseClipboard.Call()
	return BOOLResult(ret)
} //                                                              CloseClipboard

// CreateCaret library: user32.dll
func CreateCaret(
	hWnd HWND,
	hBitmap HBITMAP,
	nWidth INT,
	nHeight INT,
) BOOL {
	var ret, _, _ = userCreateCaret.Call(
		uintptr(hWnd),
		uintptr(hBitmap),
		uintptr(nWidth),
		uintptr(nHeight))
	return BOOLResult(ret)
} //                                                                 CreateCaret

// CreateMenu library: user32.dll
func CreateMenu() HMENU {
	var ret, _, _ = userCreateMenu.Call()
	return HMENU(ret)
} //                                                                  CreateMenu

// CreateWindowEx library: user32.dll
func CreateWindowEx(
	dwExStyle DWORD,
	ClassName string,
	WindowName string,
	dwStyle DWORD,
	x INT,
	y INT,
	nWidth INT,
	nHeight INT,
	hWndParent HWND,
	hMenu HMENU,
	HInstance HINSTANCE,
	lpParam LPVOID,
) HWND {
	var ret, _, _ = userCreateWindowExW.Call(
		uintptr(dwExStyle),
		UintptrFromString(&ClassName),
		UintptrFromString(&WindowName),
		uintptr(dwStyle),
		uintptr(x),
		uintptr(y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(HInstance),
		uintptr(unsafe.Pointer(lpParam)))
	return HWND(ret)
} //                                                              CreateWindowEx

// DefWindowProc library: user32.dll
func DefWindowProc(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var ret, _, _ = userDefWindowProcW.Call(
		uintptr(hWnd),
		uintptr(Msg),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
} //                                                               DefWindowProc

// DestroyCaret library: user32.dll
func DestroyCaret() BOOL {
	var ret, _, _ = userDestroyCaret.Call()
	return BOOLResult(ret)
} //                                                                DestroyCaret

// DestroyMenu library: user32.dll
func DestroyMenu(hMenu HMENU) BOOL {
	var ret, _, _ = userDestroyMenu.Call(uintptr(hMenu))
	return BOOLResult(ret)
} //                                                                 DestroyMenu

// DestroyWindow library: user32.dll
func DestroyWindow(hWnd HWND) BOOL {
	var ret, _, _ = userDestroyWindow.Call(uintptr(hWnd))
	return BOOLResult(ret)
} //                                                               DestroyWindow

// DialogBoxIndirectParam library: user32.dll
func DialogBoxIndirectParam(
	hInstance HINSTANCE,
	hDialogTemplate *DLGTEMPLATE,
	hWndParent HWND,
	lpDialogFunc DLGPROC,
	dwInitParam LPARAM,
) INT_PTR {
	var ret, _, _ = userDialogBoxIndirectParamW.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(hDialogTemplate)),
		uintptr(hWndParent),
		uintptr(lpDialogFunc),
		uintptr(dwInitParam),
	)
	return INT_PTR(ret)
} //                                                      DialogBoxIndirectParam

// DispatchMessage library: user32.dll
func DispatchMessage(lpmsg *MSG) LRESULT {
	var ret, _, _ = userDispatchMessageW.Call(
		uintptr(unsafe.Pointer(lpmsg)),
	)
	return LRESULT(ret)
} //                                                             DispatchMessage

// DrawText library: user32.dll
func DrawText(
	hDC HDC,
	lpString string,
	nCount INT,
	lpRect *RECT,
	uFormat UINT,
) INT {
	var ret, _, _ = userDrawTextW.Call(
		uintptr(hDC),                    // handle to DC
		UintptrFromString(&lpString),    // text to draw
		uintptr(nCount),                 // text length
		uintptr(unsafe.Pointer(lpRect)), // formatting dimensions
		uintptr(uFormat),                // text-drawing options
	)
	return INT(ret)
} //                                                                    DrawText

// EmptyClipboard library: user32.dll
func EmptyClipboard() BOOL {
	var ret, _, _ = userEmptyClipboard.Call()
	return BOOLResult(ret)
} //                                                              EmptyClipboard

// EndDialog library: user32.dll
func EndDialog(hDlg HWND, nResult INT_PTR) BOOL {
	var ret, _, _ = userEndDialog.Call(
		uintptr(hDlg),
		uintptr(nResult),
	)
	return BOOLResult(ret)
} //                                                                   EndDialog

// EndPaint library: user32.dll
func EndPaint(hWnd HWND, lpPaint *PAINTSTRUCT) BOOL {
	var ret, _, _ = userEndPaint.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpPaint)),
	)
	return BOOLResult(ret)
} //                                                                    EndPaint

// FillRect library: user32.dll
func FillRect(hDC HDC, lprc *RECT, hbr HBRUSH) INT {
	var ret, _, _ = userFillRect.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lprc)),
		uintptr(hbr),
	)
	return INT(ret)
} //                                                                    FillRect

// GetActiveWindow library: user32.dll
func GetActiveWindow() HWND {
	var ret, _, _ = userGetActiveWindow.Call()
	return HWND(ret)
} //                                                             GetActiveWindow

// GetClientRect library: user32.dll
func GetClientRect(hWnd HWND, lpRect *RECT) BOOL {
	var ret, _, _ = userGetClientRect.Call(
		uintptr(hWnd),
		uintptr(unsafe.Pointer(lpRect)),
	)
	return BOOLResult(ret)
} //                                                               GetClientRect

// GetClipboardData library: user32.dll
func GetClipboardData(uFormat UINT) HANDLE {
	var ret, _, _ = userGetClipboardData.Call(uintptr(uFormat))
	return HANDLE(ret)
} //                                                            GetClipboardData

// GetDC library: user32.dll
func GetDC(hWnd HWND) HDC {
	var ret, _, _ = userGetDC.Call(uintptr(hWnd))
	return HDC(ret)
} //                                                                       GetDC

// GetDlgItem library: user32.dll
func GetDlgItem(hDlg HWND, nIDDlgItem int) HWND {
	var ret, _, _ = userGetDlgItem.Call(
		uintptr(hDlg),
		uintptr(nIDDlgItem),
	)
	return HWND(ret)
} //                                                                  GetDlgItem

// GetDlgItemText library: user32.dll
func GetDlgItemText(
	hDlg HWND,
	nIDDlgItem INT,
	lpString LPWSTR,
	nMaxCount INT,
) UINT {
	var ret, _, _ = userGetDlgItemTextW.Call(
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		uintptr(unsafe.Pointer(lpString)),
		uintptr(nMaxCount),
	)
	return UINT(ret)
} //                                                              GetDlgItemText

// GetKeyState library: user32.dll
func GetKeyState(nVirtKey INT) SHORT {
	var ret, _, _ = userGetKeyState.Call(uintptr(nVirtKey)) // [in] int
	return SHORT(ret)
} //                                                                 GetKeyState

// GetMenu library: user32.dll
func GetMenu(hWnd HWND) HMENU {
	var ret, _, _ = userGetMenu.Call(uintptr(hWnd))
	return HMENU(ret)
} //                                                                     GetMenu

// GetMessage library: user32.dll
func GetMessage(
	lpMsg *MSG,
	hWnd HWND,
	wMsgFilterMin UINT,
	wMsgFilterMax UINT,
) BOOL {
	var ret, _, _ = userGetMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
	)
	return BOOLResult(ret)
} //                                                                  GetMessage

// HideCaret library: user32.dll
func HideCaret(hWnd HWND) BOOL {
	var ret, _, _ = userHideCaret.Call(uintptr(hWnd)) // [in] HWND
	return BOOLResult(ret)
} //                                                                   HideCaret

// InvalidateRect library: user32.dll
func InvalidateRect(hWnd HWND, lpRect *RECT, bErase BOOL) BOOL {
	var ret, _, _ = userInvalidateRect.Call(
		uintptr(hWnd),                   // handle to window
		uintptr(unsafe.Pointer(lpRect)), // rectangle coordinates
		uintptr(bErase),                 // erase state
	)
	return BOOLResult(ret)
} //                                                              InvalidateRect

// IsClipboardFormatAvailable library: user32.dll
func IsClipboardFormatAvailable(format UINT) BOOL {
	var ret, _, _ = userIsClipboardFormatAvailable.Call(uintptr(format))
	return BOOLResult(ret)
} //                                                  IsClipboardFormatAvailable

// IsMenu library: user32.dll
func IsMenu(hMenu HMENU) BOOL {
	var ret, _, _ = userIsMenu.Call(uintptr(hMenu))
	return BOOLResult(ret)
} //                                                                      IsMenu

// IsWindow library: user32.dll
func IsWindow(hWnd HWND) BOOL {
	var ret, _, _ = userIsWindow.Call(uintptr(hWnd))
	return BOOLResult(ret)
} //                                                                    IsWindow

// LoadAccelerators library: user32.dll
func LoadAccelerators(hInstance HINSTANCE, TableName string) HACCEL {
	var ret, _, _ = userLoadAcceleratorsW.Call(
		uintptr(hInstance),
		UintptrFromString(&TableName),
	)
	return HACCEL(ret)
} //                                                            LoadAccelerators

// LoadCursor library: user32.dll
func LoadCursor(hInstance HINSTANCE, lpCursorName LPCWSTR) HCURSOR {
	var ret, _, _ = userLoadCursorW.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpCursorName)),
	)
	return HCURSOR(ret)
} //                                                                  LoadCursor

// LoadIcon library: user32.dll
func LoadIcon(hInstance HINSTANCE, lpIconName LPCWSTR) HICON {
	var ret, _, _ = userLoadIconW.Call(
		uintptr(hInstance),
		uintptr(unsafe.Pointer(lpIconName)),
	)
	return HICON(ret)
} //                                                                    LoadIcon

// MapVirtualKey library: user32.dll
func MapVirtualKey(nCode, uMapType uint) uint {
	var ret, _, _ = userMapVirtualKeyW.Call(
		uintptr(nCode),
		uintptr(uMapType),
	)
	return uint(ret)
} //                                                               MapVirtualKey

// MessageBox library: user32.dll
func MessageBox(hWnd HWND, title, caption string, flags uint) int {
	var ret, _, _ = userMessageBoxW.Call(
		uintptr(hWnd),
		UintptrFromString(&title),
		UintptrFromString(&caption),
		uintptr(flags),
	)
	return int(ret)
} //                                                                  MessageBox

// OpenClipboard library: user32.dll
func OpenClipboard(hWnd HWND) BOOL {
	var ret, _, _ = userOpenClipboard.Call(uintptr(hWnd))
	return BOOLResult(ret)
} //                                                               OpenClipboard

// PeekMessage library: user32.dll
func PeekMessage(
	lpMsg *MSG,
	hWnd HWND,
	wMsgFilterMin UINT,
	wMsgFilterMax UINT,
	wRemoveMsg UINT,
) BOOL {
	var ret, _, _ = userPeekMessageW.Call(
		uintptr(unsafe.Pointer(lpMsg)),
		uintptr(hWnd),
		uintptr(wMsgFilterMin),
		uintptr(wMsgFilterMax),
		uintptr(wRemoveMsg),
	)
	return BOOLResult(ret)
} //                                                                 PeekMessage

// PostQuitMessage library: user32.dll
func PostQuitMessage(nExitCode INT) {
	userPostQuitMessage.Call(uintptr(nExitCode))
} //                                                             PostQuitMessage

// unused for now:
//    // RedrawWindow library: user32.dll
//    func RedrawWindow(
//          hWnd HWND,
//          lprcUpdate *RECT,
//          hrgnUpdate HRGN,
//          flags UINT,
//    ) BOOL {
//          var ret, _, _ = userRedrawWindow.Call(
//                uintptr(hWnd),
//                uintptr(unsafe.Pointer(lprcUpdate)),
//                uintptr(hrgnUpdate),
//                uintptr(flags),
//          )
//          return BOOLResult(ret)
//    }

// RegisterClassEx library: user32.dll
func RegisterClassEx(lpWndClass *WNDCLASSEX) ATOM {
	var ret, _, _ = userRegisterClassExW.Call(
		uintptr(unsafe.Pointer(lpWndClass)),
	)
	return ATOM(ret)
} //                                                             RegisterClassEx

// ReleaseDC library: user32.dll
func ReleaseDC(hWnd HWND, hDC HDC) INT {
	var ret, _, _ = userReleaseDC.Call(
		uintptr(hWnd),
		uintptr(hDC),
	)
	return INT(ret)
} //                                                                   ReleaseDC

// SendMessage library: user32.dll
func SendMessage(hWnd HWND, Msg UINT, wParam WPARAM, lParam LPARAM) LRESULT {
	var ret, _, _ = userSendMessageW.Call(
		uintptr(hWnd),
		uintptr(Msg),
		uintptr(wParam),
		uintptr(lParam),
	)
	return LRESULT(ret)
} //                                                                 SendMessage

// SetCaretPos library: user32.dll
func SetCaretPos(x, y INT) BOOL {
	var ret, _, _ = userSetCaretPos.Call(uintptr(x), uintptr(y))
	return BOOLResult(ret)
} //                                                                 SetCaretPos

// SetClipboardData library: user32.dll
func SetClipboardData(uFormat UINT, hMem HANDLE) HANDLE {
	var ret, _, _ = userSetClipboardData.Call(uintptr(uFormat), uintptr(hMem))
	return HANDLE(ret)
} //                                                            SetClipboardData

// SetCursor library: user32.dll
func SetCursor(hCursor HCURSOR) HCURSOR {
	var ret, _, _ = userSetCursor.Call(uintptr(hCursor))
	return HCURSOR(ret)
} //                                                                   SetCursor

// SetDlgItemText library: user32.dll
func SetDlgItemText(
	hDlg HWND,
	nIDDlgItem INT,
	Text string,
) BOOL {
	var ret, _, _ = userSetDlgItemTextW.Call(
		uintptr(hDlg),
		uintptr(nIDDlgItem),
		UintptrFromString(&Text),
	)
	return BOOLResult(ret)
} //                                                              SetDlgItemText

// SetMenu library: user32.dll
func SetMenu(hWnd HWND, hMenu HMENU) BOOL {
	var ret, _, _ = userSetMenu.Call(
		uintptr(hWnd),
		uintptr(hMenu),
	)
	return BOOLResult(ret)
} //                                                                     SetMenu

// SetMenuItemInfo library: user32.dll
func SetMenuItemInfo(
	hMenu HMENU,
	uItem UINT,
	fByPosition BOOL,
	lpmii *MENUITEMINFO,
) BOOL {
	var ret, _, _ = userSetMenuItemInfoW.Call(
		uintptr(hMenu),
		uintptr(uItem),
		uintptr(fByPosition),
		uintptr(unsafe.Pointer(lpmii)),
	)
	return BOOLResult(ret)
} //                                                             SetMenuItemInfo

// SetRect library: user32.dll
func SetRect(lprc *RECT, xLeft, yTop, xRight, yBottom INT) BOOL {
	var ret, _, _ = userSetRect.Call(
		uintptr(unsafe.Pointer(lprc)),
		uintptr(xLeft),
		uintptr(yTop),
		uintptr(xRight),
		uintptr(yBottom),
	)
	return BOOLResult(ret)
} //                                                                     SetRect

// SetWindowPos library: user32.dll
func SetWindowPos(
	hWnd HWND,
	hWndInsertAfter HWND,
	X INT,
	Y INT,
	cx INT,
	cy INT,
	uFlags UINT,
) BOOL {
	var ret, _, _ = userSetWindowPos.Call(
		uintptr(hWnd),
		uintptr(hWndInsertAfter),
		uintptr(X),
		uintptr(Y),
		uintptr(cx),
		uintptr(cy),
		uintptr(uFlags),
	)
	return BOOLResult(ret)
} //                                                                SetWindowPos

// SetWindowText library: user32.dll
func SetWindowText(hWnd HWND, Text string) BOOL {
	var ret, _, _ = userSetWindowTextW.Call(
		uintptr(hWnd),
		UintptrFromString(&Text),
	)
	return BOOLResult(ret)
} //                                                               SetWindowText

// ShowCaret library: user32.dll
func ShowCaret(hWnd HWND) BOOL {
	var ret, _, _ = userShowCaret.Call(uintptr(hWnd))
	return BOOLResult(ret)
} //                                                                   ShowCaret

// ShowWindow library: user32.dll
func ShowWindow(hWnd HWND, nCmdShow INT) BOOL {
	var ret, _, _ = userShowWindow.Call(
		uintptr(hWnd),
		uintptr(nCmdShow),
	)
	return BOOLResult(ret)
} //                                                                  ShowWindow

// TranslateAccelerator library: user32.dll
func TranslateAccelerator(hWnd HWND, hAccTable HACCEL, lpMsg *MSG) INT {
	var ret, _, _ = userTranslateMessage.Call(
		uintptr(hWnd),
		uintptr(hAccTable),
		uintptr(unsafe.Pointer(lpMsg)),
	)
	return INT(ret)
} //                                                        TranslateAccelerator

// TranslateMessage library: user32.dll
func TranslateMessage(lpmsg *MSG) BOOL {
	var ret, _, _ = userTranslateMessage.Call(uintptr(unsafe.Pointer(lpmsg)))
	if ret == 0 {
		return FALSE
	}
	return TRUE
	// If message is translated (posted to thread's message que) returns TRUE
	// Always returns TRUE for WM_KEYDOWN, WM_KEYUP, WM_SYSKEYDOWN, WM_SYSKEYUP
} //                                                            TranslateMessage

// UnregisterHotKey library: user32.dll
func UnregisterHotKey(hWnd HWND, id INT) BOOL {
	var ret, _, _ = userUnregisterHotKey.Call(
		uintptr(hWnd),
		uintptr(id),
	)
	return BOOLResult(ret)
} //                                                            UnregisterHotKey

// UpdateWindow library: user32.dll
func UpdateWindow(hWnd HWND) BOOL {
	var ret, _, _ = userUpdateWindow.Call(uintptr(hWnd))
	return BOOLResult(ret)
} //                                                                UpdateWindow

// WindowFromDC library: user32.dll
func WindowFromDC(hDC HDC) HWND {
	var ret, _, _ = userWindowFromDC.Call(uintptr(hDC))
	return HWND(ret)
} //                                                                WindowFromDC

//end
