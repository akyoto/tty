// +build appengine js nacl

package tty

// IsTerminal returns true if the file descriptor is a terminal which
// is always false on js and appengine classic which is a sandboxed PaaS.
func IsTerminal(fd uintptr) bool {
	return false
}
