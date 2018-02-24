// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-02-24 01:44:00 6EB730                  [zr_win/dll_gdi32_windows.go]
// -----------------------------------------------------------------------------

package win

import "syscall"      // standard
import "unicode/utf8" // standard
import "unsafe"       // standard

import "github.com/balacode/zr" // Zirconium

var gdi32 = syscall.NewLazyDLL("gdi32.dll")
var gdiBitBlt = gdi32.NewProc("BitBlt")
var gdiCreateCompatibleBitmap = gdi32.NewProc("CreateCompatibleBitmap")
var gdiCreateCompatibleDC = gdi32.NewProc("CreateCompatibleDC")
var gdiCreateFontIndirectW = gdi32.NewProc("CreateFontIndirectW")
var gdiCreateFontW = gdi32.NewProc("CreateFontW")
var gdiCreateRectRgn = gdi32.NewProc("CreateRectRgn")
var gdiCreateSolidBrush = gdi32.NewProc("CreateSolidBrush")
var gdiDeleteDC = gdi32.NewProc("DeleteDC")
var gdiDeleteObject = gdi32.NewProc("DeleteObject")
var gdiGetDeviceCaps = gdi32.NewProc("GetDeviceCaps")
var gdiGetTextExtentPoint32W = gdi32.NewProc("GetTextExtentPoint32W")
var gdiGetTextMetricsW = gdi32.NewProc("GetTextMetricsW")
var gdiSelectClipRgn = gdi32.NewProc("SelectClipRgn")
var gdiSelectObject = gdi32.NewProc("SelectObject")
var gdiSetBkColor = gdi32.NewProc("SetBkColor")
var gdiSetMapMode = gdi32.NewProc("SetMapMode")
var gdiSetTextColor = gdi32.NewProc("SetTextColor")
var gdiSetWindowOrgEx = gdi32.NewProc("SetWindowOrgEx")
var gdiTextOutW = gdi32.NewProc("TextOutW")

// not used:
// var gdiGetStockObject      = gdi32.NewProc("GetStockObject")
// var gdiIntersectClipRect   = gdi32.NewProc("IntersectClipRect")
// var gdiSetDCBrushColor     = gdi32.NewProc("SetDCBrushColor")

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
	var ret, _, _ = gdiBitBlt.Call(
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
	var ret, _, _ = gdiCreateCompatibleBitmap.Call(
		uintptr(hDC), // HDC: handle to DC
		uintptr(cx),  // int: width of bitmap, in pixels
		uintptr(cy),  // int: height of bitmap, in pixels
	)
	return HBITMAP(ret)
} //                                                      CreateCompatibleBitmap

