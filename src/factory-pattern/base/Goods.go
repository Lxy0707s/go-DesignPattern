package base

import "fmt"

type GoodsInfo struct {
	userName             string // 用户姓名
	userPhone            string // 用户手机
	sku                  string // 商品SKU
	orderId              string // 订单ID
	consigneeUserName    string // 收货人姓名
	consigneeUserPhone   string // 收货人手机
	consigneeUserAddress string // 收获人地址
}

var goodsInstance *GoodsInfo

func NewGoods() *GoodsInfo {
	if instance == nil {
		goodsInstance = &GoodsInfo{
			userName:             "",
			userPhone:            "",
			sku:                  "",
			orderId:              "",
			consigneeUserName:    "",
			consigneeUserPhone:   "",
			consigneeUserAddress: "",
		}
	}
	return goodsInstance
}

func (g *GoodsInfo) SetUserName(name string) {
	g.userName = name
}

func (g *GoodsInfo) SetUserPhone(phone string) {
	g.userPhone = phone
}

func (g *GoodsInfo) SetSku(sku string) {
	g.sku = sku
}

func (g *GoodsInfo) SetOrderId(orderId string) {
	g.orderId = orderId
}

func (g *GoodsInfo) SetConsigneeUserName(consigneeUserAddress string) {
	g.consigneeUserAddress = consigneeUserAddress
}

func (g *GoodsInfo) SetConsigneeUserPhone(consigneeUserPhone string) {
	g.consigneeUserPhone = consigneeUserPhone
}

func (g *GoodsInfo) SetConsigneeUserAddress(consigneeUserAddress string) {
	g.consigneeUserAddress = consigneeUserAddress
}

func (g *GoodsInfo) getUserName() string {
	return g.userName
}

func (g *GoodsInfo) getUserPhone() string {
	return g.userPhone
}

func (g *GoodsInfo) getSku() string {
	return g.sku
}

func (g *GoodsInfo) getOrderId() string {
	return g.orderId
}

func (g *GoodsInfo) getConsigneeUserName() string {
	return g.consigneeUserAddress
}

func (g *GoodsInfo) getConsigneeUserPhone() string {
	return g.consigneeUserPhone
}

func (g *GoodsInfo) getConsigneeUserAddress() string {
	return g.consigneeUserAddress
}

func (g *GoodsInfo) DeliverGoods(good GoodsInfo) bool {
	res := fmt.Sprint(good)
	fmt.Println("模拟发货实物商品一个：", res)
	return true
}
