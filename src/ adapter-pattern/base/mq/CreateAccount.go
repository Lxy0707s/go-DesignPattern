package mq

import "fmt"

type CreateAccount struct {
	number      string // 开户编号
	address     string // 开户地
	accountDate string // 开户时间
	desc        string // 开户描述
}

var AccountInstance *CreateAccount

func NewAccount() *CreateAccount {
	if AccountInstance == nil {
		AccountInstance = &CreateAccount{
			number:      "0",
			address:     "beijing",
			accountDate: "0-1-0-1",
			desc:        "测试用户",
		}
	}
	return AccountInstance
}

func (c *CreateAccount) SetNumber(number string) {
	c.number = number
}

func (c *CreateAccount) SetAddress(address string) {
	c.address = address
}

func (c *CreateAccount) SetAccountDate(accountDate string) {
	c.accountDate = accountDate
}

func (c *CreateAccount) SetDesc(desc string) {
	c.desc = desc
}

func (c *CreateAccount) GetNumber() string {
	return c.number
}

func (c *CreateAccount) GetAddress() string {
	return c.address
}

func (c *CreateAccount) GetAccountDate() string {
	return c.accountDate
}

func (c *CreateAccount) GetDesc() string {
	return c.desc
}

func (c *CreateAccount) ToString() string {
	return fmt.Sprint(c)
}
