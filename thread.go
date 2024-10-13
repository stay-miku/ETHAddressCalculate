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

func calculateKey(regs []*regexp.Regexp) {
	pri, add := GenKeyWallet()
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

func calculatePhrase(regs []*regexp.Regexp, len int) {
	phrase, add := GenPhraseWallet(len)
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

func threadWithPhrase(ctx context.Context, wg *sync.WaitGroup, id int) {
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
			calculatePhrase(regs, length)
		}
	}
}

func threadWithKey(ctx context.Context, wg *sync.WaitGroup, id int) {
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
			calculateKey(regs)
		}
	}
}
