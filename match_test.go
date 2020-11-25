package main

import (
	"math/rand"
	"testing"
)

/*
test sample
cmd:
	go test -bench=.
 */
func BenchmarkAddOrder(B *testing.B) {
	for i:=0;i<B.N;i++ {
		ordr := &Order{
			trdngAcntCd: 100000 + rand.Intn(100000),
			bondCd:      1 + rand.Intn(PRDUCT_NUM - 1),
			price:       90 + rand.Intn(10),
			//price: 90,
			vol: 1000000 + 1000000*rand.Intn(5),
			dir: rand.Intn(2)}
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		AddOrder(ordr)
	}
}
