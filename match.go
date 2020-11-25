// match
package main

import (
	"container/list"
	"fmt"

	"github.com/emirpasic/gods/maps/treemap"
)

// type prcLinks struct {
// 	m   map[int]list.List
// 	key []int
// }

type Trade struct {
	trdCd           string
	buyTrdngAcntCd  int //买方交易账户代码
	selltrdngAcntCd int //卖方交易账户代码
	bondCd          int
	price           int //价格
	vol             int //量
	dir             int //方向
}

const (
	BUY        = 0
	SELL       = 1
	PRDUCT_NUM = 20 // 合约数量
)

var prcLinks [2]map[int]*treemap.Map // int.List

//var prcLinks *treemap.Map

func AddOrder(ordr *Order) {
	match(ordr)
	if (ordr.vol > 0) {
		addOrdrToSameDir(ordr)
	}else {
		//to do trd txnlog msg
		//to do del txnlog msg
	}

	ShowPrcLink(BUY, ordr.bondCd)
	ShowPrcLink(SELL, ordr.bondCd)
}

func addOrdrToSameDir(ordr *Order) {
	//fmt.Printf("Add order: %v\n", *ordr)
	if nil == prcLinks[ordr.dir][ordr.bondCd] {
		prcLinks[ordr.dir][ordr.bondCd] = treemap.NewWithIntComparator()
	}

	link := prcLinks[ordr.dir][ordr.bondCd]
	pValue, r := link.Get(ordr.price)
	if false == r {
		l := list.New()
		l.PushBack(ordr)
		link.Put(ordr.price, l)
	} else {
		pList := pValue.(*list.List)
		pList.PushBack(ordr)
		//fmt.Println(pList)
	}
	//fmt.Println(link)
}

func CheckPriceCanDeal(ordrDir int, ordrPrc int, bkPrc int) bool {
	if ordrDir == BUY {
		return ordrPrc >= bkPrc
	} else if ordrDir == SELL {
		return ordrPrc <= bkPrc
	} else {
		fmt.Println("Error direction!")
		return false
	}
	return false
}
func CheckAccountCanDeal(ordrAcnt int,bkAcnt int) bool{
	if ordrAcnt==bkAcnt {return false}
	return true
}

func match(ordr *Order) {
	var trd Trade
	var key int
	var li *list.List
	d := reverseDir(ordr.dir)

	prcLink := prcLinks[d][ordr.bondCd]

	if d == BUY { // 最优价
		k, v := prcLink.Max()
		if k == nil {
			return
		}
		key = k.(int)
		li = v.(*list.List)
	} else {
		k, v := prcLink.Min()
		if k == nil {
			return
		}
		key = k.(int)
		li = v.(*list.List)
	}

	for {
		if CheckPriceCanDeal(ordr.dir, ordr.price, key) {
			pList := li
			//isTrade = true
			for node := pList.Front(); node != nil; node = node.Next() {
				o := node.Value.(*Order)
                if !CheckAccountCanDeal(ordr.trdngAcntCd,o.trdngAcntCd) {
                	fmt.Printf("same trading account,skip\n")
                	continue
                }
				if o.vol >= ordr.vol { // 订单完全成交
					trd.vol += ordr.vol
					o.vol -= ordr.vol
					ordr.vol = 0
					if o.vol == 0 {
						pList.Remove(node) //
					}

					//trade(ordr, o, &trd) // to do
					fmt.Printf("Trade: %v\n", trd)
					return
				} else {
					trd.vol += ordr.vol
					ordr.vol -= o.vol
					o.vol = 0
					pList.Remove(node)
				}
			}

			// 此价格成交完
		} else {
			return
		}
		key, li = getOrdrBkNextPriceLinkByInOrdrDir(ordr.dir, key, prcLink)
		if li == nil {
			return
		}
	}

	// pValue, r := prcLink.Get(ordr.price)
	// if true == r {

	// }

	return
}

func getOrdrBkNextPriceLinkByInOrdrDir(dir int, prc int, m *treemap.Map) (int, *list.List) {
	var dir_ int
	if (dir == BUY) {
		dir_ = SELL
	}else {
		dir_ = BUY
	}
	return getNextPriceLink(dir_, prc, m)
}

func getNextPriceLink(dir int, prc int, m *treemap.Map) (int, *list.List) {
	var key int
	var li *list.List
	if dir == SELL {
		k, v := m.Ceiling(prc + 1)
		if nil == k {
			return 0, nil
		}
		key = k.(int)
		li = v.(*list.List)
	} else {
		k, v := m.Floor(prc - 1)
		if nil == k {
			return 0, nil
		}
		key = k.(int)
		li = v.(*list.List)
	}

	return key, li
}

func reverseDir(d int) int {
	if d == BUY {
		return SELL
	} else {
		return BUY
	}
}

func trade(ordr *Order, bkOrdr *Order, trd *Trade) {
	// to do
}

func init() {
	//prcLinks[1][1000] = treemap.NewWithIntComparator()
	for i := 0; i < 2; i++ {
		prcLinks[i] = make(map[int]*treemap.Map)
		for j := 0; j < PRDUCT_NUM; j++ {
			prcLinks[i][j] = treemap.NewWithIntComparator()
		}
	}
}

func testmain() {
	m := treemap.NewWithIntComparator() // empty (keys are of type int)
	m.Put(3, "y")
	m.Put(1, "x") // 1->x
	v, _ := m.Get(1)
	fmt.Println(v)
	k, v := m.Ceiling(2)
	if k != nil {
		fmt.Println(k, v)
	}
	fmt.Printf("%T", m)
}

func ShowPrcLink(dir int, bondCd int) {
	prcLink := prcLinks[dir][bondCd]
	var key int
	var li *list.List

	if dir == BUY { // 最优价
		k, v := prcLink.Max()
		if k == nil {
			return
		}
		key = k.(int)
		li = v.(*list.List)
		ShowOrdrLink(li)
	} else {
		k, v := prcLink.Min()
		if k == nil {
			return
		}
		key = k.(int)
		li = v.(*list.List)
		ShowOrdrLink(li)
	}
	for {
		if key, li = getNextPriceLink(dir, key, prcLink); li == nil {
			return
		}
		ShowOrdrLink(li)
	}
}

func ShowOrdrLink(node *list.List) {
	head := node.Front()
	if head != nil {
		o := head.Value.(*Order)
		PrintPriceLeader(o)
	}else {
		return
	}

	for {
		if (head == nil) {
			fmt.Println()
			break
		}
		o := head.Value.(*Order)
		PrintOrdr(o)
		head = head.Next()
	}
}

func PrintPriceLeader(ordr *Order) {
	var dir string;
	if ordr.dir == 1 {
		dir = "Sell"
	}else {
		dir = "Buy"
	}
	fmt.Printf("{bondCd:%d, prc:%d, dir:%s}\n", ordr.bondCd, ordr.price, dir)
}

func PrintOrdr(ordr *Order) {
	fmt.Printf("[AcntCd:%d, Vol:%d]", ordr.trdngAcntCd, ordr.vol)//to do ordrCd
}