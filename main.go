package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {
	log.Println("Start")
	log.Println("Max Go Thread: ", Config.GoThread)
	log.Println("ETH Phrase Thread: ", Config.ETHPhraseThread)
	log.Println("ETH Key Thread: ", Config.ETHKeyThread)
	log.Println("TRON Phrase Thread: ", Config.TRONPhraseThread)
	log.Println("TRON Key Thread: ", Config.TRONKeyThread)
	log.Println("Run time: ", Config.RunTime)
	log.Println("Output file: ", Config.Output)
	log.Println("Length: ", Config.Length)
	log.Println("Reg:")
	for _, r := range Config.Reg {
		log.Println(r)
	}
	log.Println("----------------------")
	err := initOutput()
	if err != nil {
		panic(err)
	}

	var ctx context.Context
	var cancel context.CancelFunc
	if Config.RunTime < 0 {
		ctx, cancel = context.WithCancel(context.Background())
	} else {
		ctx, cancel = context.WithTimeout(context.Background(), time.Duration(Config.RunTime)*time.Second)
	}

	// set go thread
	if Config.GoThread > 0 {
		log.Println("Set go thread to", Config.GoThread)
		runtime.GOMAXPROCS(Config.GoThread)
	}

	var wg sync.WaitGroup
	i := 0
	for j := 0; j < Config.ETHPhraseThread; j++ {
		go threadWithETHPhrase(ctx, &wg, i)
		i++
		wg.Add(1)
	}
	for j := 0; j < Config.ETHKeyThread; j++ {
		go threadWithETHKey(ctx, &wg, i)
		i++
		wg.Add(1)
	}
	for j := 0; j < Config.TRONPhraseThread; j++ {
		go threadWithTronPhrase(ctx, &wg, i)
		i++
		wg.Add(1)
	}
	for j := 0; j < Config.TRONKeyThread; j++ {
		go threadWithTronKey(ctx, &wg, i)
		i++
		wg.Add(1)
	}
	time.Sleep(1 * time.Second)
	log.Println("Running...  Press Ctrl+C to stop")
	log.Println("Result will display here and save to", Config.Output)

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
