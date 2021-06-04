package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	finish := make(chan int)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() {
		for {
			_, ok := <-ctx.Done()

			if !ok {
				//time.Sleep(2 * time.Second)
				close(finish)
				return
			}
		}
	}()

	sig := <-sigs
	fmt.Println(sig)
	if sig == os.Interrupt {
		cancel()
		timeoutChan := time.NewTicker(1 * time.Second)

		select {
		case _, ok := <-finish:
			if !ok {
				fmt.Println("Горутина закончилась вовремя")
			}
		case <-timeoutChan.C:
			fmt.Println("Вышло время")
		}
	}
}
