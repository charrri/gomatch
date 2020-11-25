package main

import (
	"math/rand"
	"sync/atomic"
	"time"
)

type Order struct {
	ordrCd      string //订单编号
	trdngAcntCd int    //交易账户代码
	bondCd      int
	price       int //价格
	vol         int //量
	dir         int //方向
}

var ch = make(chan *Order, 10000)
var bondChs [PRDUCT_NUM]chan *Order
var ops uint64 = 0

func orderSubmit(i int) {
	ordr := &Order{
		trdngAcntCd: 100001,
		bondCd:      5,
		//price:       90 + i%10,
		price: 90,
		vol:   1000000,
		dir:   (i + 1) % 2}
	ch <- ordr
}

func AddOrdrMpi(in chan *Order) {
	for {
		ordr := <-in
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		AddOrder(ordr)
		atomic.AddUint64(&ops, 1)
		//fmt.Println(ops)
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

	for i := 0; i < ORDR_NUM; i++ {
		ordr := &Order{
			trdngAcntCd: 100001 + rand.Intn(10),
			bondCd:      1 + rand.Intn(PRDUCT_NUM-1),
			price:       90 + rand.Intn(10),
			//price: 90,
			vol: 1000000 + 1000000*rand.Intn(5),
			dir: rand.Intn(2)}

		bondChs[ordr.bondCd-1] <- ordr
		//fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		//AddOrder(ordr)
	}

	for {
		if ops == ORDR_NUM {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
	for i := range [PRDUCT_NUM]int{} {
		ShowPrcLink(BUY, i)
		ShowPrcLink(SELL, i)
	}
}
