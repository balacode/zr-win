// -----------------------------------------------------------------------------
// ZR Library: Windows 32 API                      zr-win/[dll_gdi32_windows.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package win

import (
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/balacode/zr"
)

var (
	gdi32 = syscall.NewLazyDLL("gdi32.dll")

	gdiBitBlt                 = gdi32.NewProc("BitBlt")
	gdiCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
	gdiCreateCompatibleDC     = gdi32.NewProc("CreateCompatibleDC")
	gdiCreateFontIndirectW    = gdi32.NewProc("CreateFontIndirectW")
	gdiCreateFontW            = gdi32.NewProc("CreateFontW")
	gdiCreateRectRgn          = gdi32.NewProc("CreateRectRgn")
	gdiCreateSolidBrush       = gdi32.NewProc("CreateSolidBrush")
	gdiDeleteDC               = gdi32.NewProc("DeleteDC")
	gdiDeleteObject           = gdi32.NewProc("DeleteObject")
	gdiGetDeviceCaps          = gdi32.NewProc("GetDeviceCaps")
	gdiGetTextExtentPoint32W  = gdi32.NewProc("GetTextExtentPoint32W")
	gdiGetTextMetricsW        = gdi32.NewProc("GetTextMetricsW")
	gdiSelectClipRgn          = gdi32.NewProc("SelectClipRgn")
	gdiSelectObject           = gdi32.NewProc("SelectObject")
	gdiSetBkColor             = gdi32.NewProc("SetBkColor")
	gdiSetMapMode             = gdi32.NewProc("SetMapMode")
	gdiSetTextColor           = gdi32.NewProc("SetTextColor")
	gdiSetWindowOrgEx         = gdi32.NewProc("SetWindowOrgEx")
	gdiTextOutW               = gdi32.NewProc("TextOutW")

	// not used:
	// gdiGetStockObject      = gdi32.NewProc("GetStockObject")
	// gdiIntersectClipRect   = gdi32.NewProc("IntersectClipRect")
	// gdiSetDCBrushColor     = gdi32.NewProc("SetDCBrushColor")
)

// BitBlt library: gdi32.dll
func BitBlt(
	hdcDest HDC,
	nXDest INT,
	nYDest INT,
	nWidth INT,
	nHeight INT,
	hdcSrc HDC,
	nXSrc INT,
	nYSrc INT,
	dwRop DWORD,
) BOOL {
	ret, _, _ := gdiBitBlt.Call(
		uintptr(hdcDest), // handle to destination DC
		uintptr(nXDest),  // x-coord of destination upper-left corner
		uintptr(nYDest),  // y-coord of destination upper-left corner
		uintptr(nWidth),  // width of destination rectangle
		uintptr(nHeight), // height of destination rectangle
		uintptr(hdcSrc),  // handle to source DC
		uintptr(nXSrc),   // x-coordinate of source upper-left corner
		uintptr(nYSrc),   // y-coordinate of source upper-left corner
		uintptr(dwRop),   // raster operation code
	)
	return BOOLResult(ret)
} //                                                                      BitBlt

// CreateCompatibleBitmap library: gdi32.dll
func CreateCompatibleBitmap(hDC HDC, cx, cy INT) HBITMAP {
	ret, _, _ := gdiCreateCompatibleBitmap.Call(
		uintptr(hDC), // HDC: handle to DC
		uintptr(cx),  // int: width of bitmap, in pixels
		uintptr(cy),  // int: height of bitmap, in pixels
	)
	return HBITMAP(ret)
} //                                                      CreateCompatibleBitmap

// CreateCompatibleDC library: gdi32.dll
func CreateCompatibleDC(hDC HDC) HDC {
	ret, _, _ := gdiCreateCompatibleDC.Call(uintptr(hDC)) // HDC:handle to DC
	return HDC(ret)
} //                                                          CreateCompatibleDC

