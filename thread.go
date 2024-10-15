package main

import (
	"context"
	"encoding/hex"
	"log"
	"regexp"
	"sync"
)

//var regs []*regexp.Regexp = make([]*regexp.Regexp, 10)

func getReg(type_ string) []*regexp.Regexp {
	if type_ == "eth" {
		regs := make([]*regexp.Regexp, 0)
		for _, r := range Config.ETHReg {
			regs = append(regs, regexp.MustCompile(r))
		}
		//log.Println("Regs: ", regs)
		return regs
	} else if type_ == "tron" {
		regs := make([]*regexp.Regexp, 0)
		for _, r := range Config.TronReg {
			regs = append(regs, regexp.MustCompile(r))
		}
		//log.Println("Regs: ", regs)
		return regs
	} else {
		panic("Invalid reg type")
	}
}

func GetMatcher(type_ string) []Matcher {
	if type_ == "eth" {
		matchers := make([]Matcher, 2)
		matchers[0] = NewETHMatcher(Config.ETHPrefix, Config.ETHSuffix)
		matchers[1] = NewETHMatcher(Config.ETHPS[0], Config.ETHPS[1])
		return matchers
	} else if type_ == "tron" {
		matchers := make([]Matcher, 2)
		matchers[0] = NewTronMatcher(Config.TronPrefix, Config.TronSuffix)
		matchers[1] = NewTronMatcher(Config.TronPS[0], Config.TronPS[1])
		return matchers
	} else {
		panic("Invalid reg type")
	}
}

func calculateETHKey(regs []*regexp.Regexp, matcher []Matcher) {
	pri, add := GenKeyETHWallet()
	if matcher[0].MatchOne(add) || matcher[1].MatchAll(add) {
		addStr := hex.EncodeToString(add)
		log.Println("Find ETH Key: ", addStr)
		err := writeResult(pri, addStr)
		if err != nil {
			log.Println(err)
		}
		return
	}

	addStr := hex.EncodeToString(add)
	for _, r := range regs {
		if r.MatchString(addStr) {
			log.Println("Find ETH Key: ", addStr)
			err := writeResult(pri, addStr)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func calculateETHPhrase(regs []*regexp.Regexp, matcher []Matcher, len int) {
	phrase, add := GenPhraseETHWallet(len)
	if matcher[0].MatchOne(add) || matcher[1].MatchAll(add) {
		addStr := hex.EncodeToString(add)
		log.Println("Find ETH Phrase: ", addStr)
		err := writeResult(phrase, addStr)
		if err != nil {
			log.Println(err)
		}
		return
	}

	addStr := hex.EncodeToString(add)
	for _, r := range regs {
		if r.MatchString(addStr) {
			log.Println("Find ETH Phrase: ", addStr)
			err := writeResult(phrase, addStr)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func threadWithETHPhrase(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Println("ETH Phrase Thread", id, "started")
	var i uint64 = 0
	length := Config.Length * 11 * 32 / 33
	regs := getReg("eth")
	matcher := GetMatcher("eth")
	for {
		select {
		case <-ctx.Done():
			log.Println("ETH Phrase Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateETHPhrase(regs, matcher, length)
		}
	}
}

func threadWithETHKey(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Println("ETH Key Thread", id, "started")
	var i uint64 = 0
	regs := getReg("eth")
	matcher := GetMatcher("eth")
	for {
		select {
		case <-ctx.Done():
			log.Println("ETH Key Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateETHKey(regs, matcher)
		}
	}
}

func calculateTronKey(regs []*regexp.Regexp, matcher []Matcher) {
	pri, add := GenKeyTronWallet()
	if matcher[0].MatchOne(add) || matcher[1].MatchAll(add) {
		addStr := string(add)
		log.Println("Find Tron Key: ", addStr)
		err := writeResult(pri, addStr)
		if err != nil {
			log.Println(err)
		}
		return
	}

	addStr := string(add)
	for _, r := range regs {
		if r.MatchString(addStr) {
			log.Println("Find Tron Key: ", addStr)
			err := writeResult(pri, addStr)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func calculateTronPhrase(regs []*regexp.Regexp, matcher []Matcher, len int) {
	phrase, add := GenPhraseTronWallet(len)
	if matcher[0].MatchOne(add) || matcher[1].MatchAll(add) {
		addStr := string(add)
		log.Println("Find Tron Phrase: ", addStr)
		err := writeResult(phrase, addStr)
		if err != nil {
			log.Println(err)
		}
		return
	}

	addStr := string(add)
	for _, r := range regs {
		if r.MatchString(addStr) {
			log.Println("Find Tron Phrase: ", addStr)
			err := writeResult(phrase, addStr)
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}

func threadWithTronPhrase(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Println("Tron Phrase Thread", id, "started")
	var i uint64 = 0
	length := Config.Length * 11 * 32 / 33
	regs := getReg("tron")
	matcher := GetMatcher("tron")
	for {
		select {
		case <-ctx.Done():
			log.Println("Tron Phrase Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateTronPhrase(regs, matcher, length)
		}
	}
}

func threadWithTronKey(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()
	log.Println("Tron Key Thread", id, "started")
	var i uint64 = 0
	regs := getReg("tron")
	matcher := GetMatcher("tron")
	for {
		select {
		case <-ctx.Done():
			log.Println("Tron Key Thread ", id, " exited, calculate: ", i)
			return
		default:
			i++
			calculateTronKey(regs, matcher)
		}
	}
}
