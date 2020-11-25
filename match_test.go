package main

import (
	"math/rand"
	"testing"
	"time"
)

/*
test sample
cmd:
	go test -bench=.
 */
func BenchmarkAddOrder(B *testing.B) {
	for i := range bondChs {
		bondChs[i] = make(chan *Order, 10000)
		go AddOrdrMpi(bondChs[i])
	}
	var cnt uint64 = 0

	for i:=0;i<B.N;i++ {
		ordr := &Order{
			trdngAcntCd: 100000 + rand.Intn(100000),
			bondCd:      1 + rand.Intn(PRDUCT_NUM - 1),
			price:       90 + rand.Intn(10),
			//price: 90,
			vol: 1000000 + 1000000*rand.Intn(5),
			dir: rand.Intn(2)}
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		//AddOrder(ordr)
		bondChs[ordr.bondCd-1] <- ordr
		cnt++
	}

	for {
		if (ops >= cnt) {
			break
		}
		time.Sleep(time.Millisecond*100)
	}

	//time.Sleep(time.Millisecond * 10000)
}
