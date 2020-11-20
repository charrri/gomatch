package main
import (
    "fmt"
)
type Order struct {
	ordrCd string//订单编号
	trdngAcntCd int//交易账户代码
	bondCd string
	price int//价格
	vol int//量
	dir int//方向
}
var ch = make(chan *Order, 10000)
func orderSubmit(i int){
    ordr := &Order{
		trdngAcntCd: 100001,
		bondCd:"123123",
		price: 90+i%10,
		vol:1000000,
		dir:(i+1)%2}
	ch<-ordr
}


func main() {
    for i := 0; i < 3; i++ {
		go  orderSubmit(i)
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息
		fmt.Println("new order price=", msg.price,"vol=",msg.vol)
	}
}