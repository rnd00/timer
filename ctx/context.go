package ctx

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func New() (context.Context, context.CancelFunc) {
	return signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
}
