package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, syscall.SIGTERM)

	go func() {
		<-signalChanel
		os.Exit(1)
	}()

	for {
		fmt.Println("sleeping...")
		time.Sleep(5 * time.Second)
	}
}
