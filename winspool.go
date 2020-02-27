// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"syscall"
	"unsafe"
)

// EnumPrinters flags
const (
	PRINTER_ENUM_DEFAULT     = 0x00000001
	PRINTER_ENUM_LOCAL       = 0x00000002
	PRINTER_ENUM_CONNECTIONS = 0x00000004
	PRINTER_ENUM_FAVORITE    = 0x00000004
	PRINTER_ENUM_NAME        = 0x00000008
	PRINTER_ENUM_REMOTE      = 0x00000010
	PRINTER_ENUM_SHARED      = 0x00000020
	PRINTER_ENUM_NETWORK     = 0x00000040
)

// Printer access flags
const (
	PRINTER_ACCESS_ADMINISTER     = 0x00000004
	PRINTER_ACCESS_USE            = 0x00000008
	PRINTER_ACCESS_MANAGE_LIMITED = 0x00000040
	PRINTER_ALL_ACCESS            = (STANDARD_RIGHTS_REQUIRED | PRINTER_ACCESS_ADMINISTER | PRINTER_ACCESS_USE)
)

type DOC_INFO_1 struct {
	PDocName    *uint16
	POutputFile *uint16
	PDatatype   *uint16
}

type PRINTER_INFO_4 struct {
	PPrinterName *uint16
	PServerName  *uint16
	Attributes   uint32
}

type PRINTER_DEFAULTS struct {
	PDatatype     *uint16
	LPDevMode     *DEVMODE
	DesiredAccess ACCESS_MASK
}

var (
	// Library
	libwinspool uintptr

	// Functions
	deviceCapabilities uintptr
	documentProperties uintptr
	enumPrinters       uintptr
	getDefaultPrinter  uintptr
	openPrinter        uintptr
	closePrinter       uintptr
	startDocPrinter    uintptr
	endDocPrinter      uintptr
	startPagePrinter   uintptr
	endPagePrinter     uintptr
	writePrinter       uintptr
)

func init() {
	// Library
	libwinspool = MustLoadLibrary("winspool.drv")

	// Functions
	deviceCapabilities = MustGetProcAddress(libwinspool, "DeviceCapabilitiesW")
	documentProperties = MustGetProcAddress(libwinspool, "DocumentPropertiesW")
	enumPrinters = MustGetProcAddress(libwinspool, "EnumPrintersW")
	getDefaultPrinter = MustGetProcAddress(libwinspool, "GetDefaultPrinterW")
	openPrinter = MustGetProcAddress(libwinspool, "OpenPrinterW")
	closePrinter = MustGetProcAddress(libwinspool, "ClosePrinter")
	startDocPrinter = MustGetProcAddress(libwinspool, "StartDocPrinterW")
	endDocPrinter = MustGetProcAddress(libwinspool, "EndDocPrinter")
	startPagePrinter = MustGetProcAddress(libwinspool, "StartPagePrinter")
	endPagePrinter = MustGetProcAddress(libwinspool, "EndPagePrinter")
	writePrinter = MustGetProcAddress(libwinspool, "WritePrinter")
}

func DeviceCapabilities(pDevice, pPort *uint16, fwCapability uint16, pOutput *uint16, pDevMode *DEVMODE) uint32 {
	ret, _, _ := syscall.Syscall6(deviceCapabilities, 5,
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(pPort)),
		uintptr(fwCapability),
		uintptr(unsafe.Pointer(pOutput)),
		uintptr(unsafe.Pointer(pDevMode)),
		0)

	return uint32(ret)
}

func DocumentProperties(hWnd HWND, hPrinter HANDLE, pDeviceName *uint16, pDevModeOutput, pDevModeInput *DEVMODE, fMode uint32) int32 {
	ret, _, _ := syscall.Syscall6(documentProperties, 6,
		uintptr(hWnd),
		uintptr(hPrinter),
		uintptr(unsafe.Pointer(pDeviceName)),
		uintptr(unsafe.Pointer(pDevModeOutput)),
		uintptr(unsafe.Pointer(pDevModeInput)),
		uintptr(fMode))

	return int32(ret)
}

func EnumPrinters(Flags uint32, Name *uint16, Level uint32, pPrinterEnum *byte, cbBuf uint32, pcbNeeded, pcReturned *uint32) bool {
	ret, _, _ := syscall.Syscall9(enumPrinters, 7,
		uintptr(Flags),
		uintptr(unsafe.Pointer(Name)),
		uintptr(Level),
		uintptr(unsafe.Pointer(pPrinterEnum)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(pcbNeeded)),
		uintptr(unsafe.Pointer(pcReturned)),
		0,
		0)

	return ret != 0
}

func GetDefaultPrinter(pszBuffer *uint16, pcchBuffer *uint32) bool {
	ret, _, _ := syscall.Syscall(getDefaultPrinter, 2,
		uintptr(unsafe.Pointer(pszBuffer)),
		uintptr(unsafe.Pointer(pcchBuffer)),
		0)

	return ret != 0
}

func OpenPrinter(pPrinterName *uint16, phPrinter *HANDLE, pDefault *PRINTER_DEFAULTS) bool {
	ret, _, _ := syscall.Syscall(openPrinter, 3,
		uintptr(unsafe.Pointer(pPrinterName)),
		uintptr(unsafe.Pointer(phPrinter)),
		uintptr(unsafe.Pointer(pDefault)))

	return ret != 0
}

func ClosePrinter(hPrinter HANDLE) bool {
	ret, _, _ := syscall.Syscall(closePrinter, 1,
		uintptr(hPrinter), 0, 0)

	return ret != 0
}

func StartDocPrinter(hPrinter HANDLE, level uint32, pDocInfo *DOC_INFO_1) uint32 {
	ret, _, _ := syscall.Syscall(startDocPrinter, 3,
		uintptr(hPrinter),
		uintptr(level),
		uintptr(unsafe.Pointer(pDocInfo)))

	return uint32(ret)
}

func EndDocPrinter(hPrinter HANDLE) bool {
	ret, _, _ := syscall.Syscall(endDocPrinter, 1,
		uintptr(hPrinter), 0, 0)

	return ret != 0
}

func StartPagePrinter(hPrinter HANDLE) bool {
	ret, _, _ := syscall.Syscall(startPagePrinter, 1,
		uintptr(hPrinter), 0, 0)

	return ret != 0
}

func EndPagePrinter(hPrinter HANDLE) bool {
	ret, _, _ := syscall.Syscall(endPagePrinter, 1,
		uintptr(hPrinter), 0, 0)

	return ret != 0
}

func WritePrinter(hPrinter HANDLE, pBuf *byte, cbBuf uint32, pcWritten *uint32) bool {
	ret, _, _ := syscall.Syscall6(writePrinter, 4,
		uintptr(hPrinter),
		uintptr(unsafe.Pointer(pBuf)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(pcWritten)),
		0,
		0)

	return ret != 0
}
