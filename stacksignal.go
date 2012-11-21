/*
Package stacksignal registers a signal handler for SIGUSR1.
Whenever the importing program receives SIGUSR1, a full
stacktrace of all goroutines will be printed to StdErr.

Since there are no functions to use, the import line should
be

	import _ "github.com/surma/stacksignal"
*/
package stacksignal

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	VERSION = "1.1.1"
)

func init() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGUSR1)
		n := 512
		for _ = range c {
			buf := make([]byte, n)
			for {
				m := runtime.Stack(buf, true)
				if m < n {
					buf = buf[:m]
					break
				}
				n *= 2
				buf = make([]byte, n)
			}
			fmt.Fprintln(os.Stderr, "=== Begin stack trace")
			fmt.Fprintln(os.Stderr, string(buf))
			fmt.Fprintln(os.Stderr, "=== End stack trace")
		}
	}()
}
