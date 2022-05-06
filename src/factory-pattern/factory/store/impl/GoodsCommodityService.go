package impl

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/base"
)

type GoodsCommodityService struct{}

var goodsServiceInstance *GoodsCommodityService

func NewGoodsServiceInstance() *GoodsCommodityService {
	if goodsServiceInstance == nil {
		goodsServiceInstance = &GoodsCommodityService{}
	}
	return goodsServiceInstance
}

func (g *GoodsCommodityService) SendCommodity(p *base.Params) {
	// 模拟注入
	deliverReq := base.NewGoods()
	deliverReq.SetUserName(g.queryUserName(p.UId))
	deliverReq.SetUserPhone(g.queryUserPhoneNumber(p.UId))
	deliverReq.SetSku(p.CommodityId)
	deliverReq.SetOrderId(p.BizId)
	deliverReq.SetConsigneeUserName(p.ExtMap["consigneeUserName"])
	deliverReq.SetConsigneeUserPhone(p.ExtMap["consigneeUserPhone"])
	deliverReq.SetConsigneeUserAddress(p.ExtMap["consigneeUserAddress"])

	isSuccess := deliverReq.DeliverGoods(*deliverReq)
	fmt.Println("请求参数[优惠券] => uId：{} commodityId：{} bizId：{} extMap：{}", p.UId, p.CommodityId, p.BizId, fmt.Sprint(p.ExtMap))
	fmt.Println("测试结果[优惠券]：{}", isSuccess)
	if !isSuccess {
		fmt.Println("error:", "实物商品发放失败")
	}
}

func (g *GoodsCommodityService) queryUserName(uId string) string {
	return "花花"
}

func (g *GoodsCommodityService) queryUserPhoneNumber(uId string) string {
	return "15200101232"
}
