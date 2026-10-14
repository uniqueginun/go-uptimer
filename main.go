package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
	"uptimer/bootstrap"
)

func main() {
	app := bootstrap.Boot()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	// app.Start blocks until shutdown, so the ticker must run in its own goroutine.
	go func() {
		tk := time.NewTicker(5 * time.Second)
		defer tk.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-tk.C:
				_, _ = fmt.Fprintf(os.Stdout, "running at: [%s]\n", time.Now().Format(time.RFC3339))
			}
		}
	}()

	app.Start()
}
