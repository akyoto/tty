package tty_test

import (
	"fmt"
	"os"

	"github.com/akyoto/tty"
)

func Example() {
	if tty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("Is a terminal")
	} else {
		fmt.Println("Is not a terminal")
	}
}
