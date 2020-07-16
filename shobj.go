// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	CLSID_TaskbarList    = CLSID{0x56FDF344, 0xFD6D, 0x11d0, [8]byte{0x95, 0x8A, 0x00, 0x60, 0x97, 0xC9, 0xA0, 0x90}}
	CLSID_FileOpenDialog = CLSID{0xDC1C5A9C, 0xE88A, 0x4DDE, [8]byte{0xA5, 0xA1, 0x60, 0xF8, 0x2A, 0x20, 0xAE, 0xF7}}
	IID_ITaskbarList3    = IID{0xea1afb91, 0x9e28, 0x4b86, [8]byte{0x90, 0xe9, 0x9e, 0x9f, 0x8a, 0x5e, 0xef, 0xaf}}
	IID_IFileOpenDialog  = IID{0xD57C7288, 0xD4AD, 0x4768, [8]byte{0xBE, 0x02, 0x9D, 0x96, 0x95, 0x32, 0xD9, 0x60}}
	IID_IShellItem       = IID{0x43826D1E, 0xE718, 0x42EE, [8]byte{0xBC, 0x55, 0xA1, 0xE2, 0x61, 0xC3, 0x7B, 0xFE}}
)

//TBPFLAG
const (
	TBPF_NOPROGRESS    = 0
	TBPF_INDETERMINATE = 0x1
	TBPF_NORMAL        = 0x2
	TBPF_ERROR         = 0x4
	TBPF_PAUSED        = 0x8
)

const (
	FOS_OVERWRITEPROMPT          = 0x2
	FOS_STRICTFILETYPES          = 0x4
	FOS_NOCHANGEDIR              = 0x8
	FOS_PICKFOLDERS              = 0x20
	FOS_FORCEFILESYSTEM          = 0x40
	FOS_ALLNONSTORAGEITEMS       = 0x80
	FOS_NOVALIDATE               = 0x100
	FOS_ALLOWMULTISELECT         = 0x200
	FOS_PATHMUSTEXIST            = 0x800
	FOS_FILEMUSTEXIST            = 0x1000
	FOS_CREATEPROMPT             = 0x2000
	FOS_SHAREAWARE               = 0x4000
	FOS_NOREADONLYRETURN         = 0x8000
	FOS_NOTESTFILECREATE         = 0x10000
	FOS_HIDEMRUPLACES            = 0x20000
	FOS_HIDEPINNEDPLACES         = 0x40000
	FOS_NODEREFERENCELINKS       = 0x100000
	FOS_OKBUTTONNEEDSINTERACTION = 0x200000
	FOS_DONTADDTORECENT          = 0x2000000
	FOS_FORCESHOWHIDDEN          = 0x10000000
	FOS_DEFAULTNOMINIMODE        = 0x20000000
	FOS_FORCEPREVIEWPANEON       = 0x40000000
	FOS_SUPPORTSTREAMABLEITEMS   = 0x80000000
)

const (
	SIGDN_NORMALDISPLAY               = 0x0
	SIGDN_PARENTRELATIVEPARSING       = 0x80018001
	SIGDN_DESKTOPABSOLUTEPARSING      = 0x80028000
	SIGDN_PARENTRELATIVEEDITING       = 0x80031001
	SIGDN_DESKTOPABSOLUTEEDITING      = 0x8004c000
	SIGDN_FILESYSPATH                 = 0x80058000
	SIGDN_URL                         = 0x80068000
	SIGDN_PARENTRELATIVEFORADDRESSBAR = 0x8007c001
	SIGDN_PARENTRELATIVE              = 0x80080001
	SIGDN_PARENTRELATIVEFORUI         = 0x80094001
)

type ITaskbarList3Vtbl struct {
	QueryInterface        uintptr
	AddRef                uintptr
	Release               uintptr
	HrInit                uintptr
	AddTab                uintptr
	DeleteTab             uintptr
	ActivateTab           uintptr
	SetActiveAlt          uintptr
	MarkFullscreenWindow  uintptr
	SetProgressValue      uintptr
	SetProgressState      uintptr
	RegisterTab           uintptr
	UnregisterTab         uintptr
	SetTabOrder           uintptr
	SetTabActive          uintptr
	ThumbBarAddButtons    uintptr
	ThumbBarUpdateButtons uintptr
	ThumbBarSetImageList  uintptr
	SetOverlayIcon        uintptr
	SetThumbnailTooltip   uintptr
	SetThumbnailClip      uintptr
}

type ITaskbarList3 struct {
	LpVtbl *ITaskbarList3Vtbl
}

type IFileOpenDialogVtbl struct {
	QueryInterface      uintptr
	AddRef              uintptr
	Release             uintptr
	Show                uintptr
	SetFileTypes        uintptr
	SetFileTypeIndex    uintptr
	GetFileTypeIndex    uintptr
	Advise              uintptr
	Unadvise            uintptr
	SetOptions          uintptr
	GetOptions          uintptr
	SetDefaultFolder    uintptr
	SetFolder           uintptr
	GetFolder           uintptr
	GetCurrentSelection uintptr
	SetFileName         uintptr
	GetFileName         uintptr
	SetTitle            uintptr
	SetOkButtonLabel    uintptr
	SetFileNameLabel    uintptr
	GetResult           uintptr
	AddPlace            uintptr
	SetDefaultExtension uintptr
	Close               uintptr
	SetClientGuid       uintptr
	ClearClientData     uintptr
	SetFilter           uintptr
	GetResults          uintptr
	GetSelectedItems    uintptr
}

type IFileOpenDialog struct {
	LpVtbl *IFileOpenDialogVtbl
}

type IShellItem struct {
	LpVtbl *IShellItemVtbl
}

type IShellItemVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr
	GetAttributes  uintptr
	Compare        uintptr
}

func (obj *ITaskbarList3) SetProgressState(hwnd HWND, state int) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetProgressState, 3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(state))
	return HRESULT(ret)
}

func (obj *IFileOpenDialog) GetOptions(fos *int) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetOptions, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(fos)),
		0)
	return HRESULT(ret)
}

func (obj *IFileOpenDialog) SetOptions(fos int) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetOptions, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fos),
		0)
	return HRESULT(ret)
}

func (obj *IFileOpenDialog) SetFolder(ppsi *IShellItem) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.SetFolder, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppsi)),
		0)
	return HRESULT(ret)
}

func (obj *IFileOpenDialog) GetResult(ppsi **IShellItem) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetResult, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(ppsi)),
		0)
	return HRESULT(ret)
}

func (obj *IFileOpenDialog) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IFileOpenDialog) Show(hwnd HWND) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Show, 2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		0)
	return HRESULT(ret)
}

func (obj *IShellItem) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.Release, 1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0)
	return uint32(ret)
}

func (obj *IShellItem) GetDisplayName(sigdnName uintptr, ppszName **uint16) HRESULT {
	ret, _, _ := syscall.Syscall(obj.LpVtbl.GetDisplayName, 3,
		uintptr(unsafe.Pointer(obj)),
		sigdnName,
		uintptr(unsafe.Pointer(ppszName)))

	return HRESULT(ret)
}
