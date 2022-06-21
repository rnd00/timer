package tick

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/rnd00/timer/notify"
)

type ticker struct {
	Context       context.Context
	WaitGroup     *sync.WaitGroup
	TickerObject  *time.Ticker
	MessageObject notify.Message
}

type Ticker interface {
	Ticking()
}

func New(ctx context.Context, wg *sync.WaitGroup, tickerObj *time.Ticker, msgObj notify.Message) Ticker {
	return &ticker{
		Context:       ctx,
		WaitGroup:     wg,
		TickerObject:  tickerObj,
		MessageObject: msgObj,
	}
}

func (t *ticker) Ticking() {
	defer t.WaitGroup.Done()
	log.Printf("Ticking Start")
	for {
		select {
		case <-t.Context.Done():
			log.Println("Breaking the ticking loop")
			t.TickerObject.Stop()
			return
		case <-t.TickerObject.C:
			log.Printf("SendingNotification:\n\tTITLE: %s\n\tTEXT: %s", t.MessageObject.GetTitle(), t.MessageObject.GetText())
			if err := t.MessageObject.Notify(); err != nil {
				log.Printf("Error while sending message; continue to the next loop")
				continue
			}
			log.Println("Notification printed without error")
		}
	}
}
