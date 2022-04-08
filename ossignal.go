package ossignals

import (
	"fmt"
	"os"
	"os/signal"
)

// Actions are a type of map to specify what signals you need to treat and execute your function
type Actions map[os.Signal]func() bool

var (
	signals       = make([]os.Signal, 0)
	currentSignal = make(chan os.Signal, 1)

	// ErrHaventActions is an error to show when you do not define Actions
	ErrHaventActions = fmt.Errorf("haven't actions to listen to signals")

	// Close is a channel to handle when finish thread
	Close = make(chan struct{}, 1)
)

// On is a listen to signals that we received from the operating system
// then we go to validate if the signal is treated and execute instruction created by you.
func On(acts Actions) {
	if len(acts) < 1 {
		panic(ErrHaventActions)
	}

	signal.Notify(currentSignal, signals...)
	for {
		if fn, ok := acts[<-currentSignal]; ok {
			if fn() {
				break
			}
		}
	}

	Close <- struct{}{}
}
