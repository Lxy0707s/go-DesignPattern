package base

import (
	"fmt"
	"sync"
)

var once sync.Once

type IQiYiCard struct {
	desc string
}

var instance *IQiYiCard

func NewIQiYiCard() *IQiYiCard {
	if instance == nil {
		instance = &IQiYiCard{
			desc: "iqiyi",
		}
	}
	return instance
}

func (i *IQiYiCard) GrantToken(bindMobileNumber, cardId string) {
	fmt.Println("模拟发放爱奇艺会员卡一张：", bindMobileNumber, ",", cardId)
}
