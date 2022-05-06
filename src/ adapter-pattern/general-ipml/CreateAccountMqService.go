package general_ipml

import "go-DesignPattern/src/ adapter-pattern/base/mq"

func OnMessage(message string) {
	mq := mq.NewAccount()
	mq.GetNumber()
	mq.GetAccountDate()
	
	// ... 处理自己的业务
}
