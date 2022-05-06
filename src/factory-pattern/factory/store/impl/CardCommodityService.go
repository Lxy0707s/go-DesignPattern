package impl

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/base"
)

type CardCommodityService struct {
}

var cardServiceInstance *CardCommodityService

func NewCardServiceInstance() *CardCommodityService {
	if cardServiceInstance == nil {
		cardServiceInstance = &CardCommodityService{}
	}
	return cardServiceInstance
}

func (c *CardCommodityService) SendCommodity(p *base.Params) {
	// 模拟注入
	iQiYiCardService := base.NewIQiYiCard()
	mobile := c.queryUserMobile(p.UId)
	iQiYiCardService.GrantToken(mobile, p.BizId)
	fmt.Println("请求参数[爱奇艺兑换卡] => uId：{} commodityId：{} bizId：{} extMap：{}", p.UId, p.CommodityId, p.BizId, fmt.Sprint(p.ExtMap))
	fmt.Println("测试结果[爱奇艺兑换卡]：success")
}

func (c *CardCommodityService) queryUserMobile(uId string) string {
	return "15200101232"
}
