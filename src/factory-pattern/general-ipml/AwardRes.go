package general_ipml

import "fmt"

type AwardRes struct {
	info string
	code string
}

var awardResInstance *AwardRes

func NewAwardRes() *AwardRes {
	if awardResInstance == nil {
		once.Do(
			func() {
				awardResInstance = &AwardRes{
					info: "优惠券",
					code: "",
				}
			})
		return awardResInstance
	}
	return awardResInstance
}

func (a *AwardRes) AwardRes(code, info string) *AwardRes {
	a.code = code
	a.info = info
	return a
}

func (a *AwardRes) setInfo(info string) {
	a.info = info
}

func (a *AwardRes) setCode(code string) {
	a.code = code
}

func (a *AwardRes) getInfo() string {
	if a == nil {
		return ""
	}
	return a.info
}

func (a *AwardRes) getCode() string {
	if a == nil {
		return ""
	}
	return a.code
}

func (a *AwardRes) sendCoupon(uId, uuid, couponNumber string) *AwardRes {
	fmt.Println("模拟发放优惠券一张：" + uId + "," + couponNumber + "," + uuid)
	return a.AwardRes("00000", "发放成功")
}
