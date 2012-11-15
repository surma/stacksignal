package stacksignal

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"
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
