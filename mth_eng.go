package main

import (
	. "./common_order"
	. "./match"
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

//var ch = make(chan *Order, 10000)
var bondChs [PRDUCT_NUM]chan *Order
var gAddedOrderCnt uint64 = 0

func AddOrdrMpi(in chan *Order) {
	for {
		ordr := <-in
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		AddOrder(ordr)
		atomic.AddUint64(&gAddedOrderCnt, 1)
		//fmt.Println(ops)
	}
}

func PutOrdr2Queue(ordr *Order) {
	bondChs[ordr.BondCd-1] <- ordr
}

func Monitor() {
	for {
		var curCnt, uintCnt uint64 = 10000, 10000
		if (gAddedOrderCnt >= curCnt) {
			fmt.Printf("order:%v, trade:%v\n", gAddedOrderCnt, GTradeCnt)
			for curCnt < gAddedOrderCnt {
				curCnt += uintCnt
			}
		}
		time.Sleep(time.Millisecond*20)
	}
}

func main() {

	for i := range bondChs {
		bondChs[i] = make(chan *Order, 10000)
		go AddOrdrMpi(bondChs[i])
	}

	// for i := 0; i < 10; i++ {
	// 	go orderSubmit(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	msg := <-ch // 等待信道返回消息
	// 	fmt.Printf("new order dir=%d, price=%d, vol=%d\n", msg.dir, msg.price, msg.vol)
	// 	AddOrder(msg)
	// }
	go Monitor()

	for i := 0; i < ORDR_NUM; i++ {
		ordr := &Order{
			TrdngAcntCd: 100001+ rand.Intn(10),
			BondCd:      1 + rand.Intn(PRDUCT_NUM-1),
			Price:       90 + rand.Intn(10),
			//price: 90,
			Vol: 1000000 + 1000000*rand.Intn(5),
			Dir: rand.Intn(2)}

		PutOrdr2Queue(ordr)
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		//AddOrder(ordr)
	}

	for {
		if (gAddedOrderCnt == ORDR_NUM) {
			break
		}
		time.Sleep(time.Millisecond*20)
	}
	/*
	for i:= range [PRDUCT_NUM]int{} {
		ShowPrcLink(BUY, i)
		ShowPrcLink(SELL, i)
	}
	 */
}