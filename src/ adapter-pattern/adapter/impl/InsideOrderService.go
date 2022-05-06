package impl

import "go-DesignPattern/src/ adapter-pattern/base/service"

type InsideOrderService struct{}

var InsideOrderServiceInstance *InsideOrderService

func NewInsideOrderService() *InsideOrderService {
	if InsideOrderServiceInstance == nil {
		InsideOrderServiceInstance = &InsideOrderService{}
	}
	return InsideOrderServiceInstance
}

func (i *InsideOrderService) IsFirst(uId string) bool {
	return service.QueryUserOrderCount(uId) <= 1
}
