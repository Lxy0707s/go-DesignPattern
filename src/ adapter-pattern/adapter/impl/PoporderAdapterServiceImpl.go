package impl

import "go-DesignPattern/src/ adapter-pattern/base/service"

type POPOrderAdapterServiceImpl struct{}

var pOPOrderAdapterServiceImplInstance *POPOrderAdapterServiceImpl

func NewPOPOrderAdapterServiceImpl() *POPOrderAdapterServiceImpl {
	if pOPOrderAdapterServiceImplInstance != nil {
		pOPOrderAdapterServiceImplInstance = &POPOrderAdapterServiceImpl{}
	}
	return pOPOrderAdapterServiceImplInstance
}

func (p *POPOrderAdapterServiceImpl) IsFirst(uId string) bool {
	return service.IsFirstOrder(uId)
}
