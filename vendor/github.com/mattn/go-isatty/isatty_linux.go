// +build linux
<<<<<<< HEAD
// +build !appengine,!ppc64,!ppc64le
=======
// +build !appengine
>>>>>>> b5201c34e840e2ec911a64aedeb052cd36fcd58a

package isatty

import (
	"syscall"
	"unsafe"
)

const ioctlReadTermios = syscall.TCGETS

// IsTerminal return true if the file descriptor is terminal.
func IsTerminal(fd uintptr) bool {
	var termios syscall.Termios
	_, _, err := syscall.Syscall6(syscall.SYS_IOCTL, fd, ioctlReadTermios, uintptr(unsafe.Pointer(&termios)), 0, 0, 0)
	return err == 0
}
