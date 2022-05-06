package base

import "fmt"

type CouponInfo struct {
	info string
	code string
}

var couponInstance *CouponInfo

func NewCoupon() *CouponInfo {
	if instance == nil {
		couponInstance = &CouponInfo{
			info: "优惠券",
			code: "",
		}
	}
	return couponInstance
}

func (c *CouponInfo) CouponResult(info, code string) *CouponInfo {
	c.code = code
	c.info = info
	return c
}

func (c *CouponInfo) setInfo(info string) {
	c.info = info
}

func (c *CouponInfo) setCode(code string) {
	c.code = code
}

func (c *CouponInfo) GetInfo() string {
	if c == nil {
		return ""
	}
	return c.info
}

func (c *CouponInfo) GetCode() string {
	if c == nil {
		return ""
	}
	return c.code
}

func (c *CouponInfo) SendCoupon(uId, uuid, couponNumber string) *CouponInfo {
	fmt.Println("模拟发放优惠券一张：" + uId + "," + couponNumber + "," + uuid)
	return c.CouponResult("发放成功", "00000")
}
