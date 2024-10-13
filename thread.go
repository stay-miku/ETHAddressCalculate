package main

import (
	"context"
	"fmt"
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

func calculate(regs []*regexp.Regexp) {
	pri, add := GenWallet()
	//log.Println("add", add)
	for _, r := range regs {
		//log.Println("add", add, "reg", r)
		if r.MatchString(add) {
			log.Println("Find: ", add, "reg: ", r)
			err := writeResult(pri, add)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
	}
}

func thread(ctx context.Context, wg *sync.WaitGroup, id int) {
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
			calculate(regs)
		}
	}
}
