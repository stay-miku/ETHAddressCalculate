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
	log.Println("Start")
	log.Println("Thread number: ", Config.Thread)
	log.Println("Run time: ", Config.RunTime)
	log.Println("Output file: ", Config.Output)
	log.Println("Chain: ", Config.Chain)
	log.Println("Type: ", Config.Type)
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
		if Config.Chain == "eth" {
			if Config.Type == "private key" {
				go threadWithETHKey(ctx, &wg, i)
			} else if Config.Type == "secret phrase" {
				go threadWithETHPhrase(ctx, &wg, i)
			} else {
				panic("Invalid type")
			}
		} else if Config.Chain == "tron" {
			if Config.Type == "private key" {
				go threadWithTronKey(ctx, &wg, i)
			} else if Config.Type == "secret phrase" {
				go threadWithTronPhrase(ctx, &wg, i)
			} else {
				panic("Invalid type")
			}
		} else {
			panic("Invalid chain")
		}
		wg.Add(1)
	}

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
