package service

import "fmt"

func QueryUserOrderCount(userId string) int {
	fmt.Println("自营商家，查询用户的订单是否为首单：{}", userId)
	return 10
}
