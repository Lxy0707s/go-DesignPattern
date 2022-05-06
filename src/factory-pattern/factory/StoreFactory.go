package factory

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/base"
	"go-DesignPattern/src/factory-pattern/factory/store/impl"
)

type StoreFactory struct{}

var storeFactoryInstance *StoreFactory

func NewFactoryInstance() *StoreFactory {
	if storeFactoryInstance == nil {
		storeFactoryInstance = &StoreFactory{}
	}
	return storeFactoryInstance
}

func (s *StoreFactory) GetCommodityService(commodityType int, p *base.Params) {
	if p == nil {
		fmt.Println("商品信息不全，无法处理")
	}
	if commodityType == 1 {
		instance := impl.NewCouponServiceInstance()
		instance.SendCommodity(p)
		return
	}
	if commodityType == 2 {
		instance := impl.NewGoodsServiceInstance()
		instance.SendCommodity(p)
		return
	}
	if commodityType == 3 {
		instance := impl.NewCardServiceInstance()
		instance.SendCommodity(p)
		return
	}
	fmt.Println("不存在的商品服务类型")
}
