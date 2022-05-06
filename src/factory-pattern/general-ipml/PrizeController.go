package general_ipml

import (
	"fmt"
	"go-DesignPattern/src/factory-pattern/base"
)

func (a *AwardRes) awardToUser(req AwardReq) *AwardRes {
	reqJson := fmt.Sprint(req)
	awardRes := &AwardRes{}

	fmt.Println("奖品发放开始{}。req:{}", req.getUId(), reqJson)
	if req.getAwardType() == 1 {
		couponService := base.NewCoupon()
		couponResult := couponService.SendCoupon(req.getUId(), req.getAwardNumber(), req.getBizId())
		if couponResult.GetCode() == "000" {
			awardRes = a.AwardRes("0000", "发放成功")
		} else {
			awardRes = a.AwardRes("0001", couponResult.GetInfo())
		}
	} else if req.getAwardType() == 2 {
		goodsService := base.NewGoods()
		goodsService.SetUserName(queryUserName(req.getUId()))
		goodsService.SetUserPhone(queryUserPhoneNumber(req.getUId()))
		goodsService.SetSku(req.getAwardNumber())
		goodsService.SetOrderId(req.getBizId())
		goodsService.SetConsigneeUserName(req.getExtMap()["consigneeUserName"])
		goodsService.SetConsigneeUserPhone(req.getExtMap()["consigneeUserPhone"])
		goodsService.SetConsigneeUserAddress(req.getExtMap()["consigneeUserAddress"])
		isSuccess := goodsService.DeliverGoods(*goodsService)
		if isSuccess {
			awardRes = a.AwardRes("0000", "发放成功")
		} else {
			awardRes = a.AwardRes("0001", "发放失败")
		}
	} else if req.getAwardType() == 3 {
		bindMobileNumber := queryUserPhoneNumber(req.getUId())
		iQiYiCardService := base.NewIQiYiCard()
		iQiYiCardService.GrantToken(bindMobileNumber, req.getAwardNumber())
		awardRes = a.AwardRes("0000", "发放成功")
	}
	return awardRes
}

func queryUserName(uId string) string {
	return "花花"
}

func queryUserPhoneNumber(uId string) string {
	return "15200101232"
}
