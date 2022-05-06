package mq

import "fmt"

type POPOrderDelivered struct {
	uId       string // 用户ID
	orderId   string // 订单号
	orderTime string // 下单时间
	sku       string // 商品
	skuName   string // 商品名称
	decimal   string // 金额
}

var POPOrderInstance *POPOrderDelivered

func NewPOPOrder() *POPOrderDelivered {
	if POPOrderInstance == nil {
		POPOrderInstance = &POPOrderDelivered{
			uId:       "001",
			orderId:   "001",
			orderTime: "1-1-1-1",
			sku:       "TEST",
			skuName:   "TTTT",
			decimal:   "------",
		}
	}
	return POPOrderInstance
}

func (p *POPOrderDelivered) Setuid(uid string) {
	p.uId = uid
}

func (p *POPOrderDelivered) SetOrderId(orderId string) {
	p.orderId = orderId
}

func (p *POPOrderDelivered) SetOrderTime(orderTime string) {
	p.orderTime = orderTime
}

func (p *POPOrderDelivered) SetSku(sku string) {
	p.sku = sku
}

func (p *POPOrderDelivered) SetSkuName(skuName string) {
	p.skuName = skuName
}

func (p *POPOrderDelivered) SetDecimal(decimal string) {
	p.decimal = decimal
}

func (p *POPOrderDelivered) GetOrderId() string {
	return p.orderId
}

func (p *POPOrderDelivered) GetOrderTime() string {
	return p.orderTime
}

func (p *POPOrderDelivered) GetSku() string {
	return p.sku
}

func (p *POPOrderDelivered) GetSkuName() string {
	return p.skuName
}

func (p *POPOrderDelivered) GetDecimal() string {
	return p.decimal
}

func (p *POPOrderDelivered) Getuid() string {
	return p.uId
}

func (p *POPOrderDelivered) ToString() string {
	return fmt.Sprint(p)
}
