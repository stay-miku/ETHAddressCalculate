package main

import (
	"context"
	"log"
	"regexp"
	"sync"
)

//var regs []*regexp.Regexp = make([]*regexp.Regexp, 10)

func getReg() []*regexp.Regexp {
	regs := make([]*regexp.Regexp, 0)
	for _, r := range Config.Reg {
		regs = append(regs, regexp.MustCompile(r))
	}
	//log.Println("Regs: ", regs)
	return regs
}

func calculateETHKey(regs []*regexp.Regexp) {
	pri, add := GenKeyETHWallet()
	//log.Println("add", add)
	for _, r := range regs {
		//log.Println("add", add, "reg", r)
		if r.MatchString(add) {
			log.Println("Find: ", add)
			err := writeResult(pri, add)
			if err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func calculateETHPhrase(regs []*regexp.Regexp, len int) {
	phrase, add := GenPhraseETHWallet(len)
	for _, r := range regs {
		if r.MatchString(add) {
			log.Println("Find: ", add)
			err := writeResult(phrase, add)
			if err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func threadWithETHPhrase(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var i uint64 = 0
	length := Config.Length * 11 * 32 / 33
	regs := getReg()
	for {
		select {
		case <-ctx.Done():
			log.Println("Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateETHPhrase(regs, length)
		}
	}
}

func threadWithETHKey(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var i uint64 = 0
	regs := getReg()
	for {
		select {
		case <-ctx.Done():
			log.Println("Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateETHKey(regs)
		}
	}
}

func calculateTronKey(regs []*regexp.Regexp) {
	pri, add := GenKeyTronWallet()
	for _, r := range regs {
		if r.MatchString(add) {
			log.Println("Find: ", add)
			err := writeResult(pri, add)
			if err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func calculateTronPhrase(regs []*regexp.Regexp, len int) {
	phrase, add := GenPhraseTronWallet(len)
	for _, r := range regs {
		if r.MatchString(add) {
			log.Println("Find: ", add)
			err := writeResult(phrase, add)
			if err != nil {
				log.Println(err)
			}
			break
		}
	}
}

func threadWithTronPhrase(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var i uint64 = 0
	length := Config.Length * 11 * 32 / 33
	regs := getReg()
	for {
		select {
		case <-ctx.Done():
			log.Println("Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateTronPhrase(regs, length)
		}
	}
}

func threadWithTronKey(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	var i uint64 = 0
	regs := getReg()
	for {
		select {
		case <-ctx.Done():
			log.Println("Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateTronKey(regs)
		}
	}
}
