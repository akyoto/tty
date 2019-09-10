// +build !windows

package tty

import (
	"os"
	"testing"
)

func TestTerminal(t *testing.T) {
	// test for non-panic
	IsTerminal(os.Stdout.Fd())
}