// CreateCompatibleDC library: gdi32.dll
func CreateCompatibleDC(hDC HDC) HDC {
	var ret, _, _ = gdiCreateCompatibleDC.Call(uintptr(hDC)) // HDC:handle to DC
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
	var ret, _, _ = gdiCreateFontW.Call(
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
	var ret, _, _ = gdiCreateFontIndirectW.Call(uintptr(unsafe.Pointer(lplf)))
	return HFONT(ret)
} //                                                          CreateFontIndirect

// CreateRectRgn library: gdi32.dll
func CreateRectRgn(
	nLeftRect INT,
	nTopRect INT,
	nRightRect INT,
	nBottomRect INT,
) HRGN {
	var ret, _, _ = gdiCreateRectRgn.Call(
		uintptr(nLeftRect),   // x-coord of upper-left corner
		uintptr(nTopRect),    // y-coord of upper-left corner
		uintptr(nRightRect),  // x-coord of lower-right corner
		uintptr(nBottomRect), // y-coord of lower-right corner
	)
	return HRGN(ret)
} //                                                               CreateRectRgn

// CreateSolidBrush library: gdi32.dll
func CreateSolidBrush(color COLORREF) HBRUSH {
	var ret, _, _ = gdiCreateSolidBrush.Call(uintptr(color))
	return HBRUSH(ret)
} //                                                            CreateSolidBrush

// DeleteDC library: gdi32.dll
func DeleteDC(hDC HDC) BOOL {
	var ret, _, _ = gdiDeleteDC.Call(uintptr(hDC))
	return BOOLResult(ret)
} //                                                                    DeleteDC

// DeleteObject library: gdi32.dll
func DeleteObject(hObj HGDIOBJ) BOOL {
	var ret, _, _ = gdiDeleteObject.Call(uintptr(hObj))
	return BOOLResult(ret)
} //                                                                DeleteObject

// GetDeviceCaps library: gdi32.dll
func GetDeviceCaps(
	hDC HDC, //     handle to DC
	nIndex INT, //  index of capability
) INT {
	var ret, _, _ = gdiGetDeviceCaps.Call(uintptr(hDC), uintptr(nIndex))
	return INT(ret)
} //                                                               GetDeviceCaps

/*UNUSED
// GetStockObject library: gdi32.dll
func GetStockObject(fnObject INT) HGDIOBJ {
    var ret, _, _ = gdiGetStockObject.Call(
        uintptr(fnObject)) // int: stock object type
    )
    return HGDIOBJ(ret)
} //                                                           GetStockObject
*/

// GetTextExtentPoint32 library: gdi32.dll
func GetTextExtentPoint32(
	hDC HDC, //       handle to DC
	Text string, //   text string
	c INT, //         number of characters in string
	lpSize *SIZE, //  string size
) BOOL {
	var blankStr = c == 0 || len(Text) == 0
	if blankStr {
		Text, c = " ", 1
	}
	var charCount = utf8.RuneCountInString(Text)
	if charCount != int(c) {
		zr.Error("GetTextExtentPoint32() has inconsistent arguments:",
			"runes in text", charCount, "!= c ", c)
		c = INT(charCount)
	}
	var ret, _, _ = gdiGetTextExtentPoint32W.Call(
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
	var ret, _, _ = gdiGetTextMetricsW.Call(
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
//      var ret, _, _ = gdiIntersectClipRect.Call(
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
	var ret, _, _ = gdiSelectClipRgn.Call(
		uintptr(hDC),  // handle to DC
		uintptr(hRgn), // handle to region
	)
	return INT(ret)
} //                                                               SelectClipRgn

// SelectObject library: gdi32.dll
func SelectObject(hDC HDC, hGdiObj HGDIOBJ) HGDIOBJ {
	var ret, _, _ = gdiSelectObject.Call(
		uintptr(hDC),     // HDC: handle to DC
		uintptr(hGdiObj), // HGDIOBJ: handle to object
	)
	return HGDIOBJ(ret)
} //                                                                SelectObject

// SetBkColor library: gdi32.dll
func SetBkColor(hDC HDC, color COLORREF) COLORREF {
	var ret, _, _ = gdiSetBkColor.Call(
		uintptr(hDC),   // HDC: handle to DC
		uintptr(color), // COLORREF: background color value
	)
	return COLORREF(ret)
} //                                                                  SetBkColor

/*UNUSED
// SetDCBrushColor library: gdi32.dll
func SetDCBrushColor(hDC HDC, crColor COLORREF) COLORREF {
    var ret, _, _ = gdiSetDCBrushColor.Call(
        uintptr(hDC),     // HDC: handle to DC
        uintptr(crColor), // COLORREF: new brush color
    )
    return COLORREF(ret)
} //                                                             SetDCBrushColor
*/

// SetMapMode library: gdi32.dll
func SetMapMode(hDC HDC, fnMapMode INT) INT {
	var ret, _, _ = gdiSetMapMode.Call(
		uintptr(hDC),       // HDC: handle to device context
		uintptr(fnMapMode), // int: new mapping mode
	)
	return INT(ret)
} //                                                                  SetMapMode

// SetTextColor library: gdi32.dll
func SetTextColor(hDC HDC, color COLORREF) COLORREF {
	var ret, _, _ = gdiSetTextColor.Call(
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
	var ret, _, _ = gdiSetWindowOrgEx.Call(
		uintptr(hDC),                     // HDC: handle to device context
		uintptr(X),                       // int: new x-coord of window origin
		uintptr(Y),                       // int: new y-coord of window origin
		uintptr(unsafe.Pointer(lpPoint)), // LPPOINT : original window origin
	)
	return BOOLResult(ret)
} //                                                              SetWindowOrgEx

// TextOut library: gdi32.dll
func TextOut(hDC HDC, x, y INT, Text string, c INT) BOOL {
	var ret, _, _ = gdiTextOutW.Call(
		uintptr(hDC),             // HDC: handle to DC
		uintptr(x),               // int: x-coordinate of starting position
		uintptr(y),               // int: y-coordinate of starting position
		UintptrFromString(&Text), // LPCTSTR: character string
		uintptr(c),               // int: number of characters
	)
	return BOOLResult(ret)
} //                                                                     TextOut

//end
