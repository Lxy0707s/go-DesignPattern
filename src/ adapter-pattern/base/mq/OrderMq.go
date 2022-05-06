package mq

import "fmt"

type OrderMq struct {
	uid             string // 用户ID
	sku             string // 商品
	orderId         string // 订单ID
	createOrderTime string // 下单时间
}

var OrderMqInstance *OrderMq

func NewOrderMq() *OrderMq {
	if OrderMqInstance == nil {
		OrderMqInstance = &OrderMq{
			uid:             "000",
			sku:             "test",
			orderId:         "000",
			createOrderTime: "0-0-0-0",
		}
	}
	return OrderMqInstance
}

func (o *OrderMq) Setuid(uid string) {
	o.uid = uid
}

func (o *OrderMq) SetSku(sku string) {
	o.sku = sku
}

func (o *OrderMq) SetOrderId(orderId string) {
	o.orderId = orderId
}

func (o *OrderMq) SetCreateOrderTime(createOrderTime string) {
	o.createOrderTime = createOrderTime
}

func (o *OrderMq) Getuid() string {
	return o.uid
}

func (o *OrderMq) GetSku() string {
	return o.sku
}

func (o *OrderMq) GetOrderId() string {
	return o.orderId
}

func (o *OrderMq) GetCreateOrderTime() string {
	return o.createOrderTime
}

func (o *OrderMq) ToString() string {
	return fmt.Sprint(o)
}
