package orderbook

import (
	"container/list"
)

// bid - buy, ask - sell

func Remove(l *list.List, elems []*list.Element) {
	for i := 0; i < len(elems); i++ {
		l.Remove(elems[i])
	}
}

type Orderbook struct {
	bidOrders *list.List
	askOrders *list.List
}

func (o1 *Order) CmpAsk(o2 *Order) int {
	if o1.Price > o2.Price {
		return 1
	}
	if o1.Price < o2.Price {
		return -1
	}
	return o1.CmpVolume(o2)
}

func (o1 *Order) CmpBid(o2 *Order) int {
	if o1.Price > o2.Price {
		return 1
	}
	if o1.Price < o2.Price {
		return -1
	}
	return o1.CmpVolumeReverse(o2)
}

func (o1 *Order) CmpVolume(o2 *Order) int {
	if o1.Volume > o2.Volume {
		return 1
	}
	if o1.Volume < o2.Volume {
		return -1
	}
	return 0
}

func (o1 *Order) CmpVolumeReverse(o2 *Order) int {
	if o1.Volume > o2.Volume {
		return -1
	}
	if o1.Volume < o2.Volume {
		return 1
	}
	return 0
}

func New() *Orderbook {
	return &Orderbook{
		bidOrders: list.New(),
		askOrders: list.New(),
	}
}

func (ob *Orderbook) InsertSortedBid(order *Order) {
	e := ob.bidOrders.Back()

	for ; e != nil && order.CmpBid(e.Value.(*Order)) < 0; e = e.Prev() {
	}

	if e == nil {
		ob.bidOrders.PushFront(order)
		return
	}

	ob.bidOrders.InsertAfter(order, e)
}

func (ob *Orderbook) InsertSortedAsk(order *Order) {
	e := ob.askOrders.Front()

	for ; e != nil && order.CmpAsk(e.Value.(*Order)) > 0; e = e.Next() {
	}

	if e == nil {
		ob.askOrders.PushBack(order)
		return
	}

	ob.askOrders.InsertBefore(order, e)
}

func (ob *Orderbook) Cancel(id int) bool {
	for e := ob.askOrders.Front(); e != nil; e = e.Next() {
	}

	return true
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	if order.Kind == KindLimit {
		if order.Side == SideBid {
			return orderbook.matchLimitBid(order)
		} else {
			return orderbook.matchLimitAsk(order)
		}
	} else {
		if order.Side == SideBid {
			return orderbook.matchMarketBid(order)
		} else {
			return orderbook.matchMarketAsk(order)
		}
	}
}

func fillOrders(askOrder *Order, bidOrder *Order, price uint64) *Trade {
	trade := &Trade{
		Bid:   bidOrder,
		Ask:   askOrder,
		Price: price,
	}

	if askOrder.Volume > bidOrder.Volume {
		trade.Volume = bidOrder.Volume
		askOrder.Volume -= bidOrder.Volume
		bidOrder.Volume = 0
	} else if askOrder.Volume < bidOrder.Volume {
		trade.Volume = askOrder.Volume
		bidOrder.Volume -= askOrder.Volume
		askOrder.Volume = 0
	} else {
		trade.Volume = askOrder.Volume
		askOrder.Volume = 0
		bidOrder.Volume = 0
	}

	return trade
}

func (ob *Orderbook) matchMarketBid(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	restingOrdersForRemove := []*list.Element{}

	for e := ob.askOrders.Front(); e != nil && order.Volume > 0; e = e.Next() {
		askOrder := e.Value.(*Order)

		trades = append(trades, fillOrders(askOrder, order, askOrder.Price))

		if askOrder.Volume == 0 {
			restingOrdersForRemove = append(restingOrdersForRemove, e)
		}
	}

	Remove(ob.askOrders, restingOrdersForRemove)

	if len(trades) == 0 {
		trades = nil
	}

	if order.Volume > 0 {
		return trades, order
	}
	return trades, nil
}

func (ob *Orderbook) matchMarketAsk(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	restingOrdersForRemove := []*list.Element{}

	for e := ob.bidOrders.Back(); e != nil && order.Volume > 0; e = e.Prev() {
		bidOrder := e.Value.(*Order)

		trades = append(trades, fillOrders(order, bidOrder, bidOrder.Price))

		if bidOrder.Volume == 0 {
			restingOrdersForRemove = append(restingOrdersForRemove, e)
		}
	}

	Remove(ob.bidOrders, restingOrdersForRemove)

	if len(trades) == 0 {
		trades = nil
	}

	if order.Volume > 0 {
		return trades, order
	}
	return trades, nil
}

func (ob *Orderbook) matchLimitBid(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	restingOrdersForRemove := []*list.Element{}

	for e := ob.askOrders.Front(); e != nil && order.Volume > 0; e = e.Next() {
		askOrder := e.Value.(*Order)

		if askOrder.Price <= order.Price {
			trades = append(trades, fillOrders(askOrder, order, askOrder.Price))
		}

		if askOrder.Volume == 0 {
			restingOrdersForRemove = append(restingOrdersForRemove, e)
		}
	}

	Remove(ob.askOrders, restingOrdersForRemove)

	if len(trades) == 0 {
		trades = nil
	}

	if order.Volume > 0 {
		ob.InsertSortedBid(order)
	}
	return trades, nil
}

func (ob *Orderbook) matchLimitAsk(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}
	restingOrdersForRemove := []*list.Element{}

	for e := ob.bidOrders.Back(); e != nil && order.Volume > 0; e = e.Prev() {
		bidOrder := e.Value.(*Order)

		if bidOrder.Price < order.Price {
			break
		}

		trades = append(trades, fillOrders(bidOrder, order, bidOrder.Price))

		if bidOrder.Volume == 0 {
			restingOrdersForRemove = append(restingOrdersForRemove, e)
		}
	}

	Remove(ob.bidOrders, restingOrdersForRemove)

	if len(trades) == 0 {
		trades = nil
	}

	if order.Volume > 0 {
		ob.InsertSortedAsk(order)
	}
	return trades, nil
}
