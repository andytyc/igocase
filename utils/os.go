package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func NotifySignal() (c chan os.Signal) {
	c = make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	return
}
