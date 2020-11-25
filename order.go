package main
import (
	"fmt"
	"math/rand"
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

func main() {

	// for i := 0; i < 10; i++ {
	// 	go orderSubmit(i)
	// }
	// for i := 0; i < 10; i++ {
	// 	msg := <-ch // 等待信道返回消息
	// 	fmt.Printf("new order dir=%d, price=%d, vol=%d\n", msg.dir, msg.price, msg.vol)
	// 	AddOrder(msg)
	// }

	for i := 0; i < 20; i++ {
		ordr := &Order{
			trdngAcntCd: 100001+ rand.Intn(10),
			bondCd:      5,
			price:       90 + rand.Intn(10),
			//price: 90,
			vol: 1000000 + 1000000*rand.Intn(5),
			dir: rand.Intn(2)}

		fmt.Printf("new order dir=%d, price=%d, vol=%d, addr=%p\n", ordr.dir, ordr.price, ordr.vol, ordr)
		AddOrder(ordr)
	}

	//ShowPrcLink(BUY, 5)
	//ShowPrcLink(SELL, 5)
}
