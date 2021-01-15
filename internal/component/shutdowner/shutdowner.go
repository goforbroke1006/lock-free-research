package shutdowner

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForInterrupt() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
}
