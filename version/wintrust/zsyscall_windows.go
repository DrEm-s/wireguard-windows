// Code generated by 'go generate'; DO NOT EDIT.

package wintrust

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modwintrust = windows.NewLazySystemDLL("wintrust.dll")
	modcrypt32  = windows.NewLazySystemDLL("crypt32.dll")

	procWinVerifyTrust   = modwintrust.NewProc("WinVerifyTrust")
	procCryptQueryObject = modcrypt32.NewProc("CryptQueryObject")
)

func WinVerifyTrust(hWnd windows.Handle, actionId *windows.GUID, data *WinTrustData) (err error) {
	r1, _, e1 := syscall.Syscall(procWinVerifyTrust.Addr(), 3, uintptr(hWnd), uintptr(unsafe.Pointer(actionId)), uintptr(unsafe.Pointer(data)))
	if r1 != 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func cryptQueryObject(objectType uint32, object uintptr, expectedContentTypeFlags uint32, expectedFormatTypeFlags uint32, flags uint32, msgAndCertEncodingType *uint32, contentType *uint32, formatType *uint32, certStore *windows.Handle, msg *windows.Handle, context *uintptr) (err error) {
	r1, _, e1 := syscall.Syscall12(procCryptQueryObject.Addr(), 11, uintptr(objectType), uintptr(object), uintptr(expectedContentTypeFlags), uintptr(expectedFormatTypeFlags), uintptr(flags), uintptr(unsafe.Pointer(msgAndCertEncodingType)), uintptr(unsafe.Pointer(contentType)), uintptr(unsafe.Pointer(formatType)), uintptr(unsafe.Pointer(certStore)), uintptr(unsafe.Pointer(msg)), uintptr(unsafe.Pointer(context)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
