package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	log.Println("Start with thread number: ", Config.Thread, " run time: ", Config.RunTime, "s", " output: ", Config.Output, " reg: ", Config.Reg)
	err := initOutput()
	if err != nil {
		panic(err)
	}

	if Config.Thread < 1 {
		panic("Invalid thread number")
	}
	var ctx context.Context
	var cancel context.CancelFunc
	if Config.RunTime < 0 {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(Config.RunTime)*time.Second)
	}

	var wg sync.WaitGroup
	for i := 0; i < Config.Thread; i++ {
		go thread(ctx, &wg, i)
		wg.Add(1)
	}

	//go func() {
	//
	//}()

	//wg.Wait()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	select {
	case <-sigs:
		log.Println("Received signal, exit...")
		cancel()
		wg.Wait()
		return
	case <-ctx.Done():
		log.Println("Time out, exit...")
		wg.Wait()
		return
	}
}
