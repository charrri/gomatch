package main

import (
	. "./common_order"
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
			TrdngAcntCd: 100000 + rand.Intn(100000),
			BondCd:      1 + rand.Intn(PRDUCT_NUM - 1),
			Price:       90 + rand.Intn(10),
			//price: 90,
			Vol: 1000000 + 1000000*rand.Intn(5),
			Dir: rand.Intn(2)}
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		PutOrdr2Queue(ordr)
		//bondChs[ordr.BondCd-1] <- ordr
		cnt++
	}

	for {
		if (gAddedOrderCnt >= cnt) {
			break
		}
		time.Sleep(time.Millisecond*100)
	}

	//time.Sleep(time.Millisecond * 10000)
}