// CreateFont library: gdi32.dll
func CreateFont(
	cHeight INT,
	cWidth INT,
	cEscapement INT,
	cOrientation INT,
	cWeight INT,
	bItalic DWORD,
	bUnderline DWORD,
	bStrikeOut DWORD,
	iCharSet DWORD,
	iOutPrecision DWORD,
	iClipPrecision DWORD,
	iQuality DWORD,
	iPitchAndFamily DWORD,
	pszFaceName string,
) HFONT {
	ret, _, _ := gdiCreateFontW.Call(
		uintptr(cHeight),                // int: height of font
		uintptr(cWidth),                 // int: average character width
		uintptr(cEscapement),            // int: angle of escapement
		uintptr(cOrientation),           // int: base-line orientation angle
		uintptr(cWeight),                // int: font weight
		uintptr(bItalic),                // DWORD: italic attribute option
		uintptr(bUnderline),             // DWORD: underline attribute option
		uintptr(bStrikeOut),             // DWORD: strikeout attribute option
		uintptr(iCharSet),               // DWORD: character set identifier
		uintptr(iOutPrecision),          // DWORD: output precision
		uintptr(iClipPrecision),         // DWORD: clipping precision
		uintptr(iQuality),               // DWORD: output quality
		uintptr(iPitchAndFamily),        // DWORD: pitch and family
		UintptrFromString(&pszFaceName), // LPCTSTR: typeface name
	)
	return HFONT(ret)
} //                                                                  CreateFont

// CreateFontIndirect library: gdi32.dll
func CreateFontIndirect(lplf *LOGFONT) HFONT {
	ret, _, _ := gdiCreateFontIndirectW.Call(uintptr(unsafe.Pointer(lplf)))
	return HFONT(ret)
} //                                                          CreateFontIndirect

// CreateRectRgn library: gdi32.dll
func CreateRectRgn(
	nLeftRect INT,
	nTopRect INT,
	nRightRect INT,
	nBottomRect INT,
) HRGN {
	ret, _, _ := gdiCreateRectRgn.Call(
		uintptr(nLeftRect),   // x-coord of upper-left corner
		uintptr(nTopRect),    // y-coord of upper-left corner
		uintptr(nRightRect),  // x-coord of lower-right corner
		uintptr(nBottomRect), // y-coord of lower-right corner
	)
	return HRGN(ret)
} //                                                               CreateRectRgn

// CreateSolidBrush library: gdi32.dll
func CreateSolidBrush(color COLORREF) HBRUSH {
	ret, _, _ := gdiCreateSolidBrush.Call(uintptr(color))
	return HBRUSH(ret)
} //                                                            CreateSolidBrush

// DeleteDC library: gdi32.dll
func DeleteDC(hDC HDC) BOOL {
	ret, _, _ := gdiDeleteDC.Call(uintptr(hDC))
	return BOOLResult(ret)
} //                                                                    DeleteDC

// DeleteObject library: gdi32.dll
func DeleteObject(hObj HGDIOBJ) BOOL {
	ret, _, _ := gdiDeleteObject.Call(uintptr(hObj))
	return BOOLResult(ret)
} //                                                                DeleteObject

// GetDeviceCaps library: gdi32.dll
func GetDeviceCaps(
	hDC HDC, //     handle to DC
	nIndex INT, //  index of capability
) INT {
	ret, _, _ := gdiGetDeviceCaps.Call(uintptr(hDC), uintptr(nIndex))
	return INT(ret)
} //                                                               GetDeviceCaps

//  UNUSED
//  // GetStockObject library: gdi32.dll
//  func GetStockObject(fnObject INT) HGDIOBJ {
//      ret, _, _ := gdiGetStockObject.Call(
//          uintptr(fnObject)) // int: stock object type
//      )
//      return HGDIOBJ(ret)
//  } //                                                          GetStockObject

// GetTextExtentPoint32 library: gdi32.dll
func GetTextExtentPoint32(
	hDC HDC, //       handle to DC
	Text string, //   text string
	c INT, //         number of characters in string
	lpSize *SIZE, //  string size
) BOOL {
	blankStr := c == 0 || len(Text) == 0
	if blankStr {
		Text, c = " ", 1
	}
	charCount := utf8.RuneCountInString(Text)
	if charCount != int(c) {
		zr.Error("GetTextExtentPoint32() has inconsistent arguments:",
			"runes in text", charCount, "!= c ", c)
		c = INT(charCount)
	}
	ret, _, _ := gdiGetTextExtentPoint32W.Call(
		uintptr(hDC),
		UintptrFromString(&Text),
		uintptr(c),
		uintptr(unsafe.Pointer(lpSize)),
	)
	if blankStr {
		(*lpSize).Cx = 0
	}
	return BOOLResult(ret)
} //                                                        GetTextExtentPoint32

