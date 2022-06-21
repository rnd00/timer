package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/rnd00/timer/ctx"
	"github.com/rnd00/timer/notify"
	"github.com/rnd00/timer/tick"
)

func main() {
	// context
	ctx, stop := ctx.New()
	defer stop()
	var wg sync.WaitGroup

	// notification msg
	title := "Hourly Reminder"
	msg := fmt.Sprintf("Istirahat, cek pesan")

	notification := notify.New(title, msg)

	// run ticker
	wg.Add(1)
	// duration := time.Hour / 2
	duration := time.Second * 2
	ticker := time.NewTicker(duration)
	tickerObj := tick.New(ctx, &wg, ticker, notification)
	go tickerObj.Ticking()
	// Ticking(ctx, &wg, ticker, notification)

	// wait
	wg.Wait()
	log.Println("Ticker stopped, done triggered")
}
