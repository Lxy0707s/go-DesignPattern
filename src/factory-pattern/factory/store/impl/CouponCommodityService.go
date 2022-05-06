package impl

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/base"
)

type CouponCommodityService struct {
}

var couponServiceInstance *CouponCommodityService

func NewCouponServiceInstance() *CouponCommodityService {
	if couponServiceInstance == nil {
		couponServiceInstance = &CouponCommodityService{}
	}
	return couponServiceInstance
}

func (c *CouponCommodityService) SendCommodity(p *base.Params) {
	// 模拟注入
	CouponService := base.NewCoupon()
	couponResult := CouponService.SendCoupon(p.UId, p.CommodityId, p.BizId)
	fmt.Println("请求参数[优惠券] => uId：{} commodityId：{} bizId：{} extMap：{}", p.UId, p.CommodityId, p.BizId, fmt.Sprint(p.ExtMap))
	fmt.Println("测试结果[优惠券]：{}", fmt.Sprint(couponResult))
	if couponResult.GetCode() != "00000" {
		fmt.Println("error", couponResult.GetCode())
	}
}
