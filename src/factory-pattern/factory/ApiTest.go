package factory

import "go-DesignPattern/src/factory-pattern/base"

func Test() {
	storeFactory := NewFactoryInstance()
	// 1. 优惠券
	p1 := &base.Params{
		UId:         "10001",
		CommodityId: "EGM1023938910232121323432",
		BizId:       "791098764902132",
		ExtMap:      nil,
	}
	storeFactory.GetCommodityService(1, p1)

	// 2. 实物商品
	extMap := make(map[string]string)
	extMap["consigneeUserName"] = "谢飞机"
	extMap["consigneeUserPhone"] = "15200292123"
	extMap["consigneeUserAddress"] = "吉林省.长春市.双阳区.XX街道.檀溪苑小区.#18-2109"
	p2 := &base.Params{
		UId:         "10001",
		CommodityId: "9820198721311",
		BizId:       "1023000020112221113",
		ExtMap:      extMap,
	}
	storeFactory.GetCommodityService(2, p2)

	// 3. 第三方兑换卡(爱奇艺)
	p3 := &base.Params{
		UId:         "10001",
		CommodityId: "AQY1xjkUodl8LO975GdfrYUio",
		BizId:       "",
		ExtMap:      nil,
	}
	storeFactory.GetCommodityService(3, p3)
}
