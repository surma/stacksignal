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
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

const (
	VERSION = "1.0.0"
)

func init() {
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGUSR1)
		buf := make([]byte, 1024*16)
		for _ = range c {
			runtime.Stack(buf, true)
			os.Stderr.Write(buf)
		}
	}()
}
