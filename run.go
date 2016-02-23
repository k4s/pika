package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/k4s/pika/pika"
)

func main() {
	pika.Flags()
	pika.PikaRun()
	pika.WebRun()
	forStop()

}

func forStop() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	fmt.Println("Got signal:", s)
}
