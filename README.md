# ossignals-go
This package helps us to use signals in Golang

## How to install
```bash
$ go install github.com/renanbastos93/ossignals
```

## How to use
```go
package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/renanbastos93/ossignals"
)

func main() {
	// Show PID to you can send signals for thread main of this app
	fmt.Println("PID:", os.Getpid())

	// Create a new thread to listen signals received by operation system
	go ossignals.On(

		// Create actions to each signal
		ossignals.Actions{
			syscall.SIGTERM: func() bool {
				fmt.Println("I'm dying...")
				return true
			},
			syscall.SIGHUP: func() bool {
				// Here we can do reload a config, reload connections, or anything...
				fmt.Println("process anything...")
				return false
			},
			syscall.SIGINT: func() bool {
				fmt.Println("close by CTRL + C")
				return true
			},
		},
	)
	
	// When receiving true in the function defined in your actions thread main is finished.
	<-ossignals.Close
}
```

Case you don't know about signals in the operating system you can [read this doc.](https://man7.org/linux/man-pages/man7/signal.7.html)

