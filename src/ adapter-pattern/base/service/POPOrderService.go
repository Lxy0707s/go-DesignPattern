package service

import "fmt"

func IsFirstOrder(uId string) bool {
	fmt.Println("POP商家，查询用户的订单是否为首单：{}", uId)
	return true
}
