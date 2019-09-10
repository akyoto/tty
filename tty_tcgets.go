// +build linux aix
// +build !appengine
// +build !android

package tty

import "golang.org/x/sys/unix"

// IsTerminal return true if the file descriptor is a terminal.
func IsTerminal(fd uintptr) bool {
	_, err := unix.IoctlGetTermios(int(fd), unix.TCGETS)
	return err == nil
}
