package general_ipml

import "fmt"

func (a *AwardRes) TestPrizeController() {
	ctl := awardResInstance
	fmt.Println("模拟发放优惠券测试")
	// 模拟发放优惠券测试
	req01 := NewAwardReq()
	req01.setUId("10001")
	req01.setAwardType(1)
	req01.setAwardNumber("EGM1023938910232121323432")
	req01.setBizId("791098764902132")
	awardRes01 := ctl.awardToUser(*req01)

	fmt.Println("请求参数：{}", fmt.Sprint(req01))
	fmt.Println("测试结果：{}", fmt.Sprint(awardRes01))

	fmt.Println("\r\n模拟方法实物商品")
	// 模拟方法实物商品
	req02 := NewAwardReq()
	req02.setUId("10001")
	req02.setAwardType(2)
	req02.setAwardNumber("9820198721311")
	req02.setBizId("1023000020112221113")
	req02.setExtMap(
		map[string]string{
			"consigneeUserName":    "谢飞机",
			"consigneeUserPhone":   "15200292123",
			"consigneeUserAddress": "吉林省.长春市.双阳区.XX街道.檀溪苑小区.#18-2109",
		})

	awardRes02 := ctl.awardToUser(*req02)
	fmt.Println("请求参数：{}", fmt.Sprint(req02))
	fmt.Println("测试结果：{}", fmt.Sprint(awardRes02))

	fmt.Println("\r\n第三方兑换卡(爱奇艺)")
	req03 := NewAwardReq()
	req03.setUId("10001")
	req03.setAwardType(3)
	req03.setAwardNumber("AQY1xjkUodl8LO975GdfrYUio")

	awardRes03 := ctl.awardToUser(*req03)
	fmt.Println("请求参数：{}", fmt.Sprint(req03))
	fmt.Println("测试结果：{}", fmt.Sprint(awardRes03))
}
