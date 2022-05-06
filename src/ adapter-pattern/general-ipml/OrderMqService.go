package general_ipml

import "go-DesignPattern/src/ adapter-pattern/base/mq"

func OnOrderMqMessage(message string) {
	// 传入的是订单信息，转换细节略
	mq := mq.NewOrderMq()

	mq.Getuid()
	mq.GetOrderId()
	mq.GetCreateOrderTime()
	// ... 处理自己的业务
}
