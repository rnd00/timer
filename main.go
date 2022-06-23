package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/rnd00/timer/ctx"
	"github.com/rnd00/timer/logger"
	"github.com/rnd00/timer/notify"
	"github.com/rnd00/timer/tick"
)

func main() {
	logger.Println("Starting")

	// context
	logger.Println("Building essentials (context and waitgroup)")
	ctx, stop := ctx.New()
	defer stop()
	var wg sync.WaitGroup

	// logger
	logger.Println("Initiating logger instance")
	cl := logger.New(ctx, &wg)
	wg.Add(1)
	go cl.Run()

	// notification msg
	cl.Send("Building notification msg")
	title := "Hourly Reminder"
	msg := fmt.Sprintf("Istirahat, cek pesan")
	notification := notify.New(title, msg)

	// run ticker
	cl.Send("Initiating Ticker instance")
	wg.Add(1)
	duration := time.Hour
	// duration := time.Hour / 2
	// duration := time.Second * 2
	cl.Send(fmt.Sprintf("Duration for ticker will be for %+v", duration))

	ticker := time.NewTicker(duration)
	tickerObj := tick.New(ctx, &wg, cl, ticker, notification)
	go tickerObj.Ticking()
	// Ticking(ctx, &wg, ticker, notification)

	// wait on the end
	wg.Wait()
	logger.Println("Ticker stopped, done triggered")
}
