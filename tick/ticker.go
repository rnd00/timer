package tick

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/rnd00/timer/logger"
	"github.com/rnd00/timer/notify"
)

type ticker struct {
	Context       context.Context
	WaitGroup     *sync.WaitGroup
	TickerObject  *time.Ticker
	MessageObject notify.Message
	LoggerObject  logger.Logger
}

type Ticker interface {
	Ticking()
}

// New will make a new struct of ticker
// which later can be used to run reminders
func New(ctx context.Context, wg *sync.WaitGroup, cl logger.Logger, tickerObj *time.Ticker, msgObj notify.Message) Ticker {
	return &ticker{
		Context:       ctx,
		WaitGroup:     wg,
		TickerObject:  tickerObj,
		MessageObject: msgObj,
		LoggerObject:  cl,
	}
}

// Ticking will run and blocking to wait for channels
func (t *ticker) Ticking() {
	defer t.WaitGroup.Done()
	t.LoggerObject.Send("Ticking Start")
	for {
		select {
		case <-t.Context.Done():
			// add the empty line first
			logger.Println("Ticker: Done triggered")
			logger.Println("Breaking the ticking loop")
			t.TickerObject.Stop()
			return
		case <-t.TickerObject.C:
			t.LoggerObject.Send(fmt.Sprintf("SendingNotification:\n\tTITLE: %s\n\tTEXT: %s", t.MessageObject.GetTitle(), t.MessageObject.GetText()))
			if err := t.MessageObject.Notify(); err != nil {
				t.LoggerObject.Send(fmt.Sprintf("Error while sending message; continue to the next loop"))
				continue
			}
			t.LoggerObject.Send("Notification printed without error")
		}
	}
}
