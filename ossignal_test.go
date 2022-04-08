package ossignals

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestOn(t *testing.T) {
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal(err)
	}
	sig := make(chan os.Signal, 1)

	acts := Actions{
		syscall.SIGTERM: func() bool {
			<-sig
			signal.Stop(sig)
			return true
		},
		syscall.SIGHUP: func() bool {
			return false
		},
	}
	go On(acts)

	proc.Signal(syscall.SIGHUP)
	time.Sleep(time.Second * 2)

	proc.Signal(syscall.SIGTERM)
	time.Sleep(time.Second * 2)

}