// GetTextMetrics library: gdi32.dll
func GetTextMetrics(
	hDC HDC, //           handle to DC
	lptm *TEXTMETRIC, //  text metrics
) BOOL {
	ret, _, _ := gdiGetTextMetricsW.Call(
		uintptr(hDC),
		uintptr(unsafe.Pointer(lptm)),
	)
	return BOOLResult(ret)
} //                                                              GetTextMetrics

//  // IntersectClipRect library: gdi32.dll
//  func IntersectClipRect(
//      hDC HDC,
//      nLeftRect INT,
//      nTopRect INT,
//      nRightRect INT,
//      nBottomRect INT,
//  ) INT {
//      ret, _, _ := gdiIntersectClipRect.Call(
//          uintptr(hDC),           // handle to DC
//          uintptr(nLeftRect),     // x-coord of upper-left corner
//          uintptr(nTopRect),      // y-coord of upper-left corner
//          uintptr(nRightRect),    // x-coord of lower-right corner
//          uintptr(nBottomRect),   // y-coord of lower-right corner
//      )
//      return INT(ret)
//  } //                                                       IntersectClipRect

// SelectClipRgn library: gdi32.dll
func SelectClipRgn(hDC HDC, hRgn HRGN) INT {
	ret, _, _ := gdiSelectClipRgn.Call(
		uintptr(hDC),  // handle to DC
		uintptr(hRgn), // handle to region
	)
	return INT(ret)
} //                                                               SelectClipRgn

// SelectObject library: gdi32.dll
func SelectObject(hDC HDC, hGdiObj HGDIOBJ) HGDIOBJ {
	ret, _, _ := gdiSelectObject.Call(
		uintptr(hDC),     // HDC: handle to DC
		uintptr(hGdiObj), // HGDIOBJ: handle to object
	)
	return HGDIOBJ(ret)
} //                                                                SelectObject

// SetBkColor library: gdi32.dll
func SetBkColor(hDC HDC, color COLORREF) COLORREF {
	ret, _, _ := gdiSetBkColor.Call(
		uintptr(hDC),   // HDC: handle to DC
		uintptr(color), // COLORREF: background color value
	)
	return COLORREF(ret)
} //                                                                  SetBkColor

//  UNUSED
//  // SetDCBrushColor library: gdi32.dll
//  func SetDCBrushColor(hDC HDC, crColor COLORREF) COLORREF {
//      ret, _, _ := gdiSetDCBrushColor.Call(
//          uintptr(hDC),     // HDC: handle to DC
//          uintptr(crColor), // COLORREF: new brush color
//      )
//      return COLORREF(ret)
//  } //                                                         SetDCBrushColor

// SetMapMode library: gdi32.dll
func SetMapMode(hDC HDC, fnMapMode INT) INT {
	ret, _, _ := gdiSetMapMode.Call(
		uintptr(hDC),       // HDC: handle to device context
		uintptr(fnMapMode), // int: new mapping mode
	)
	return INT(ret)
} //                                                                  SetMapMode

// SetTextColor library: gdi32.dll
func SetTextColor(hDC HDC, color COLORREF) COLORREF {
	ret, _, _ := gdiSetTextColor.Call(
		uintptr(hDC),   // HDC: handle to DC
		uintptr(color), // COLORREF: text color
	)
	return COLORREF(ret)
} //                                                                SetTextColor

// SetWindowOrgEx library: gdi32.dll
func SetWindowOrgEx(
	hDC HDC,
	X INT,
	Y INT,
	lpPoint *POINT,
) BOOL {
	ret, _, _ := gdiSetWindowOrgEx.Call(
		uintptr(hDC),                     // HDC: handle to device context
		uintptr(X),                       // int: new x-coord of window origin
		uintptr(Y),                       // int: new y-coord of window origin
		uintptr(unsafe.Pointer(lpPoint)), // LPPOINT : original window origin
	)
	return BOOLResult(ret)
} //                                                              SetWindowOrgEx

// TextOut library: gdi32.dll
func TextOut(hDC HDC, x, y INT, Text string, c INT) BOOL {
	ret, _, _ := gdiTextOutW.Call(
		uintptr(hDC),             // HDC: handle to DC
		uintptr(x),               // int: x-coordinate of starting position
		uintptr(y),               // int: y-coordinate of starting position
		UintptrFromString(&Text), // LPCTSTR: character string
		uintptr(c),               // int: number of characters
	)
	return BOOLResult(ret)
} //                                                                     TextOut

// end
