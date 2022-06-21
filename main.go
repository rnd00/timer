package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/rnd00/timer/notify"
)

func main() {
	// sigint
	ctrlc := make(chan os.Signal, 1)
	signal.Notify(ctrlc, os.Interrupt)

	// ticker done
	var wg sync.WaitGroup
	done := make(chan bool)

	// notification msg
	title := "Hourly Reminder"
	msg := fmt.Sprintf("Istirahat, cek pesan")

	// run ticker
	wg.Add(1)
	ticker := Ticking(done, title, msg)

	// check sigint
	go CaptureSigint(ctrlc, ticker, done, &wg)

	// wait
	wg.Wait()
	log.Println("Ticker stopped, done triggered")
}

func Ticking(done chan bool, title, message string) *time.Ticker {
	ticker := time.NewTicker(time.Hour / 2)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				log.Println("Sending notification;", title, message)
				notify.Notify(title, message)
			}
		}
	}()

	return ticker
}

func CaptureSigint(ctrlc chan os.Signal, ticker *time.Ticker, trigger chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for range ctrlc {
		fmt.Println("\nSIGINT Triggered!")
		ticker.Stop()
		trigger <- true
		return
	}
}
