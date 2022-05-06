package general_ipml

import "go-DesignPattern/src/ adapter-pattern/base/mq"

func OnPOPOrderMessage(message string) {
	mq := mq.NewPOPOrder()
	mq.Getuid()
	mq.GetOrderId()
	mq.GetOrderTime()

	// ... 处理自己的业务
}
