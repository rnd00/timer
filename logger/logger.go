package logger

import (
	"context"
	"io"
	"log"
	"os"
	"sync"
)

type logger struct {
	Context   context.Context
	WaitGroup *sync.WaitGroup
	Message   chan string
}

type Logger interface {
	Run()
	Send(msg string)
}

func New(ctx context.Context, wg *sync.WaitGroup) Logger {
	// make outputfile
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open/create log file", err)
	}
	// set multiwriter
	multiwriter := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multiwriter)

	return &logger{
		Context:   ctx,
		WaitGroup: wg,
		Message:   make(chan string),
	}
}

func (l *logger) Run() {
	defer l.WaitGroup.Done()
	for {
		select {
		case <-l.Context.Done():
			// return here, add empty line
			log.Println("Logger: Done triggered")
			log.Println("Breaking logger loop")
			return
		case msg, ok := <-l.Message:
			if ok {
				log.Println(msg)
			}
		}
	}
}

func (l *logger) Send(msg string) {
	l.Message <- msg
}

// Shorthand for log.Println()
func Println(a ...any) {
	log.Println(a...)
}
