package common_order

const (
	BUY        = 0
	SELL       = 1
	PRDUCT_NUM = 20 // 合约数量
	ORDR_NUM   = 1000000
)

type Order struct {
	ordrCd      string //订单编号
	TrdngAcntCd int    //交易账户代码
	BondCd      int
	Price       int //价格
	Vol         int //量
	Dir         int //方向
}

var ch = make(chan *Order, 10000)

func orderSubmit(i int) {
	ordr := &Order{
		TrdngAcntCd: 100001,
		BondCd:      5,
		//price:       90 + i%10,
		Price: 90,
		Vol:   1000000,
		Dir:   (i + 1) % 2}
	ch <- ordr
}

/*
func main_order() {

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
			trdngAcntCd: 100001+ rand.Intn(10),
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
		if (ops == ORDR_NUM) {
			break
		}
		time.Sleep(time.Millisecond*100)
	}
	for i:= range [PRDUCT_NUM]int{} {
		ShowPrcLink(BUY, i)
		ShowPrcLink(SELL, i)
	}
}
*/